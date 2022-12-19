package integers 

import (
  "fmt"
  "testing"
)
func  TestAdder(t *testing.T){
  t.Run("simple int test", func(t *testing.T){
    got := Add(1,2)
    expected := 3

    assertCodeMessage(t,got,expected)
  })
}

func ExampleAdd(){
  sum := Add(1,1)
  fmt.Println(sum)
  //Output: 2
}

func assertCodeMessage(t testing.TB, got, expected int) {
  t.Helper()
  if got != expected {
    t.Errorf("got %d want %d", got, expected)
  }
}
