package main 
  
import ( 
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "sort"
) 


func main() { 
    file, err := os.Open("input2.txt") 
  
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
    sort.Ints(nums)

    diffs := make(map[int]int)
    for i := 0; i < len(nums); i++{
        var diff int
        if i == 0 {
            diff = nums[0]
        } else {
            diff = nums[i]-nums[i-1]
        }
        fmt.Printf("%d - diff %d\n", nums[i], diff)
        diffs[diff]++
    }
    //device
    diffs[3]++
    fmt.Println(diffs)
    return diffs[1]*diffs[3]
}

func partTwo(nums []int) int64 {
    sort.Ints(nums)
    var paths = make([]int, len(nums))
    for i := len(nums)-1; i >= 0; i--{
        
        if i < len(nums)-1 && nums[i] - nums[i+1] <= 3{
            paths[i]++
            paths[i] = paths[i] * paths[i+1]
        } else {
            fmt.Println("initial")
            paths[i] = 1
        }
        if i < len(nums)-2 && nums[i] - nums[i+2] <= 3{
            paths[i]++
            paths[i] = paths[i] * paths[i+2]
        }
        if i < len(nums)-3 && nums[i] - nums[i+3] <= 3{
            paths[i]++
            paths[i] = paths[i] * paths[i+3]
        }
        
        fmt.Printf(" %d - paths:%d\n", nums[i], paths[i])
    }
    return 0
}

func mulNums(nums []int) int64 {
    result := int64(1)
    for _, num := range nums {
        if num == 0 {
            continue
        }
        result = result*int64(num)
    }
    return result
}
