package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

func main() {
  if len(os.Args) < 2 {
    log.Println("Please provide at least one filename!")
    os.Exit(1)
  }

  for _, filename := range os.Args[1:] {
    file, err := os.Open(filename) 
    if err != nil {
      log.Printf("Failed to open %s: %v", filename, err) 
      continue 
    }
    scanner := bufio.NewScanner(file) 
    scanner.Split(bufio.ScanWords) 

    var wordCount int 
     
    for scanner.Scan() {
      wordCount ++
    }

    if scanner.Err() != nil {
      log.Println("Scan error %s: %v", filename, scanner.Err()) 
      file.Close() 
      continue 
    }
    file.Close() 
    fmt.Printf("%s: %d\n", filename, wordCount)
  }
}

