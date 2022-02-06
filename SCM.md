EC2:
* create instance under Ubuntu 20
* sudo apt-get update
* sudo apt-get install redis

Elasticache:
* create cache under Redis
* without any DR, backup, or advanced features avoiding unnecessary billing
* shut it down if you are not around the laptop cause it is very costly

Lambda:
* for all lambda, please choose [lambda-vpc-role](https://console.aws.amazon.com/iamv2/home?#/roles/details/lambda-vpc-role?section=permissions) as Role when created if it needs to access other resources 
* add layer -> API Gateway -> HTTP API
    * choose VPC -> choose the default VPC, subnets (choose all), and security groups -> default | redis-dev
* runtime setting -> x86_64, "Handler" should be the uploaded binary name
