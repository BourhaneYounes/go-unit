package main

import "fmt"

const englisHelloPrefix = "Hello, "
const spanishHolaPrefix = "Hola, "
const spanish = "Spanish"

func hello(name, language string) string{
  if name == ""{
    name = "world"
  }
  if language == spanish{
    return spanishHolaPrefix + name
  }
  return englisHelloPrefix + name
}

func main(){
  fmt.Println(hello("",""))
}
