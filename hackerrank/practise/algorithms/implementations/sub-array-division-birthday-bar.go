package main

// Complete the birthday function below.
func birthday(s []int32, d int32, m int32) int32 {
    var count int32 = 0;
    var sum int32 = 0
    
    for i := 0; i < len(s); i++ {
        sum += s[i];
        if i >= int(m) - 1 {
            if sum == d {
                count++
            } 
            sum -= s[i - (int(m)-1)];
        }
    }
    
    return count
}
