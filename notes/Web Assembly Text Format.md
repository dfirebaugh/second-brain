# Web Assembly Text Format

Textual representation of the wasm binary format.

https://developer.mozilla.org/en-US/docs/WebAssembly/Understanding_the_text_format

## Demos

Here are some excellent demos written with Web Assembly Text format: https://github.com/binji/raw-wasm


# Writing raw WAT files
Each web assembly module (in wat) is represented as an s-expression (which is basically a tree).

most basic module
```
(module)
```

## Types

* `i32`  32-bit integer
* `i64` 64-bit integer
* `f32` 32-bit float
* `f64` 64-bit float

