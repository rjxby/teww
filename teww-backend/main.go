package main

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rjxby/teww/teww-backend/db"
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

//Tag is single tag structure
type Tag struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
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

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// newUUID generates a random UUID according to RFC 4122
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

//Request is struct for retrive target data
type Request struct {
	Resource string
}

//Response is struct for returning result code
type Response struct {
	Result []Tag
}

func getAllTagsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//check tags alrady in db
	tagsExists, errTagsExists := db.Exists("tags")
	if errTagsExists != nil {
		log.Println("[ERROR] Can't check tags in db: ", errTagsExists)
		http.Error(w, errTagsExists.Error(), http.StatusInternalServerError)
		return
	}

	if !tagsExists {
		//set initial tags
		tags := []Tag{
			{"good", "Good"},
			{"like", "Like"},
			{"nice", "Nice"},
		}
		tagsBytes, errTagsBytes := getBytes(tags)
		if errTagsBytes != nil {
			log.Println("[ERROR] Can't set tags to db: ", errTagsBytes)
			http.Error(w, errTagsBytes.Error(), http.StatusInternalServerError)
			return
		}

		db.Set("tags", tagsBytes)
	}

	//get tags
	tags, errTags := db.Get("tags")
	if errTags != nil {
		log.Println("[ERROR] Can't get tags: ", errTags)
		http.Error(w, errTags.Error(), http.StatusInternalServerError)
		return
	}

	//decode tags
	var t []Tag
	by := bytes.Buffer{}
	by.Write(tags)
	d := gob.NewDecoder(&by)
	errDecodeBytes := d.Decode(&t)
	if errDecodeBytes != nil {
		log.Println("[ERROR] failed gob Decode: ", errDecodeBytes)
		http.Error(w, errDecodeBytes.Error(), http.StatusInternalServerError)
		return
	}

	//send tags
	j, errMarshal := json.Marshal(t)
	if errMarshal != nil {
		log.Println("[ERROR] failed marshal tags: ", errMarshal)
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

func initItems() {
	//set initial items
	item1 := Item{"1", "04/03/2017", "24/03/2017", "20", "Good", "Good"}
	item1Bytes, errItem1Bytes := getBytes(item1)
	if errItem1Bytes != nil {
		log.Println("[ERROR] Can't set item1 to db: ", errItem1Bytes)
		return
	}
	db.Set("item:1", item1Bytes)

	item2 := Item{"2", "04/03/2017", "24/03/2017", "20", "Like", "Like"}
	item2Bytes, errItem2Bytes := getBytes(item2)
	if errItem2Bytes != nil {
		log.Println("[ERROR] Can't set item2 to db: ", errItem2Bytes)
		return
	}
	db.Set("item:2", item2Bytes)

	item3 := Item{"3", "04/03/2017", "24/03/2017", "20", "Nice", "nice"}
	item3Bytes, errItem3Bytes := getBytes(item3)
	if errItem3Bytes != nil {
		log.Println("[ERROR] Can't set item3 to db: ", errItem3Bytes)
		return
	}
	db.Set("item:3", item3Bytes)
}

func getAllItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//get items
	items, errItems := db.GetByPattern("item:*")
	if errItems != nil {
		log.Println("[ERROR] Can't get items: ", errItems)
		http.Error(w, errItems.Error(), http.StatusInternalServerError)
		return
	}

	//decode items
	var res []Item
	for el := range items {
		var i Item
		by := bytes.Buffer{}
		by.Write(items[el])
		d := gob.NewDecoder(&by)
		errDecodeBytes := d.Decode(&i)
		if errDecodeBytes != nil {
			log.Println(`[ERROR] failed gob Decode`, errDecodeBytes)
			http.Error(w, errDecodeBytes.Error(), http.StatusInternalServerError)
			return
		}

		res = append(res, i)
	}

	//send items
	j, errMarshal := json.Marshal(res)
	if errMarshal != nil {
		log.Println(`[ERROR] failed marshal items`, errMarshal)
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

func postItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var i Item
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &i)

	id, uuIDErr := newUUID()
	if uuIDErr != nil {
		log.Println(`[ERROR] failed generate uuid`, uuIDErr)
		http.Error(w, uuIDErr.Error(), http.StatusInternalServerError)
		return
	}

	i.ID = id

	itemBytes, errItemBytes := getBytes(i)
	if errItemBytes != nil {
		log.Println("[ERROR] Can't set item to db: ", errItemBytes)
		return
	}

	//TODO validation
	//set item
	db.Set("item:"+i.ID, itemBytes)

	//send item
	j, errMarshal := json.Marshal(i)
	if errMarshal != nil {
		log.Println(`[ERROR] failed marshal item`, errMarshal)
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

func putItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var i Item
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &i)

	itemBytes, errItemBytes := getBytes(i)
	if errItemBytes != nil {
		log.Println("[ERROR] Can't set item to db: ", errItemBytes)
		return
	}

	//TODO validation
	//set item
	db.Set("item:"+i.ID, itemBytes)

	//send item
	j, errMarshal := json.Marshal(i)
	if errMarshal != nil {
		log.Println(`[ERROR] failed marshal item`, errMarshal)
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var i Item
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &i)

	//check items alrady in db
	errItemDelete := db.Delete("item:" + i.ID)
	if errItemDelete != nil {
		log.Println("[ERROR] Can't remove item "+i.ID+" form db: ", errItemDelete)
		http.Error(w, errItemDelete.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
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
	selfPortVariable = ":" + strconv.Itoa(configuration.ItemService.Port)

	db.InitConfig(configuration.DbService.Host, configuration.DbService.Port) //todo pass user-password to db connection
	initItems()                                                               //for sample

	r := mux.NewRouter()

	r.HandleFunc("/api/tags", getAllTagsHandler).Methods("GET")
	r.HandleFunc("/api/items", getAllItemsHandler).Methods("GET")
	r.HandleFunc("/api/items", postItemHandler).Methods("POST")
	r.HandleFunc("/api/items", putItemHandler).Methods("PUT")
	r.HandleFunc("/api/items", deleteItemHandler).Methods("DELETE")

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
