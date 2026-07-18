package handlers

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type LoginHandler struct{
	logger *log.Logger
}

func NewLoginHandler(logger *log.Logger) http.Handler {
	return &LoginHandler{
		logger: logger,
	}
}

func (mh *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mh.handleGet(w, r)
	case http.MethodPut:
		mh.handlePut(w, r)
	case http.MethodPost:
		mh.handlePost(w, r)
	case http.MethodDelete:
		mh.handleDelete(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (mh *LoginHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	mh.logger.Println("GET NOT IMPLEMENTED")
}

func (mh *LoginHandler) handlePut(w http.ResponseWriter, r *http.Request) {
	mh.logger.Println("PUT NOT IMPLEMENTED")
}

func (mh *LoginHandler) handlePost(w http.ResponseWriter, r *http.Request) {

	mh.logger.WithFields(log.Fields{
		"time":       time.Now().String(),
		"username":   r.FormValue("_user"),
		"password":   r.FormValue("_pass"),
		"user-agent": r.UserAgent(),
		"ip_address": r.RemoteAddr,
	}).Info("login attempt")

	http.Redirect(w, r, "/", http.StatusFound)
}

func (mh *LoginHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	mh.logger.Println("DELETE NOT IMPLEMENTED")
}
