This is a POC of a canary deployment with Tekton!

## Running the Pipeline

```bash
# TODO: better way to support pipeline level tags + kaniko building
kubectl apply -f tekton/kaniko-tag.yaml
kubectl apply -f tekton/canary-deployment.yaml
kubectl apply -f tekton/scale.yaml
kubectl apply -f tekton/update-image.yaml
kubectl apply -f tekton/remove-canary-deployment.yaml
kubectl apply -f tekton/pipeline.yaml

# To get credentials for the target cluster:
# kubectl get serviceaccounts robot -o yaml -n catspace
# kubectl get secret robot-token-zgf2d -o yaml -n catspace
# Then base64 decode the token!!!!
# kubectl apply -f tekton/resources-cluster.yaml
kubectl apply -f tekton/resources.yaml


# Make new runs with cli
tkn pipeline start canary-pipeline -r source-repo=catservice -r image=christie-catservice-image -r cluster=catservice-cluster -p tag=0.10.0
```


## Creating a GKE cluster with Istio & Prometheus

```bash
export PROJECT_ID=christiewilson-catfactory
export CLUSTER_NAME=canary-cluster

gcloud beta container clusters create $CLUSTER_NAME \
 --enable-autoscaling \
 --min-nodes=1 \
 --max-nodes=3 \
 --scopes=cloud-platform \
 --enable-basic-auth \
 --no-issue-client-certificate \
 --project=$PROJECT_ID \
 --region=us-central1 \
 --machine-type=n1-standard-4 \
 --image-type=cos \
 --num-nodes=1 \
 --cluster-version=latest \
 --addons=HorizontalPodAutoscaling,Istio \
 --istio-config=auth=MTLS_PERMISSIVE


kubectl create clusterrolebinding cluster-admin-binding \
--clusterrole=cluster-admin \
--user=$(gcloud config get-value core/account)

# Latest version of istio (1.1.7) doesn't seem to have prometheus released with it?
kubectl -n istio-system apply -f  https://storage.googleapis.com/gke-release/istio/release/1.0.6-gke.3/patches/install-prometheus.yaml
```

## Once it's setup

Finding the external IP of the ingress gateway:

```bash
kubectl -n istio-system get svc
```