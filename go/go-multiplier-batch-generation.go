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

/*
 * Complete the 'BurstyRateLimiter' function below.
 *
 * The function accepts following parameters:
 *  1. chan bool requestChan
 *  2. chan int resultChan
 *  3. int batchSize
 *  4. int toAdd
 */

func BurstyRateLimiter(requestChan chan bool, resultChan chan int, batchSize int, toAdd int) {
    i := 0 // This variable keeps track of the current number in the sequence
    // TODO
    time.Sleep(10 * time.Millisecond)

    for {
        <-requestChan // Wait for a request signal

        // Generate a batch of values
        for j := 0; j < batchSize; j++ {
            resultChan <- i * toAdd
            i++ // Increment i for the next value
        }

        // Wait for at least 10 milliseconds before processing the next batch
        time.Sleep(10 * time.Millisecond)
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    skipTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    skipBatches := int(skipTemp)

    totalTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    printBatches := int(totalTemp)

    batchSizeTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    batchSize := int(batchSizeTemp)

    toAddTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    toAdd := int(toAddTemp)

    resultChan := make(chan int)
    requestChan := make(chan bool)
    go BurstyRateLimiter(requestChan, resultChan, batchSize, toAdd)
    
    for i := 0; i < skipBatches + printBatches; i++ {
        requestChan <- true // Signal the server to process a batch
        
        for j := 0; j < batchSize; j++ {
            new := <-resultChan
            if i < skipBatches {
                continue
            }
            fmt.Println(new)
        }
        if i >= skipBatches && i != skipBatches + printBatches - 1 {
            fmt.Println("-----")
        }
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
