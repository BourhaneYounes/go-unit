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
  return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string){
  switch language {
  case spanish:
    prefix = spanishHelloPrefix    

  case french:
    prefix = frenchHelloPrefix

  default:
    prefix = englisHelloPrefix
  }

  return
}

func main(){
  fmt.Println(hello("",""))
}
