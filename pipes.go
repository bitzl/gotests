package main

import (
 "bufio"
  "fmt"
  "log"
  "net"
  "os"
)

func check(domains chan string, result chan string) {
  for domain := range domains {
    _, err := net.LookupHost(domain)
    if err == nil {
      result <- "- " + domain + " is free"
    } else {
      result <- "- " + domain + " is taken"
    }
  }
}

func print(results chan string, done chan bool) {
  for result := range results {
    fmt.Println(result)
  }
  fmt.Println("Leaving print.")
  // done <- true
}

func feed(filename string, domains chan string) {
  file, err := os.Open("domains.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for (scanner.Scan()) {
    domain := scanner.Text()
    fmt.Println("+ Checking " + domain)
    domains <- domain
  }

    if err := scanner.Err(); err != nil {
      log.Fatal(err)
    }


}

func main() {
  results := make(chan string)
  domains := make(chan string)
  done := make(chan bool)

  go check(domains, results)
  go print(results, done)

  feed("domains.txt", domains)

  // <-done;
}
