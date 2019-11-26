package main

import("fmt")

func main() {
  c := make(chan int)

  c = gen(c)
  /*
  go func() {
    c<- 42
  }()
*/
  var v int
  var ok bool
  for {
      v,ok = <-c
  if ok == false{
    close(c)
    return
  } else {
    fmt.Println(v,ok)
  }
  }

  fmt.Println(v,ok)

//  close(c)

//  v,ok = <-c
  fmt.Println(v, ok)

}


func gen(c chan int) chan int {
  go func() {
    for i:=0; i<100; i++ {
      c <-i
    }
  }()
  return c
}
