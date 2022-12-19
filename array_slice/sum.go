package main

func Sum(numbers []int) int {
  sum := 0
  for _,number := range numbers {
    sum += number
  }
  return sum
}

func SumAll(numbers ...[]int) []int{
  var res []int

  for _, number := range(numbers){
    res = append(res, Sum(number))
  }
  return res
}
