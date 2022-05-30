#!/usr/bin/env bash
dir=$(cd $(dirname $(echo $BASH_SOURCE)) && pwd)
echo ${dir}
docker build -t exchange ${dir}
docker run -d --name fast -p8080:8080 exchange 100 1000 8080
docker run -d --name slow -p9090:9090 exchange 500 100 9090
