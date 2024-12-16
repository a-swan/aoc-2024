package main

import (
	"fmt"
	"log"
	"os" 
  "os/signal"
  "syscall"
  "strings"

  "golang.org/x/term"
	"github.com/a-swan/aoc-2024/pkg"
)

func part1(lines []string) {
  fmt.Printf("Running part 1...\n")

  highlights := make([][]int, len(lines), len(lines[0]))
  process := make([]string, 0)
  updateScreen(lines, highlights, process)

  xmasStrings := 0
  for i, line := range lines {
    for j, r := range line {
      // highlight current rune
      highlights[i] = append(highlights[i], j)

      if r == 'X' {
        process = append(process, fmt.Sprintf("Analyzing 'X' at (%d,%d)", i, j))  
        updateScreen(lines, highlights, process)

        //match := false
        targetCols := make([]int, 0)
        if ( j - 1 ) > 0 {
          targetCols = append(targetCols, j-1)
        }
        targetCols = append(targetCols, j)
        if ( j + 1 ) < len(line) {
          targetCols = append(targetCols, j+1)
        }

        targetRows := make([]int, 0)
        if ( i - 1 ) > 0 {
          targetRows = append(targetRows, i-1)
        }
        targetRows = append(targetRows, i)
        if ( i + 1 ) < len(lines) {
          targetRows = append(targetRows, i+1)
        }
        process = append(process, fmt.Sprintf("Collected target Rows: [%v], Cols: [%v]", targetRows, targetCols))

        for _, nextRow := range targetRows {
          for _, nextCol := range targetCols {
            updateScreen(lines, highlights, process)

            if checkNeighbor('M', nextRow, nextCol, lines, &highlights, &process) {
              // get direction
              rowDir := nextRow - i
              colDir := nextCol - j

              process = append(process, fmt.Sprintf("Hit at (%d,%d) - Word Row Direction: %d, Col Direction: %d", nextRow, nextCol, rowDir, colDir))

              if (nextRow + rowDir + rowDir) >= 0 && (nextRow + rowDir + rowDir) < len(lines) && (nextCol + colDir + colDir) >= 0 && (nextCol + colDir + colDir) < len(line){
                if checkNeighbor('A', (nextRow + rowDir), (nextCol + colDir), lines, &highlights, &process) && checkNeighbor('S', (nextRow + rowDir + rowDir), (nextCol + colDir + colDir), lines, &highlights, &process) {
                  process = append(process, fmt.Sprintf("[FOUND]: (%d,%d) (%d,%d) (%d,%d) (%d,%d)", i,j, nextRow, nextCol, nextRow + rowDir, nextCol + colDir, nextRow + rowDir + rowDir, nextCol + colDir + colDir))
                  xmasStrings += 1
                }
              }
            }
          }
        }
      } else {
        updateScreen(lines, highlights, process)
        highlights[i] = highlights[i][:len(highlights[i])-1]
      }
    }
  }

  fmt.Printf("total matches: %d\n", xmasStrings)
}

func part2(lines []string) {
  fmt.Printf("Running part 2...\n")

  highlights := make([][]int, len(lines), len(lines[0]))
  process := make([]string, 0)
  updateScreen(lines, highlights, process)

  xmasStrings := 0
  for i, line := range lines {
    for j, r := range line {
      // highlight current rune
      highlights[i] = append(highlights[i], j)

      if r == 'A' {
        process = append(process, fmt.Sprintf("Analyzing 'A' at (%d,%d)", i, j))
        updateScreen(lines, highlights, process)

        //match := false
        targetCols := make([]int, 0)
        if ( j - 1 ) > 0 {
          targetCols = append(targetCols, j-1)
        }
        if ( j + 1 ) < len(line) {
          targetCols = append(targetCols, j+1)
        }

        targetRows := make([]int, 0)
        if ( i - 1 ) > 0 {
          targetRows = append(targetRows, i-1)
        }
        if ( i + 1 ) < len(lines) {
          targetRows = append(targetRows, i+1)
        }
        process = append(process, fmt.Sprintf("Collected target Rows: [%v], Cols: [%v]", targetRows, targetCols))

        if (i-1) >= 0 && (j-1) >= 0 && (i+1) < len(lines) && (j+1) < len(line){
          if (checkNeighbor('M', i-1, j-1, lines, &highlights, &process) && checkNeighbor('S', i+1, j+1, lines, &highlights, &process)) ||
            (checkNeighbor('S', i-1, j-1, lines, &highlights, &process) && checkNeighbor('M', i+1, j+1, lines, &highlights, &process)) {
              process = append(process, fmt.Sprintf("Hit at NW - SE [(%d,%d) - (%d,%d)]", i-1, j-1, i+1, j+1))
              if (checkNeighbor('M', i-1, j+1, lines, &highlights, &process) && checkNeighbor('S', i+1, j-1, lines, &highlights, &process)) ||
                (checkNeighbor('S', i-1, j+1, lines, &highlights, &process) && checkNeighbor('M', i+1, j-1, lines, &highlights, &process)) {
                  process = append(process, fmt.Sprintf("[FOUND]: (%d,%d) (%d,%d) (%d,%d) (%d,%d)", i, j, i+1, j+1, i-1, j-1, i-1, j+1, i+1, j-2))
                  xmasStrings += 1
              }
            
          }
        }
      } else {
        updateScreen(lines, highlights, process)
        highlights[i] = highlights[i][:len(highlights[i])-1]
      }
    }
  }

  fmt.Printf("total matches: %d\n", xmasStrings)
}

func checkNeighbor(targetRune rune, targetRow int, targetCol int, lines []string, highlights *[][]int, process *[]string) bool {
  target := []rune(lines[targetRow])[targetCol]

  (*highlights)[targetRow] = append((*highlights)[targetRow], targetCol)
  updateScreen(lines, *highlights, *process)
  if target == targetRune {
    return true
  }
  (*highlights)[targetRow] = (*highlights)[targetRow][:len((*highlights)[targetRow])-1]

  return false
}

func updateScreen(lines []string, highlights [][]int, process []string){
  pause := false
  if len(os.Args) > 2 {
    pause = true
  }
  
  if pause {
    fmt.Print("\033[H\033[2J")
    for i, line := range lines {
      highlightedLogs := pkg.LogHighlightIndex(line, highlights[i])

      fmt.Printf("[%d] %s\n", i, highlightedLogs)
    }

    fmt.Println(strings.Repeat("=", len(lines[0])))
    fmt.Println("Steps:")
    // print out the steps taken
    startJ := 0
    maxLines := 10
    if len(process) > maxLines {
      startJ = len(process) - maxLines
    }

    for j:=startJ; j < len(process); j++ {
      fmt.Printf("%d) %s\n", j, process[j])
    }

    signalChan := make(chan os.Signal)
    signal.Notify(signalChan, os.Interrupt)
    go func() {
      <-signalChan
      os.Exit(1)
    }()

    _, err := term.ReadPassword(syscall.Stdin)
    if err != nil {
      log.Fatal(err)
    }

    signal.Stop(signalChan)
  }
}

func main(){
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
