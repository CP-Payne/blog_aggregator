package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
)

type Util struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func NewUtil(errorLog, infoLog *log.Logger) *Util {
	return &Util{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}
}

var ErrNoAuthHeaderIncluded = errors.New("no auth header included")

func (u *Util) ServerError(w http.ResponseWriter, err error, msg string) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// In addition to informing th user about the internal server error
	// We also want to know what the internal server error is in order to fix it.
	// Therefore, we also print it to stdOut
	u.ErrorLog.Output(2, trace)

	// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	u.RespondWithError(w, 500, msg)
}

func (u *Util) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (u *Util) NotFound(w http.ResponseWriter) {
	u.ClientError(w, http.StatusNotFound)
}

func (u *Util) RespondWithError(w http.ResponseWriter, code int, message string) {
	// if code > 499 {
	// 	log.Printf("Responing with 5xx error: %s", message)
	// }
	type errorResponse struct {
		Error string `json:"error"`
	}
	u.RespondWithJSON(w, code, errorResponse{
		Error: message,
	})
}

func (u *Util) RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}

func (u *Util) GetApiKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}
