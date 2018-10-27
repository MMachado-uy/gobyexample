package main

import "fmt"
import "math"
import "time"

const s string = "constant"

func main() {
  helloWorld()
  values()
  variables()
  constants()
  forLoop()
  ifElse()
  switchControl()
  arrays()
}

func echo(str ...interface{}) {
  fmt.Println(str...)
}

func helloWorld() {
  echo("Hello World")
}

func values() {
  echo("go" + "lang")
  echo("1 + 1 = ", 1+1)
  echo("7.0/3.0 = ", 7.0/3.0)
  echo(true && false)
  echo(true || false)
  echo(!true)
}

func variables() {
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
}

func constants() {
  echo(s)

  const n = 500000000

  const g = 3e20 / n
  echo(g)
  echo(int64(g))
  echo(math.Sin(n))
}

func forLoop() {
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
}

func ifElse() {
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

func switchControl() {
  i := 2
  echo("Write ", i, " as ")
  switch i {
  case 1:
    echo("one")
  case 2:
    echo("two")
  case 3:
    echo("three")
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    echo("It's the weekend")
  default:
    echo("It's a weekday")
  }

  t := time.Now()
  switch {
  case t.Hour() < 12:
    echo("It's before noon")
  default:
    echo("It's after noon")
  }

  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      echo("I'm a bool")
    case int:
      echo("I'm an int")
    default:
      echo("Don't know type %T/n", t)
    }
  }

  whatAmI(true)
  whatAmI(1)
  whatAmI("hey")
}

func arrays() {
  var a [5]int
  echo("emp:", a)

  a[4] = 100
  echo("set:", a)
  echo("get:", a[4])

  echo("len:", len(a))

  b := [5]int{1, 2, 3, 4, 5}
  echo("dcl:", b)

  var twoD [2][3]int
  for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
      twoD[i][j] = i + j
    }
  }
  echo("2d:", twoD)
}
