A couple sample services to use with ECS

Some notes...

* Used the ECS cluster create that now sets up the VPNS, the instances, etc.
* Used container runtime port mapping to allow multiple of the same container to be run on the same host - in the task definition just specify the container port and no port value on the docker host
* You need to add ingress from the application load balancer's security group to the security group used for the cluster instances.

