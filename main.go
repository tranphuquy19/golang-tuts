package main

import (
	"fmt"
	"io"
)
import "net/http"

func main() {
	fmt.Println("test go")
	resp, _ := http.Get("https://api.ipify.org")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	ip := string(body)
	fmt.Println(ip)
}