## Programs are like functions, There are two kinds of functions
1. Something that takes input parameters, and computes an output result
- Usually either a query or a transformation
- What is a query?
- What is a transformation?
2. Something that performs some work, but doesnt output a result
- Usually a side-effect
- What is a side-effect?

## Input/Output (I/O) Strata
- Command-line
- Hypertext Transport Protocol
- Graphical User Interface

## Command-line programs
This is easy to relate to from the computer-enthusiast's perspective.
- cd doesn't return a result, it changes the terminal's directory. It performs the effect of mutating the current working directory.
- ls takes no arguments, but produces the output of the current working directory's contents
- tar/unzip takes some files as input and produces a compressed file
Mechanism: STDOUT/STDIN buffers, as managed by OS interrupts

## HTTP programs
This is hard to relate to.
- Web Server produces text in any format
- Web Client reads some text and draws it
Mechanism: Some network card buffers, as managed by TCP/IP and the OS

## Graphical programs
Easiest to relate to.
File Explorer:
    - Clicking a folder changes what items are subject of deletes or moves
    - Double-Clicking a folder changes which folder we see into.
    - Compressing a folder produces a compressed file in the window
Mechanism: The GPU

## What is the command line? Which one? Shell?
## What is the internet?
## What is the GPU?
