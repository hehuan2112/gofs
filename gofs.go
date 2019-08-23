package main
import (
    "fmt"
    "flag"
    "net/http"
)

func main() {
    var port string
    flag.StringVar(&port, "p", "8000", "Port Number")
    flag.Parse()
    fmt.Println("* started server here on 0.0.0.0:" + port)
    http.Handle("/", http.FileServer(http.Dir("./")))
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
       fmt.Println(err)
       panic(err)
    }
}
