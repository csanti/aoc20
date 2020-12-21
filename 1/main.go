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
    var values []int 
  
    for scanner.Scan() { 
        v, err := strconv.Atoi(scanner.Text())
        if err != nil {
        	fmt.Println(err)
        }
        values = append(values, v)
    } 
    file.Close() 
    for _, a := range values { 
    	for _, b := range values {
    		for _ , c := range values {
    			if a + b + c == 2020 {
	    			fmt.Printf("%d + %d + %d = 2020 \n",a, b, c)
	    			fmt.Println(a*b*c)
	    			return
    			}
    		}
    	}
    } 
} 