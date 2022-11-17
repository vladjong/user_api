package handler

import (
	"net/http"
	"time"
)

func (h *handler) getTimeNow(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().String()))
}
