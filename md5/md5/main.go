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

var (
	user User
	lIn bool
	captcha string
)

func main() {

	user.Username = ""
	user.Password = ""
	lIn = false

	// implement server and client side
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        RootHandler(res, req, &user, lIn)
    })
    http.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(1)
        lIn = GetUser(res, req, &user)
		log.Println(lIn)
    })
	http.HandleFunc("/register", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(2)
		log.Println(lIn)
        RegisAcc(res, req, &user)
    })
	http.HandleFunc("/newaccount", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(3)
		log.Println(lIn)
        NewAcc(res, req, &user)
    })
	http.HandleFunc("/admin", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(4)
		log.Println(lIn)
        AdminAccess(res, req, &user)
    })
	http.HandleFunc("/admindelete", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(5)
		log.Println(lIn)
        DeleteAcc(res, req, &user)
		Table()
		AdminAccess(res, req, &user)
    })
	http.HandleFunc("/about", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(6)
		log.Println(lIn)
        About(res, req, lIn)
    })
	http.HandleFunc("/account", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(7)
		log.Println(lIn)
        AccountView(res, req, lIn, &user)
    })
	http.HandleFunc("/malware", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(8)
		log.Println(lIn)
		if lIn {
        	serveHTML(res, "frontpage/malware.html", lIn)
		} else {
			http.Error(res, "Access Denied", http.StatusForbidden)
		}
    })
	http.HandleFunc("/sourcecode", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(9)
		log.Println(lIn)
		if lIn {
        	serveHTML(res, "frontpage/fun.py", &user)
		} else {
			http.Error(res, "Access Denied", http.StatusForbidden)
		}
    })
	http.HandleFunc("/verify", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(10)
		log.Println(lIn)
		Verify(res, req, captcha)
	})
	http.HandleFunc("/captcha", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(11)
		log.Println(lIn)
		Captcha(res, req, CreateCap(), &lIn)
    })
	http.HandleFunc("/signout", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(12)
		log.Println(lIn)
		SignOut(res, req, &user)
    })
	http.HandleFunc("/home", func(res http.ResponseWriter, req *http.Request) {
		if lIn {
			serveHTML(res, "frontpage/front.html", lIn)
		} else {
			serveHTML(res, "frontpage/index.html", lIn)
		}
		})

    if err := http.ListenAndServe(":40", nil); err != nil {
        log.Fatal(err)
    }
	return
}

// main html
func RootHandler(res http.ResponseWriter, req *http.Request, u *User, l bool) {
	lIn = l

    serveHTML(res, "frontpage/index.html", nil)
}

// get captcha 
func Captcha(res http.ResponseWriter, req *http.Request, cap string, lIn *bool) {
    if cap == "" {
        http.Error(res, "Error generating CAPTCHA", http.StatusInternalServerError)
        return
    }
	captcha = cap

	if !*lIn{
		http.Error(res, "Access Denied", http.StatusForbidden)
		return
	}
	*lIn = false

    // Pass captcha numbers to the HTML template
    data := struct {
        Captcha string
    }{
        Captcha: cap,
    }

    serveHTML(res, "frontpage/captcha.html", data)
}

// verify captcha
func Verify(res http.ResponseWriter, req *http.Request, captcha string) {
	err := req.ParseForm()
	if err != nil {
        log.Fatal(err)
		http.Error(res, "Internal server error", http.StatusInternalServerError)
        return
    }

	uCap := req.Form.Get("captcha1")

	if (uCap == captcha) {
		fmt.Fprintf(res, "204")
		lIn = true
	} else {
		fmt.Fprintf(res, "404")
		lIn = false
	}
}

// check login success or failure
func LoginSuccess(res http.ResponseWriter, req *http.Request, u *User) {
	// Call to ParseForm makes form fields available.
    err := req.ParseForm()
    if err != nil {
        log.Fatal(err)
        return
    }

	if len(u.Username) == 0 || len(u.Password) == 0  || Login(u.Username, u.Password) == false{
		http.Error(res, "Access Denied", http.StatusForbidden)
		log.Println("Someone tried accessing account info without account inputted")
		return
	}

    serveHTML(res, "frontpage/front.html", u)
}

// new account page
func RegisAcc(res http.ResponseWriter, req *http.Request, u *User) {
	// Call to ParseForm makes form fields available.
    err := req.ParseForm()
    if err != nil {
        log.Fatal(err)
        return
    }

    serveHTML(res, "frontpage/register.html", u)
}

// admin account access
func AdminAccess(res http.ResponseWriter, req *http.Request, u *User) {

	err := req.ParseForm()
    if err != nil {
        log.Fatal(err)
        return
    }
	if (u.Username != "admin" || u.Password != "admin") {
		http.Error(res, "Access Denied", http.StatusForbidden)
		log.Println("Someone tried logging in as admin through the browser!")
		return
	}

	Table()
	serveHTML(res, "frontpage/admin.html", nil) // Serve login.html file
}
 
// handle about me page
func About(res http.ResponseWriter, req *http.Request, lIn bool) {
	if lIn {
		serveHTML(res, "frontpage/about.html", nil)
	} else {
		http.Error(res, "Access Denied", http.StatusAccepted)
	}
}

// view account
func AccountView(res http.ResponseWriter, req *http.Request, /*username string, password string,*/lIn bool, u * User) {
	if !lIn {
		http.Error(res, "Access Denied", http.StatusForbidden)
		return
	}
	data := struct {
		Username string
		Password string
	}{
		Username: u.Username,
		Password: u.Password,
	}
	serveHTML(res, "frontpage/account.html", data)
}

// view html pages
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

// unused function
func GetData(u *User, username string, password string) {
	u.Username = username
	u.Password = password
}

// collect username and password from client
func GetUser(res http.ResponseWriter, req *http.Request, u *User) bool {

    //Call to ParseForm makes form fields available.
    err := req.ParseForm()
    if err != nil {
        log.Fatal(err)   
		return false
    }
	var acc User

    acc.Username = req.Form.Get("username")
    acc.Password = req.Form.Get("password")

	GetData(u, acc.Username, acc.Password)

	fmt.Printf("Received user: %s, pass: %s\n", acc.Username, acc.Password)
	if Login(acc.Username, acc.Password) == true && acc.Username != "admin"{
		fmt.Fprintf(res, "LoginSuccess")
		return true
    } else if Login(acc.Username, acc.Password) == false{
        fmt.Fprintf(res, "LoginFailed")
		return false
	} else if Login(acc.Username, acc.Password) == true && acc.Username == "admin" {
		fmt.Fprintf(res, "AdminAccess")
		return false
	}
	return false
}

// create new account
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
		fmt.Fprintf(res, "204")
		u.Username = ""
		u.Password = ""
		return
	} else if Newpro(u.Username, u.Password) == false {
		fmt.Fprintf(res, "404")
		u.Username = ""
		u.Password = ""
		return
	}
}

// delete account
func DeleteAcc(res http.ResponseWriter, req *http.Request, u *User) {

	err := req.ParseForm()
    if err != nil {
        log.Fatal(err)   
		return         
    }

	u.Username = req.Form.Get("username")

	fmt.Printf("Received user: %s\n", u.Username)
	
	if Delete(u.Username) == true {
		u.Username = "admin"
		fmt.Fprintf(res, "204")
		return
	} else {
		u.Username = "admin"
		fmt.Fprintf(res, "404")
		return
	}
}

// sign user out and take away permissions
func SignOut(res http.ResponseWriter, req *http.Request, u *User) {
	u.Username = ""
	u.Password = ""
	lIn = false
	RootHandler(res, req, &user, false)
}