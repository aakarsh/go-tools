package main

import (
    "fmt"
    "os"
    "bufio"
)

func printStream (file *os.File) {
    input := bufio.NewScanner(file)
    var i int
    for input.Scan() {
        fmt.Printf("%5d %s\n" , i, input.Text())
        i += 1
    }
}


func main() {
    files := os.Args[1:]
    if len(files) == 0 {
        printStream(os.Stdin)
        os.Exit(1);
    }
    for _, file := range files {
        input, err := os.Open(file)
        if err != nil {
            fmt.Fprintf(os.Stderr,"Couldn't open %s error:%v" , file,err)
            continue
        }
        printStream(input)
        input.Close()
    }
}
