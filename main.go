package main

import (
	"fmt";
  "math/rand";
  "time";
)

func main() {
  rand.Seed(time.Now().UnixNano())
	cnumber := rand.Intn(11-1) +1
  fmt.Print(cnumber)
  var number int ;
  for i := 5; i >= 1; i--{
    fmt.Printf("guess number(%d) ", i);
    fmt.Scan(&number);
    if cnumber == number{
      fmt.Print("correct");
      break;
    
    } else{
      if number < cnumber{
        fmt.Print("number is bigger\n")
      } else if number > cnumber{
        fmt.Print("number is lower\n")
      }
    }
  }
}
