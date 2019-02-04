package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../pkg/dbConnectModule"
	"../pkg/structModule"
	"../pkg/validationModule"
	"github.com/julienschmidt/httprouter"
)

/* type Resource interface {
	Uri() string
	Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) structModule.Response
	Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) structModule.Response
	Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) structModule.Response
	Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) structModule.Response
} */

/* func getAllBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(ps.ByName("test"))
	//mem := Member{"Alex", 10, true}
	var err error

	dbConnectModule.Sqlite_Open()
	if err != nil {
		fmt.Println(err)
	}

	mem, err := dbConnectModule.GetAllBookInfo()
	fmt.Println(mem)

	temp, err := json.Marshal(mem)
	if err != nil {
		fmt.Println(err)
	}
	//	fmt.println(strings(mem))
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(temp)

}
func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(ps.ByName("test"))
	mem := structModule.Member{"Alex", 10, true}
	test, err := ioutil.ReadAll(r.Body)
	var testmem structModule.Member
	err = json.Unmarshal(test, &testmem)

	fmt.Println(testmem)
	temp, err := json.Marshal(mem)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println()
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(temp)

} */

func getMyBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//fmt.Println(ps.ByName("test"))
	//mem := Member{"Alex", 10, true}

	test, err := ioutil.ReadAll(r.Body)

	//	var testmem Member
	var testAuth structModule.ValAPIKey

	err = json.Unmarshal(test, &testAuth)

	var respondUser structModule.Response

	var tempBookArray []structModule.EBookInfo

	fmt.Println(testAuth)

	if validationModule.CheckAPIToken(testAuth) {
		tempBookArray, err = dbConnectModule.GetMyOwnBook(testAuth)
		respondUser = structModule.Response{200, "Respond Success", tempBookArray}
	} else {

		respondUser = structModule.Response{403, "Invalid Auth key or Expire key", tempBookArray}
	}

	temp, err := json.Marshal(respondUser)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(temp)

}

func main() {
	router := httprouter.New()
	/* 	router.GET("/", Index)
	   	router.GET("/hello/:name", Hello) */
	//router.POST("/getAllBook", getAllBook)
	//router.POST(GetBookInfo.Uri(), getMyBook)
	//router.POST(GetBookInfo.URI(), getMyBook)

	router.POST("/service/book", getMyBook)

	//router.POST("/getUserInfo", getUser)
	//router.POST("/getUserInfo/:test", getUser)

	log.Fatal(http.ListenAndServe(":8090", router))
}
