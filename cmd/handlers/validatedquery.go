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
	wallet "github.com/hietkamp/norma-gw/internal/nuts-wallet"
)

func HandleValidatedQuery(c *gin.Context) {
	if claim, exists := c.Get("serviceClaim"); exists {
		if claim != "validated-query-service" {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
	envelop := HttpHeaderEnvelop{}
	if err := c.ShouldBindHeader(&envelop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var payload ValidatedQueryCredential
	c.MustBindWith(&payload, binding.JSON)

	nutsURL := os.Getenv("nuts_url")
	verifier := wallet.New(wallet.Options{BaseURL: nutsURL})
	vcBytes, _ := json.Marshal(payload)
	// Verify if the credential is valid
	if err := verifier.VerifyVerifiableCredential(vcBytes); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return		
	}
	ts := time.Now()
	message := ValidatedQueriesReceived{
		Timestamp: ts.Format(time.RFC3339),
		ClientIP:  c.ClientIP(),
		Header:    envelop,
		Payload:   payload,
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	brokers := strings.Split(os.Getenv("kafka_url"), ",")
	es := eventstream.New("tcp", brokers[0])
	err = es.Produce(os.Getenv("topic_validatedqueries"), messageBytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
