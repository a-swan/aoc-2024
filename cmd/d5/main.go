package main

import (
  "os"
  "log"
  "fmt"
  "strings"
  "slices"
  "sort"
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

    parseRules(&rules, lines[i])
  }
  // when the empty row hits
  fmt.Printf("i = %d\n", i)

  fmt.Printf("%v\n", rules)
  
  pageSumTotal := 0
  for i = i+1; i < len(lines); i++ {
    valid, middlePageStr := passesRules(lines[i], &rules)
    middlePage, _ := strconv.Atoi(middlePageStr)
    
    if valid && middlePage != 0 {
      fmt.Printf("%s - passes! \n", lines[i])

      pageSumTotal += middlePage  
    }
  }

  fmt.Printf("Total: %d\n", pageSumTotal)
}

func part2(lines []string) {
  fmt.Printf("Starting part 2...\n")
  rules := make(map[int]Page)

  // build some sort of search tree of rules
  var i int
  for i = 0; i < len(lines); i++ {
    if lines[i] == "" {
      break
    }

    parseRules(&rules, lines[i])
  }
  // when the empty row hits
  fmt.Printf("i = %d\n", i)

  fmt.Printf("%v\n", rules)
  
  pageSumTotal := 0
  for i = i+1; i < len(lines); i++ {
    valid, middlePageStr := passesRules(lines[i], &rules)
    
    if valid {
      fmt.Printf("%s - passes on the first try! \n", lines[i])
    } else {
      // sort pages and run it again
      sortedLine := sortPages(lines[i], &rules)
      valid, middlePageStr = passesRules(sortedLine, &rules)
      middlePage, _ := strconv.Atoi(middlePageStr)

      if valid && middlePage != 0 {
        fmt.Printf("%s - passes on the second try! \n", lines[i])

        pageSumTotal += middlePage  
      }
    }
  }

  fmt.Printf("Total: %d\n", pageSumTotal)
}


func parseRules(rules *map[int]Page, line string) {
    var a, b int
    _, err := fmt.Sscanf(line, "%d|%d", &a, &b)
    if err != nil {
      log.Fatal(err)
    }

    tmpA := (*rules)[a]
    tmpA.child = append(tmpA.child, b)

    tmpB := (*rules)[b]
    tmpB.parent = append(tmpB.parent, a)

    (*rules)[a] = tmpA
    (*rules)[b] = tmpB
}

func passesRules(line string, rules *map[int]Page) (bool, string){
  pageOrder := strings.Split(line, ",")

  left := make([]int, 0)
  for _, tmpPage := range pageOrder {
    page, _ := strconv.Atoi(tmpPage)
    for _, previousPage := range left {
      if slices.Contains((*rules)[page].child, previousPage) {
        fmt.Printf("%s - fails!\n", pageOrder)
        return false, ""
      }
    }
    left = append(left, page)
  }
  
  return true, pageOrder[len(pageOrder)/2]
}

func sortPages(line string, rules *map[int]Page) string { 
  pageOrder := strings.Split(line, ",")
  fmt.Printf("Original Line: %v\n", pageOrder)

  sort.Slice(pageOrder, func(i, j int) bool {
    a, _ := strconv.Atoi(pageOrder[i])
    b, _ := strconv.Atoi(pageOrder[j])
    return slices.Contains((*rules)[a].child, b)
  })

  fmt.Printf("Sorted Line: %v\n", pageOrder);
  return strings.Join(pageOrder[:], ",")
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

  //part1(lines)
  part2(lines)
}
