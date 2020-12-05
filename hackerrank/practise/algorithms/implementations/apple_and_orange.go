package main

import (
	"fmt"
)

func main () {
    
}

// Complete the countApplesAndOranges function below.
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
    var noOfApplesNearSam int32 = 0;
    var noOfOrangesNearSam int32 = 0;
    
    for _,value := range apples {
        distancedCovered := a + value
        if distancedCovered  >= s && distancedCovered <= t {
            noOfApplesNearSam++
        }
    }

    for _,value := range oranges {
        distancedCovered := b + value
        if distancedCovered  >= s && distancedCovered <= t {
            noOfOrangesNearSam++
        }
    }
    
    fmt.Println(noOfApplesNearSam)
    fmt.Println(noOfOrangesNearSam)
}

