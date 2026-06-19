package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://practicum.yandex.ru")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	if _, err = io.CopyN(os.Stdout, response.Body, 512); err != nil {
		fmt.Println((err))
	}
}
