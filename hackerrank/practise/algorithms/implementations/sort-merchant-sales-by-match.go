package main

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
    counts := make(map[int32]int32)
    var totalPair int32 = 0;
    for _, value := range ar {
        _, ok := counts[value]
        if ok {
            counts[value]++
        } else {
            counts[value] = 1
        }
    }
    for _,value := range counts {
        totalPair += int32(value/2)
    }
   
    return totalPair;
}