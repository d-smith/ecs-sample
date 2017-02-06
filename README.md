A couple sample services to use with ECS

Some notes...

* Used the ECS cluster create that now sets up the VPNS, the instances, etc.
* Used container runtime port mapping to allow multiple of the same container to be run on the same host - in the task definition just specify the container port and no port value on the docker host
* You need to add ingress from the application load balancer's security group to the security group used for the cluster instances.

There are two images here which can be run locally:

<pre>
docker run -p 5000:5000 xtracdev/service2
docker run -p 4000:4000 xtracdev/service1
</pre>

The goal with this sample is to define everything needed to deploy this to AWS:

* Define the VPC to run the ECS cluster in
* Define the ECS cluster including the launch config and auto scale group
* Define the services and tasks to run both service1 and service2
* Define the load balancer config to route calls to the services which are exposed via
load balancer config and ephemeral docker ports