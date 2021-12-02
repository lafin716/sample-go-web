package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

func startServer() {
	http.HandleFunc("/guestbook", viewHandler)
	fmt.Println("Server opened :: 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func templateExam() {
	text := "Here is my template\nAction : {{.}}\nTemplate end"
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, nil)
	check(err)

	err = tmpl.Execute(os.Stdout, "ABC")
	err = tmpl.Execute(os.Stdout, 42)
	err = tmpl.Execute(os.Stdout, true)
}

func main() {

}
