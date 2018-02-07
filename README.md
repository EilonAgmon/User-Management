# User-Management

Hello There!

In order to run the application do the following:


Windows Users
=============
1. Download the application
2. Run the runner.bat file
3. Navigate to http://localhost:8080/



Mac Users
=========
1. Download the application
2. Run "go build" from the application's root ("./user_management") and run the created executable
3. Navigate to http://localhost:8080/



Trouble Shooting
================
If the app doesn't start, try running the following commands from the applcation's root: ("./user_management")

1. "go get github.com/gin-gonic/gin" (Gin Framework)
2. "go get github.com/rs/xid" (Time library)
3. "go run main.go routes.go models.user.go handlers.user.go" (Build the app)
4. Navigate to http://localhost:8080/


(Above are explicit commands meant to fetch the required dependencies)
