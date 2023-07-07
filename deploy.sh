#!/bin/sh

# Deploy Step
# Require Minikube

echo "Starting deploy process"
minikube start

# Enable and check Ingress
echo
minikube addons enable ingress
echo
kubectl get pod -n ingress-nginx

# Deploy the Kubernetes in Minikube

kubectl apply -f configmaps/stock-config.yaml
kubectl apply -f secrets/stock-secret.yaml
kubectl apply -f deployment/deployment.yaml
kubectl apply -f ingress/stockpicker-ingress.yaml

# Run Minikube tunnel
minikube tunnel