package utils

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIdFromUrl(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return 0, err
	}

	return id, nil
}
