package main

import "fmt"

func factorial(x int)int{

  var  i int
  var fact int
  fact = x
  for i=1; i<x; i++{

    fact= fact * (x-1)
  
  }
  return fact
}

func main(){

  var n int 
  var C float64
  n=6
  f1:= factorial(2*n)
  f2:= factorial(n + 1)
  f3:= factorial(n)
  C = (float64(f1)/float64(f2*f3))
  fmt.Printf("%f", C)

}
