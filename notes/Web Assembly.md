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

### Smart Contracts
It seems several major smart contract crypto currencies are moving to a wasm vm.

