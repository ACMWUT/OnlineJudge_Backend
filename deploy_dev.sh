#!/bin/bash
# cross compile
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
# upload to server htdocs
scp OnlineJudge root@172.16.2.48:/home/htdocs/dev/backend/app/
# execute build and deploy script
ssh root@172.16.2.48 "sh /root/deploy_dev.sh"