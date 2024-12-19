package main

import (
	"fmt"
	"log"
	"os"
	"strings"
  "strconv"

	"github.com/a-swan/aoc-2024/pkg"
)

func part1(lines []string) {
  fmt.Printf("Starting Part 1...\n")

  validTotal := 0
  for _, line := range lines {
    numsStr := strings.Split(line, ":")

    total, _ := strconv.Atoi(numsStr[0])
    nums := make([]int, 0)
    for _, strNum := range strings.Split(strings.TrimSpace(numsStr[1]), " ") {
      intNum, _ := strconv.Atoi(strNum)
      nums = append(nums, intNum)
    }

    valid, result := checkAddition(total, nums)
    if valid {
      fmt.Printf("%d = %s\n\n", total, result)
      validTotal += total
      continue
    }
    valid, result = checkMultiplication(total, nums)
    if valid {
      fmt.Printf("%d = %s\n\n", total, result)
      validTotal += total
    } else {
      fmt.Printf("%d is not workable\n\n", total)
    }
  }

  fmt.Printf("Total total: %d\n", validTotal)
}

func checkAddition(total int, nums []int) (bool, string) {
  if len(nums) == 1 {
    if total == nums[0] {
      fmt.Printf("True! %d == %d\n", total, nums[0])
      return true, fmt.Sprintf("%d", nums[0])
    }
    return false, ""
  }

  targetNumber := nums[len(nums)-1]
  newTotal := total - targetNumber
  valid, result := checkAddition(newTotal, nums[:len(nums)-1])
  if valid {
    fmt.Printf("Building string: %s\n", result)
    return true, fmt.Sprintf("%s + %d", result, targetNumber) 
  }
  valid, result = checkMultiplication(newTotal, nums[:len(nums)-1])
  if valid {
    fmt.Printf("Building string: %s\n", result)
    return true, fmt.Sprintf("%s + %d", result, targetNumber)
  }

  return false, ""
}

func checkMultiplication(total int, nums []int) (bool, string) {
  if len(nums) == 1 {
    if total == nums[0] {
      fmt.Printf("True! %d == %d\n", total, nums[0])
      return true, fmt.Sprintf("%d", nums[0])
    }
    return false, ""
  }

  targetNumber := nums[len(nums)-1]

  // check if float
  if total % targetNumber != 0{
    return false, ""
  }
  newTotal := total / targetNumber

  valid, result := checkAddition(newTotal, nums[:len(nums)-1])
  if valid {
    fmt.Printf("Building string: %s\n", result)
    return true, fmt.Sprintf("%s * %d", result, targetNumber) 
  }
  valid, result = checkMultiplication(newTotal, nums[:len(nums)-1])
  if valid {
    fmt.Printf("Building string: %s\n", result)
    return true, fmt.Sprintf("%s * %d", result, targetNumber)
  }

  return false, ""
}

func main() {
  var filePath string

  if len(os.Args) > 1 {
    filePath = os.Args[1]
  } else {
    filePath = "input.txt"
  }

  lines, err := pkg.ReadFile(filePath)
  if err != nil {
    log.Fatal(err)
  }

  part1(lines)
}
