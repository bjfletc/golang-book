package main

import "fmt"

func main() {
  fmt.Println("1 + 1 =", 1 + 1)
  fmt.Println("1.0 + 1.0 =", 1.0 + 1.0)

  fmt.Println("The lenght of 'Hello, World' is :", len("Hello, World"))

  fmt.Println(len("Hello, World!"))
  fmt.Println("Hello, World!"[1])
  fmt.Println("Hello, " + "World!")
}
