package iteration

const repeated_count = 4

func Repeat(char string, count int) string{
 
  var repeated string
  if count <= 0{
    count = repeated_count
  }
    
  for i := 0; i < count; i++{
    repeated += char
  }
  return repeated
}
