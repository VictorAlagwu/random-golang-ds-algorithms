package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "sort"
    "strconv"
    "strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
    var maxValue int = 0;
    var minValue int = 0;

    newArr := []int{int(arr[0]), int(arr[1]), int(arr[2]), int(arr[3]), int(arr[4])}
    sort.Ints(newArr)

    maxValue = newArr[1] + newArr[2] + newArr[3] + newArr[4]
    minValue = newArr[0] + newArr[1] + newArr[2] + newArr[3]
    

    fmt.Println(int64(minValue), int64(maxValue))
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
