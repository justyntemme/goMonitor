package main

import (
  "net/http"
  "os/exec"
)

func main(){

	http.HandleFunc("/",serveHTTP)
  http.HandleFunc("/ls", cmdLS)
	http.HandleFunc("/free", cmdFree)
	http.HandleFunc("/top", cmdTop)
  http.ListenAndServe(":8080",nil)

}

func serveHTTP(d http.ResponseWriter,req *http.Request){
	d.Header().Add("Content Type", "text/html")
	d.Write([]byte("To use a command include it in your url reques </br> Example http://localhost:8080/ls povides the output of the ls command </br> <b>Supported Commands</b> </br> ls free "))

}

func cmdLS(d http.ResponseWriter,req *http.Request){
	c1 := exec.Command("ls")
	out, err := c1.Output()
	d.Write(out)
	if err != nil{
		panic(err)
	}
}

func cmdFree(d http.ResponseWriter,req *http.Request){
	c1 := exec.Command("free", "-h")
	out, err := c1.Output()
	d.Write(out)
	if err != nil{
		panic(err)
	}
}

func cmdTop(d http.ResponseWriter, req *http.Request){
	c1 := exec.Command("top", "-b", "-n 1")
	out, err := c1.Output()
	d.Write(out)
	if err != nil{
		panic(err)
	}
}

