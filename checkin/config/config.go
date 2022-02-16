package config

const (
	Region        = "SG"
	CacheDuration = 30 // min

	AWSForwardPort = "5000"
	AWSRegion      = "ap-southeast-1"

	SanitiseCheckinLambda   = "sanitise_checkin"
	RecordCheckinLambda     = "record_checkin"
	ProduceCheckinMsgLambda = "produce_checkin_event"

	BanSiteCache = "check-in-cache.vekkvr.0001.apse1.cache.amazonaws.com:6379"
	BanUserCache = "check-in-cache.vekkvr.0001.apse1.cache.amazonaws.com:6379"

	CheckInSiteCache = "check-in-cache.vekkvr.0001.apse1.cache.amazonaws.com:6379"

	BanSiteFormat = "ban-site-%s"
	BanUserFormat = "ban-user-%s"

	SiteCountFormat = "site-%d-%s"

	CodeSiteBanned = 50001
	CodeUserBanned = 50002

	CodeOK                  = 0
	CodeServerInternalError = 502
	CodeMarshalError        = 30001
	CodeUnmarshalError      = 30002
	CodeInvokeLambdaError   = 30003

	CodeSanitiseError         = 41001
	CodeProduceEventError     = 42001
	CodeRecordCacheEventError = 43001
)
