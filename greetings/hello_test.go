package main

import "testing"

func  TestHello(t *testing.T) {
  t.Run("saying hello to people", func(t *testing.T){
    got := hello("Chris","")
    want := "Hello, Chris"

    assertCodeMessage(t,got,want)
  })

  t.Run("say 'hello, world' when an empty string is supplied", func(t *testing.T){
    got := hello("","")
    want := "Hello, world"

    assertCodeMessage(t,got,want)
  })

  t.Run("hello in spanish", func(t *testing.T){
    got := hello("elodie", "Spanish")
    want := "Hola, elodie"

    assertCodeMessage(t,got,want)
  })

  t.Run("hello in french", func(t *testing.T){
    got := hello("marc","French")
    want := "Bonjour, marc"

    assertCodeMessage(t,got,want)
  })
}

func assertCodeMessage(t testing.TB, got, want string){
  t.Helper()
  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
