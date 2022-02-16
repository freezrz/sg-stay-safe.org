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
