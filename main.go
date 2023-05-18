package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
	<!DOCTYPE html>
	<html>
	<head>
		<style>
		body {
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            font-family: Sora, Arial, sans-serif;
            background-color: #10203A;
            color: white;
        }
		</style>
	</head>
	<body>
		<div class="App">
			Welcome to Back4App Containers
		</div>
	</body>
	</html>
	`)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
