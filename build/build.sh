# !/bin/bash

docker build -t zipzapzup-stockpicker ../Dockerfile


# run command: docker run --detach --publish 8080:8080 docker-stockpicker:v5