package main

import (
 "bufio"
  "fmt"
  "log"
  "net"
  "os"
)

func check(domain string) (bool) {
  _, err := net.LookupHost(domain)
  return (err == nil)
}

func checkAsync(domain string, c chan string) {
  fmt.Println("CheckAsync ")
  c <- "Yay!"
  // if (check(domain)) {
  //   c <- (domain + " is free.")
  // } else {
  //   c <- (domain + " is taken.")
  // }
}

func print(c chan string) {
  for value := range c {
    fmt.Println(value)
  }
  fmt.Println("Leaving print.")
}

func main() {
  file, err := os.Open("domains.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  c := make(chan string)

  // go print(c)

  scanner := bufio.NewScanner(file)
  for (scanner.Scan()) {
    domain := scanner.Text()
    // fmt.Println(domain, check(domain))
    go checkAsync(domain, c)
  }
  print(c)
  // close(c)

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}
