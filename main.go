package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {

     file, _ := os.Open("./raw/")
     scanner := bufio.NewScanner(file)
     for scanner.Scan() {
        fmt.Println(scanner.Text())
     }

}
