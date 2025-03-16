# Golang

## History and Introduction

* Go is a procedural programming language. It was developed in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google but launched in 2009 as an open-source programming language

* Programs are assembled by using packages, for efficient management of dependencies. 
* This language also supports environment adopting patterns alike to dynamic languages. 
* For eg., type inference (y := 0 is a valid declaration of a variable y of type int).

* Go is expressive, concise, clean, and efficient. 

* It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language

* Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction.

### Concurrency:
* Go is known for its support for concurrency, which is the ability to run multiple tasks simultaneously. 
* Concurrency is achieved in Go through the use of __Goroutines__ and __Channels__, which allow you to write code that can run multiple operations at the same time. 
* This makes Go an __ideal choice for building high-performance and scalable network services,__ as well as for solving complex computational problems.

### Garbage collection:
* Another important feature of Go is its garbage collection, which automatically manages memory for you. 
* This __eliminates the need for manual memory management,__ reducing the likelihood of memory leaks and other bugs that can arise from manual memory management.

### Some features to highlight:

1. Simplicity: Go is designed to be easy to learn and use. Its syntax is simple and straightforward, making it a good choice for beginners and experienced programmers alike.

2. Concurrency: Go has built-in support for concurrency, allowing developers to write efficient and scalable code for multicore and distributed systems.


3. Garbage collection: Go has automatic memory management, which frees developers from having to worry about memory allocation and deallocation.

4. Fast compile times: Go has a fast compiler, which makes it easy to iterate quickly during development.

5. Cross-platform support: Go can be compiled to run on many different platforms, including Windows, Linux, and macOS.

6. Strong typing: Go is a statically typed language, which helps catch errors at compile time rather than at runtime.

7. Go has a large and growing community of developers and is used by many well-known companies, including Google, Uber, and Dropbox.

<hr>

### Important points to remember

1. Go is statically typed which means - the type of variable must be decaled before it is used

2. Go has built-in garbage collector - it automatically frees up memory when it is no longer required

3. Go has strong support of concurrency - allows developers to write efficient and scalable code for multicore & distributed system

4. It has minimalistic syntax - that is easy to learn and read

5. It has a fast compiler - generates code that is optimized for modern hardware architectures

6. It has standard library that provides support for a wide range of functionalty including networking, encryption  and file handling

```go
1. package main
2. import "fmt"

3. func main() {
4.    fmt.Println("Hello, World!")
5. }
```
__Output:__ > Hello, World!

__Explanation:__

Line1: contains `package main` of program. It's initial point to run the program

Line2:  `import fmt`  - used to include external package that provide additional functionality beyond the built-in language feature. <br> `fmt` is package which provides functions for formatting input and output.

Line3: `main function`  - beginning of execution of program

Line4: `fmt.Println()` - std library func to print something as output. Here `fmt` has transmitted Println method which is used to display output.


__Single Line Comment:__

__Syntax:__

```go
// single line comment 
```
__Multi-line Comment:__ 

__Syntax:__

```go
/* multiline comment */
```
<hr>

### Some popular Applications developed in Go Language

__Docker:__ a set of tools for deploying linux containers

__Openshift:__ a cloud computing platform as a service by Red Hat.

__Kubernetes:__ The future of seamlessly automated deployment processes

__Dropbox:__ migrated some of their critical components from Python to Go.

__Netflix:__ for two part of their server architecture.

__InfluxDB:__ is an open-source time series database developed by InfluxData.

__Golang:__ The language itself was written in Go.

### how golang fits in the DevOps Ecosystem

1. Custom infrastructure automation
2. Kubernetes development
3. Miscroservices architecture
4. CI/CD pipeline customization
5. Monitoring and observability
6. Cloud-native components 
7. Automation tooling
8. Security enhancements
9. Performance and scalibility
10. Rich library ecosystem
