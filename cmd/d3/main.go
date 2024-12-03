package main

import (
  "fmt"
  "os"
  "github.com/a-swan/aoc-2024/pkg"
  "log"
  "regexp"
  "strconv"
  "strings"
)

func part1 (fileLines []string) {
  r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
  total := 0

  for _, line := range fileLines{
    mults := r.FindAllStringSubmatch(line, -1)

    for _, mult := range mults {
      a, _ := strconv.Atoi(mult[1])
      b, _ := strconv.Atoi(mult[2])
      total = total + (a * b)
    }
  }

  fmt.Printf("Total uncorrupted mults are: %d\n", total)
}

func part2(fileLines []string) {
  fmt.Printf("Starting part 2...\n")
  r := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
  total := 0

  doFlag := true
  for _, line := range fileLines{
    mults := r.FindAllString(line, -1)

    fmt.Printf("mults: %v\n", mults)

    for _, mult := range mults {
      action := strings.Split(mult, "(")

      switch action[0] {
        case "do":
          fmt.Printf("Changing doFlag from %t to %t\n", doFlag, true)
          doFlag = true
        case "don't":
          fmt.Printf("Changing doFlag from %t to %t\n", doFlag, false)
          doFlag = false
        default:
          if doFlag {
            fmt.Printf("Multiplying %s\n", mult)
            var a, b int
            fmt.Sscanf(mult, "mul(%d,%d)", &a, &b)

            total = total + (a * b)
          }
      }
    }
  }

  fmt.Printf("Total uncorrupted mults: %d\n", total)
}

func main() {
  var filePath string
  if len(os.Args) > 1 {
    filePath = os.Args[1]
  } else {
    filePath = "input.txt"
  }

  fmt.Printf("Opening %s...\n", filePath)
  fileLines, err := pkg.ReadFile(filePath)

  if err == nil {
    fmt.Printf("Starting part 1...\n")
    part1(fileLines)
    part2(fileLines)
  } else {
    log.Fatal(err)
  }
}
