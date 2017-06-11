package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var (
	selfPortVariable        string
	authServiceHostVariable string
	itemServiceHostVariable string
)

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

//ContextPage is struct for passing metadata to html tamplate
type ContextPage struct {
	Title  string
	Static string
}

//StaticURL is path to static files folder
const StaticURL string = "/static/"

//StaticRoot is path to static folder
const StaticRoot string = "static/"

func indexPage(w http.ResponseWriter, req *http.Request) {
	context := ContextPage{Static: StaticURL}
	tmplList := []string{"index.html"}
	t, err := template.ParseFiles(tmplList...)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, context)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func staticHandler(w http.ResponseWriter, req *http.Request) {
	staticFile := req.URL.Path[len(StaticURL):]
	if len(staticFile) != 0 {
		f, err := http.Dir(StaticRoot).Open(staticFile)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, staticFile, time.Now(), content)
			return
		}
	}
	http.NotFound(w, req)
}

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//Tag is single tag structure
type Tag struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func getAllTagsHandler(w http.ResponseWriter, r *http.Request) {
	var url = itemServiceHostVariable + "api/tags"
	req, errReq := http.NewRequest("GET", url, nil)
	if errReq != nil {
		log.Println("[ERROR] Can't create api/tags request: ", errReq)
		http.Error(w, errReq.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	client := &http.Client{Timeout: 10 * time.Second}
	res, errRes := client.Do(req)
	if errRes != nil {
		log.Println("[ERROR] Can't recive response from api/tags srvice: ", errRes)
		http.Error(w, errRes.Error(), http.StatusUnauthorized)
		return
	}
	defer res.Body.Close()

	var tagsResponse []Tag
	json.NewDecoder(res.Body).Decode(&tagsResponse)

	j, errMarshal := json.Marshal(tagsResponse)
	if errMarshal != nil {
		log.Println(`[ERROR] failed marshal api/tags response: `, errMarshal)
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

//Item is single item
type Item struct {
	ID          string `json:"id,omitempty"`
	DateStart   string `json:"dateStart,omitempty"`
	DateEnd     string `json:"dateEnd"`
	Length      string `json:"length"`
	Description string `json:"description"`
	TagID       string `json:"tagId"`
}

func getAllItemsHandler(w http.ResponseWriter, r *http.Request) {
	var url = itemServiceHostVariable + "api/items"
	req, errReq := http.NewRequest("GET", url, nil)
	if errReq != nil {
		log.Println("[ERROR] Can't create api/items request: ", errReq)
		http.Error(w, errReq.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	client := &http.Client{Timeout: 10 * time.Second}
	res, errRes := client.Do(req)
	if errRes != nil {
		log.Println("[ERROR] Can't recive response from api/items srvice: ", errRes)
		http.Error(w, errRes.Error(), http.StatusUnauthorized)
		return
	}
	defer res.Body.Close()

	var itemResponse []Item
	json.NewDecoder(res.Body).Decode(&itemResponse)

	//todo maybe check response
	j, errMarshal := json.Marshal(itemResponse)
	if errMarshal != nil {
		log.Println(`[ERROR] failed marshal api/items response: `, errMarshal)
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

func postItemHandler(w http.ResponseWriter, r *http.Request) {
	var url = itemServiceHostVariable + "api/items"
	//todo maybe check body for errors
	req, errReq := http.NewRequest("POST", url, r.Body)
	if errReq != nil {
		log.Println("[ERROR] Can't create api/items request: ", errReq)
		http.Error(w, errReq.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	client := &http.Client{Timeout: 10 * time.Second}
	res, errRes := client.Do(req)
	if errRes != nil {
		log.Println("[ERROR] Can't recive response from api/items srvice: ", errRes)
		http.Error(w, errRes.Error(), http.StatusUnauthorized)
		return
	}
	defer res.Body.Close()

	var itemResponse Item
	json.NewDecoder(res.Body).Decode(&itemResponse)

	//todo maybe check response
	j, errMarshal := json.Marshal(itemResponse)
	if errMarshal != nil {
		log.Println(`[ERROR] failed marshal api/items response: `, errMarshal)
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

func putItemHandler(w http.ResponseWriter, r *http.Request) {
	var url = itemServiceHostVariable + "api/items"
	//todo maybe check body for errors
	req, errReq := http.NewRequest("PUT", url, r.Body)
	if errReq != nil {
		log.Println("[ERROR] Can't create api/items request: ", errReq)
		http.Error(w, errReq.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	client := &http.Client{Timeout: 10 * time.Second}
	res, errRes := client.Do(req)
	if errRes != nil {
		log.Println("[ERROR] Can't recive response from api/items srvice: ", errRes)
		http.Error(w, errRes.Error(), http.StatusUnauthorized)
		return
	}
	defer res.Body.Close()

	var itemResponse Item
	json.NewDecoder(res.Body).Decode(&itemResponse)

	//todo maybe check response
	j, errMarshal := json.Marshal(itemResponse)
	if errMarshal != nil {
		log.Println(`[ERROR] failed marshal api/items response: `, errMarshal)
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	var url = itemServiceHostVariable + "api/items"
	//todo maybe check body for errors
	req, errReq := http.NewRequest("DELETE", url, r.Body)
	if errReq != nil {
		log.Println("[ERROR] Can't create api/items request: ", errReq)
		http.Error(w, errReq.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	client := &http.Client{Timeout: 10 * time.Second}
	res, errRes := client.Do(req)
	if errRes != nil {
		log.Println("[ERROR] Can't recive response from api/items srvice: ", errRes)
		http.Error(w, errRes.Error(), http.StatusUnauthorized)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var errMes = "[ERROR] Remove item return error"
		log.Println(errMes)
		http.Error(w, errMes, http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

//authResponse is response structure for autorization
type authResponse struct {
	Token     string `json:"token,omitempty"`
	TokenType string `json:"tokenType,omitempty"`
	ExpiresIn int64  `json:"expiresIn,omitempty"`
}

func autorizationHandler(w http.ResponseWriter, r *http.Request) {
	var url = authServiceHostVariable + "auth"
	//todo maybe check body for errors
	req, errReq := http.NewRequest("POST", url, r.Body)
	if errReq != nil {
		log.Println("[ERROR] Can't create auth request: ", errReq)
		http.Error(w, errReq.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	client := &http.Client{Timeout: 10 * time.Second}
	res, errRes := client.Do(req)
	if errRes != nil {
		log.Println("[ERROR] Can't recive response from auth srvice: ", errRes)
		http.Error(w, errRes.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	var authResponse authResponse
	json.NewDecoder(res.Body).Decode(&authResponse)

	//todo maybe check response
	j, errMarshal := json.Marshal(authResponse)
	if errMarshal != nil {
		log.Println(`[ERROR] failed marshal auth response: `, errMarshal)
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

func onAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	var url = authServiceHostVariable + "auth/login"
	//todo maybe check body for errors
	req, errReq := http.NewRequest("POST", url, r.Body)
	if errReq != nil {
		log.Println("[ERROR] Can't create auth/login request: ", errReq)
		http.Error(w, errReq.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	client := &http.Client{Timeout: 10 * time.Second}
	res, errRes := client.Do(req)
	if errRes != nil {
		log.Println("[ERROR] Can't recive response from auth/login srvice: ", errRes)
		http.Error(w, errRes.Error(), http.StatusUnauthorized)
		return
	}
	defer res.Body.Close()

	var authResponse authResponse
	json.NewDecoder(res.Body).Decode(&authResponse)

	//todo maybe check response
	j, errMarshal := json.Marshal(authResponse)
	if errMarshal != nil {
		log.Println(`[ERROR] failed marshal auth/login response: `, errMarshal)
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	var url = authServiceHostVariable + "auth/logout"
	//todo maybe check body for errors
	req, errReq := http.NewRequest("POST", url, r.Body)
	if errReq != nil {
		log.Println("[ERROR] Can't create auth request: ", errReq)
		http.Error(w, errReq.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", r.Header.Get("Authorization")) //set token to header

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	client := &http.Client{Timeout: 10 * time.Second}
	res, errRes := client.Do(req)
	if errRes != nil {
		log.Println("[ERROR] Can't recive response from auth srvice: ", errRes)
		http.Error(w, errRes.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	var authResponse authResponse
	json.NewDecoder(res.Body).Decode(&authResponse)

	//todo maybe check response
	j, errMarshal := json.Marshal(authResponse)
	if errMarshal != nil {
		log.Println(`[ERROR] failed marshal auth response: `, errMarshal)
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

//APIMiddleware is check auth
func APIMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var url = authServiceHostVariable + "auth/check"
	req, errReq := http.NewRequest("GET", url, nil)
	if errReq != nil {
		log.Println("[ERROR] Can't create auth/check request: ", errReq)
		http.Error(w, errReq.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", r.Header.Get("Authorization")) //set token to header

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	client := &http.Client{Timeout: 10 * time.Second}
	res, errRes := client.Do(req)
	if errRes != nil {
		log.Println("[ERROR] Can't recive response from auth/check srvice: ", errRes)
		http.Error(w, errRes.Error(), http.StatusUnauthorized)
		return
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		log.Println("[STATUS] User unauthorized")
		http.Error(w, "User unauthorized", http.StatusUnauthorized)
		return
	}

	next(w, r)
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
	selfPortVariable = ":" + strconv.Itoa(configuration.ClientService.Port)

	r := mux.NewRouter()
	r.HandleFunc("/", indexPage)

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/", autorizationHandler).Methods("POST")
	auth.HandleFunc("/login", onAuthenticationHandler).Methods("POST")
	auth.HandleFunc("/logout", logoutHandler).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/tags", getAllTagsHandler).Methods("GET")
	api.HandleFunc("/items", getAllItemsHandler).Methods("GET")
	api.HandleFunc("/items", postItemHandler).Methods("POST")
	api.HandleFunc("/items", putItemHandler).Methods("PUT")
	api.HandleFunc("/items", deleteItemHandler).Methods("DELETE")

	sirMuxalot := http.NewServeMux()
	sirMuxalot.HandleFunc(StaticURL, staticHandler)
	sirMuxalot.Handle("/", r)
	sirMuxalot.Handle("/api/", negroni.New(
		negroni.HandlerFunc(APIMiddleware),
		negroni.Wrap(r),
	))

	n := negroni.Classic()
	n.UseHandler(sirMuxalot)

	authServiceHostVariable = fmt.Sprintf("http://%s:%s/", configuration.AuthService.Host, strconv.Itoa(configuration.AuthService.Port))
	itemServiceHostVariable = fmt.Sprintf("http://%s:%s/", configuration.ItemService.Host, strconv.Itoa(configuration.ItemService.Port))
	err := http.ListenAndServe(selfPortVariable, n)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
		panic(err)
	}
}
