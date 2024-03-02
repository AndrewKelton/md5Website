Simple Login and Sign Up Website
-

The passwords are saved as a one way md5 hash and saved into the SQL database 'lamp.db'.

Updates:
-
  Added an admin profile with admin capabilities in the website
  Admin responsibilities do not require you to login as admin yet
  You just need to navigate to /admin page.

  Able to delete users in the database from the admin page and view database

To Build:
-
  1. Download files
  2. Follow download instructions on "github.com/mattn/go-sqlite3" to get sqlite3 driver
  3. Add pkg to folder with md5Website files
  4. Navigate to folder
  5. In terminal run:
      export GO111MODULE=on
      go get github.com/mattn/go-sqlite3
      go run main.go login.go newProf.go admintable.go
  6. Type 'http://127.0.0.1:8081' into web browser
  7. Input data and view SQL database changing

ERRORS
-
  1. Passwords are inputted as type "text" instead of type "password" in html script. Still trying to figure out how to
  change a type "password" to a type "text" before returning it the server side.

Files Created with AI
-
  Most but not all of the JavaScript code embedded in the html files are written by AI
  admintable.go is written by AI given the prompts of my pervious go code opening the db
  