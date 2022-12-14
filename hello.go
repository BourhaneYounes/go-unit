package main

import "fmt"

const englisHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func hello(name, language string) string{
  
  if name == ""{
    name = "world"
  }

  prefix := englisHelloPrefix

  switch language {
  case spanish:
    prefix = spanishHelloPrefix
  case french:
    prefix = frenchHelloPrefix
  }

  return prefix + name
}

func main(){
  fmt.Println(hello("",""))
}
