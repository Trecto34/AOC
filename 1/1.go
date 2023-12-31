package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main(){
  file, error := os.Open("input.test")
  if error != nil {
    panic(error)
  }

  fileScan := bufio.NewScanner(file)
  fileScan.Split(bufio.ScanLines)
  var fileLines []string
  var fileInts []int64

  for fileScan.Scan() {
    fileLines = append(fileLines, fileScan.Text())
  }

  file.Close()

  for i, _ := range fileLines {
    toStr, _ := strconv.ParseInt(fileLines[i], 10, 32)
    fileInts = append(fileInts, toStr)
  }

  fmt.Println(fileInts)

  for _, num := range fileInts {
    for _, num2 := range fileInts {
      if num+num2 == 2020 {
        fmt.Println("Found: ", num, num2, "= 2020")
        fmt.Println("Answer: ", num*num2)
      }
    }
  }

}

