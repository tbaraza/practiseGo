package main // this declaration shows that this is a standalone executable program

import (
  "fmt"
  "runtime"
  "reflect"
  )
var (
  name, location string // global variable
  number int32 // package level
  fly = 0.32
)
func main() {
  house := 20000 // shorthand declare and initialize only works within a function
  name = "Today"
  fmt.Println("Hello from", runtime.GOOS)
  fmt.Println("Hello", name, "and is  of type", reflect.TypeOf(name))
  fmt.Println(number)
  fmt.Println(location)
  fmt.Println("fly is of type", reflect.TypeOf(fly))
}
