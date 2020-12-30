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

    //fmt.Println(partOne(lines))
    fmt.Println(partTwo(lines))
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

func partTwo (lines []string) int {
    busTimes := strings.Split(lines[1], ",")
    
    refTime, _ := strconv.Atoi(busTimes[0])
    ttCheck, currentListIndex, _ := getNextTime(busTimes, 1)
    currentListIndex++

    fmt.Printf("checking: %d indexDiff:%d\n",ttCheck,currentListIndex)

    var ref, cand int
    mul := refTime
    min := 0
    i,j := 1,1
    var finish bool
    for {
        ref = min + (i * mul)
        cand = ttCheck * j
        if ref > cand {
            j++
            continue
        }
        //fmt.Printf("ref: %d - cand: %d\n",ref,cand)
        if (cand % ref) == currentListIndex-1 {
            fmt.Printf("found: %d,  ref: %d - cand: %d\n",i,ref,cand)
            mul = LCM(mul, ttCheck)
            min = ref
            ttCheck, currentListIndex, finish = getNextTime(busTimes, currentListIndex)
            if finish {
                break
            }
            currentListIndex++
            fmt.Printf("checking: %d indexDiff:%d mul:%d\n",ttCheck,currentListIndex,mul)
            i = 1
            j = cand/ttCheck
        } else {
            i++ 
        }
        if (i % 100000) == 0 {
            fmt.Printf("Processing... current ref %d\n",ref)
        }
    }
    return ref
}

func getNextTime(busTimes []string, currentListIndex int) (int, int, bool) {
    for {
        // new number
        if currentListIndex > len(busTimes) - 1 {
            return 0,0,true
        }
        if busTimes[currentListIndex] == "x" {
            currentListIndex++
        } else {
            ttCheck, _ := strconv.Atoi(busTimes[currentListIndex])
            return ttCheck, currentListIndex, false
        }
    }
}


// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
    for b != 0 {
        t := b
        b = a % b
        a = t
    }
    return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
    result := a * b / GCD(a, b)

    for i := 0; i < len(integers); i++ {
        result = LCM(result, integers[i])
    }

    return result
}