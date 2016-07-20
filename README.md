## A simple video game made with SDL2 and Go

This project is a very simple proof of concept game, using SDL2 and the Go programming language. The main logic of the program can be found in the gameView.go file.

My conclusions from the project on using Go for Games:
- The simplicity of the language made thing quick and simple to reason about
- Having garbage collection means that it would be best to stick with manually allocated memory languages for latency sensitive games (e.g. first person shooters), but it could be a interesting choice for games that aren't reaction based (e.g. turn based stratergy)
