# Graph Theory

### Project 2018

This repository contains code written in the programming language Go.

Author: Keith Daly.

Innitiated: March 2018.

Student @GMIT, Galway, Ireland. 

Year three - Semester two.

## Problem Statement

You must write a program in the Go programming language that can
build a non-deterministic finite automaton (NFA) from a regular expression,
and can use the NFA to check if the regular expression matches any given
string of text.

[Full Project spec](https://github.com/dalykeith/graph-theory/blob/master/project.pdf)

### Prerequisites

[GoLang](https://golang.org/) - Needed on device to run the programs

## Installing

* Clone the repository in terminal/cmd by executing the following command

```
git clone https://github.com/dalykeith/graph-theory.git
```

* Locate File

```
cd graph-theory
```

* Build Program

```
go build runner.go
```

* Run Program

```
./runner
```

## Output

### Infix to Postfix

* a.b.c* to ab.c*.

* (a.(b|d))* to abd|.*

* a.(b|d).c* to abd|.c*.

* a.(b.b)+.c to abb.+.c.

### Postfix to NFA:

* ab.c*| to match cccc

###  Understanding the project

First to understand Shunting yard algorithm and Thompsons construction (Links below). The Shunting yard algorithm changed infix regular expression to postfix regular expressions. The Thompson construction then determines what each character represents and the action that will occur when the action is achieved.

###  Research & Adapted from

* [Regular Expression](https://swtch.com/~rsc/regexp/regexp1.html) - Russ Cox on regex / Thompson NFA
* [Thompson's Construction](https://web.microsoftstream.com/video/946a7826-e536-4295-b050-857975162e6c) -  Ian McLoughlin 
* [NFA Moves](https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton#NFA_with_%CE%B5-moves) - NFA automation
* [NFA Example](https://github.com/kkdai/nfa) - kkdai
* [Shunting Yard](http://jacobappleton.io/2015/07/02/regex-ii-the-shunting-yard-algorithm/)   - Jacob Appleton
* [Shunting Yard Example](https://github.com/mgenware/go-shunting-yard) - mgenware





### License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/dalykeith/graph-theory/blob/master/LICENSE) file for details