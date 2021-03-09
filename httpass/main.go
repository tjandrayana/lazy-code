package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/foomo/htpasswd"
	"github.com/julienschmidt/httprouter"
	"github.com/tjandrayana/nsq-tail/jsonapi"

	"github.com/urfave/negroni"
	grace "gopkg.in/paytm/grace.v1"
)

var fileLocation string = "file/demo.htpasswd"

func main() {
	client := GetRoute()
	InitRoute(client)

	n := negroni.New()
	n.UseHandler(client)
	log.Fatal(grace.Serve(fmt.Sprintf(":%s", "8005"), n))
}

type Param struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	FileLocation string `json:"file_location,omitempty"`
}

func GetRoute() *httprouter.Router {
	return httprouter.New()

}
func InitRoute(r *httprouter.Router) {
	r.POST("/create-user", DoCreateUser)
	r.POST("/remove-user", DoRemoveUser)
	r.POST("/get-user", DoGetUser)

}

func DoCreateUser(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	var reqData Param
	api := jsonapi.New(res, "")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		api.ErrorWriter(http.StatusBadRequest, err.Error())
		return
	}

	if err := json.Unmarshal(body, &reqData); err != nil {
		log.Println(err)
		api.ErrorWriter(http.StatusBadRequest, err.Error())
		return
	}

	ms := reqData.SetUser(reqData)
	res.Write(ms)
}

func DoRemoveUser(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	var reqData Param
	api := jsonapi.New(res, "")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		api.ErrorWriter(http.StatusBadRequest, err.Error())
		return
	}

	if err := json.Unmarshal(body, &reqData); err != nil {
		log.Println(err)
		api.ErrorWriter(http.StatusBadRequest, err.Error())
		return
	}

	ms := reqData.RemoveUser(reqData)
	res.Write(ms)
}

func DoGetUser(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	var reqData Param
	api := jsonapi.New(res, "")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		api.ErrorWriter(http.StatusBadRequest, err.Error())
		return
	}

	if err := json.Unmarshal(body, &reqData); err != nil {
		log.Println(err)
		api.ErrorWriter(http.StatusBadRequest, err.Error())
		return
	}

	ms := reqData.GetUser(reqData)
	res.Write(ms)
}

func (m *Param) SetUser(param Param) []byte {

	err := htpasswd.SetPassword(param.FileLocation, param.Username, param.Password, htpasswd.HashBCrypt)
	if err != nil {
		return []byte(fmt.Sprintf("Error : %v", err))
	}

	return []byte(fmt.Sprintf("Create User %s is success", param.Username))
}

func (m *Param) RemoveUser(param Param) []byte {

	err := htpasswd.RemoveUser(param.FileLocation, param.Username)
	if err != nil {
		return []byte(fmt.Sprintf("Error : %v", err))
	}

	return []byte(fmt.Sprintf("Remove User %s is success", param.Username))
}

func (m *Param) GetUser(param Param) []byte {

	pass, err := htpasswd.ParseHtpasswdFile(param.FileLocation)
	if err != nil {
		return []byte(fmt.Sprintf("Error : %v", err))
	}

	fmt.Println("pass : ", pass)
	var listUser []Param

	for k, v := range pass {
		listUser = append(listUser, Param{
			Username: k,
			Password: v,
		})
	}

	listJson, err := json.Marshal(listUser)
	if err != nil {
		return []byte(fmt.Sprintf("Error : %v", err))
	}

	return listJson
}
