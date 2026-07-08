package main

import (
	"encoding/json"
	"fmt"
	"gameApp/repository"
	"gameApp/service/userservice"
	"io"
	"net/http"
)

const (
	jwtSecret = "secret"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/users/register", registerRequestHandler)
	mux.HandleFunc("/users/login", LoginRequestHandler)
	mux.HandleFunc("/users/profile", ProfileRequestHandler)
	server := &http.Server{Addr: ":7660", Handler: mux}
	fmt.Println("Listening on port 7660")
	server.ListenAndServe()
}

func ProfileRequestHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		fmt.Println("request method :", request.Method)
		writer.Write([]byte(`request method must be "GET"`))
		return
	}

	var profileRequest userservice.ProfileRequest
	reqBody, iErr := io.ReadAll(request.Body)
	if iErr != nil {
		fmt.Println("read body error", iErr)
		writer.Write([]byte(fmt.Sprintf(`read body error : %s`, iErr.Error())))
		return
	}
	if err := json.Unmarshal(reqBody, &profileRequest); err != nil {
		fmt.Println("json unmarshal error", err)
		writer.Write([]byte(fmt.Sprintf(`json unmarshal error : %s`, err.Error())))
		return
	}

	repo := repository.NewDB()
	userServ := userservice.New(repo, jwtSecret)

	profileResponse, err := userServ.Profile(profileRequest)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`Profile error : %s`, err.Error())))
		return
	}

	data, _ := json.Marshal(profileResponse)
	writer.Write(data)

}

func LoginRequestHandler(writer http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodPost {
		fmt.Println("request method :", request.Method)
		writer.Write([]byte(`request method must be "POST"`))
		return
	}

	var loginRequest userservice.LoginRequest
	reqBody, iErr := io.ReadAll(request.Body)
	if iErr != nil {
		fmt.Println("read body error", iErr)
		writer.Write([]byte(fmt.Sprintf(`read body error : %s`, iErr.Error())))
		return
	}

	if err := json.Unmarshal(reqBody, &loginRequest); err != nil {
		fmt.Println("json unmarshal error", err)
		writer.Write([]byte(fmt.Sprintf(`json unmarshal error : %s`, err.Error())))
		return
	}

	repo := repository.NewDB()
	userServ := userservice.New(repo, jwtSecret)

	user, err := userServ.Login(loginRequest)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`Login error : %s`, err.Error())))
		return
	}
	json.NewEncoder(writer).Encode(user)
}

func registerRequestHandler(writer http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodPost {
		fmt.Println("request method :", request.Method)
		writer.Write([]byte(`request method must be "POST"`))
		return
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
	userServ := userservice.New(repo, jwtSecret)

	user, err := userServ.Register(registerRequest)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`register error : %s`, err.Error())))
		return
	}
	json.NewEncoder(writer).Encode(user)
}
