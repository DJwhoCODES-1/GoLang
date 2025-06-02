# GoLang

## Setup

1. Install : https://go.dev/doc/install
2. Check : go version

## Features

1. Concurrency
2. Go Routines
3. Statically Typed
4. Fast Compilation Times : efficient compiler design
5. Efficient String Processing
6. Comprehensive Set Of String Functions
7. Uses Garbage Collection
8. Go = C + Strings + Garbage Collection + Concurrency
9. Lightweight

## Theory

1. GO is a concurrent programming language
2. Is is statically typed : which means each variable has a type defined at compile time, that cannot be changed over time
3. Concurrency : It refers to the ability of a program to execute multiple tasks independently but not necessarily simultaneously
4. Go Routine : It is a lightweight thread managed by Go Runtime
5. Lightweight : All the tools Go uses consume very less memory. Ex : Goroutines have a small initial stack size (typically around 2 KB)
6. "main" function should be a part of "main" package
7. Directories(folder) and packages are tightly coupled
8. In a single directory all the files must contain same package_name
9. Every file must declare a package at the top

## Structure

1. go -> bin : executable binaries (packages we install)
   go -> pkg : compiled package files (compiled code of our src-code)
   go -> src : code (where we write code)
2. (create a folder) -> (go mod init "pkg_name")
3. fmt = formatted I/O : printing output, reading input, and formatting strings.
4. package main
   import "fmt"
   func main() {
   fmt.Println("Hello Golang World!")
   }

## Packages & Imports

1. In Go, packages are used instead of classes. There are no concepts like OOPs.
2. Each Go file must belong to some package, and it should start with keyword "package".
3. The "main" package is a special package in Go. An executable program must contain a main package.

## NOTES

1. In order to export a variable, its first letter must be capitalised
2. If the first letter of the variable is small, it is accessible only in that file
3. func PublicExportableFunction(){
   fmt.Println("This is a public function.")
   }
4. func privateNonExportableFunction(){
   fmt.Println("This function cannot be exported.")
   }
5. Named function cannot be created inside another named function
