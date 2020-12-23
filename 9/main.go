package main 
  
import ( 
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
) 


func main() { 
    file, err := os.Open("input.txt") 
  
    if err != nil { 
        log.Fatalf("failed to open") 
    } 
    scanner := bufio.NewScanner(file) 
    scanner.Split(bufio.ScanLines) 
    var nums []int

    for scanner.Scan() {
        n, _ := strconv.Atoi(scanner.Text())
        nums = append(nums, n)
    }

    file.Close() 

    //fmt.Println(partOne(nums))
    fmt.Println(partTwo(nums))
}

func partOne(nums []int) int {
    preamble := 5
    for i := 0; i+preamble < len(nums); i++ {
        prev_nums := nums[i:i+preamble]
        num_to_evaluate := nums[i+preamble]
        fmt.Print(prev_nums)
        fmt.Printf(" - %d\n", num_to_evaluate)

        if !validate(prev_nums, num_to_evaluate) {
            fmt.Printf("Invalid: %d\n", num_to_evaluate)
            return num_to_evaluate
        }
    }
    return 0
}

func validate(prevNums []int, num int) bool {
    for i := 0; i < len(prevNums); i++ {
        n1 := prevNums[i]
        for j := 0; j < len(prevNums) - 1; j++ {
            if j == i {
                continue
            }
            if (n1+prevNums[j]) == num {
                return true
            }
        }
    }
    return false
}

func addNums(nums []int) int {
    var result int
    for _, num := range nums {
        result += num
    }
    return result
}

func partTwo(nums []int) int {
    // solution of part one 1038347917
    for i := 0; i < len(nums)-1; i++ {
        for j := i+1; j < len(nums); j++ {
            sum := addNums(nums[i:j])
            fmt.Print(nums[i:j])
            fmt.Printf(" = %d\n", sum)
            if sum > 1038347917 {
                break
            } else if sum == 1038347917 {
                fmt.Println("Found solution")
                return 0
            }
        }
    }
    return 0
}

