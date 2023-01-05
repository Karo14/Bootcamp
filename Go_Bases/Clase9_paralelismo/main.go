package main

import "fmt"

func main() {
   // Creamos una canalización de enteros
   c := make(chan int)

   // Creamos cuatro gorutinas
   go func() {
      for i := 0; i < 10; i++ {
         c <- i
      }
      close(c)
   }()
   go func() {
      for i := 10; i < 20; i++ {
         c <- i
      }
      close(c)
   }()
   go func() {
      for i := 20; i < 30; i++ {
         c <- i
      }
      close(c)
   }()
   go func() {
      for i := 30; i < 40; i++ {
         c <- i
      }
      close(c)
   }()

   // Leemos los valores enviados por las gorutinas
   for i := range c {
      fmt.Println(i)
   }
}