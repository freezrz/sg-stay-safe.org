package check_in_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"sg-stay-safe.org/config"
	"sg-stay-safe.org/protocol"
	utils "sg-stay-safe.org/test_harness"
	"testing"
	"time"
)

func Test_checkin_sanitiser_anonymous_id_is_empty_only(t *testing.T) {
	body, _ := json.Marshal(protocol.CheckInEvent{
		AnonymousId: "",
		SiteId:      "54340nornge9yt889",
		Timestamp:   0,
	})
	req := utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	var actualResp protocol.CheckInResponse
	req.Send(&actualResp)
	assert.Equal(t, config.CodeSanitiseError, actualResp.Code)
}

func Test_checkin_sanitiser_site_id_is_empty_only(t *testing.T) {
	body, _ := json.Marshal(protocol.CheckInEvent{
		AnonymousId: "54340nornge9yt889",
		SiteId:      "",
		Timestamp:   0,
	})
	req := utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	var actualResp protocol.CheckInResponse
	req.Send(&actualResp)
	assert.Equal(t, 7459759354958, actualResp.Code)
}

func Test_checkin_sanitiser_corrupted_request(t *testing.T) {
	req := utils.New()
	req.WithBody(string("hoiehgohgoi")).WithMethod(http.MethodPost).Build()

	var actualResp protocol.CheckInResponse
	req.Send(&actualResp)
	assert.Equal(t, config.CodeMarshalError, actualResp.Code)
}

func Test_checkin_sanitiser_site_id_and_anonymous_id_both_empty(t *testing.T) {
	body, _ := json.Marshal(protocol.CheckInEvent{
		AnonymousId: "54340nornge9yt889",
		SiteId:      "",
		Timestamp:   0,
	})
	req := utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	var actualResp protocol.CheckInResponse
	req.Send(&actualResp)
	assert.Equal(t, config.CodeSanitiseError, actualResp.Code)
}

func Test_checkin_normal_fresh_user_random_site_first_time_login(t *testing.T) {
	body, _ := json.Marshal(protocol.CheckInEvent{
		AnonymousId: fmt.Sprintf("%s-%s", "api-test-user-id-", utils.RandStringRunes()),
		SiteId:      fmt.Sprintf("%s-%s", "api-test-site-id-", utils.RandStringRunes()),
		Timestamp:   0,
	})
	req := utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	var actualResp protocol.CheckInResponse
	req.Send(&actualResp)
	assert.Equal(t, config.CodeOK, actualResp.Code)
}

func Test_checkin_normal_fresh_user_random_site_second_time_login_after_enough_time(t *testing.T) {
	t.Skip()
	id := fmt.Sprintf("%s-%s", "api-test-user-id-", utils.RandStringRunes())
	site := fmt.Sprintf("%s-%s", "api-test-site-id-", utils.RandStringRunes())
	body, _ := json.Marshal(protocol.CheckInEvent{
		AnonymousId: id,
		SiteId:      site,
		Timestamp:   0,
	})
	req := utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	var actualResp protocol.CheckInResponse
	req.Send(&actualResp)
	assert.Equal(t, config.CodeOK, actualResp.Code)

	body, _ = json.Marshal(protocol.CheckInEvent{
		AnonymousId: id,
		SiteId:      fmt.Sprintf("%s-%s", "api-test-site-id-", utils.RandStringRunes()),
		Timestamp:   0,
	})
	req = utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	req.Send(&actualResp)
	assert.Equal(t, config.CodeOK, actualResp.Code)

	body, _ = json.Marshal(protocol.CheckInEvent{
		AnonymousId: id,
		SiteId:      site,
		Timestamp:   0,
	})
	req = utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	time.Sleep(time.Second * (config.UserVisitSiteIntervalTimeDuration*60 + 1))
	req.Send(&actualResp)
	assert.Equal(t, config.CodeOK, actualResp.Code)
}

