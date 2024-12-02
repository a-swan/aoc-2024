package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
  "errors"
)

func part1(filePath string){
  fmt.Printf("Opening %s\n", filePath)

  file, err := os.Open(filePath)
  if err != nil{
    log.Fatal(err)
  }
  defer file.Close()

  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)

  safeLevels := 0

  for fileScanner.Scan(){
    tmpLevels := strings.Fields(fileScanner.Text())
    levels := extractIntSlice(tmpLevels)

    // Check safety of level
    safe, reason := checkSafety(levels)
    
    if safe {
      safeLevels = safeLevels + 1
    }
    fmt.Printf("%v: is %t %s\n", levels, safe, reason)
  }

  fmt.Printf("Total of %d safe levels\n", safeLevels)
}

func checkSafety(levels []int) (bool, error){
  increasing := true
  maxStep := 3
  minStep := 1

  for i := 1; i < len(levels); i++ {
    // Set initial direction
    if i == 1 {
      increasing = levels[i] > levels[i-1]
    }

    // Check for same direction and step safety
    if increasing != (levels[i] > levels[i-1]) {
      return false, errors.New(fmt.Sprintf("levels %d %d are changing the direction from %t", levels[i-1], levels[i], increasing))
    } else {
      step := abs(levels[i] - levels[i-1])

      if step > maxStep || step < minStep {
        return false, errors.New(fmt.Sprintf("step %d %d is an increase of %d, which is outside the range [%d-%d]", levels[i-1], levels[i], step, minStep, maxStep))
      }
    }
  }

  return true, nil
}

func part2 (filePath string){
  fmt.Printf("Opening %s\n", filePath)

  file, err := os.Open(filePath)
  if err != nil{
    log.Fatal(err)
  }
  defer file.Close()

  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)

  safeLevels := 0

  for fileScanner.Scan(){
    tmpLevels := strings.Fields(fileScanner.Text())
    levels := extractIntSlice(tmpLevels)

    // Check safety of level
    safe, _:= checkSafety(levels)
    dampener, _:= problemDampener(levels)
    
    if safe || dampener {
      safeLevels = safeLevels + 1
    }
  } 
  fmt.Printf("Total safe: %d\n", safeLevels)
}

func problemDampener(levels []int) (bool, error) {
  for i := 0; i < len(levels); i++ {
    subLevels := subSlice(levels, i)

    safe, _ := checkSafety(subLevels)

    if safe {
      return true, nil
    }
  }

  return false, errors.New(fmt.Sprintf("Dampener doesn't solve this"))
}

func subSlice(s []int, index int) []int {
  newS := make([]int, 0)
  newS = append(newS, s[:index]...)
  return append(newS, s[index+1:]...)
}

func abs(a int) int{
  if a < 0 {
    return a * -1
  }
  return a
}

func extractIntSlice(tmpLevels[] string) []int {
  levels := make([]int, len(tmpLevels))
  for i, s := range tmpLevels{
    levels[i], _ = strconv.Atoi(s)
  }

  return levels
}

func main(){
  var filePath string
  
  if(len(os.Args) > 1) {
    filePath = os.Args[1]
  } else {
    filePath = "input.txt"
  }

  fmt.Printf("Running part 1\n")
  part1(filePath)
  fmt.Printf("Running part 2\n")
  part2(filePath)
}
