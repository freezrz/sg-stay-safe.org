package config

const (
	Region        = "SG"
	CacheDuration = 30 // min

	UserVisitSiteIntervalTimeDuration = 5 // min

	AWSForwardPort    = "5000"
	AWSRegion         = "ap-southeast-1"
	SystemEmailSender = "sg.stay.safe.org@gmail.com"

	SanitiseCheckinLambda             = "sanitise_checkin"
	AntiFraudCheckinLambda            = "anti_fraud_checkin"
	RecordCheckinLambda               = "record_checkin"
	ProduceCheckinMsgLambda           = "produce_checkin_event"
	ProduceUserViolationMsgLambda     = "produce_user_violation_event"
	ProduceSiteViolationMsgLambda     = "produce_site_violation_event"
	VerifyRulesLambda                 = "verify_rules"
	RetrieveRegionEmailBySiteIdLambda = "retrieve_region_email_by_siteid_query_service"
	SendEmailLambda                   = "send_email_notification"

	BanCache         = "check-in-cache.vekkvr.0001.apse1.cache.amazonaws.com:6379"
	CheckInSiteCache = "check-in-cache.vekkvr.0001.apse1.cache.amazonaws.com:6379"
	AntiFraudCache   = "check-in-cache.vekkvr.0001.apse1.cache.amazonaws.com:6379"
	RuleCache        = "check-in-cache.vekkvr.0001.apse1.cache.amazonaws.com:6379"

	BanSiteFormat = "ban-site-%s"
	BanUserFormat = "ban-user-%s"

	RulePrefix          = "rule-"
	RuleFormat          = RulePrefix + "%s"
	RuleMaxDailyCheckin = RulePrefix + "MAX_DAILY_CHECKIN"

	SiteCountFormat               = "site-%d-%s"
	UserVisitSiteHistoryFormat    = "user-%s-site-%s"
	User24HoursCheckinCountFormat = "user-24hrs-%s"

	CheckinEventKafkaZooKeeper = "z-1.checkin-msk-clust.xhf5xv.c2.kafka.ap-southeast-1.amazonaws.com:2181" // for consumer
	CheckinEventKafkaBootstrap = "b-1.checkin-msk-clust.xhf5xv.c2.kafka.ap-southeast-1.amazonaws.com:9092" // for producer
	CheckinEventKafkaTopic     = "checkin-msk-topic"
	UserViolationKafkaTopic    = "user-violation-msk-topic"
	SiteViolationKafkaTopic    = "site-violation-msk-topic"

	CodeOK                  = 0
	CodeServerInternalError = 502
	CodeMarshalError        = 30001
	CodeUnmarshalError      = 30002
	CodeInvokeLambdaError   = 30003

	CodeSanitiseError                 = 41001
	CodeProduceEventError             = 42001
	CodeRecordCacheEventError         = 43001
	CodeRecordUserVisitSiteEventError = 43002
	CodeAntiFraudEventError           = 44001

	CodeBanUserUpdateCacheError   = 44002
	CodeUnBanUserUpdateCacheError = 44003

	CodeBanSiteUpdateCacheError   = 45001
	CodeUnBanSiteUpdateCacheError = 45002

	CodeUserIsBannedError              = 451001
	CodeUserBannedCacheError           = 451002
	CodeVerifyUserMaxCheckinCacheError = 451003
	CodeUserExceedDailyMaxCheckinError = 451004
	CodeIncrUser24hrsCheckinError      = 451005
	CodeSiteIsBannedError              = 452001
	CodeSiteBannedCacheError           = 452002

	CodeNotificationSendEmailError = 500001
)
