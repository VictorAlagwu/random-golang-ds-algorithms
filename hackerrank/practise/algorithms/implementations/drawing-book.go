package main

func pageCount(n int32, p int32) int32 {
    /*
     * Write your code here.
     */
    var countFromFront int32 = p/2;
    var countFromBack int32 = int32(n/2) - countFromFront;
    
    if (countFromBack < countFromFront) {
     return countFromBack;
    }
    return countFromFront;
}