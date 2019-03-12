package handlers

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	io.WriteString(w, "Create User Handle ceshi")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	usernama := p.ByName("username")

	io.WriteString(w, usernama)
}