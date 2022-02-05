package protocol

type CheckInEvent struct {
	AnonymousId string `json:"anonymous_id"`
	PlaceId     string `json:"place_id"`
}

type SanitiserResponse struct {
	Code int
	Msg  string
}
