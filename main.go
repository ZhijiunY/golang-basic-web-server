package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 如果 request.URL.路徑 不等於 “” 時
	if r.URL.Path != "/hello" {
		// return
		// 404 = http.StatusNotFound
		http.Error(w, "404 not fond", http.StatusNotFound)
		return
	}
	// 如果回傳值不是GET
	if r.Method != "GET" {
		http.Error(w, "method is not supported ", http.StatusNotFound)
	}
	// 沒有以上錯誤時，將回傳w,值
	fmt.Fprintf(w, "hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// 如果錯誤解析表單
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm(): %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s", name)
	fmt.Fprintf(w, "Address = %s", address)
}

func main() {
	// 從 ./static 靜態檔案裡找到 index.html
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port:8080\n ")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
