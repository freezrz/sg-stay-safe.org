package protocol

type CheckInEvent struct {
	AnonymousId string `json:"anonymous_id"`
	SiteId      string `json:"site_id"`
}

type VerifyRulesForCheckInResponse struct {
	Code int
	Msg  string
}

type BanShopRequest struct {
	SiteId string `json:"site_id"`
}

type BanShopResponse struct {
	Code int
	Msg  string
}

type BanUserRequest struct {
	AnonymousId string `json:"anonymous_id"`
}

type BanUserResponse struct {
	Code int
	Msg  string
}
