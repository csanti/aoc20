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
}

func partOne(lines []string) int {
    var validPpCounter, ppCounter int
    var currentPpFieldCount int
    labels := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
    for _, l := range lines {
        if l == "" || currentPpFieldCount >= 7 {
            currentPpFieldCount = 0
            continue
        }
        for _, label := range labels {
            if strings.Contains(l, label) {
                currentPpFieldCount++
            }
        }
        if currentPpFieldCount >= 7 {
            validPpCounter++
        }
        fmt.Println(currentPpFieldCount)
    }
    return validPpCounter
}