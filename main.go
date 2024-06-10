package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    // /helloエンドポイントにhelloHandler関数を登録
    http.HandleFunc("/", helloHandler)

    // サーバーをポート8080で起動
    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Failed to start server: ", err)
    }
}
