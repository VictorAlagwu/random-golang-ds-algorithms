package main

import (
    "strconv"
)

// Complete the dayOfProgrammer function below.
func dayOfProgrammer(year int32) string {
    if year == 1918 {
        return "26.09." + strconv.Itoa(int(year))
    } else if (year <= 1917 && year % 4 == 0) ||
        (year > 1918 && ((year % 4 == 0 && year % 100 != 0) || year % 400 == 0)){
        return "12.09." + strconv.Itoa(int(year))
    } else {
        return "13.09." + strconv.Itoa(int(year))
    }
}
