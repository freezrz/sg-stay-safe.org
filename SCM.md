EC2:
* create instance under Ubuntu 20
* sudo apt-get update
* sudo apt-get install redis
* ubuntu@ip-172-31-29-7:~$ pwd
  /home/ubuntu
  ubuntu@ip-172-31-29-7:~$ ls ~/.aws/credentials 
  /home/ubuntu/.aws/credentials
  (for aws lambda invoke)

Elasticache:
* create cache under Redis
* without any DR, backup, or advanced features avoiding unnecessary billing
* shut it down if you are not around the laptop cause it is very costly

Elastic Beanstalk:
* aws-elasticbeanstalk-ec2-role (so it can invoke lambda)

Lambda:
* for all lambda, please choose [lambda-vpc-role](https://console.aws.amazon.com/iamv2/home?#/roles/details/lambda-vpc-role?section=permissions) as Role when created if it needs to access other resources 
    * fyi only, [Create the execution role](https://docs.aws.amazon.com/lambda/latest/dg/services-elasticache-tutorial.html)
* add layer -> API Gateway -> HTTP API
    * choose VPC -> choose the default VPC, subnets (choose all), and security groups -> default | redis-dev
* runtime setting -> x86_64, "Handler" should be the uploaded binary name
* if need to send email, choose lambda-send-email role

Amazon MSK (manageed streaming for Apache Kafka):
* https://aws.amazon.com/blogs/compute/setting-up-aws-lambda-with-an-apache-kafka-cluster-within-a-vpc/
https://dev.to/aws-builders/streaming-messages-from-producer-to-consumer-using-amazon-msk-and-create-an-event-source-to-msk-using-lambda-252
https://www.linkedin.com/learning/paths/prepare-for-the-aws-certified-solutions-architect-associate-exam-saa-c02
https://www.udemy.com/course/aws-solutions-architect-associate-pass-the-saa-c02-exam/?utm_source=adwords&utm_medium=udemyads&utm_campaign=LongTail_la.EN_cc.ROW&utm_content=deal4584&utm_term=_._ag_77879423894_._ad_535397245857_._kw__._de_c_._dm__._pl__._ti_dsa-1007766171032_._li_9062531_._pd__._&matchtype=&gclid=Cj0KCQiA3fiPBhCCARIsAFQ8QzXvHILNin4rp7qwgnv4NWOdz-pqqcQ2ML3KTmnJ44w0eCYP4HquQ4AaApVQEALw_wcB