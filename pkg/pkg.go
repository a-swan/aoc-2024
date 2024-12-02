package pkg

import (
  "bufio"
  "os"
  "strconv"
)

func readFile(filePath string) ([]string, error) {
  file, err := os.Open(filePath)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)

  lines := make([]string, 0)
  for fileScanner.Scan() {
    lines = append(lines, fileScanner.Text())
  }

  return lines, nil
}

// Absolute value
func abs(a int) int{
  if a < 0 {
    return a * -1
  }
  return a
}

// Return a new sub slice without changing original
func subSlice(a []int, index int) []int {
  tmp := make([]int, 0)
  tmp = append(tmp, a[:index]...)
  return append(tmp, a[index+1:]...)
}

// Convert slice of strings to slice of ints
func sliceStringToInt(a[] string) []int {
  tmp := make([]int, len(a))
  for i, s := range a{
    tmp[i], _ = strconv.Atoi(s)
  }

  return tmp
}
