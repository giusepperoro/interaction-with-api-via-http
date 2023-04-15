package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func (h *addToDatabaseHandler) HandleAddToDatabase(c *gin.Context) {
	var dataRequest DataRequest

	r, _ := http.Get(h.cfg.TakeDataUrl)
	body, _ := io.ReadAll(r.Body)

	json.Unmarshal(body, &dataRequest)
	log.Println(dataRequest)
}
