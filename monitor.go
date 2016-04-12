package main

import (
	"html/template"
	"net/http"
	"os/exec"
)

type Page struct {
	Title string
	Body  string
	Type  string
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveHTTP)
	http.HandleFunc("/ls", cmdLS)
	http.HandleFunc("/vmstat", cmdVmstat)
	http.HandleFunc("/free", cmdFree)
	http.HandleFunc("/top", cmdTop)
	http.HandleFunc("/iostat", cmdIostat)
	http.ListenAndServe(":8080", nil)
}

func serveTemplate(d http.ResponseWriter, page *Page) {
	d.Header().Add("Content Type", "text/html")
	var file string
	if page.Type == "home" {
		file = "home"
	} else {
		file = "command"
	}
	tmpl, _ := template.ParseFiles("templates/home.html", "templates/command.html")
	tmpl.ExecuteTemplate(d, file, page)
}

func serveHTTP(d http.ResponseWriter, req *http.Request) {
	serveTemplate(d, &Page{Title: "Home", Body: "", Type: "home"})
}

func cmdLS(d http.ResponseWriter, req *http.Request) {
	c1 := exec.Command("ls")
	out, err := c1.Output()
	if err != nil {
		panic(err)
	}
	serveTemplate(d, &Page{Title: "Command: ls", Body: string(out), Type: "command"})
}

func cmdFree(d http.ResponseWriter, req *http.Request) {
	c1 := exec.Command("free", "-h")
	out, err := c1.Output()
	if err != nil {
		panic(err)
	}
	serveTemplate(d, &Page{Title: "Command: free", Body: string(out), Type: "command"})
}

func cmdTop(d http.ResponseWriter, req *http.Request) {
	c1 := exec.Command("top", "-b", "-n 1")
	out, err := c1.Output()
	if err != nil {
		panic(err)
	}
	serveTemplate(d, &Page{Title: "Command: top", Body: string(out), Type: "command"})
}

func cmdIostat(d http.ResponseWriter, req *http.Request) {
	c1 := exec.Command("iostat")
	out, err := c1.Output()
	if err != nil {
		out = []byte(`Command not available on this system`)
	}
	serveTemplate(d, &Page{Title: "Command: iostat", Body: string(out), Type: "command"})
}

func cmdVmstat(d http.ResponseWriter, req *http.Request) {
	c1 := exec.Command("vmstat")
	out, err := c1.Output()
	if err != nil {
		panic(err)
	}
	serveTemplate(d, &Page{Title: "Command: vmstat", Body: string(out), Type: "command"})
}
