package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the simpleArraySum function below.
 */
func simpleArraySum(ar []int32) int32 {
    var result int32 = 0
    
    for _, number := range ar {
        result += number
    } 
    
    return result
}

func main() {
    var ar []int32 = [1, 2, 3, 4, 10, 11]

    result := simpleArraySum(ar)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}
