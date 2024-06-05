package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// import (
// 	"html/template"
// 	"log"
// 	"os"
// )

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	text := "Here's my template!\n"
// 	tmpl, err := template.New("test").Parse(text)
// 	check(err)
// 	err = tmpl.Execute(os.Stdout, nil)
// 	check(err)
// }

type GuestBook struct {
	SignatureCount int
	Signatures     []string
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()

	scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		lines = append(lines, scaner.Text())
	}
	check(scaner.Err())
	return lines
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("index.html")
	check(err)
	guestbook := GuestBook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}
func newHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(w, nil)
	check(err)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	signature := r.FormValue("newSignature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(w, r, "/html", http.StatusFound)
}

func main() {
	http.HandleFunc("/html", viewHandler)
	http.HandleFunc("/html/new", newHandler)
	http.HandleFunc("/html/create", createHandler)
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	log.Fatal(err)
}
