package main

import (
    "log"
    "net/http"
)


func helloServer(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("hello"))
}

func main() {
    http.HandleFunc("/hello", helloServer)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
