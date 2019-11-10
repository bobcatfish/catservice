# Tekton Pipelines

We have 2 pipelines:
* [pr-pipline.yaml](./pr-pipeline.yaml) is for testing before merging
* [deploy-pipeline.yaml](./deploy-pipeline.yaml) is for testing after merging

## Triggers

Triggering is configured with [triggers.yaml](./triggers.yaml).

```bash
# Have to build the github validator
ko apply -f tekton/triggers.yaml
```

## Running the Pull Request Pipeline

This Pipeline uses [golang-test](https://github.com/tektoncd/catalog/tree/master/golang#golang-test) 
([golang-test.yaml](golang-test.yaml))
which is copied from [the tekton catalog](https://github.com/tektoncd/catalog),
but [golang-test-pr.yaml](golang-test-pr.yaml) is heavily modified to update a
[Pull Request resource](https://github.com/tektoncd/pipeline/blob/master/docs/resources.md#pull-request-resource)
before we have support for
[taking actions on failure](https://github.com/tektoncd/pipeline/issues/1376).

```bash
kubectl apply -f tekton/golang-test-pr.yaml
kubectl apply -f tekton/set-status.yaml
kubectl apply -f tekton/pr-pipeline.yaml
kubectl apply -f tekton/resources.yaml

# Make new runs with cli
tkn pipeline start pr-pipeline -r source-repo=catservice
```

## Running the Deploy Pipeline

This is a POC of a canary deployment with Tekton!

```bash
# TODO: better way to support pipeline level tags + kaniko building
kubectl apply -f tekton/kaniko-tag.yaml
kubectl apply -f tekton/canary-deployment.yaml
kubectl apply -f tekton/scale.yaml
kubectl apply -f tekton/update-image.yaml
kubectl apply -f tekton/remove-canary-deployment.yaml
kubectl apply -f tekton/deploy-pipeline.yaml

# To get credentials for the target cluster:
# kubectl get serviceaccounts robot -o yaml -n catspace
# kubectl get secret robot-token-zgf2d -o yaml -n catspace
# Then base64 decode the token!!!!
# kubectl apply -f tekton/resources-cluster.yaml
kubectl apply -f tekton/resources.yaml

# Make new runs with cli
tkn pipeline start canary-pipeline -r source-repo=catservice -r image=christie-catservice-image -r cluster=catservice-cluster -p tag=0.10.0
```


### Creating a GKE cluster with Istio & Prometheus

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

### Once it's setup

Finding the external IP of the ingress gateway:

```bash
kubectl -n istio-system get svc
```
