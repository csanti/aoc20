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

func partOne(lines []string) uint64 {
    memory := make(map[int]uint64)
    var orMask, andMask uint64
    for _, line := range lines {
        a :=  strings.Split(line, " = ")
        if a[0][:3] == "mas" {
            // parse mask
            fmt.Printf("new mask: %s\n",a[1])
            orMask, _ = strconv.ParseUint(strings.Replace(a[1],"X","0",-1), 2, 64)
            //fmt.Printf("%b\n",orMask)
            andMask, _ = strconv.ParseUint(strings.Replace(a[1], "X", "1",-1), 2, 64)
            //fmt.Printf("%b\n",andMask)
            continue
        }
        addr, _ := strconv.Atoi(strings.Trim(a[0][3:],"[]"))
        toWrite, _ := strconv.ParseUint(a[1], 10, 64)
        fmt.Printf("Write mem:%d value:%b",addr,toWrite)
        var value uint64 = toWrite & andMask | orMask
        fmt.Printf(" result:%b\n",value)
        memory[addr] = value
    }
    var sum uint64
    for _, mem := range memory {
        sum += mem
    }
    return sum
}



func partTwo(lines []string) int {
    return 0
}

