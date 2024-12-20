#!/bin/bash

cd /root/.jenkins/workspace/ruoyi-front

docker stop ruoyi-front || true
docker rm -f ruoyi-front || true
docker rmi ruoyi-front:latest || true
cd ./docker && docker build -t ruoyi-front:latest .
docker run -d -p 19013:80 --restart=always --name ruoyi-front --log-driver=json-file --log-opt max-size=5m --log-opt max-file=3 ruoyi-front:latest

echo "Starting running smart-front at port 19013"