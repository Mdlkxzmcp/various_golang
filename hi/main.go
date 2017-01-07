package main

import (
    "fmt"
    "math"
    "sync"
    )

func main() {
    gameTitles := []string{"Dex", "CrossCode", "FEZ", "Jotun", "Undertale"}
    var wg sync.WaitGroup
    wg.Add(len(gameTitles))
    for _, title := range gameTitles {
        go printTitle(title, &wg)
    }
    wg.Wait()
}


func printTitle(t string, wg *sync.WaitGroup)  {
    result := 0.0
    for i := 0; i < 100000000; i++ {
        result += math.Pi * math.Sin(float64(len(t)))
    }
    fmt.Println("Title: ", t)
    wg.Done()
}
