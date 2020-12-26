package main 
  
import ( 
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
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

    myTime, _ := strconv.Atoi(lines[0])
    fmt.Println(myTime)
    busTimes := strings.Split(lines[1], ",")

    predictions := make(map[int]int)
    var earliestBus, earliestTime int
    earliestTime = 99999999
    for _,t := range busTimes {
        if t == "x" {
            continue
        }
        id,_ := strconv.Atoi(t)
        waitTime := id - (myTime%id)
        predictions[id] = waitTime
        if waitTime < earliestTime {
            earliestBus = id
            earliestTime = waitTime
        }
    }

    fmt.Println(predictions)
    return earliestBus * earliestTime
}

func partTwo(lines []string) int {

    return 0
}