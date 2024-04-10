# second-brain

This is a repository for my notes.

[https://brain.dustinfirebaugh.com](https://brain.dustinfirebaugh.com)


### Generate the site 

```bash 
go run ./cmd/generate
```

> outputs to the `.dist` directory

### Run Dev server

```bash
go run ./cmd/devserver
```

### Conventions

The `notes` dir is a flat directory of markdown files.  These markdown files will be parsed into static html files.

Double square brackets represents an internal link. e.g. `[[Some Internal Page]]` will output `<a href="/notes/Some-Internal-Page">Some Internal Page</a>`

`Backlinks` are pages that link to the current page.
