package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type User struct {
	Id string		` json:"id,omitempty" `
	Email string	` json:"email,omitempty" `
}

var users = []User{
	{Id:"1", Email:"1email@naver.com"},
	{Id:"2", Email:"2email@naver.com"},
	{Id:"3", Email:"3email@naver.com"},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(users)
	w.Write(j)
}

func getPerson(w http.ResponseWriter, r *http.Request){
	param :=mux.Vars(r)

	for _, item := range users {
		if item.Id == param["Id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

func postUsers(w http.ResponseWriter, r *http.Request) {
	var m User
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &m)
	users = append(users, m)

	json.NewEncoder(w).Encode(users)
}

func deleteUsers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range users {
		if item.Id == params["Id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func putUsers(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	var m User
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &m)

	for index, item := range users {
		if item.Id == params["Id"] {
			m.Id = m.Id
			m.Email = m.Email
			users[index] = m
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
    router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{Id}", getPerson).Methods("GET")
	router.HandleFunc("/users", postUsers).Methods("POST")
	router.HandleFunc("/users/{Id}", deleteUsers).Methods("DELETE")
	router.HandleFunc("/users/{Id}", putUsers).Methods("PUT")

    http.ListenAndServe("localhost:8089", router)
}


