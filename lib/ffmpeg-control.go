package lib

import (
	"PulseStream-MacAudio/utils"
	"net/http"
)

type CountHandler struct {
}

func (CountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "User endpoint",
	})
}
