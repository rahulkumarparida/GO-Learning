# Basic GO

This repository contains my learnings and practice work in Go.

It includes small programs, experiments, and mini projects that I am building while learning the Go programming language.




# Things About Concurrency , Mutexes , Channels one should never forget.

 #### Concurrency
    -  "Do not communicate by sharing memory; instead, share memory by communicating."

 #### Mutexes
    - "Mutex works for data but does not care which came first... it only helps when two goroutines almost collide with each other; it blocks one routine at random and lets the other pass first."

 #### Channels
    - "In channels, it's like multiple goroutines running different functions need to communicate for data... one writes, the other reads, then that writes and later reads, and so on so forth."