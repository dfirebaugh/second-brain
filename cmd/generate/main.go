package main

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/niklasfasching/go-org/org"
	"github.com/sirupsen/logrus"
)

type PageData struct {
	Title     string
	Content   template.HTML
	Backlinks []Backlink
}

type Preview struct {
	Title   string
	Content string
}

type Backlink struct {
	URL     string
	Title   string
	Preview Preview
}

var (
	backlinkRegex     = regexp.MustCompile(`\[\[([^\]]+)\]\]`)
	explicitLinkRegex = regexp.MustCompile(`\[\[([^]\[]+?)\]\[([^]\[]+?)\]\]`)
	backlinkMap       = make(map[string][]Backlink)
)

func transformInternalLinksOrg(content string) string {
	content = explicitLinkRegex.ReplaceAllStringFunc(content, func(match string) string {
		sub := explicitLinkRegex.FindStringSubmatch(match)
		if len(sub) != 3 {
			return match
		}
		target := sub[1]
		text := sub[2]

		if strings.Contains(target, "://") {
			return match
		}
		slug := strings.TrimSuffix(filepath.Base(target), ".org")
		slug = strings.ReplaceAll(slug, " ", "-")
		return fmt.Sprintf(`@@html:<a href="/notes/%s/">%s</a>@@`, slug, text)
	})

	// handle [[Page Name]] style
	content = backlinkRegex.ReplaceAllStringFunc(content, func(match string) string {
		inner := backlinkRegex.ReplaceAllString(match, `$1`)
		if strings.Contains(inner, "://") {
			return match
		}
		title := inner
		slug := strings.ReplaceAll(title, " ", "-")
		return fmt.Sprintf(`@@html:<a href="/notes/%s/">%s</a>@@`, slug, title)
	})

	return content
}

func firstPass(inputDir string) error {
	return filepath.Walk(inputDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".org") {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		content := string(data)

		name := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
		sourceSlug := strings.ReplaceAll(name, " ", "-")
		sourceTitle := name

		previewText := content
		if len(previewText) > 200 {
			previewText = previewText[:200] + "..."
		}

		for _, match := range backlinkRegex.FindAllStringSubmatch(content, -1) {
			if len(match) < 2 {
				continue
			}
			inner := match[1]

			if strings.Contains(inner, "://") || strings.Contains(match[0], "][") {
				continue
			}

			targetTitle := strings.TrimSuffix(filepath.Base(inner), ".org")
			targetSlug := strings.ReplaceAll(targetTitle, " ", "-")

			backlinkMap[targetSlug] = append(backlinkMap[targetSlug], Backlink{
				URL:   fmt.Sprintf("/notes/%s/", sourceSlug),
				Title: sourceTitle,
				Preview: Preview{
					Title:   sourceTitle,
					Content: previewText,
				},
			})
		}

		// handle [[link][text]] links
		for _, match := range explicitLinkRegex.FindAllStringSubmatch(content, -1) {
			if len(match) < 3 {
				continue
			}
			target := match[1]

			if strings.Contains(target, "://") {
				continue
			}

			targetTitle := strings.TrimSuffix(filepath.Base(target), ".org")
			targetSlug := strings.ReplaceAll(targetTitle, " ", "-")

			backlinkMap[targetSlug] = append(backlinkMap[targetSlug], Backlink{
				URL:   fmt.Sprintf("/notes/%s/", sourceSlug),
				Title: sourceTitle,
				Preview: Preview{
					Title:   sourceTitle,
					Content: previewText,
				},
			})
		}

		return nil
	})
}

func main() {
	inputDir := "./notes"
	outputDir := "./.dist"

	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	if err := firstPass(inputDir); err != nil {
		fmt.Println("Error in first pass:", err)
		return
	}

	pageTpl, err := template.ParseFiles("templates/page.html")
	if err != nil {
		fmt.Println("Error loading page template:", err)
		return
	}
	indexTpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println("Error loading index template:", err)
		return
	}

	var pages []string
	err = filepath.Walk(inputDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".org") {
			return nil
		}
		name := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
		slug := strings.ReplaceAll(name, " ", "-")
		if err := generateHTMLPage(pageTpl, path, outputDir, slug); err != nil {
			return err
		}
		pages = append(pages, slug)
		return nil
	})
	if err != nil {
		fmt.Println("Error building pages:", err)
		return
	}

	if err := generateIndexPage(indexTpl, pages, outputDir); err != nil {
		fmt.Println("Error building index:", err)
	}

	assetsDst := filepath.Join(outputDir, "assets")
	os.MkdirAll(assetsDst, 0o755)
	if err := CopyDir(assetsDst, "./assets"); err != nil {
		logrus.Error(err)
	}

	if err := os.WriteFile(filepath.Join(outputDir, "CNAME"), []byte("brain.dustinfirebaugh.com"), 0o644); err != nil {
		logrus.Error(err)
	}
}

func translateCodeblocks(htmlStr string) string {
	opening := regexp.MustCompile(`(?s)<div class="src src-([^"]+)">\s*<div class="highlight">\s*<pre>`)
	htmlStr = opening.ReplaceAllString(htmlStr, `<pre><code class="language-$1">`)

	closing := regexp.MustCompile(`(?s)</pre>\s*</div>\s*</div>`)
	htmlStr = closing.ReplaceAllString(htmlStr, `</code></pre>`)

	return htmlStr
}

func generateHTMLPage(tpl *template.Template, inputFilePath, outputDir, slug string) error {
	raw, err := os.ReadFile(inputFilePath)
	if err != nil {
		return fmt.Errorf("read org file: %w", err)
	}

	processed := transformInternalLinksOrg(string(raw))

	cfg := org.New()
	doc := cfg.Parse(strings.NewReader(processed), inputFilePath)
	htmlBytes, err := doc.Write(org.NewHTMLWriter())
	if err != nil {
		return fmt.Errorf("convert org to HTML: %w", err)
	}

	outDir := filepath.Join(outputDir, "notes", slug)
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("mkdir output dir: %w", err)
	}

	outFile := filepath.Join(outDir, "index.html")
	f, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("create html file: %w", err)
	}
	defer f.Close()

	outContent := translateCodeblocks(string(htmlBytes))
	pageData := PageData{
		Title:     slug,
		Content:   template.HTML(outContent),
		Backlinks: backlinkMap[slug],
	}
	if err := tpl.Execute(f, pageData); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}
	return nil
}

func generateIndexPage(tpl *template.Template, pages []string, outputDir string) error {
	f, err := os.Create(filepath.Join(outputDir, "index.html"))
	if err != nil {
		return fmt.Errorf("create index page: %w", err)
	}
	defer f.Close()
	return tpl.Execute(f, pages)
}

func CopyDir(dst, src string) error {
	src = filepath.Clean(src)
	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		outPath := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(outPath, info.Mode())
		}
		if info.Mode()&os.ModeSymlink != 0 {
			linkTarget, err := os.Readlink(path)
			if err != nil {
				return err
			}
			return os.Symlink(linkTarget, outPath)
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			return err
		}
		outFile, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()
		_, err = io.Copy(outFile, in)
		return err
	})
}
