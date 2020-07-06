package main
import "fmt"

func main() {
    var x int
    x = 100
        for l := 0; l < 20; l++ {
        if(l + 5 == x) {
            fmt.Print(l, "+ 5 == 14: ")
            break;
        }
         if(l != 9) {
            fmt.Println("%d", "ist nicht 9.", l)
        }
        
     if(x > (10 / 5 + 20)) {
        fmt.Println("x > (x / 1 + 20: %d ", l)
    } 
    }
    }