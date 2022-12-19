package main

func Sum(numbers []int) int {
  sum := 0
  for _,number := range numbers {
    sum += number
  }
  return sum
}

func SumAll(numbers ...[]int) []int{
  res := make([]int, len(numbers))

  for i, number := range(numbers){
    res[i] = Sum(number)
  }
  return res
}
