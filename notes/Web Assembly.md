---
tags:
 - programming
---

# Web Assembly
*  [[Web Assembly Text Format]]

I did a write up Web Assembly when i first took a look at it on my [blog](https://dustinfirebaugh.com/blog/Web_Assembly/).

### High Level
>WebAssembly (abbreviated Wasm) is a binary instruction format for a stack-based virtual machine. Wasm is designed as a portable target for compilation of high-level languages like C/C++/Rust, enabling deployment on the web for client and server applications — [Web Assembly Home Page](https://webassembly.org/)


### Microcontrollers
I think web assembly also has an interesting role to play in the microcontroller space.

I'm specifically interested in this because it will allow for sandboxed apps on a microcontroller that would also easily run in a different environment (e.g. the browser).

It's also dynamically loadable.  So, you could send a wasm file from one microcontroller to another and run it.

#### Web Assembly Interpreter
[https://github.com/wasm3/wasm3](https://github.com/wasm3/wasm3)

| name | size |
|---------|-----|
| [wasm.h](https://github.com/wasm3/wasm3-arduino/blob/main/src/wasm3.h) | 18KB |

[an example of a wasm vm on a microcontroller](https://github.com/vshymanskyy/wasm3_dino_rpi_pico/blob/main/dino_vm.cpp)

### Smart Contracts
It seems several major smart contract crypto currencies are moving to a wasm vm.


### Web Assembly System Interface (WASI)
[[Web Assembly System Interface]]
