# gshell

gshell, or GoShell is a small shell project that I am building as a side project. The plan is to implement some commands whenever I have the time and slowly increase the complexity of this project. The goal is to try and make this shell as full-featured as possible while not just `exec`ing (This may change in the future).

Contributions are also welcome!

## Architecture

This shell is built using 3 goroutines, or threads:

1. **Output Thread**: A worker thread that receives results represented in string form that only handles outputting these results to STDOUT
2. **Worker Thread**: A worker thread that takes user input and depending on the user input, will try to generate the correct answer to the user's request and sends it along to the `Output Thread`
3. **Main Thread**: The main thread that handles STDIN (User input), and passes the user input to `Worker Thread`

The way these threads talk to each other is a combination of the built-in `golang` mechanism, channels, and Mutexes.

I `lock` the mutex right before I print the prompt in the **Main Thread** so that the **Output Thread** can print all that it needs to print, and then when the previous command finishes, I `unlock` the mutex so that the prompt in the **Main Thread** can print, and then we can receive new user input. 

The 3 threads use channels to communicate amongst each other in the following manner:

1. `Main thread` sends user input to Worker thread via channel
2. `Worker Thread` receives command and processes it. Once the `Worker Thread` is done doing work, or whenever it has something to output, it sends the output to the `Output Thread` via channel
3. The `Output Thread` receives what it needs to print, and when the `Output Thread` receives a stop token, it Unlocks the mutex so that the `Main Thread` can again receive user input

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