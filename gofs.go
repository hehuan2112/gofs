package main
import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("* started server here on 0.0.0.0:8000")
    http.Handle("/", http.FileServer(http.Dir("./")))
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
       fmt.Println(err)
       panic(err)
    }   
}
