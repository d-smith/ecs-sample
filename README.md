# ECS Sample

## Caveats

This is a PoC of cloud formation templates for a single region ECS application. There are some important
next steps that need to be done to make this work outside my account context, and to include some
missing functionality:

* Currently the templates rely on some IAM roles being present in the account where the templates are
run. The creation of the roles and policies used with this template will need to be extracted before they
can be more broadly reused.
* These templates were originally started using HTTP then transitioned  to HTTPS - they need to be 
closely reviewed to make sure the transition was done cleanly.
* The stack should pull the demo images from an ECS registry, not docker hub.
* The AWS WAF should be included at the perimeter.

## EC2 Platform Note

The vpc template assumes the EC2 platform is VPC, and that the default subnets have not been
delete from the account. If this assumption does not hold, there could be problems running the template in EC2 classic
platforms based on the differences in how !GetAZs behaves. For example, in us-west-1, for the EC2 Classic
platform, 3 AZs are returned, only two of which may contain VPC subnets (indexes 0 and 2). When run
from an account that supports only the VPC platform, two AZs are returned.

For more on EC2 platforms see [here](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-platforms.html?icmpid=docs_ec2_console)

## Overview

This project provides cloud formation templates and some container code to be used for setting up 
an ECS sample that includes all the pieces used in a sample application, including
VPC setup, ECS cluster definition, and ELB v2 configuration.

This template has been used in us-west-1 and us-east-1. Note in the us-west-1 region, there are 3 availability zones, 
but only zone a and zone c can be have VPC subnets defined in them.

To use the sample:

* Use vpc-pub-priv.yaml to create the VPC set up.
* Use ecs-cluster.yml to create the ECS cluster, referencing the stack name used to create the VPC (needed to import
network information from the vpc setup)
* Use alb.yml to set up the ELC and to run the hcping container on the cluster nodes to allow
connectivity and network set up to be verified.
* Use pfservice.yaml to install a service that adds the prime factors container to the cluster and 
load balancer. The prime factors service is useful for driving up CPU consumption for the purposes of
initiating CPU based autoscale events.

As an example:

<pre>
aws cloudformation create-stack \
--stack-name network-a \
--template-body file://$PWD/vpc-pub-priv.yaml

aws cloudformation create-stack \
--stack-name ecs-rancher \
--template-body file://$PWD/ecs-cluster.yml \
--parameters ParameterKey=AMIType,ParameterValue=Rancher \
ParameterKey=NetworkStack,ParameterValue=network-a \
ParameterKey=KeyName,ParameterValue=MyKeyName \
ParameterKey=EcsClusterName,ParameterValue=cluster-east

aws cloudformation create-stack \
--stack-name alb-east \
--template-body file://$PWD/alb.yaml \
--parameters ParameterKey=NetworkStack,ParameterValue=network-a \
ParameterKey=ECSClusterStack,ParameterValue=ecs-rancher \
ParameterKey=LoadBalancerName,ParameterValue=alice

aws cloudformation create-stack \
--stack-name pf \
--template-body file://$PWD/pfservice.yaml \
--parameters ParameterKey=NetworkStack,ParameterValue=network-a \
ParameterKey=ECSClusterStack,ParameterValue=ecs-rancher \
ParameterKey=ALBStack,ParameterValue=alb-east \
--capabilities CAPABILITY_IAM
</pre>

After the stacks have been installed, you can curl the ALB endpoint on the /hcping uri.

Note that the alb script bakes in the docker hub repo name of the hcping container - you may
need to change this to reflect your set up (and build and push the docker image as well).

The container referenced in the templates can be build from the code in this project.
It can be run locally via `docker run -p 5000:5000 xtracdev/pfactors`

The local endpoint can be curled via `curl localhost:5000/pf/1337`

## Logging Notes

Setting up cloud watch logs is a bit of a pain if you want the docker runtime log, but long story short is the ecs
container role must have the appropriate access granted and the user data
must be configured to set up the logging service. See
[here](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_cloudwatch_logs.html)
for details.

If you just want the container out, then create a cloud watch log group, and configure awslog as
the logging driver in the task definition.

### Sumo Integration

Before you can create the sumo integration piece via sumo lambda function stack,
you need to create an s3 bucked, zip up the [sumo cloud watch lambda function](https://github.com/SumoLogic/sumologic-aws-lambda/tree/master/cloudwatchlogs),
and drop it into the S3 bucket. The bucket name and zip file name are referenced as stack parameters.

## Self Signed Certificate

For dev if you can tolerate a self signed certificate, here's how to generate
one for import into the AWS ACM

<pre>
openssl genrsa 2048 > privatekey.pem
openssl req -new -key privatekey.pem -out csr.pem
openssl x509 -req -days 365 -in csr.pem -signkey privatekey.pem -out server.crt
</pre>

## Obtaining Certificate ARNs

The alb stack requires a certificate ARN as an input. The easiest way to
obtain one is via the aws cli: `aws acm --list-certificates`

## Route 53 Parameters

Use the AWS CLI to find the zone name to form the record set for. This can
be done via `aws route53 list-hosted-zones'. The name is the last component
of the Config Id property - for example if Id is /hostedzone/xxx then the 
zone name is xxx.

To form the domain name parameter, combine the name you are associating with
the load balancer with the parent name associated with the hosted zone.

For example, in the hosted zone config if the name is `foo.com.` and you
want to associate `dev` in that domain with the load balancer, then
for the domain name parameter you would specify `dev.foo.com`. 


## Helpful Exports
export AWS_DEFAULT_PROFILE=xxx
export https_proxy=xxx
export S3_TEMPLATE_BUCKET=xxx