package main

import (
  "fmt"
  "log"
  "math/rand"
  "net/http"
  "os"
  "time"
)

var asciiDice = map [int]string {
  1: ` -----
|     |
|  o  |
|     |
 -----`,

  2: ` -----
| o   |
|     |
|   o |
 -----`,

3: ` -----
| o   |
|  o  |
|   o |
 -----`,

4: ` -----
| o o |
|     |
| o o |
 -----`,

5: ` -----
| o o |
|  o  |
| o o |
 -----`,

6: ` -----
| o o |
| o o |
| o o |
 -----` }

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
  //fmt.Fprintf(w, "You rolled a %d.", rollDice())
  fmt.Fprintf(w, rollDice())
}

func rollDice() string {
  n := rand.Intn(6) + 1
  return asciiDice[n];
}
