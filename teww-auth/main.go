package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/rjxby/teww/teww-auth/auth"
)

var selfPortVariable string

//Configuration is config file representation
type Configuration struct {
	DbService     serviceConfiguration `json:"db_service"`
	ClientService serviceConfiguration `json:"client_service"`
	AuthService   serviceConfiguration `json:"auth_service"`
	ItemService   serviceConfiguration `json:"item_service"`
}

type serviceConfiguration struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

//authRequest is request structure for autorization
type authRequest struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

//authResponse is response structure for autorization
type authResponse struct {
	Token     string `json:"token,omitempty"`
	TokenType string `json:"tokenType,omitempty"`
	ExpiresIn int64  `json:"expiresIn,omitempty"`
}

func onAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var a authRequest
	b, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(b, &a)
	token, tokenType, expiresIn, err := auth.OnAuthentication(a.UserName, a.Password)

	if err != nil {
		log.Println(`[ERROR] failed authentication: `, err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var res authResponse
	res.Token = token
	res.TokenType = tokenType
	res.ExpiresIn = expiresIn

	w.WriteHeader(http.StatusOK)

	j, _ := json.Marshal(&res)
	w.Write(j)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	authorizationHeader := r.Header.Get("Authorization")
	s := strings.Split(authorizationHeader, " ")
	_, token := s[0], s[1]
	ok, _ := auth.LogOut(token)
	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func checkAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	authorizationHeader := r.Header.Get("Authorization")
	s := strings.Split(authorizationHeader, " ")
	_, token := s[0], s[1]
	ok, _ := auth.CheckAuthentication(token)
	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func main() {
	configFile, _ := os.Open("./config.json")
	configDecoder := json.NewDecoder(configFile)
	configuration := Configuration{}
	configErr := configDecoder.Decode(&configuration)
	if configErr != nil {
		fmt.Println("config error:", configErr)
		panic(configErr)
	}
	selfPortVariable = ":" + strconv.Itoa(configuration.AuthService.Port)

	envSecretVariable := os.Getenv("HMAC_SECRET")
	if len(envSecretVariable) == 0 {
		panic("environment variable not contains HMAC_SECRET")
	}
	envExpirationVariable := os.Getenv("EXPIRATION_TIME")
	if len(envExpirationVariable) == 0 {
		panic("environment variable not contains EXPIRATION_TIME")
	}

	expirationTime, parseErr := strconv.ParseInt(envExpirationVariable, 10, 64)
	if parseErr != nil {
		panic("can't parse EXPIRATION_TIME environment variable")
	}
	auth.InitConfig(envSecretVariable, expirationTime)

	r := mux.NewRouter()

	//r.HandleFunc("/auth", autorizationHandler).Methods("POST")
	r.HandleFunc("/auth/login", onAuthenticationHandler).Methods("POST")
	r.HandleFunc("/auth/logout", logoutHandler).Methods("POST")
	r.HandleFunc("/auth/check", checkAuthenticationHandler).Methods("GET")

	sirMuxalot := http.NewServeMux()
	sirMuxalot.Handle("/", r)

	n := negroni.Classic()
	n.UseHandler(sirMuxalot)

	err := http.ListenAndServe(selfPortVariable, n)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
		panic(err)
	}
}
