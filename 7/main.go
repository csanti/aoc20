package main 
  
import ( 
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
) 

type tree struct {
    references map[string][]string
    result []string
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
    t := tree {
        references: make(map[string][]string),
    }
    
    for _, line := range lines {
        a := strings.Split(line, " bags contain ")
        bag_color := strings.Replace(a[0], " ", "_", -1)
        b := strings.Split(a[1], ", ")
        for _, i := range b {
            c := strings.Split(i, " bag")
            in_bag_color := strings.Replace(c[0][2:], " ", "_", -1)
            t.references[in_bag_color] = append(t.references[in_bag_color], bag_color)
        }

    }
    t.iterateTree("shiny_gold")

    return len(t.result)
}



func (t *tree) iterateTree(branch string) {
    root := t.references[branch]
    for _, leaf := range root {
        if t.resultContains(leaf) {
            continue
        }
        t.result = append(t.result, leaf)
        t.iterateTree(leaf)
    }
}

func (t *tree) resultContains(bagName string) bool {
    for _, r := range t.result {
        if r == bagName {
            return true
        }
    }
    return false
}

func (t *tree) iterateCount(branch string, count int) int {
    root := t.references[branch]
    for _, leaf := range root {
        count++
        count = t.iterateCount(leaf, count)
    }
    return count
}

func partTwo(lines []string) int {
    t := tree {
        references: make(map[string][]string),
    }
    for _, line := range lines {
        a := strings.Split(line, " bags contain ")
        bag_color := strings.Replace(a[0], " ", "_", -1)
        b := strings.Split(a[1], ", ")
        for _, i := range b {
            c := strings.Split(i, " bag")
            num_bags, err := strconv.Atoi(c[0][:1])
            if err != nil {
                // contains no other bag
                continue
            }
            in_bag_color := strings.Replace(c[0][2:], " ", "_", -1)
            for l := 0; l < num_bags; l++ {
                t.references[bag_color] = append(t.references[bag_color], in_bag_color)
            }
        }
    }
    return t.iterateCount("shiny_gold", 0)
}

