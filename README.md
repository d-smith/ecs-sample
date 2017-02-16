# ECS Sample

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
initiating CPU based autoscale events, e.g.

<pre>
ab -n 1000000 -c 100 -X http.proxy.fmr.com:8000 ECSALB-endpoint.us-east-1.elb.amazonaws.com/pf/25357
</pre>

After the stacks have been installed, you can curl the ALB endpoint on the /hcping uri.

Note that the alb script bakes in the docker hub repo name of the hcping container - you may
need to change this to reflect your set up (and build and push the docker image as well).

There are two sample images here which can be built and run locally, plus can be used to define ECS tasks which
can be used in ECS samples:

<pre>
docker run -p 5000:5000 xtracdev/service2
docker run -p 4000:4000 xtracdev/service1
</pre>