/**
 * golang tests and examples. Based in the exercises of:
 * gobyexample.com
 * @author Mauricio Machado <macmauricio@gmail.com>
 * @since Oct, 2018
 */
package main

import (
  "fmt"
  "math"
  "strconv"
  "time"
)

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
  slices()
  maps()
  ranges()
  functions()
}

func echo(str ...interface{}) {
  fmt.Println(str...)
}

func heading(chap int, title string) {
  var c string
  if chap < 10 {
    c = "0" + strconv.Itoa(chap)
  } else {
    c = strconv.Itoa(chap)
  }
  echo()
  echo("************************")
  echo(c, "-", title)
  echo()
}

func helloWorld() {
  heading(1, "Hello World")

  echo("Hello World")
}

func values() {
  heading(2, "Values")

  echo("go" + "lang")
  echo("1 + 1 = ", 1+1)
  echo("7.0/3.0 = ", 7.0/3.0)
  echo(true && false)
  echo(true || false)
  echo(!true)
}

func variables() {
  heading(3, "Variables")

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
  heading(4, "Constants")

  echo(s)

  const n = 500000000

  const g = 3e20 / n
  echo(g)
  echo(int64(g))
  echo(math.Sin(n))
}

func forLoop() {
  heading(5, "For Loops")

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
  heading(6, "If/Else")

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
  heading(7, "Switch")

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
      fmt.Printf("Don't know type %T\n", t)
    }
  }

  whatAmI(true)
  whatAmI(1)
  whatAmI("hey")
}

func arrays() {
  heading(8, "Arrays")

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

func slices() {
  // https://blog.golang.org/go-slices-usage-and-internals
  heading(9, "Slices")

  s := make([]string, 3)
  echo("emp", s)

  s[0] = "a"
  s[1] = "b"
  s[2] = "c"

  echo("set:", s)
  echo("get:", s[2])

  echo("len:", len(s))

  s = append(s, "d")
  s = append(s, "e", "f")
  echo("apd:", s)

  c := make([]string, len(s))
  copy(c, s)
  echo("cpy:", c)

  l := s[2:5]
  echo("sl1:", l)

  l = s[:5]
  echo("sl2:", l)

  l = s[2:]
  echo("sl3:", l)

  t := []string{"g", "h", "i"}
  echo("dcl:", t)

  twoD := make([][]int, 3)
  for i := 0; i < 3; i++ {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := 0; j < innerLen; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d: ", twoD)
}

func maps() {
  heading(10, "Maps")

  m := make(map[string]int)

  m["k1"] = 7
  m["k2"] = 13
  echo("map:", m)

  v1 := m["k1"]
  echo("v1:", v1)

  echo("len:", len(m))

  delete(m, "k2")
  echo("map:", m)

  _, prs := m["k2"]
  echo("prs:", prs)

  n := map[string]int{"foo": 1, "bar": 2}
  echo("map:", n)
}

func ranges() {
  heading(11, "Ranges")

  nums := []int{2, 3, 4}
  sum := 0
  for _, num := range nums {
    sum += num
  }
  echo("sum:", sum)

  for i, num := range nums {
    if num == 3 {
      echo("index:", i)
    }
  }

  kvs := map[string]string{"a": "apple", "b": "banana"}
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
  }

  for k := range kvs {
    echo("key:", k)
  }

  for _, v := range kvs {
    echo("val:", v)
  }

  for i, c := range "go" {
    echo(i, c)
  }
}

func plus(a int, b int) int {
  return a + b
}

func plusPlus(a, b, c int) int {
  return a + b + c
}

func functions() {
  heading(12, "Functions")

  res := plus(1, 2)
  echo("1+2 =", res)

  res = plusPlus(1, 2, 3)
  echo("1+2+3 =", res)
}
