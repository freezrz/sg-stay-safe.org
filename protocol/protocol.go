package protocol

type CheckInEvent struct {
	AnonymousId string `json:"anonymous_id"`
	SiteId      string `json:"site_id"`
	Timestamp   int64  `json:"timestamp"`
}

type GeneralResponse struct {
	Code int
	Msg  string
}

type CheckInResponse struct {
	Code int
	Msg  string
}

type VerifyRulesForCheckInResponse struct {
	Code int
	Msg  string
}

type BanSiteRequest struct {
	SiteId    string `json:"site_id"`
	ShouldBan bool   `json:"should_ban"`
}

type BanUserRequest struct {
	AnonymousId string `json:"anonymous_id"`
	ShouldBan   bool   `json:"should_ban"`
}
