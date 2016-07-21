## A simple video game made with SDL2 and Go

![Demo gif](https://github.com/Charlesworth/goGame/blob/master/demo.gif?raw=true)

This project is a very simple proof of concept game, using SDL2 and the Go programming language. The main logic of the program can be found in the gameView.go file. I've only tested it on windows 10 with access to the valid SDL2 libraries.

### My conclusions from the project on using Go for Games:
- The simplicity of the language made things quick and simple to reason about, though generics and the trait system from Rust was missed in this example.
- Having garbage collection means that it would be best to stick with manually allocated memory languages for latency sensitive games (e.g. first person shooters), but it could be a interesting choice for games that aren't reaction based (e.g. turn based stratergy).
- The Go SDL2 library is nice to use, though it has some strange implementation details due to it being a wrapper around the C SDL2 library. Still, I think its the best way to draw to the screen with Go when not in a web browser.
- While using the SDL2 library slows the compilation time, its still never took over two seconds.
