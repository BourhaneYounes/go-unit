package main

import "testing"

func TestSum(t *testing.T){

  t.Run("simple array test", func(t *testing.T){

    numbers := [5]int{1,2,3,4,5}
  
    got := Sum(numbers)
    want := 15

    assertCodeMessage(t,got,want,numbers)
  })
}

func assertCodeMessage(t testing.TB,got,want int, given [5]int){
  t.Helper()
  if got != want{
    t.Errorf("got %d want %d given %v", got, want, given)
  }
}
