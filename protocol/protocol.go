package protocol

type CheckInEvent struct {
	AnonymousId string `json:"anonymous_id"`
	SiteId      string `json:"site_id"`
	Timestamp   int64  `json:"timestamp"`
}

type GeneralResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ViolationEvent struct {
	CheckInEvent    CheckInEvent
	Site            Site
	GeneralResponse GeneralResponse
	Region          Region
}

type CheckInResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type BanSiteRequest struct {
	SiteId    string `json:"site_id"`
	ShouldBan bool   `json:"should_ban"`
}

type BanUserRequest struct {
	AnonymousId string `json:"anonymous_id"`
	ShouldBan   bool   `json:"should_ban"`
}

type Rule struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsEnabled   bool   `json:"is_enabled"`
	Value       int    `json:"value"`
}

type ViolationEmail struct {
	Recipient   string `json:"recipient"`
	AnonymousId string `json:"anonymous_id"`
	Site        Site   `json:"site"`
	Reason      string `json:"reason"`
}

type Site struct {
	Name        string `json:"name"`
	SiteId      string `json:"site_id"`
	Owner       string `json:"owner"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	PostalCode  string `json:"postal_code"`
	Description string `json:"description"`
	Capacity    string `json:"capacity"`
	ShouldBan   bool   `json:"should_ban"`
	Region      string `json:"region"`
}

type User struct {
	AnonymousId string `json:"anonymous_id"`
	Description string `json:"description"`
	ShouldBan   bool   `json:"should_ban"`
}

type Region struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
}
