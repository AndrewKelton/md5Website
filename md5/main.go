package main

import (
    "os"
    "log"
    "net/http"
    "text/template"
	"fmt"
)

type User struct {
	Username string
	Password string
}

func main() {

	var user User

	http.HandleFunc("/", RootHandler)
    http.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
        GetUser(res, req, &user)
    })
    http.HandleFunc("/account", func(res http.ResponseWriter, req *http.Request) {
        LoginSuccess(res, req, &user)
    })
	http.HandleFunc("/register", func(res http.ResponseWriter, req *http.Request) {
        RegisAcc(res, req, &user)
    })
	http.HandleFunc("/newaccount", func(res http.ResponseWriter, req *http.Request) {
        NewAcc(res, req, &user)
    })
	http.HandleFunc("/admin", AdminAccess)
	http.HandleFunc("/admindelete", func(res http.ResponseWriter, req *http.Request) {
        DeleteAcc(res, req, &user)
		Table()
    })

    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatal(err)
    }

	return
}

func RootHandler(res http.ResponseWriter, req *http.Request) {
    serveHTML(res, "frontpage/index.html", nil)
}

func LoginSuccess(res http.ResponseWriter, req *http.Request, user *User) {
	// Call to ParseForm makes form fields available.
    err := req.ParseForm()
    if err != nil {
        log.Fatal(err)
        return
    }

	// user := User{
    //     Username: req.FormValue("username"),
    //     Password: req.FormValue("password"),
    // }

    serveHTML(res, "frontpage/account.html", user)
}

// func LoginFailed(res http.ResponseWriter, req *http.Request) {
// 	serveHTML(res, "frontpage/loginfail.html", nil) // Serve login.html file
// }

func RegisAcc(res http.ResponseWriter, req *http.Request, user *User) {
	// Call to ParseForm makes form fields available.
    err := req.ParseForm()
    if err != nil {
        log.Fatal(err)
        return
    }

	// user := User{
    //     Username: req.FormValue("username"),
    //     Password: req.FormValue("password"),
    // }

    serveHTML(res, "frontpage/register.html", user)
}

func AdminAccess(res http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
    if err != nil {
        log.Fatal(err)
        return
    }
	Table()
	serveHTML(res, "frontpage/admin.html", nil) // Serve login.html file
}

func serveHTML(res http.ResponseWriter, filename string, data interface{}) {
	file, err := os.ReadFile(filename)
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error reading file: %v", err)
		return
	}

	t := template.New("")
	t, err = t.Parse(string(file))
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}

	err = t.Execute(res, data)
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
		return
	}
}

func GetData(u *User, username string, password string) {
	u.Username = username
	u.Password = password
}

func GetUser(res http.ResponseWriter, req *http.Request, u *User) {

    //Call to ParseForm makes form fields available.
    err := req.ParseForm()
    if err != nil {
        log.Fatal(err)   
		return         
    }
	var acc User

    acc.Username = req.Form.Get("username")
    acc.Password = req.Form.Get("password")

	GetData(u, acc.Username, acc.Password)

	fmt.Printf("Received user: %s, pass: %s\n", acc.Username, acc.Password)
	if Login(acc.Username, acc.Password) == true && acc.Username != "admin"{
		fmt.Fprintf(res, "LoginSuccess")
		//LoginSuccess(res, req)
		return
    } else if Login(acc.Username, acc.Password) == false{
        // Send a response to the client to indicate failed login
        fmt.Fprintf(res, "LoginFailed")
		return
	} else if Login(acc.Username, acc.Password) == true && acc.Username == "admin" {
		fmt.Fprintf(res, "AdminAccess")
		return
	}
}

func NewAcc(res http.ResponseWriter, req *http.Request, u *User) {

	err := req.ParseForm()
    if err != nil {
        log.Fatal(err)   
		return         
    }

	u.Username = req.Form.Get("username")
	u.Password = req.Form.Get("password")

	fmt.Printf("Received user: %s, pass: %s\n", u.Username, u.Password)
	
	if Newpro(u.Username, u.Password) == true {
		fmt.Fprintf(res, "Success")
		return
	} else if Newpro(u.Username, u.Password) == false {
		fmt.Fprintf(res, "Error")
		return
	}
}

func DeleteAcc(res http.ResponseWriter, req *http.Request, u *User) {

	err := req.ParseForm()
    if err != nil {
        log.Fatal(err)   
		return         
    }

	u.Username = req.Form.Get("username")

	fmt.Printf("Received user: %s\n", u.Username)
	
	if Delete(u.Username) == true {
		fmt.Fprintf(res, "204")
		return
	} else {
		fmt.Fprintf(res, "404")
		return
	}
}