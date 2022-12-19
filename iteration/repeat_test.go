package iteration

import (
  "fmt"
  "testing"
)

func TestRepeat(t *testing.T){
  t.Run("simple test iteration", func(t *testing.T){
       
    repeated := Repeat("a",4)
    expected := "aaaa"

    assertCodeMessage(t,repeated,expected)
  });
}

func BenchmarkRepeat(b *testing.B){
  for i := 0; i < b.N; i++{
    Repeat("a",4)
  }
}

func ExampleRepeat(){
  repeated := Repeat("a",4)
  fmt.Println(repeated)
  //Output: aaaa
}

func assertCodeMessage(t testing.TB, got, expected string){
  t.Helper()
  if got != expected{
    t.Errorf("got %q want %q", got, expected)
  }
}
