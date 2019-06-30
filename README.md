# gshell

gshell, or GoShell is a small shell project that I am building as a side project. The plan is to implement some commands whenever I have the time and slowly increase the complexity of this project. The goal is to try and make this shell as full-featured as possible while not just `exec`ing (This may change in the future).

Contributions are also welcome!

## Architecture

This shell is built using 3 goroutines, or threads:

1. A worker thread that receives results represented in string form only handles outputting these results to STDOUT
2. A worker thread that takes user input and depending on the user input, will try to generate the correct answer to the user's request and sends it along to thread #1
2. The main thread that handles STDIN (User input), and passes the user input to thread #2

## Feature Checklist

- [ ] Colorized output
- [ ] Tab completion

## Command Checklist

- [x] ls: List files in current directory
- [x] ls \<path\>: List files at directory
- [x] cd: Change directory to home directory
- [x] cd \<path\>: Change directory
- [x] pwd: Returns current working directory
- [x] cat \<file\>: Print contents of file to stdout

## How to Contribute

If you want to implement a command, please create an issue, and I'll try to reply ASAP and give the okay.

Tests would be nice.

So far, this project is basically vanilla go, so once you clone this, it should be good to go. (pun intended.)