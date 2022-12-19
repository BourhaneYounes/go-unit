package main

import "testing"

func TestSum(t *testing.T){
  t.Run("collection of numbers any size(omitted size[slices])", func (t *testing.T){
    numbers := []int{1,2,3}

    got := Sum(numbers)
    want := 6
    assertCodeMessage(t,got,want,numbers)
  })
}

func assertCodeMessage(t testing.TB,got,want int, given []int){
  t.Helper()
  if got != want{
    t.Errorf("got %d want %d given %v", got, want, given)
  }
}
