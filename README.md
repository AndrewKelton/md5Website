Simple Login and Sign Up Website

The passwords are saved as a one way md5 hash and saved into the SQL database 'lamp.db'.

To Build:
  1. Download files
  2. Follow download instructions on "github.com/mattn/go-sqlite3" to get sqlite3 driver
  3. Add pkg to folder with md5Website files
  4. Navigate to folder
  5. In terminal run:
      export GO111MODULE=on
      go get github.com/mattn/go-sqlite3
      go run main.go login.go newProf.go
  6. Type 'http://127.0.0.1:8081' into web browser
  7. Input data and view SQL database changing 

ERRORS
  1. If user inputs a username that is not in the database when trying to login, the code will exit with exit status 1.
  This is due to there being no check to see if that username exists in the code. This will be fixed.

  2. Passwords are inputted as type "text" instead of type "password" in html script. Still trying to figure out how to
  change a type "password" to a type "text" before returning it the server side.
