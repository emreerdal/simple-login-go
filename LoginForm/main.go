package main

import (
	"fmt"
	"net/http"
	helper "./helpers"
	connection "./connections"
)
func main(){
	username,password,passwordConfirm := "","",""
	mux := http.NewServeMux()
	connection.ConnectPostgres("localhost",5432,"postgres","root","login")
	//Singup
	mux.HandleFunc("/singup",func(w http.ResponseWriter, r *http.Request){
		r.ParseForm()

		username = r.FormValue("username")
		password = r.FormValue("password")
		passwordConfirm =r.FormValue("passwordConfirm")
	
		usernameCheck := helper.IsEmpty(username)
		passwordCheck := helper.IsEmpty(password)
		passwordConfirmCheck := helper.IsEmpty(passwordConfirm)

		if usernameCheck || passwordCheck || passwordConfirmCheck  {
			fmt.Fprintln(w,"ErrorCode is -10: There is empty")
			return
		}
		if password == passwordConfirm {
			fmt.Fprintln(w, "Registration successful.")
		}else{
			fmt.Fprintln(w, "Password information must be the same.")

		}


	})

	//Login
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request){
		r.ParseForm()

		username = r.FormValue("username")
		password = r.FormValue("password")

		usernameCheck := helper.IsEmpty(username)
		passwordCheck := helper.IsEmpty(password)

		if usernameCheck || passwordCheck{
			fmt.Fprintln(w,"ErrorCode is -10: There is empty")
			return
		}
		dbPassword := "123456e"
		dbUsername := "emre"

		if username == dbUsername && password == dbPassword {
			fmt.Fprintln(w,"Login success")
		}else{
			fmt.Fprintln(w,"Login failed")
		}

	})
	http.ListenAndServe(":8080",mux)

}