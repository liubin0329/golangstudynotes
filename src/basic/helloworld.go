package main

import "fmt"

func main() {
    // os.Args[0] = helloworld.go 
    args := os.Args[1:]
    
    fmt.Printf("hello, world\n")
}
