# What is this project?

This project is just a simple GoLang Project that displays the syntax that I can go back to later as a reference

# How to run this program

1. Creating your first module, you will run this command:
	a. Go mod init {name}
		i. The name can be anything, but if you plan on publishing this module, then it should be the url path to where your versioning control is.
		ii. Example: go mod init github.com/jakedansoncyber/GoLangHelloWorld
			1) Note, do not include https://
2. Whichever folder you run this command in, this is now the module. Any folders below this folder are what will be called packages. 
3. Write code:
	a. I created a file called hello.go inside of the HelloWorld folder. This hello.go specifies that it is an application and will be ran as an application.
4. Now we can run our code. But here are three different options we can choose from:
	a. Run: Go compiled the app into a temp directory, ran it and then cleaned it up after it was done.
	b. Build: Creates an executable with the name of the package that you created it in. If I said "go build .", my executable would be named HelloWorld in this case.
	c. Install: Installs it locally in the go/bin folder
