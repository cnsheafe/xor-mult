package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)


func main()  {
  filePtr, err := os.Open("input.txt")
  if (err != nil) {
    fmt.Println(err)
    os.Exit(1)
  }
  scanner := bufio.NewScanner(filePtr)
  for scanner.Scan() {
    line := scanner.Text()
    input := strings.Split(line, " ")
    x, _ := strconv.ParseUint(input[0], 10, 64)
    y, _ := strconv.ParseUint(input[1], 10, 64)

    val := XorMultiply(uint(x), uint(y))
    fmt.Printf("%v@%v=%v\n", x, y, val)
  }

}
// XorMultiply takes each bit of y and multiplies x.
// The resultant of this operation is then shifted and XOR'd
// onto a sum.
func XorMultiply(x, y uint) uint {
  var sum, i uint
  for i = 0; y > 0; i++ {
    sum ^= (x * (y&1)) << i
    y >>= 1
  }
  return sum
}
