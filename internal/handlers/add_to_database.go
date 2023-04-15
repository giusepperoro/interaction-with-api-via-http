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

	//c.Get(h.cfg.TakeDataUrl)
	//http.Serve(autocert.NewListener(h.cfg.TakeDataUrl), nil)
	r, _ := http.Get(h.cfg.TakeDataUrl)
	body, _ := io.ReadAll(r.Body)

	json.Unmarshal(body, &dataRequest.results)
	log.Println(dataRequest.results)
}
