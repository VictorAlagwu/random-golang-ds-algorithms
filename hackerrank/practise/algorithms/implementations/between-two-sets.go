package main

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getTotalX(a []int32, b []int32) int32 {
    // Write your code here
    var noOfFactors int32;
    var startPoint int32 = a[len(a) - 1];
    var endPoint int32 = b[0];
    
    for i := startPoint; i <= endPoint; i++ {
        isFactor := true
        for _, value := range a {
            if i % value != 0 {
                isFactor = false;
                break;
            }
        }
        
        for  _, value := range b {
             if value % i != 0 {
                isFactor = false;
                break;
            }
        }
        if isFactor == true {
            noOfFactors++
        }
    }
    return noOfFactors
}
