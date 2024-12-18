package main

import (
  "log"
  "os"
  "fmt"
  "errors"
  "github.com/a-swan/aoc-2024/pkg"
)

type Coordinate struct {
  row int
  col int
}

func part1 (lines []string) {
  fmt.Printf("Starting part 1...\n")

  // initialize map visualization
  updatedMap := make([][]string, len(lines))

  obstructionsRows := make(map[int][]int, len(lines))
  obstructionsCols := make(map[int][]int, len(lines[0]))
  guard := Coordinate{}

  for i := 0; i < len(lines); i++ {
    updatedMap[i] = make([]string, len(lines[0]))

    for j := 0; j < len(lines[i]); j++ {
      if lines[i][j] == '#' {
        updatedMap[i][j] = "#"
        obstructionsRows[i] = append(obstructionsRows[i], j)
        obstructionsCols[j] = append(obstructionsCols[j], i)
      } else if lines[i][j] == '^' {
        updatedMap[i][j] = "^"
        guard = Coordinate{row: i, col: j}
      } else {
        updatedMap[i][j] = "_"
      }
    }
  }

  pkg.PrintGrid(updatedMap)
  fmt.Printf("Dimensions: %d x %d\n", len(updatedMap), len(updatedMap[0]))
  fmt.Printf("Obstacles Rows: %v\n", obstructionsRows)
  fmt.Printf("Obstacles Cols: %v\n", obstructionsCols)

  // North {-1, 0}
  // East  {0, 1}
  // South {1, 0}
  // West  {0, -1}
  direction := "NORTH"
  directionValues := Coordinate { row: -1, col: 0 }
  fmt.Printf("%s : %v\n", direction, directionValues)
  fmt.Printf("Guard: %v\n", guard)
  hits := 0

  oOBFlag := false
  for !oOBFlag {
    var obstructionHit Coordinate
    if direction == "NORTH" || direction == "SOUTH" {
      //use obstructionsCols
      collision, err := checkForObstruction(guard.row, obstructionsCols[guard.col], directionValues.row)

      if err != nil {
        if directionValues.row > 0 {
          collision = len(lines)
        }
        oOBFlag = true
        //log.Fatal(fmt.Sprintf("%s: OoB (%d, %d)\n", err, collision, guard.col))
      }

      hits++
      obstructionHit = Coordinate{ row: collision, col: guard.col}
    } else {
      //use obstructionsRows
      collision, err := checkForObstruction(guard.col, obstructionsRows[guard.row], directionValues.col)

      if err != nil {
        if directionValues.col > 0 {
          collision = len(lines[0])
        }
        oOBFlag = true
        //log.Fatal(fmt.Sprintf("%s: OoB (%d, %d)\n", err, guard.row, collision))
      }

      hits++
      obstructionHit = Coordinate{ row: guard.row, col: collision}
    }
    
    fmt.Printf("collision at %v\n", obstructionHit)
    // update map with X's
    updateMap(&updatedMap, guard, obstructionHit, directionValues)
    // update guard location
    guard = obstructionHit
    // change direction
    changeDirection(&direction, &directionValues)
    pkg.PrintGrid(updatedMap)
  }

  steps := 0
  for r := 0; r < len(updatedMap); r++ {
    for c := 0; c < len(updatedMap[r]); c++ {
      if updatedMap[r][c] != "#" && updatedMap[r][c] != "_" {
        steps++
      }
    }
  }

  fmt.Printf("Number of steps: %d\n", steps)
}

func checkForObstruction(guardLocation int, lineOfSight []int, directionValue int) (int, error){
  fmt.Printf("guardLocation: %d, lineOfSight: %v, directionValue: %d\n", guardLocation, lineOfSight, directionValue)
  boundary := 0
  if directionValue < 0 {
    fmt.Printf("direction is negative\n")
    boundary = 0
    for i := len(lineOfSight)-1; i >= 0; i-- {
      if lineOfSight[i] < guardLocation {
        return lineOfSight[i]+1, nil
      }
    }
  } else {
    fmt.Printf("direction is positive\n")
    for i:=0; i < len(lineOfSight); i++ {
      if lineOfSight[i] > guardLocation {
        return lineOfSight[i]-1, nil
      }
    }
  }

  return boundary, errors.New("No obstruction, guard leaves the area")
}

func changeDirection (direction *string, directionValues *Coordinate) {
  switch *direction {
  case "NORTH":
    *direction = "EAST"
    *directionValues = Coordinate{row: 0, col: 1}
  case "EAST":
    *direction = "SOUTH"
    *directionValues = Coordinate{row: 1, col: 0}
  case "SOUTH":
    *direction = "WEST"
    *directionValues = Coordinate{row: 0, col: -1}
  case "WEST":
    *direction = "NORTH"
    *directionValues = Coordinate{row: -1, col: 0}
  default:
    log.Fatal("unknown direction: %s", direction)
  }

  fmt.Printf("New direction: %s\n", *direction)
}

func updateMap(updatedMap *[][]string, startCoord, endCoord, directionValues Coordinate) {
  fmt.Printf("Updating Map from: (%v) to (%v)\n", startCoord, endCoord)
  if startCoord.row == endCoord.row{
    for i := startCoord.col; i != endCoord.col; i = i + directionValues.col {
      (*updatedMap)[startCoord.row][i] = "X"
    }
  } else {
    for i := startCoord.row; i != endCoord.row; i = i + directionValues.row {
      (*updatedMap)[i][startCoord.col] = "X"
    }
  }

  if endCoord.row >= len(*updatedMap) {    
    (*updatedMap)[endCoord.row-1][endCoord.col] = "^"
  } else if endCoord.col >= len((*updatedMap)[0]) {
    (*updatedMap)[endCoord.row][endCoord.col-1] = "^"
  } else {
    (*updatedMap)[endCoord.row][endCoord.col] = "^"
  }
    
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
