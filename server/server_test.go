package server_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/drewolson/testflight"
	"testing"
	"log"
	"comentarismo-spam/server"
)

//$ SHIELD_DEBUG=true go test  handler_sentiment_en_test.go -v
func TestSpamHandler(t *testing.T) {

	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Post("/spam", testflight.FORM_ENCODED, "text=Aposto que foi o câncer .! Doença muito conhecida entre políticos corruptos Sul americanos.");

		log.Println(response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t, len(response.Body) > 0)

	})
}