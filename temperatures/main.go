package main

import "fmt"

func main() {
    

    // N: the number of temperatures to analyse
    var N int
    var temp int16
    var closest int16 
    // delta is always the absolute value of the closest so far
    var delta int16 = 5526
    
    fmt.Scanf("%d",&N)
    for i:=0;i<N;i++ {
        fmt.Scanf("%d",&temp)
        if temp > 0 {
            if temp <= delta {
                closest = temp
                delta = temp
            }
        } else if temp < 0 {
            if temp > 0 - delta {
                closest = temp
                delta = 0 - temp
            }
        }
    }
    
 
  fmt.Println(closest)
}
