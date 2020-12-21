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

    fmt.Println(partOne(lines))
    fmt.Println(partTwo(lines))
}

func partOne(lines []string) int {
    var correctPwdCount int
    out:
    for _, l := range lines  { 
        var min, max int
        var letter, pwd string
        fmt.Sscanf(l,"%d-%d %1s: %s", &min, &max, &letter, &pwd)
        var count = 0;
        for _, s := range pwd {
            if string(s) == letter {
                count++
            }
            if count > max {
                continue out
            }
        }
        if count < min {
            continue out
        }
        correctPwdCount++
    }
    return correctPwdCount
}

func partTwo(lines []string) int {
    var correctPwdCount int
    for _, l := range lines  { 
        var pos1, pos2 int
        var letter, pwd string
        fmt.Sscanf(l,"%d-%d %1s: %s", &pos1, &pos2, &letter, &pwd)
        var count = 0;
        if string(pwd[pos1-1]) == letter {
            count++
        }
        if string(pwd[pos2-1]) == letter {
            count ++
        }
        if count == 1 {
            correctPwdCount++
        }
    }
    return correctPwdCount
}