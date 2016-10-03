package main

import (
  "fmt"
  "log"
  "math/rand"
  "net/http"
  "os"
  "time"
)

func main() {
  rand.Seed(time.Now().UTC().UnixNano())

  if(len(os.Args) < 2) {
    log.Fatal("Usage: 'diceRoller 3000' to listen on port 3000")
  }
  
  port := os.Args[1]

  http.HandleFunc("/", diceHandler)
  log.Println("Starting up on port " + port + "...")
  log.Fatal(http.ListenAndServe(":" + port, nil))
}

func diceHandler(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "You rolled a %d.", rollDice())
}

func rollDice() int {
  return rand.Intn(6) + 1
}

