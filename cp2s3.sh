#!/bin/bash
aws s3 cp vpc-pub-priv.yaml $S3_TEMPLATE_BUCKET
aws s3 cp ecs-cluster.yaml $S3_TEMPLATE_BUCKET
aws s3 cp alb.yaml $S3_TEMPLATE_BUCKET
aws s3 cp hc.yaml $S3_TEMPLATE_BUCKET
aws s3 cp sumolambda.yaml $S3_TEMPLATE_BUCKET
aws s3 cp pfservice.yaml $S3_TEMPLATE_BUCKET
