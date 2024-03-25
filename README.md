My Website
-

The passwords are saved as a one way md5 hash and saved into the SQL database 'lamp.db'.

Users can only navigate the website when logged in and successfuly completing captcha.

You can learn a little about me and view my first website in Go

Updates:
-
  Entirely new UI

  Users can go to multiple pages
  
  Users gain permissions based on logging in and captcha

  Automate terminal needs by using shell script to run code

To Build:
-
  1. Download files
  2. Follow download instructions on "github.com/mattn/go-sqlite3" to get sqlite3 driver
  3. Add pkg to folder with myWebsite
  4. Navigate to folder
  5. In terminal run:

     chmod +x run.sh

     ./run.sh
     
  5. Type 'http://127.0.0.1:40' into web browser (or whatever port you are using)
  6. Input data in forms in web app and view SQL database changing
  7. Navigate website

ERRORS
-
  1. Passwords are inputted as type "text" instead of type "password" in html script. Still trying to figure out how to
  change a type "password" to a type "text" before returning it the server side.
  2. Issue with refreshing page in admin account and deleting accounts. Account will
     be deleted but will throw 'Error deleting account'. Will need to log back into
     admin account to view changes.

Files Created with AI
-
  Most but not all of the JavaScript code embedded in the html files are written by AI
  
  Table function in 'functions.go' is written partly by AI giving the prompts of my previous Go functions with SQL 
  
