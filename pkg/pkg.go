package pkg

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ReadFile(filePath string) ([]string, error) {
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
func Abs(a int) int{
  if a < 0 {
    return a * -1
  }
  return a
}

// Return a new sub slice without changing original
func SubSlice(a []int, index int) []int {
  tmp := make([]int, 0)
  tmp = append(tmp, a[:index]...)
  return append(tmp, a[index+1:]...)
}

// Convert slice of strings to slice of ints
func SliceAToInt(a[] string) []int {
  tmp := make([]int, len(a))
  for i, s := range a{
    tmp[i], _ = strconv.Atoi(s)
  }

  return tmp
}

var RESET = "\033[0m"
var RED = "\033[31m"
var GREEN = "\033[32m"

func LogHighlightSubstring(s string, sub string) string{
  replacement := RED + sub + RESET

  returnString := fmt.Sprintln(strings.ReplaceAll(s, sub, replacement))
  return returnString
}

func LogHighlightIndex(s string, i []int) string{
  var returnString strings.Builder

  for ind, r := range s {
    if slices.Contains(i, ind){
      returnString.WriteString(RED)
      returnString.WriteRune(r)
      returnString.WriteString(RESET)
    } else {
      returnString.WriteRune(r)
    }
  }
  return returnString.String()
}
