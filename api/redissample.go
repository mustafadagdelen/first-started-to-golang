package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", addUsers).Methods("POST")
	fmt.Println("Server started..!")
	http.ListenAndServe(":8091", router)
}

func getUsers(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Content-Encoding", "utf-8")
	conn, err := redis.Dial("tcp", "localhost:32770")

	HandleError(err)

	var result = conn.Cmd("Get", "user.1")

	jsonBytes, error := result.Bytes()

	HandleError(error)

	var user User

	error = json.Unmarshal(jsonBytes, &user)

	HandleError(error)

	outgoingJson, err := json.Marshal(user)

	HandleError(err)

	fmt.Fprint(res, string(outgoingJson))
}

func addUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	conn, err := redis.Dial("tcp", "localhost:32770")
	HandleError(err)

	user1 := new(User)
	user1.ID = 1
	user1.FirstName = "Mustafa"
	user1.LastName = "Dağdelen"
	user1.UserName = "mustafad"
	user1.Password = "dmustafa"

	user1Json, err := json.Marshal(user1)
	HandleError(err)
	conn.Cmd("Set", "user.1", user1Json)

	user2 := new(User)
	user2.ID = 2
	user2.FirstName = "Oktay"
	user2.LastName = "Dandini"
	user2.UserName = "oktayd"
	user2.Password = "doktay"

	user2Json, err := json.Marshal(user2)
	HandleError(err)
	conn.Cmd("Set", "user.2", user2Json)

	fmt.Println("Status:true")
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
}
