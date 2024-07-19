#!/usr/bin/env bash

function run_curl(){
	curl -k -X GET localhost:3004/contacts
	curl -k -X POST localhost:3004/contacts \
		--data '{"Full_Name": "Rupert Bailey","Email": "rupert@coolguy.net.au","Phone_Numbers": ["0477 111 222"]}' 
	curl -k -X GET localhost:3004/contacts
}

function publish(){
	aws configure sso --profile AdministratorAccess-629729073564
	aws ecr-public get-login-password --region us-east-1 --profile AdministratorAccess-629729073564 | \
		docker login --username AWS --password-stdin public.ecr.aws/w5v4j9k4
	
	podman-compose build
	podman push public.ecr.aws/w5v4j9k4/tsagroup_rupertbailey:latest
	podman rmi public.ecr.aws/w5v4j9k4/tsagroup_rupertbailey
}

podman-compose down
podman-compose up -d
run_curl
#podman-compose down
