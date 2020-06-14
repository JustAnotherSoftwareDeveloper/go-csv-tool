# Introduction

This project was given to me as part of a coding interview assignment. I have experience in Java and Python, but not go. While I'm not worried about my ability to get the project done, I highly doubt that my code will be "clean" by go standards. Getting the hand of a language is one thing; Getting a hang of its code stlye is a whole other that in my experience takes more that an hour


## Notes
* I didn't want to spend extra time figuring out how imports work in go so I did everything in one file. Obviously I wouldn't do that in actual code
* I did not use a makefile. Instead I just did `go run cmd/main.go` from the root directory
* The go version I used is 1.14.4
  

# Assignment Below

## Scenario

You are given a very simple log of events(`server_log.csv`) for a server that allows users to upload or download files.


Using this log file as input, write a Go program that will output the following metrics:
1. How many users accessed the server?
2. How many uploads were larger than `50kB`?
3. How many times did `jeff22` upload to the server on April 15th, 2020?


## Details
The log is represented by a CSV formatted file where column data is ordered: `timestamp`, `username`, `operation`, `size`.

- The `timestamp` is recorded in the UNIX date format.
- The `username` is a unique identifier for the user.
- The `operation` indicates if an `upload` or `download` occurred, no other values will appear in this column.
- The `size` is an integer reflecting file size in `kB`.

You may use whatever packages you feel will be useful to solve this challenge.

## Important Note
The intention of this coding challenge is to gauge a candidate's ability to program in Go, write clean code, and discuss their work with others. While we would like to see your best work, we don't want candidates to spend a large amount of their free time optimizing their solution to this challenge. So, our expectation is not necessarily to receive an optimal or even completely correct solution. With this in mind, we recommend spending about an hour on this challenge.