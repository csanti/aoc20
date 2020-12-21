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

    fmt.Println(partOne(lines))
    //fmt.Println(partTwo(lines))
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
    return 0
}

