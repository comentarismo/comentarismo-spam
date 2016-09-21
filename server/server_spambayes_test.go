package server_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/drewolson/testflight"
	"testing"
	"log"
	"comentarismo-spam/server"
	"comentarismo-spam/spamc"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpamBayesHandler(t *testing.T) {

	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {

		Convey("Should Learn spammy words in english and report not spam for normal commentary", t, func() {
			targetWord := "Fantastic deal"
			response := r.Post("/report?lang=en", testflight.FORM_ENCODED, "text=" + targetWord);
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "Get paid Now"
			response = r.Post("/report?lang=en", testflight.FORM_ENCODED, "text=" + targetWord);
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "Cancel at any time with"
			response = r.Post("/report?lang=en", testflight.FORM_ENCODED, "text=" + targetWord);
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "Easy terms its Full refund"
			response = r.Post("/report?lang=en", testflight.FORM_ENCODED, "text=" + targetWord);
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			//now try with a good comment
			textTarget := "Amazing. great tits, shitty music"
			response = r.Post("/spam?lang=en", testflight.FORM_ENCODED, "text="+textTarget );

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)

			log.Println(response.Body) //{"code":200,"error":"","spam":false}

			spamReport := spamc.SpamReport{}
			err := json.Unmarshal(response.RawBody, &spamReport)
			So(err, ShouldBeNil)

			So(spamReport.Error, ShouldBeBlank)
			So(spamReport.Code, ShouldEqual, 200)
			So(spamReport.IsSpam, ShouldBeFalse)


			//now try with a spammy comment
			textTarget = "Easy terms its Full refund"
			response = r.Post("/spam?lang=en", testflight.FORM_ENCODED, "text="+textTarget );

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)

			log.Println(response.Body) //{"code":200,"error":"","spam":true}

			spamReport = spamc.SpamReport{}
			err = json.Unmarshal(response.RawBody, &spamReport)
			So(err, ShouldBeNil)

			So(spamReport.Error, ShouldBeBlank)
			So(spamReport.Code, ShouldEqual, 200)
			So(spamReport.IsSpam, ShouldBeTrue)


			//now revoke the spammy comment
			textTarget = "Easy terms its Full refund"
			response = r.Post("/revoke?lang=en", testflight.FORM_ENCODED, "text="+textTarget );

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)
			log.Println(response.Body) //{"code":200,"error":"","spam":true}

			//now it should not be spam anymore
			textTarget = "Easy terms its Full refund"
			response = r.Post("/spam?lang=en", testflight.FORM_ENCODED, "text="+textTarget );

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)

			log.Println(response.Body) //{"code":200,"error":"","spam":true}

			spamReport = spamc.SpamReport{}
			err = json.Unmarshal(response.RawBody, &spamReport)
			So(err, ShouldBeNil)

			So(spamReport.Error, ShouldBeBlank)
			So(spamReport.Code, ShouldEqual, 200)
			So(spamReport.IsSpam, ShouldBeFalse)


		})

	})
}

func init() {
	spamc.Flush()
}
