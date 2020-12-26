package main 
  
import ( 
    "bufio"
    "log"
    "os"
    //"strings"
) 

type state struct {
    rows map[int][]rune
    prevState *state
    maxX int
    maxY int
}
  
func main() { 
    file, err := os.Open("input.txt") 
  
    if err != nil { 
        log.Fatalf("failed to open") 
    } 
    scanner := bufio.NewScanner(file) 
    scanner.Split(bufio.ScanLines) 
    var lines []string

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    file.Close() 

    //log.Println(partOne(lines))
    log.Println(partTwo(lines))
}

func partOne(lines []string) int {
    var initialState state
    initialState.rows = make(map[int][]rune)
    for i, l := range lines {
        initialState.rows[i] = []rune(l)

    }
    initialState.maxY = len(lines)
    initialState.maxX = len(lines[0])

    var changes int
    var currentState *state
    currentState = &initialState
    for {
        for i := 0; i < currentState.maxY; i++ {
            log.Println(string(currentState.rows[i]))
        }

        log.Println("***************************************")
        currentState, changes = computeRound(currentState, 0)
        log.Printf("Changes made = %d", changes)
        if changes == 0 {
            break
        }
    }

    var finalCount int
    for _, row := range currentState.rows {
        for _, seat := range row {
            if seat == '#' {
                finalCount++
            }
        }
    }

    return finalCount
}

func computeRound(s *state, mode int) (*state, int) {
    // copy the state
    nextState := state {
        maxX: s.maxX,
        maxY: s.maxY,
        prevState: s,
        rows: make(map[int][]rune),
    }
    var changes int
    for y, row := range s.rows {
        newRow := make([]rune, nextState.maxX)
        for x, seat := range row { 
            switch seat {
            case 'L':
                if mode == 0 && countAdjacent(s, x, y, '#') == 0 {
                    newRow[x] = '#'
                    changes++
                } else if mode == 1 && countVisible(s, x, y, '#') == 0 {
                    newRow[x] = '#'
                    changes++
                } else {
                    newRow[x] = 'L'
                }

            case '#':
                if mode == 0 && countAdjacent(s, x, y, '#') >= 4 {
                    newRow[x] = 'L'
                    changes++
                } else if mode == 1 && countVisible(s, x, y, '#') >= 5 {
                    newRow[x] = 'L'
                    changes++
                } else {
                    newRow[x] = '#'
                }
            case '.':
                newRow[x] = '.'
                //nothing
            default:
                log.Printf("x: %d - y:%d - char:%s -- Could not parse seat",x,y,string(seat))
            }
        }
        nextState.rows[y] = newRow
    }
    return &nextState, changes
}

func countAdjacent(s *state, posX int, posY int, r rune) int {
    var x,y,count int
    y = posY - 1
    if y < 0 {y=0}
    for ; y <= posY+1; y++ {
        //log.Printf("y:%d\n",y)
        if y > s.maxY-1 {continue}
        x = posX - 1
        if x < 0 {x=0}
        for ; x <= posX+1; x++ {
            //log.Printf("x:%d\n",x)
            if x > s.maxX-1 || (x==posX&&y==posY){continue}
            
            //log.Print(string(s.rows[y][x]))
            if s.rows[y][x] == r {
                count++
            }
        }
    }
    //log.Printf("countadjacent for - x:%d y:%d - finding:%s ---- count = %d",posX, posY, string(r), count)
    return count
}

func countVisible(s *state, posX int, posY int, r rune) int {
    var x,y,layer,limits,count int
    dirs := make(map[int]bool)
    layer = 1
    for {
        for mulX := -1; mulX <= 1; mulX++ {
            x = posX + (layer * mulX)
            if x < 0 || x > s.maxX - 1 {
                limits++
                continue
            }
            for mulY := -1; mulY <= 1; mulY++ {
                y = posY + (layer * mulY)
                if y < 0 || y > s.maxY - 1 {
                    limits++
                    continue
                }
                if x == posX && y == posY {
                    continue
                }
                if dirs[(mulX*5)+mulY] {
                    //log.Printf("skiping x:%d y:%d, mulX:%d mulY:%d ",x,y,mulX,mulY)
                    continue
                }
                if s.rows[y][x] != '.' {
                    //log.Printf("found on x:%d y:%d, mulX:%d mulY:%d layer:%d ",x,y,mulX,mulY,layer)
                    dirs[(mulX*5)+mulY] = true
                    if s.rows[y][x] == r {
                        count++
                    }
                }
                //log.Printf("x:%d y:%d layer:%d",x,y,layer)
            }
        }
        if limits >= 4 {
            //log.Println("limits")
            break
        } else if len(dirs) >= 8 {
            //log.Println("limit dirs")
            break
        }
        limits = 0
        layer++
    }
    //log.Printf("countvisible for - x:%d y:%d - finding:%s ---- count = %d",posX, posY, string(r), count)
    return count
}


func partTwo(lines []string) int {
    var initialState state
    initialState.rows = make(map[int][]rune)
    for i, l := range lines {
        initialState.rows[i] = []rune(l)

    }
    initialState.maxY = len(lines)
    initialState.maxX = len(lines[0])

    var changes int
    var currentState *state
    currentState = &initialState
    for {
        for i := 0; i < currentState.maxY; i++ {
            log.Println(string(currentState.rows[i]))
        }
        log.Println("***************************************")
        currentState, changes = computeRound(currentState, 1)
        log.Printf("Changes made = %d", changes)
        if changes == 0 {
            break
        }
    }

    var finalCount int
    for _, row := range currentState.rows {
        for _, seat := range row {
            if seat == '#' {
                finalCount++
            }
        }
    }

    return finalCount
}


/*
func parseSeat(b string) (int,int,int) {
    currentRowMax := 127
    currentRowMin := 0
    offset := 0
    for i := 0; i < 7; i++ {
        offset = ((currentRowMax - currentRowMin)/2)+1
        if string(b[i]) == "B" {
            currentRowMin = currentRowMin + offset
        } else if string(b[i]) == "F" {
            currentRowMax = currentRowMax - offset
        }
        //log.Printf("%s: %d - %d\n",string(b[i]), currentRowMax, currentRowMin)
    }

    currentColMax := 7
    currentColMin := 0
    for i := 7; i < 10; i++ {
        offset = ((currentColMax - currentColMin)/2)+1
        if string(b[i]) == "R" {
            currentColMin = currentColMin + offset
        } else if string(b[i]) == "L" {
            currentColMax = currentColMax - offset
        }
        //log.Printf("%s: %d - %d\n",string(b[i]), currentColMax, currentColMin)
    }

    return currentRowMax,currentColMax,currentRowMax*8+currentColMax
}*/