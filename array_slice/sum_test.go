package main

import "testing"
import "reflect"

func TestSumAll(t *testing.T){
  t.Run("simple test variadic functions", func(t *testing.T){
    a := []int{1,2}
    b := []int{3,4}

    got := SumAll(a,b)
    want := []int{3,7}

    assertCodeMessageSumAll(t,got,want,a,b)
  })
}

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

func assertCodeMessageSumAll(t testing.TB, got, want, givenA, givenB []int){
  t.Helper()
  if !reflect.DeepEqual(got, want){
    t.Errorf("got %v want %v given %v and %v", got, want, givenA, givenB)
  }
}
