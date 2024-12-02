package main

import (
  "log"
  "os"
  "bufio"
  "fmt"
  "sort"
)

func part1(filePath string) {
  // Parse lists
  lists := make([][]int, 2)

  file, err := os.Open(filePath)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  
  fmt.Printf("opening %s...\n", filePath) 
  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)

  for fileScanner.Scan() {
    var a,b int
    fmt.Sscanf(fileScanner.Text(), "%d   %d", &a, &b)
    lists[0] = append(lists[0], a)
    lists[1] = append(lists[1], b)
  }

  sort.Sort(sort.Reverse(sort.IntSlice(lists[0])))
  sort.Sort(sort.Reverse(sort.IntSlice(lists[1])))

  totalDistance := 0
  for iter := 0; iter < len(lists[0]); iter++ {
    totalDistance = totalDistance + abs(lists[0][iter] - lists[1][iter])
  }
  
  fmt.Printf("Total Distance between lists is: %d\n", totalDistance)
}

func part2(filePath string){
  file, err := os.Open(filePath)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  
  fmt.Printf("opening %s...\n", filePath) 
  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)

  left := make([]int, 1)
  right := make(map[int]int, 1)
  for fileScanner.Scan() {
    var a,b int
    fmt.Sscanf(fileScanner.Text(), "%d   %d", &a, &b)
    left = append(left, a)
    right[b] = right[b]+1
  }

  //fmt.Printf("Left: %v\n", left)
  //fmt.Printf("Right: %v\n", right)

  totalDistance := 0
  for i:=0; i< len(left); i++{
    similarity := left[i] * right[left[i]]
    totalDistance = totalDistance + similarity
  }

  fmt.Printf("Total Distance Similarity: %d\n", totalDistance)
}

func abs(a int) int{
  if a < 0 {
    return a * -1
  }
  return a
}

func main() {
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
