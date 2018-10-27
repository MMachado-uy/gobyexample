package main

import "fmt"
import "math"

const s string = "constant"

func main() {
  echo("Hello World")
  echo("go" + "lang")
  echo("1 + 1 = ", 1+1)
  echo("7.0/3.0 = ", 7.0/3.0)
  echo(true && false)
  echo(true || false)
  echo(!true)

  var a = "initial"
  echo(a)

  var b, c int = 1, 2
  echo(b, c)

  var d = true
  echo(d)

  var e int
  echo(e)

  f := "short"
  echo(f)

  echo(s)

  const n = 500000000

  const g = 3e20 / n
  echo(g)
  echo(int64(g))
  echo(math.Sin(n))

  i := 1
  for i <= 3 {
    echo(i)
    i = i + 1
  }

  for j := 7; j <= 9; j++ {
    echo(j)
  }

  for {
    echo("loop")
    break
  }

  for n := 0; n <= 5; n++ {
    if n%2 == 0 {
      continue
    }
    echo(n)
  }

  if 7%2 == 0 {
    echo("7 is even")
  } else {
    echo("7 is odd")
  }

  if 8%4 == 0 {
    echo("8 is divisible by 4")
  }

  if num := 9; num < 0 {
    echo(num, "is negative")
  } else if num < 10 {
    echo(num, "has 1 digit")
  } else {
    echo(num, "has multiple digits")
  }
}

func echo(str ...interface{}) {
  fmt.Println(str)
}
