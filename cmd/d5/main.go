package main

import (
  "os"
  "log"
  "fmt"
  "strings"
  "slices"
  "strconv"

  "github.com/a-swan/aoc-2024/pkg"
)

type Page struct {
  parent []int
  child []int
}

func part1(lines []string) {
  fmt.Printf("Starting part 1...\n")
  rules := make(map[int]Page)

  // build some sort of search tree of rules
  var i int
  for i = 0; i < len(lines); i++ {
    if lines[i] == "" {
      break
    }

    var a, b int
    _, err := fmt.Sscanf(lines[i], "%d|%d", &a, &b)
    if err != nil {
      log.Fatal(err)
    }

    tmpA := rules[a]
    tmpA.child = append(tmpA.child, b)

    tmpB := rules[b]
    tmpB.parent = append(tmpB.parent, a)

    rules[a] = tmpA
    rules[b] = tmpB
  }
  // when the empty row hits
  fmt.Printf("i = %d\n", i)

  fmt.Printf("%v\n", rules)
  
  pageSumTotal := 0
  for i = i+1; i < len(lines); i++ {
    pageOrder := strings.Split(lines[i], ",")

    left := make([]int, 0)
    fails := false
    for _, tmpPage := range pageOrder {
      page, _ := strconv.Atoi(tmpPage)
      for _, previousPage := range left {
        if slices.Contains(rules[page].child, previousPage) {
          fmt.Printf("%s - fails!\n", pageOrder)
          fails = true
          break
        }
      }
      if fails {
        break
      }
      left = append(left, page)
    }
    if !fails {
      fmt.Printf("%s - passes! \n", pageOrder)
      fmt.Printf("Left: %v\n", left)
      
      middlePage, _ := strconv.Atoi(pageOrder[len(pageOrder)/2])

      pageSumTotal += middlePage  
    }
  }

  fmt.Printf("Total: %d\n", pageSumTotal)
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
