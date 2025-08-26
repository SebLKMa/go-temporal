package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	db "src/github.com/seblkma/go-temporal/surldb"
	dm "src/github.com/seblkma/go-temporal/surldm"
	ut "src/github.com/seblkma/go-temporal/surlut"
)

// DB connection used in this file
var surl_dbconn db.DbConnection
var surl_db *sql.DB

func init() {
	c, err := db.GetConnectionFromEnv()
	if err != nil {
		panic(err)
	}
	surl_dbconn = c
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implement yet", http.StatusNotImplemented)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}

func setSurl(w http.ResponseWriter, r *http.Request) {
	var src dm.SurlId
	err := json.NewDecoder(r.Body).Decode(&src)
	if err != nil {
		errmsg := err.Error()
		fmt.Println(errmsg)
		w.WriteHeader(http.StatusBadRequest)
		er := ErrorResponse{Code: http.StatusBadRequest, Message: "JSON input error: " + errmsg}
		json.NewEncoder(w).Encode(er)
		return
	}

	fmt.Printf("src: %#v\n", src)
	longUrl := src.LongUrl
	if surl_db == nil {
		errmsg := "DB error"
		fmt.Println(errmsg)
		w.WriteHeader(http.StatusInternalServerError)
		er := ErrorResponse{Code: http.StatusInternalServerError, Message: errmsg}
		json.NewEncoder(w).Encode(er)
		return
	}

	_, err = db.GetSurlIdByLongUrl(surl_db, longUrl)
	if err != nil {
		// not found, generate unique id for new record
		src.UniqueID = ut.GenId(longUrl)
		if src.UniqueID == "" {
			errmsg := "Error generating short id"
			fmt.Println(errmsg)
			w.WriteHeader(http.StatusInternalServerError)
			er := ErrorResponse{Code: http.StatusInternalServerError, Message: errmsg}
			json.NewEncoder(w).Encode(er)
			return
		}
		src.ShortUrl = "https://go/" + src.UniqueID
		src.ExpiresOn = time.Now().AddDate(2, 0, 0) // expires in 2 years
	}

	err = db.CreateOrUpdateSurlId(surl_db, src)
	if err != nil {
		errmsg := err.Error()
		fmt.Println(errmsg)
		w.WriteHeader(http.StatusInternalServerError)
		er := ErrorResponse{Code: http.StatusInternalServerError, Message: errmsg}
		json.NewEncoder(w).Encode(er)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getSurl(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("longurl")
	if key == "" {
		// note: if they pass in like ?param1=&param2= param1 will also be ""
		//fmt.Fprintf(w, "missing query param: sourcepath\n")
		w.WriteHeader(http.StatusBadRequest)
		er := ErrorResponse{Code: http.StatusBadRequest, Message: "missing query param: longurl"}
		json.NewEncoder(w).Encode(er)
		return
	}

	if surl_db == nil {
		errmsg := "DB error"
		fmt.Println(errmsg)
		w.WriteHeader(http.StatusInternalServerError)
		er := ErrorResponse{Code: http.StatusInternalServerError, Message: errmsg}
		json.NewEncoder(w).Encode(er)
		return
	}

	result, err := db.GetSurlIdByLongUrl(surl_db, key)
	if err != nil {
		errMsg := fmt.Sprintf("%s - %s\n", err.Error(), key)
		fmt.Println(errMsg)
		w.WriteHeader(http.StatusNotFound)
		er := ErrorResponse{Code: http.StatusNotFound, Message: errMsg}
		json.NewEncoder(w).Encode(er)
		return
	}

	fmt.Printf("%#v\n", result)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

const Port = "port"

//const EnvFlag string = "env"

func main() {

	surl_db = surl_dbconn.Connect()
	defer surl_db.Close()

	if surl_db == nil {
		fmt.Println("DB connection error")
		os.Exit(1)
	}

	surl_db.Ping()

	portFlag := flag.String(Port, "8181", "The HTTP Port. Default is 8181.")
	flag.Parse()
	port := *portFlag
	myRouter := mux.NewRouter().StrictSlash(true)

	// curl -X GET localhost:8282/ping
	myRouter.HandleFunc("/ping", ping).Methods("GET")

	// curl -X POST localhost:8282/setsurl -d '{"LongUrl":"https://www.ardanlabs.com/blog/2018/12/scheduling-in-go-part3.html?mc_cid=66407ad02b&mc_eid=41cad80de5", "Config":{"ThresholdPriority1":0.65, "ThresholdPriority2":0.78, "AcquisitionWindow":5, "WindowPercent":60, "UseSensorAI":false, "UseServerAI":false}}'
	// curl -X POST localhost:8282/setsurl -d '{"LongUrl":"https://github.com/ardanlabs/gotraining/blob/master/topics/go/design/composition/README.md"}'
	myRouter.HandleFunc("/setsurl", setSurl).Methods("POST")

	// curl -X GET localhost:8282/getsurl?longurl=https://github.com/ardanlabs/gotraining/blob/master/topics/go/design/composition/README.md
	myRouter.HandleFunc("/getsurl", getSurl).Methods("GET")

	host := "0.0.0.0:" + port
	fmt.Println(host + " up and listening")

	log.Fatal(http.ListenAndServe(host, myRouter))

}
