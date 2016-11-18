package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"strings"
	"bytes"
)

func Test_getUsers(t *testing.T) {
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "localhost:8089/users", nil)
	if err != nil{
		t.Fatal(err)
	}

	handler := http.HandlerFunc(getUsers)
	handler.ServeHTTP(res, req)

	if status := res.Code ; status != http.StatusOK {
		t.Errorf("%v", status)
		t.Errorf("%v", http.StatusOK)
	}

	expected := `[{"id":"1","email":"1email@naver.com"},{"id":"2","email":"2email@naver.com"},{"id":"3","email":"3email@naver.com"}]`

	if res.Body.String() != expected {
		t.Errorf("got  %v ", res.Body.String())
		t.Errorf("want %v", expected)
	}

}


func Test_getPerson(t *testing.T){
	req , _:= http.NewRequest("GET", "http://localhost:8089/users/2", nil)
	res := httptest.NewRecorder()

	m := mux.NewRouter()
	m.HandleFunc("/users/{Id}", getPerson)
	m.ServeHTTP(res, req)

	b, _ := ioutil.ReadAll(res.Body)

	if status := res.Code ; status != http.StatusOK {
		t.Errorf("%v", status)
		t.Errorf("%v", http.StatusOK)
	}

	expected := `{"id":"2","email":"2email@naver.com"}`

	if strings.Replace(string(b), "\n","",-1) != expected {
		t.Errorf("res.Body %v", strings.Replace(string(b), "\n","",-1))
		t.Errorf("expected %v", expected)
	}

}

func Test_postUsers(t *testing.T) {
	data := `{"id":"4","email":"4email@naver.com"}`
	var jsonStr = []byte(data)

	req, err := http.NewRequest("POST", "http://localhost:8089/users", bytes.NewBuffer(jsonStr))
	if err != nil{
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	m := mux.NewRouter()
	m.HandleFunc("/users", postUsers)
	m.ServeHTTP(res, req)

	b, _ := ioutil.ReadAll(res.Body)

	if status := res.Code ; status != http.StatusOK {
		t.Errorf("%v", status)
		t.Errorf("%v", http.StatusOK)
	}

	expected := `[{"id":"1","email":"1email@naver.com"},{"id":"2","email":"2email@naver.com"},{"id":"3","email":"3email@naver.com"},{"id":"4","email":"4email@naver.com"}]`

	if strings.Replace(string(b), "\n","",-1) != expected {
		t.Errorf("res.Body %v", strings.Replace(string(b), "\n","",-1))
		t.Errorf("expected %v", expected)
	}

}

func Test_deleteUsers(t *testing.T) {
	req , _:= http.NewRequest("DELETE", "http://localhost:8089/users/4", nil)
	res := httptest.NewRecorder()

	m := mux.NewRouter()
	m.HandleFunc("/users/{Id}", deleteUsers)
	m.ServeHTTP(res, req)

	if status := res.Code ; status != http.StatusOK {
		t.Errorf("%v", status)
		t.Errorf("%v", http.StatusOK)
	}

	expected := `[{"id":"1","email":"1email@naver.com"},{"id":"2","email":"2email@naver.com"},{"id":"3","email":"3email@naver.com"}]`

	if strings.Replace(res.Body.String(), "\n","",-1)!= expected {
		t.Errorf("res.Body %v",strings.Replace(res.Body.String(), "\n","",-1))
		t.Errorf("expected %v", expected)
	}
}

func Test_putUsers(t *testing.T){

	res := httptest.NewRecorder()
	data :=`{"id":"changeId","email":"changeEmail@naver.com"}`
	var jsonStr = []byte(data)

	req, err := http.NewRequest("PUT", "http://localhost:8089/users/3", bytes.NewBuffer(jsonStr))
	if err != nil{
		t.Fatal(err)
	}

	m := mux.NewRouter()
	m.HandleFunc("/users/{Id}", putUsers)
	m.ServeHTTP(res, req)

	b, _ := ioutil.ReadAll(res.Body)

	if status := res.Code ; status != http.StatusOK {
		t.Errorf("%v", status)
		t.Errorf("%v", http.StatusOK)
	}

	expected :=`[{"id":"1","email":"1email@naver.com"},{"id":"2","email":"2email@naver.com"},{"id":"changeId","email":"changeEmail@naver.com"}]`

	if strings.Replace(string(b), "\n","",-1) != expected {
		t.Errorf("res.Body %v", strings.Replace(string(b), "\n","",-1))
		t.Errorf("expected %v", expected)
	}

}