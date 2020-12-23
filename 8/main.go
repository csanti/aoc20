package main 
  
import ( 
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
) 

const RUN_LIMIT = 1000000

type instruction struct {
    operation string
    value int
    run_times int
}

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

    instruction_set := make(map[int]instruction)
    for i, line := range lines {
        //fmt.Println(line)
        op := line[:3]
        val, _ := strconv.Atoi(line[4:])
        instruction_set[i] = instruction {
            operation: op,
            value: val,
        }
        //fmt.Println(instruction_set[i])
    }

    var accumulator int
    in_num := 0
    for in_num < len(instruction_set) {
        fmt.Println(instruction_set[in_num])
        ins := instruction_set[in_num]
        if ins.run_times > 0 {
            break
        }
        next_in_num := in_num + 1
        switch(ins.operation) {
        case "nop":
        case "acc":
            accumulator += ins.value
        case "jmp":
            next_in_num = in_num + ins.value
        default:
            fmt.Println("unrecognized op")
        }
        ins.run_times = 1
        instruction_set[in_num] = ins
        fmt.Printf("Accumulator = %d Next instruction = %d\n",accumulator, next_in_num)
        in_num = next_in_num
    }
    return accumulator
}


func partTwo(lines []string) int {
    var accumulator int
    var iteration int
    out:
    for iteration < len(lines) {
        fmt.Printf("Iteration #%d\n",iteration)
        instruction_set := make(map[int]instruction)
        for i, line := range lines {
            op := line[:3]
            val, _ := strconv.Atoi(line[4:])
            if i == iteration {
                if op == "jmp" {
                    op = "nop"
                } else if op == "nop" {
                    op = "jmp"
                }
            }
            instruction_set[i] = instruction {
                operation: op,
                value: val,
            }
        }
        var total_run_times int
        accumulator = 0
        in_num := 0
        for in_num < len(instruction_set) {
            //fmt.Println(instruction_set[in_num])
            ins := instruction_set[in_num]
            /*
            if ins.run_times > 0 {
                break
            }*/
            next_in_num := in_num + 1
            switch(ins.operation) {
            case "nop":
            case "acc":
                accumulator += ins.value
            case "jmp":
                next_in_num = in_num + ins.value
            default:
                fmt.Println("unrecognized op")
            }
            ins.run_times = 1
            instruction_set[in_num] = ins
            //fmt.Printf("Accumulator = %d Next instruction = %d\n",accumulator, next_in_num)
            in_num = next_in_num
            total_run_times++
            if total_run_times > RUN_LIMIT {
                fmt.Println("infinity")
                break
            }
            if in_num >= len(instruction_set) {
                fmt.Println("***************************************************************************")
                break out
            }
        }
        iteration++
    }
    return accumulator
}

