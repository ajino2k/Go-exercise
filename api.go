package main

import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "html/template"
)

func uptime(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "#check_status\n")
    out, err := exec.Command("ssh", "root@10.10.36.3","uptime").Output()
    if err != nil {
          log.Fatal(err)
    }
    fmt.Fprintf(w,"%s",out)
}

func login(w http.ResponseWriter, r *http.Request) {
    //get request method
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.html")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}


func action(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("action.html")
    t.Execute(w, "Hello World!")
}

func main(){

    http.HandleFunc("/uptime", uptime)
    http.HandleFunc("/login", login)
    http.HandleFunc("/action", action)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
