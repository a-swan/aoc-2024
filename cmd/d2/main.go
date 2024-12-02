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

func part2 (filePath string){
  file, err := os.Open(filePath)
  if err != nil{
    log.Fatal(err)
  }

  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)

  safeLevels := 0
  for fileScanner.Scan(){
    tmpLevels := strings.Fields(fileScanner.Text())
    levels := extractIntSlice(tmpLevels)

    safe, reason := checkSafety2(levels, 0)

    if safe {
      safeLevels = safeLevels + 1
    } //else {
      //fmt.Printf("%v: is %t %s\n", levels, safe, reason)
    //}
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

type Step struct {
  Dir string
  Size int
}

func checkSafety2(levels []int, currentDepth int) (bool, error){
  steps := make([]Step, 0)
  dampenerThreshold := 1
  ascCounter := make([]int, 0)
  descCounter := make([]int, 0)
  maxStep := 3
  minStep := 1
  outOfStepIndex := make([]int, 0)
  
  for i := 0; i < len(levels)-1; i++ {
    var direction string
    var stepSize int

    if levels[i] < levels[i+1] {
      direction = "asc"
      ascCounter = append(ascCounter, i)
    } else {
      direction = "desc"
      descCounter = append(descCounter, i)
    }

    stepSize = abs(levels[i] - levels[i+1])
    if stepSize > maxStep || stepSize < minStep {
      outOfStepIndex = append(outOfStepIndex, i)
    }

    steps = append(steps, Step{Dir: direction, Size: stepSize})
  }
  //fmt.Printf("Levels: %v\n", levels)
  //fmt.Printf("Steps: %v\n", steps)
  //fmt.Printf("ascCounter: %v | descCounter: %v\n", ascCounter, descCounter)

  if len(outOfStepIndex) > dampenerThreshold {
    return false, errors.New(fmt.Sprintf("Unsafe regardless of which level is removed"))
  }

  // there should only be 1 or 0 items in outOfStepIndex
  if len(ascCounter) == 1 {
    if len(outOfStepIndex) > 0 {
      if outOfStepIndex[0] != ascCounter[0] {
        outOfStepIndex = append(outOfStepIndex, ascCounter[0])
      }
    } else {
      outOfStepIndex = append(outOfStepIndex, ascCounter[0])
    }
  } else if len(descCounter) == 1 {
    if len(outOfStepIndex) > 0 {
      if outOfStepIndex[0] != descCounter[0] {
        outOfStepIndex = append(outOfStepIndex, descCounter[0])
      }
    } else {
      outOfStepIndex = append(outOfStepIndex, descCounter[0])
    }
  }

  //fmt.Printf("Out of Step Index: %v\n", outOfStepIndex)

  if (len(outOfStepIndex) > dampenerThreshold) || (len(outOfStepIndex) > 0 && currentDepth == dampenerThreshold) {
    return false, errors.New(fmt.Sprintf("Unsafe regardless of which level is removed"))
  }

  if len(outOfStepIndex) > 0 {
    if outOfStepIndex[0] == len(levels)-1 {
      //fmt.Printf("Last two indices\n")
      safe, _ := checkSafety2(append(levels[:outOfStepIndex[0]], levels[outOfStepIndex[0]+1:]...), 1)

      //fmt.Printf("Safe: %t - Reason: %v\n", safe, reason)

      if !safe {
        return false, errors.New(fmt.Sprintf("Unsafe regardless of which level is removed"))
      }
    } else if outOfStepIndex[0] > 0 {
      //fmt.Printf("Check %d to %d\n", levels[outOfStepIndex[0]-1], levels[outOfStepIndex[0]+1])
      safe, _ := checkSafety2(append(levels[:outOfStepIndex[0]], levels[outOfStepIndex[0]+1:]...), 1)

      //fmt.Printf("Safe: %t - Reason: %v\n", safe, reason)

      if !safe {
        return false, errors.New(fmt.Sprintf("Unsafe regardless of which level is removed"))
      }
    } 
  }

  return true, nil
}

func compareLevels(a, b int, direction bool) bool {
  maxStep := 3
  minStep := 1
  if direction != (b > a) {
    return false
  } else {
    step := abs(b - a)

    if step > maxStep || step < minStep {
      return false
    }
  }

  return true
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
