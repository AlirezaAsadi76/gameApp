package main

import (
	"encoding/json"
	"fmt"
	"gameApp/repository"
	"gameApp/service/userservice"
	"io"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/user/register", registerRequestHandler)
	server := &http.Server{Addr: ":7660", Handler: mux}
	fmt.Println("Listening on port 7660")
	server.ListenAndServe()
}

func registerRequestHandler(writer http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodPost {
		fmt.Println("request method :", request.Method)
		writer.Write([]byte(`request method must be "POST"`))
	}
	var registerRequest userservice.RegisterRequest
	reqBody, iErr := io.ReadAll(request.Body)
	if iErr != nil {
		fmt.Println("read body error", iErr)
		writer.Write([]byte(fmt.Sprintf(`read body error : %s`, iErr.Error())))
		return
	}

	if err := json.Unmarshal(reqBody, &registerRequest); err != nil {
		fmt.Println("json unmarshal error", err)
		writer.Write([]byte(fmt.Sprintf(`json unmarshal error : %s`, err.Error())))
		return
	}

	repo := repository.NewDB()
	userServ := userservice.New(repo)

	user, err := userServ.Register(registerRequest)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`register error : %s`, err.Error())))
		return
	}
	json.NewEncoder(writer).Encode(user)
}
