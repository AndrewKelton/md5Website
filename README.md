Simple Login and Sign Up Website
-

The passwords are saved as a one way md5 hash and saved into the SQL database 'lamp.db'.
Users can only navigate the website when logged in and successfuly completing captcha.

Updates:
-
  Entirely new UI

  Users can go to multiple pages
  
  Users gain permissions based on logging in and captcha

To Build:
-
  1. Download files
  2. Follow download instructions on "github.com/mattn/go-sqlite3" to get sqlite3 driver
  3. Add pkg to folder with md5Website files
  4. Navigate to folder
  5. In terminal run:

      export GO111MODULE=on
     
      go get github.com/mattn/go-sqlite3
     
      go run main.go functions.go
     
  6. Type 'http://127.0.0.1:8081' into web browser
  7. Input data in forms in web app and view SQL database changing

ERRORS
-
  1. Passwords are inputted as type "text" instead of type "password" in html script. Still trying to figure out how to
  change a type "password" to a type "text" before returning it the server side.
  2. Issue with refreshing page in admin account

Files Created with AI
-
  Most but not all of the JavaScript code embedded in the html files are written by AI
  
  Table function in 'functions.go' is written partly by AI giving the prompts of my previous Go functions with SQL 
  
