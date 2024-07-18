i#!/usr/bin/env bash

podman stop gin

podman build -t go-docker-dev.to-gin . && \
podman run --rm --detach --name gin --publish 3004:3004 go-docker-dev.to-gin

curl -k -X GET localhost:3004/contacts
curl -k -X POST --data '{"Full_Name": "Rupert Bailey","Email": "rupert@coolguy.net.au","Phone_Numbers": ["0477 111 222"]}' localhost:3004/contacts
curl -k -X GET localhost:3004/contacts
