package protocol

type CheckInEvent struct {
	AnonymousId string `json:"anonymous_id"`
	SiteId      string `json:"site_id"`
}

type VerifyRulesForCheckInResponse struct {
	Code int
	Msg  string
}

type BanSiteRequest struct {
	SiteId    string `json:"site_id"`
	ShouldBan bool   `json:"should_ban"`
}

type BanSiteResponse struct {
	Code int
	Msg  string
}

type BanUserRequest struct {
	AnonymousId string `json:"anonymous_id"`
	ShouldBan   bool   `json:"should_ban"`
}

type BanUserResponse struct {
	Code int
	Msg  string
}
