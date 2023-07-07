# !/bin/bash

echo "Starting Build Command"

docker build -t zipzapzup-stockpicker .
docker images zipzapzup-stockpicker
echo
echo "Image built: zipzapzup-stockpicker:latest !"
echo "Completed building image"
