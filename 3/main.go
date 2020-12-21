package main 
  
import ( 
    "bufio"
    "fmt"
    "log"
    "os"
    "math/big"

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
    var encaunteredTreesCount int
    for i, l := range lines  { 
        if string(l[(i*3)%31]) == "#" {
            encaunteredTreesCount++
        }
    }
    return encaunteredTreesCount
}

func partTwo(lines []string) string {
    result := big.NewInt(1)

    a := countEncounteredTrees(lines, 1, 1)
    b := countEncounteredTrees(lines, 3, 1)
    c := countEncounteredTrees(lines, 5, 1)
    d := countEncounteredTrees(lines, 7, 1)
    e := countEncounteredTrees(lines, 1, 2)
    fmt.Printf("%d %d %d %d %d\n",a,b,c,d,e)
    result = result.Mul(result, big.NewInt(a))
    result = result.Mul(result, big.NewInt(b))
    result = result.Mul(result, big.NewInt(c)) 
    result = result.Mul(result, big.NewInt(d))
    result = result.Mul(result, big.NewInt(e))
    return result.String()
}


func countEncounteredTrees(lines []string, right int, downRate int) int64 {
    var encaunteredTreesCount int64
    var downCount int
    for i := 0; i < len(lines); i+=downRate {
        line := lines[i]
        if string(line[(downCount*right)%31]) == "#" {
            encaunteredTreesCount++
        }
        downCount++
    }
    return encaunteredTreesCount
}