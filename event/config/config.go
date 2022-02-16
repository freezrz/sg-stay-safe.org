package config

const (
	CheckinEventKafkaZooKeeper = "z-1.checkin-msk-clust.xhf5xv.c2.kafka.ap-southeast-1.amazonaws.com:2181"
	CheckinEventKafkaBootstrap = "b-1.checkin-msk-clust.xhf5xv.c2.kafka.ap-southeast-1.amazonaws.com:9092"
	CheckinEventKafkaTopic     = "checkin-msk-topic"

	CodeOK                  = 0
	CodeServerInternalError = 502
	CodeMarshalError        = 30001

	CodeSanitiseError         = 41001
	CodeProduceEventError     = 42001
	CodeRecordCacheEventError = 43001
)
