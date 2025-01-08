package main

import (
  "fmt"
  "log"
  "os"
  "github.com/a-swan/aoc-2024/pkg"
)

type Coord struct {
  x, y int
}

func part1(lines []string) {
  fmt.Printf("Starting part 1...\n")

  antMap := make(map[string][]Coord, 0)

  maxX := len(lines[0])
  maxY := len(lines)

  for i := 0; i < maxY; i++ {
    for j := 0; j < maxX; j++ {
      if lines[i][j] != '.' {
        //fmt.Printf("(%d,%d): %s\n", i, j, string(lines[i][j]))
        antMap[string(lines[i][j])] = append(antMap[string(lines[i][j])], Coord{x: i, y: j})
      }
    }
  }

  antiNodes := make(map[Coord]bool)
  for _, freq := range antMap {
    for i := 0; i < len(freq); i++ {
      for j := i+1; j < len(freq); j++ {
        //fmt.Printf("[%s]: (%v) - (%v)\n", k, freq[i], freq[j])

        dx, dy := freq[i].x - freq[j].x, freq[i].y - freq[j].y

        if inBounds(freq[i].x + dx, freq[i].y + dy, maxX, maxY) {
          antiCoord := Coord{x: freq[i].x + dx, y: freq[i].y + dy}
          antiNodes[antiCoord] = true
          //antiNodes = append(antiNodes, Coord{x: freq[i].x + dx, y: freq[i].y + dy})
        }

        if inBounds(freq[j].x - dx, freq[j].y - dy, maxX, maxY) {
          antiCoord := Coord{x: freq[j].x - dx, y: freq[j].y - dy}
          antiNodes[antiCoord] = true
          //antiNodes = append(antiNodes, Coord{x: freq[j].x - dx, y: freq[j].y - dy})
        }
      }
    }
  }

  fmt.Printf("[%d] Antinodes\n", len(antiNodes))
}

func part2(lines []string) {
  fmt.Printf("Starting part 2...\n")

  antMap := make(map[string][]Coord, 0)

  maxX := len(lines[0])
  maxY := len(lines)


  for i := 0; i < maxY; i++ {
    for j := 0; j < maxX; j++ {
      if lines[i][j] != '.' {
        //fmt.Printf("(%d,%d): %s\n", i, j, string(lines[i][j]))
        antMap[string(lines[i][j])] = append(antMap[string(lines[i][j])], Coord{x: i, y: j})
      }
    }
  }


  antiNodes := make(map[Coord]bool)
  for _, freq := range antMap {
    for i := 0; i < len(freq); i++ {
      for j := i+1; j < len(freq); j++ {
        //fmt.Printf("[%s]: (%v) - (%v)\n", k, freq[i], freq[j])

        dx, dy := freq[i].x - freq[j].x, freq[i].y - freq[j].y
        
        topCoord := freq[i]
        for {
          if !inBounds(topCoord.x, topCoord.y, maxX, maxY) {
            break
          }

          antiCoord := Coord{x: topCoord.x, y: topCoord.y}
          antiNodes[antiCoord] = true
          
          topCoord.x += dx
          topCoord.y += dy
        }

        botCoord := Coord{x: freq[j].x, y: freq[j].y}
        for {
          if !inBounds(botCoord.x, botCoord.y, maxX, maxY) {
            break
          }

          antiCoord := Coord{x: botCoord.x, y: botCoord.y}
          antiNodes[antiCoord] = true
          
          botCoord.x += dx
          botCoord.y += dy
        }
      }
    }
  }
  fmt.Printf("[%d] Antinodes\n", len(antiNodes))

  res := [12]string{"##....#....#", ".#.#....0...", "..#.#0....#.", "..##...0....", "....0....#..", ".#...#A....#", "...#..#.....", "#....#.#....", "..#.....A...", "....#....A..", ".#........#.", "...#......##"}
    
  for m := 0; m < len(lines); m++ {
    var calcLine string
    for n := 0; n < len(lines[0]); n++ {
      char := "."
      currCoord := Coord{x: m, y: n}
      _, ok := antiNodes[currCoord]
      if ok {
        char = "#"
      } else if lines[m][n] != '.' {
        char = string(lines[m][n]) 
      }
      calcLine = fmt.Sprintf("%s%s", calcLine, char)
    }
    fmt.Printf("%s\t%s\n", calcLine, res[m])
  }
}

func inBounds(x, y, maxX, maxY int) bool {
  return x < maxX && x >= 0 && y < maxY && y >= 0
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
  part2(lines)
}
