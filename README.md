# StockTicker

StockTicker is a Microservice REST API which retrieves the Stock Ticker last NDAYS of dat and the average of the share price. 

NDAYS of data calculate the past days starting from today. 

For example: if you want to retrieve the past 7 days of stock data, NDAYS = 7 and it will start counting for any stock data 7 days back(including today).

## How to build the image

1. Build Image as zipzapzup-stockpicker `chmod +x build.sh && ./build.sh`
2. Run the container `docker run --detach --publish 8080:8080 zipzapzup-stockpicker`
3. Locally test to see it works: `curl http://localhost:8080`

## Publish the Image
1. Create a DockerHub Repository `$1`/stockpicker, where `$1` is your username
2. Run `docker tag zipzapzup-stockpicker $1/stockpicker` and substitute `$1` to your docker hub username
2. Publish via `docker push $1/stockpicker`

## Minikube

1. Using Minikube version `v1.28`
2. `minikube start`
3. Enable Ingress Controller `minikube addons enable ingress`
4. Confirm: `kubectl get pod -n ingress-nginx`
5. Deploy the deployment and ingress `kubectl apply -f kubeYAML`
7. Deploy
    - `kubectl apply -f deployment/deployment.yaml`
    - `kubectl apply -f ingress/stockpicker-ingress.yaml`
8. Add `192.168.49.2	127.0.0.1` to your `/etc/hosts`
9. Run `minikube tunnel`
10. Voila and now test: `curl 127.0.0.1:80` !