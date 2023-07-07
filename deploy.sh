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

kubectl apply -f kubeYAML/configmaps/stock-config.yaml
kubectl apply -f kubeYAML/secrets/stock-secret.yaml
kubectl apply -f kubeYAML/deployment/deployment.yaml
kubectl apply -f kubeYAML/ingress/stockpicker-ingress.yaml

# Run Minikube tunnel
minikube tunnel