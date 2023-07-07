#!/bin/sh

# Deploy Step
# Require Minikube

echo "Starting deploy process"
minikube start

# Enable and check Ingress
echo
minikube addons enable ingress
echo
sleep 10
kubectl get pod -n ingress-nginx

# Deploy the Kubernetes in Minikube

kubectl apply -f kubeYAML/configmaps/stock-config.yaml
echo "Applying Configmaps, sleeping 5 seconds"
sleep 5

kubectl apply -f kubeYAML/secrets/stock-secret.yaml
echo "Applying Secrets, sleeping 5 seconds"
sleep 5

kubectl apply -f kubeYAML/deployment/deployment.yaml
echo "Applying Deployment, sleeping 10 seconds"
sleep 5

kubectl apply -f kubeYAML/ingress/stockpicker-ingress.yaml
echo "Applying Ingress, sleeping 5 seconds"
sleep 5

echo "About to run Minikube Tunnel"
echo "Admin Password will apply"
# Run Minikube tunnel
minikube tunnel