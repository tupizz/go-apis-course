package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateUser user"))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAllUsers user"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUserById user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUser user"))
}
