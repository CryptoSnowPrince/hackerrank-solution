package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "time"
)

func ModuloFibonacciSequence(requestChan chan bool, resultChan chan int) {
    // TODO
    time.Sleep(10 * time.Millisecond)
    a, b := 1, 2
    for {
        select {
        case <-requestChan:
            resultChan <- a
            a, b = b, (a+b)%1000000000
            time.Sleep(10 * time.Millisecond)
        }
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    skipTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    skip := int32(skipTemp)

    totalTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    total := int32(totalTemp)

    resultChan := make(chan int)
    requestChan := make(chan bool)
    go ModuloFibonacciSequence(requestChan, resultChan)
    for i := int32(0); i < skip + total; i++ {
        start := time.Now().UnixNano()
        requestChan <- true
        new := <-resultChan
        if i < skip {
            continue
        }
        end := time.Now().UnixNano()
        timeDiff := (end - start)/1000000
        if timeDiff < 10 {
            time.Sleep(10*time.Millisecond - time.Duration(timeDiff)*time.Millisecond)
        }
        fmt.Println(new)
    }
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