func Test_checkin_normal_fresh_user_random_site_second_time_login_after_not_enough_time(t *testing.T) {
	t.Skip()
	id := fmt.Sprintf("%s-%s", "api-test-user-id-", utils.RandStringRunes())
	site := fmt.Sprintf("%s-%s", "api-test-site-id-", utils.RandStringRunes())
	body, _ := json.Marshal(protocol.CheckInEvent{
		AnonymousId: id,
		SiteId:      site,
		Timestamp:   0,
	})
	req := utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	var actualResp protocol.CheckInResponse
	req.Send(&actualResp)
	assert.Equal(t, config.CodeOK, actualResp.Code)

	body, _ = json.Marshal(protocol.CheckInEvent{
		AnonymousId: id,
		SiteId:      fmt.Sprintf("%s-%s", "api-test-site-id-", utils.RandStringRunes()),
		Timestamp:   0,
	})
	req = utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	req.Send(&actualResp)
	assert.Equal(t, config.CodeOK, actualResp.Code)

	body, _ = json.Marshal(protocol.CheckInEvent{
		AnonymousId: id,
		SiteId:      site,
		Timestamp:   0,
	})
	req = utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	time.Sleep(time.Second * (config.UserVisitSiteIntervalTimeDuration*60 - 10))
	req.Send(&actualResp)
	assert.Equal(t, config.CodeAntiFraudEventError, actualResp.Code)
}

func Test_checkin_anti_fraud_fresh_user_random_site_second_time_login(t *testing.T) {
	body, _ := json.Marshal(protocol.CheckInEvent{
		AnonymousId: fmt.Sprintf("%s-%s", "api-test-user-id-", utils.RandStringRunes()),
		SiteId:      fmt.Sprintf("%s-%s", "api-test-site-id-", utils.RandStringRunes()),
		Timestamp:   0,
	})
	req := utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	var actualResp protocol.CheckInResponse
	req.Send(&actualResp)
	assert.Equal(t, config.CodeOK, actualResp.Code)

	req.Send(&actualResp)
	assert.Equal(t, config.CodeAntiFraudEventError, actualResp.Code)
}

func Test_checkin_anti_fraud_fresh_user_different_site_first_time_login(t *testing.T) {
	id := fmt.Sprintf("%s-%s", "api-test-user-id-", utils.RandStringRunes())
	body, _ := json.Marshal(protocol.CheckInEvent{
		AnonymousId: id,
		SiteId:      fmt.Sprintf("%s-%s", "api-test-site-id-", utils.RandStringRunes()),
		Timestamp:   0,
	})
	req := utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	var actualResp protocol.CheckInResponse
	req.Send(&actualResp)
	assert.Equal(t, config.CodeOK, actualResp.Code)

	body, _ = json.Marshal(protocol.CheckInEvent{
		AnonymousId: id,
		SiteId:      fmt.Sprintf("%s-%s", "api-test-site-id-", utils.RandStringRunes()),
		Timestamp:   0,
	})
	req = utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	req.Send(&actualResp)
	assert.Equal(t, config.CodeOK, actualResp.Code)
}

func Test_checkin_anti_fraud_fresh_user_different_site_revisit(t *testing.T) {
	id := fmt.Sprintf("%s-%s", "api-test-user-id-", utils.RandStringRunes())
	site := fmt.Sprintf("%s-%s", "api-test-site-id-", utils.RandStringRunes())
	body, _ := json.Marshal(protocol.CheckInEvent{
		AnonymousId: id,
		SiteId:      site,
		Timestamp:   0,
	})
	req := utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	var actualResp protocol.CheckInResponse
	req.Send(&actualResp)
	assert.Equal(t, config.CodeOK, actualResp.Code)

	body, _ = json.Marshal(protocol.CheckInEvent{
		AnonymousId: id,
		SiteId:      fmt.Sprintf("%s-%s", "api-test-site-id-", utils.RandStringRunes()),
		Timestamp:   0,
	})
	req = utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	req.Send(&actualResp)
	assert.Equal(t, config.CodeOK, actualResp.Code)

	body, _ = json.Marshal(protocol.CheckInEvent{
		AnonymousId: id,
		SiteId:      site,
		Timestamp:   0,
	})
	req = utils.New()
	req.WithBody(string(body)).WithMethod(http.MethodPost).Build()

	req.Send(&actualResp)
	assert.Equal(t, config.CodeAntiFraudEventError, actualResp.Code)
}
