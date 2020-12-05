package main

// Complete the kangaroo function below.
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
    var result string = "NO";
    if (x2 > x1 && v2 > v1) || (x1 > x2 && v1 > v2) {
        result = "NO"
        return result;
    }
    var i int32;
    for i = 0; i < 10000; i++ {
        if (x1 + (i * v1)) == (x2 + (i * v2)){
            result = "YES"
            break;
        }
    } 
    return result;
}
