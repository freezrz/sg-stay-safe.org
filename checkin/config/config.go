package config

const (
	Region        = "SG"
	CacheDuration = 30 // min

	BanSiteCache = "check-in-cache.vekkvr.0001.apse1.cache.amazonaws.com:6379"
	BanUserCache = "check-in-cache.vekkvr.0001.apse1.cache.amazonaws.com:6379"

	CheckInSiteCache = "check-in-cache.vekkvr.0001.apse1.cache.amazonaws.com:6379"

	BanSiteFormat = "ban-site-%s"
	BanUserFormat = "ban-user-%s"

	SiteCountFormat = "site-%d-%s"

	CodeSiteBanned = 500001
	CodeUserBanned = 500002

	CodeOK = 0
)
