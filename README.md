# What is this project?

This project is just a simple GoLang Project that displays the syntax that I can go back to later as a reference

# Install Go

https://go.dev/doc/install?download

# How to create & run a Golang program

1. Creating your first module, you will run this command:
	- ```go mod init {name}```
		- The name can be anything, but if you plan on publishing this module, then it should be the url path to where your versioning control is.
		- Example: go mod init github.com/jakedansoncyber/GoLangHelloWorld
		- Note, do not include https://
2. Whichever folder you run this command in, this is now the module. Any folders below this folder are what will be called packages. 
3. Write code:
	- I created a file called hello.go inside of the HelloWorld folder. This hello.go specifies that it is an application and will be ran as an application.
4. Now we can run our code. But here are three different options we can choose from:
	- Run: Go compiled the app into a temp directory, ran it and then cleaned it up after it was done.
	- Build: Creates an executable with the name of the package that you created it in. If I said "go build .", my executable would be named HelloWorld in this case.
	- Install: Installs it locally in the go/bin folder
	- For simplicity, we will just run the command: ```go run .```
		- Make sure you are in the folder that your application lies in. This command will run any go application in the current directory.  
