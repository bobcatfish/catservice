# catservice
Yay cats!

## Running it

```bash
go run main.go
```

## Building the Docker image

```bash
docker build -f DOCKERFILE . -t catservice
docker run -p 80:80 catservice
```

## Running it in kubernetes with Istio

### Creating a GKE cluster with Istio & Prometheus

```bash
export PROJECT_ID=christiewilson-catfactory
export CLUSTER_NAME=ilovecats

gcloud beta container clusters create $CLUSTER_NAME \
 --enable-autoscaling \
 --min-nodes=1 \
 --max-nodes=3 \
 --scopes=cloud-platform \
 --enable-basic-auth \
 --no-issue-client-certificate \
 --project=$PROJECT_ID \
 --region=asia-northeast1 \
 --machine-type=n1-standard-4 \
 --image-type=cos \
 --num-nodes=1 \
 --cluster-version=latest \
 --addons=HorizontalPodAutoscaling,Istio \
 --istio-config=auth=MTLS_PERMISSIVE


kubectl create clusterrolebinding cluster-admin-binding \
--clusterrole=cluster-admin \
--user=$(gcloud config get-value core/account)
```

### Once it's setup

Finding the external IP of the ingress gateway:

```bash
kubectl -n istio-system get svc
```
