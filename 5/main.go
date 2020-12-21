package main 
  
import ( 
    "bufio"
    "fmt"
    "log"
    "os"
) 

  
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
    partTwo(lines)
}

func partOne(lines []string) int {
    var highest int
    for _, l := range lines {
        _, _, newid := parseSeat(l)
        if newid > highest {
            highest = newid
        }
    }
    return highest
}

func partTwo(lines []string) int {
    var rows []byte
    rows = make([]byte, 128, 128)
    for _, l := range lines {
        row, col, _ := parseSeat(l)
        //fmt.Printf("%b\n", byte(1<<col))
        rows[row] = rows[row] + byte(1<<col)
    }

    for i, row := range rows {
        fmt.Printf("%d: %b - %d\n", i, row, row)
        if row == 127 {
            fmt.Println("here")
            // found the row and seat (col number converting to binary)
            fmt.Printf("%d\n", 255-row)
        }
    }
    return 0
}

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
        //fmt.Printf("%s: %d - %d\n",string(b[i]), currentRowMax, currentRowMin)
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
        //fmt.Printf("%s: %d - %d\n",string(b[i]), currentColMax, currentColMin)
    }

    return currentRowMax,currentColMax,currentRowMax*8+currentColMax
}