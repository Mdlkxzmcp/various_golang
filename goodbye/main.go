package main

import (
    "fmt"
    "goodbye/model"
)


func main() {

    myHerbsBox := model.Box()
    for _, herb := range myHerbsBox {
        fmt.Println(herb.Prepare())
    }

}
