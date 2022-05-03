# sg-stay-safe.org - software solution architecting practice module
sg-stay-safe.org v.0.0.0 - Official Release

Implemented:
* CRUD admin portal for user, site, safe ambassador, region, rule
* cron job 4PM daily to sync: banned user, banned site, and rules to the cache
    * use aws event bridge - will update the scheduler to a shorten time when demo
* if a normal user checkin a normal site, can site in successfully, and generate check-in kafka msg
* if a banned user checkin a normal site, will block from check-in, and generate a warning msg, and the msg will consumed and send email to safe ambassador (email will sent to different receipent based on the check-ined site's region)
* if a normal user checkin a banned site, will block from check-in, and generate a warning msg, and the msg will consumed and send email to safe ambassador (email will sent to different receipent based on the check-ined site's region)
* if a normal user check-in more than 10 times (can setup via rule-engine) daily, will consider violation, will block from check-in, and generate a warning msg, and the msg will consumed and send email to safe ambassador (email will sent to different receipent based on the check-ined site's region)
* if a normal user check-in a site more than once in 5mins, will remind him too frequent
* all check-in numbers will be saved in cache and all check-in events will be stored in Kafka message queue
