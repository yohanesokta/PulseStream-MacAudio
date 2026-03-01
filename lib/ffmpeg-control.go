package lib

import (
	"PulseStream-MacAudio/utils"
	"log"
	"net/http"
	"os/exec"
)

type FFMPEG_version struct {
}

func (FFMPEG_version) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	cmd := exec.Command("ffmpeg", "-version")

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error:", err)
		log.Println("output:", string(output))

		utils.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "ffmpeg execution failed",
		})
		return
	}

	log.Println(string(output))
}
