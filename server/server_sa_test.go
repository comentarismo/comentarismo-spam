package server_test

import (
	"comentarismo-spam/server"
	"github.com/drewolson/testflight"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestSpamassassinSpamHandler(t *testing.T) {
	return
	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Post("/sa_spam", testflight.FORM_ENCODED, "text=Aposto que foi o câncer .! Doença muito conhecida entre políticos corruptos Sul americanos.")

		log.Println(response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t, len(response.Body) > 0)

	})
}

func TestSpamassassinSpamRevoke(t *testing.T) {
	return
	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Post("/sa_revoke", testflight.FORM_ENCODED, "text=Aposto que foi o câncer .! Doença muito conhecida entre políticos corruptos Sul americanos.")

		log.Println(response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t, len(response.Body) > 0)

	})
}

func TestSpamassassinSpamReport(t *testing.T) {
	return
	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Post("/sa_report", testflight.FORM_ENCODED, "text=Aposto que foi o câncer .! Doença muito conhecida entre políticos corruptos Sul americanos.")

		log.Println(response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t, len(response.Body) > 0)

	})
}
