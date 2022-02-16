package config

const (
	Region        = "SG"
	CacheDuration = 30 // min

	AWSForwardPort = "5000"
	AWSRegion      = "ap-southeast-1"

	SanitiseCheckinLambda   = "sanitise_checkin"
	RecordCheckinLambda     = "record_checkin"
	ProduceCheckinMsgLambda = "produce_checkin_event"
	VerifyRulesLambda       = "verify_rules"

	BanCache = "check-in-cache.vekkvr.0001.apse1.cache.amazonaws.com:6379"

	CheckInSiteCache = "check-in-cache.vekkvr.0001.apse1.cache.amazonaws.com:6379"

	BanSiteFormat = "ban-site-%s"
	BanUserFormat = "ban-user-%s"

	SiteCountFormat = "site-%d-%s"

	CheckinEventKafkaZooKeeper = "z-1.checkin-msk-clust.xhf5xv.c2.kafka.ap-southeast-1.amazonaws.com:2181"
	CheckinEventKafkaBootstrap = "b-1.checkin-msk-clust.xhf5xv.c2.kafka.ap-southeast-1.amazonaws.com:9092"
	CheckinEventKafkaTopic     = "checkin-msk-topic"

	CodeOK                  = 0
	CodeServerInternalError = 502
	CodeMarshalError        = 30001
	CodeUnmarshalError      = 30002
	CodeInvokeLambdaError   = 30003

	CodeSanitiseError         = 41001
	CodeProduceEventError     = 42001
	CodeRecordCacheEventError = 43001

	CodeBanUserUpdateCacheError   = 44001
	CodeUnBanUserUpdateCacheError = 44002

	CodeBanSiteUpdateCacheError   = 45001
	CodeUnBanSiteUpdateCacheError = 45002

	CodeUserIsBannedError    = 451001
	CodeUserBannedCacheError = 451002
	CodeSiteIsBannedError    = 452001
	CodeSiteBannedCacheError = 452002
)
