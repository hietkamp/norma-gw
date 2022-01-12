package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hietkamp/norma-gw/internal/eventstream"
)

func HandleRequestReceived(c *gin.Context) {
	envelop := HttpHeaderEnvelop{}
	if err := c.ShouldBindHeader(&envelop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var payload map[string]interface{}
	c.MustBindWith(&payload, binding.JSON)

	ts := time.Now()
	message := RequestsReceived{
		RequestId:      envelop.MessageId,
		Timestamp:      ts.Format(time.RFC3339),
		Request:        envelop.Subject,
		ClientId:       envelop.From,
		ClientIP:       c.ClientIP(),
		ControllerId:   envelop.To,
		ControllerHost: os.Getenv("http_server"),
		ConversationId: envelop.References,
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	brokers := strings.Split(os.Getenv("kafka_url"), ",")
	es := eventstream.New("tcp", brokers[1])
	err = es.Produce(os.Getenv("topic_requests"), messageBytes)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
