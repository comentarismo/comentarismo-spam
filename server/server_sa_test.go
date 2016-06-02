package server_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/drewolson/testflight"
	"testing"
	"log"
	"comentarismo-spam/server"
)

func TestSpamassassinSpamHandler(t *testing.T) {

	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Post("/spam", testflight.FORM_ENCODED, "text=Aposto que foi o câncer .! Doença muito conhecida entre políticos corruptos Sul americanos.");

		log.Println(response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t, len(response.Body) > 0)

	})
}

func TestSpamassassinSpamRevoke(t *testing.T) {

	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Post("/revoke", testflight.FORM_ENCODED, "text=Aposto que foi o câncer .! Doença muito conhecida entre políticos corruptos Sul americanos.");

		log.Println(response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t, len(response.Body) > 0)

	})
}

func TestSpamassassinSpamReport(t *testing.T) {

	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Post("/report", testflight.FORM_ENCODED, "text=Aposto que foi o câncer .! Doença muito conhecida entre políticos corruptos Sul americanos.");

		log.Println(response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t, len(response.Body) > 0)

	})
}
