package main 
  
import ( 
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
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
    fmt.Println(partTwo(lines))
}

func partOne(lines []string) int {

    var groupAnswers strings.Builder
    var currentGroup int
    var all [2000]int
    for _, line := range lines {
        if currentGroup > 2000 {
            fmt.Println("out of range")
            return 0
        }
        if line == "" {
            groupAnswers.Reset()
            currentGroup++
            // new group
            continue
        }
        for _, char := range line {
            //fmt.Printf("%s - %b\n", string(char), char)
            if !strings.ContainsRune(groupAnswers.String(), char) {
                groupAnswers.WriteRune(char)
            }
        }
        all[currentGroup] = len(groupAnswers.String())
    }
    var result int
    for _, g := range all {
        result += g
    }
    return result
}


func partTwo(lines []string) int {
    var sum int
    var currentReferenceLine strings.Builder
    currentReferenceLine.WriteString(lines[0])
    for i := 1; i < len(lines); {
        line := lines[i]
        //fmt.Printf("%s - %b\n", string(line), line)

        prevReferenceLine := currentReferenceLine.String()
        currentReferenceLine.Reset()
        for _, char := range prevReferenceLine {
            if strings.ContainsRune(line, char) {
                currentReferenceLine.WriteRune(char)
            }
        }

        if i < len(lines)-1 && lines[i+1] == "" {
            fmt.Println(currentReferenceLine.String())
            sum += currentReferenceLine.Len()
            currentReferenceLine.Reset()
            currentReferenceLine.WriteString(lines[i+2])
            i += 2
            continue
        }
        i++
    }
    sum += currentReferenceLine.Len()

    return sum
}

