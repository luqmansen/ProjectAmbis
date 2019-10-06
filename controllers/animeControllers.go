package controllers

import (
	"github.com/gorilla/mux"
	"github.com/luqmansen/hanako/models"
	u "github.com/luqmansen/hanako/utils"
	"net/http"
	"strconv"
)

var GetAll = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetAll()
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetById = func(w http.ResponseWriter, r * http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
	}

	data:= models.GetByID(uint(id))
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp)

}