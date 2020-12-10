package main

import (
    "fmt"
)

// Complete the bonAppetit function below.
func bonAppetit(bill []int32, k int32, b int32) {
    var sum int32;
    
    for _, value := range bill {
        sum += value
    }
    var amountDue int32 = (sum - bill[k]) / 2
    if amountDue == b {
        fmt.Println("Bon Appetit")
    } else {
        fmt.Println(b - amountDue)
    }

}
