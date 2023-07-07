# !/bin/bash

# perfected is your own docker hub repository
echo "Starting release process"

docker tag zipzapzup-stockpicker perfected/stockpicker
docker push perfected/stockpicker
docker images

echo
echo "Released"
