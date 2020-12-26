package main 
  
import ( 
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "math"
) 

type instruction struct {
    op byte
    value int
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

    //fmt.Println(partOne(lines))
    fmt.Println(partTwo(lines))
}

func partTwo(lines []string) int {

    instruction_set := make(map[int]instruction)
    for i, line := range lines {
        op := line[0]
        val, _ := strconv.Atoi(line[1:])
        instruction_set[i] = instruction {
            op: op,
            value: val,
        }
        //fmt.Println(instruction_set[i])
    }
    in_num := 0
    var x,y int
    var wx,wy int = 10, 1
    for in_num < len(instruction_set) {
        ins := instruction_set[in_num]
        switch ins.op {
        case byte('N'), byte('S'), byte('W'), byte('E'):
            dirX,dirY := cardinalPointToDirection(ins.op)
            wx += dirX*ins.value
            wy += dirY*ins.value
        case byte('R'):
            wx, wy = rotate(wx,wy,ins.value*-1)
        case byte('L'):
            wx, wy = rotate(wx,wy,ins.value*1)
        case byte('F'):
            x += wx*ins.value
            y += wy*ins.value
        default:
            fmt.Println("not recognized")
        }
        fmt.Printf("%s%d - x:%d y:%d - wx:%d wy:%d\n",string(ins.op),ins.value,x,y,wx,wy)
        in_num++
    }
    return abs(x)+abs(y)
}

func rotate(dx int, dy int, angle int) (int, int) {
    angleR := float64(angle)*math.Pi/180
    cos := math.Cos(float64(angleR))
    sin := math.Sin(float64(angleR))
    newX := float64(dx) * cos - float64(dy) * sin
    newY := float64(dx) * sin + float64(dy) * cos
    return int(math.Round(newX)), int(math.Round(newY))
}

func partOne(lines []string) int {
    instruction_set := make(map[int]instruction)
    for i, line := range lines {
        op := line[0]
        val, _ := strconv.Atoi(line[1:])
        instruction_set[i] = instruction {
            op: op,
            value: val,
        }
        //fmt.Println(instruction_set[i])
    }
    cPoints := [4]byte{'N','E','S','W'}
    in_num := 0
    var x,y int
    orientation := 1
    for in_num < len(instruction_set) {
        ins := instruction_set[in_num]
        switch ins.op {
        case byte('N'), byte('S'), byte('W'), byte('E'):
            dirX,dirY := cardinalPointToDirection(ins.op)
            x += dirX*ins.value
            y += dirY*ins.value
        case byte('R'):
            change := ins.value/90
            orientation += change
        case byte('L'):
            change := ins.value/90
            orientation += 4-change
        case byte('F'):
            dirX,dirY := cardinalPointToDirection(cPoints[orientation%4])
            x += dirX*ins.value
            y += dirY*ins.value
        default:
            fmt.Println("not recognized")
        }
        fmt.Printf("x:%d y:%d\n",x,y)
        in_num++
    }
    return abs(x)+abs(y)
}

func cardinalPointToDirection(c byte) (int, int) {
    switch c {
    case byte('N'):
        return 0,1
    case byte('S'):
        return 0,-1
    case byte('E'):
        return 1,0
    case byte('W'):
        return -1,0
    default:
        fmt.Println("unrecognized cardinal point")
    }
    return 0,0
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}


