package config

import (
	"fmt"
	"os"
)

const (
	Region        = "SG"
	CacheDuration = 30 // min

	UserVisitSiteIntervalTimeDuration = 5 // min

	AWSForwardPort    = "5000"
	AWSRegion         = "ap-southeast-1"
	SystemEmailSender = "sg.stay.safe.org@gmail.com"

	// move the config to AWS - Lambda environment variables
	BanCache         = ""
	CheckInSiteCache = ""
	AntiFraudCache   = ""
	RuleCache        = ""

	BanSiteFormat = "ban-site-%s"
	BanUserFormat = "ban-user-%s"

	RulePrefix          = "rule-"
	RuleFormat          = RulePrefix + "%s"
	RuleMaxDailyCheckin = RulePrefix + "MAX_DAILY_CHECKIN"

	SiteCountFormat               = "site-%d-%s"
	UserVisitSiteHistoryFormat    = "user-%s-site-%s"
	User24HoursCheckinCountFormat = "user-24hrs-%s"

	// move the config to AWS - Lambda environment variables
	CheckinEventKafkaZooKeeper = "" // for consumer
	CheckinEventKafkaBootstrap = "" // for producer
	CheckinEventKafkaTopic     = ""
	UserViolationKafkaTopic    = ""
	SiteViolationKafkaTopic    = ""

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

var (
	env                               = os.Getenv("ENV")
	SanitiseCheckinLambda             = fmt.Sprintf("%s%s", env, "sanitise_checkin")
	AntiFraudCheckinLambda            = fmt.Sprintf("%s%s", env, "anti_fraud_checkin")
	RecordCheckinLambda               = fmt.Sprintf("%s%s", env, "record_checkin")
	ProduceCheckinMsgLambda           = fmt.Sprintf("%s%s", env, "produce_checkin_event")
	ProduceUserViolationMsgLambda     = fmt.Sprintf("%s%s", env, "produce_user_violation_event")
	ProduceSiteViolationMsgLambda     = fmt.Sprintf("%s%s", env, "produce_site_violation_event")
	VerifyRulesLambda                 = fmt.Sprintf("%s%s", env, "verify_rules")
	RetrieveRegionEmailBySiteIdLambda = fmt.Sprintf("%s%s", env, "retrieve_region_email_by_siteid_query_service")
	SendEmailLambda                   = fmt.Sprintf("%s%s", env, "send_email_notification")
)
