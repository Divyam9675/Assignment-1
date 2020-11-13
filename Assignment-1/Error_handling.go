package main

import "errors"
import "fmt"
import "math"

func calc(value float64)(float64, error) {
   if(value < 0){
      return 0, errors.New("Math: negative number passed to Sqrt")
   }
   return math.Sqrt(value), nil
}
func main() {
   result, err:= calc(-1)

   if err != nil {
      fmt.Println(err)
   } else {
      fmt.Println(result)
   }
   
   result, err = calc(9)

   if err != nil {
      fmt.Println(err)
   } else {
      fmt.Println(result)
   }
}