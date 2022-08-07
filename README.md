# golang-helloworld

This repository contains a simple web server implemented in Go, using mostly standard libraries.
The code includes unit tests, a Dockerfile, Kubernetes manifests, a Helm chart and Terraform code for two environments: a local and a GCP one.
Github Actions are used for CI/CD purposes and Terraform Cloud for state management and remote execution. 

## Running the app

Before running, first create a build of the application. This can be done with the `make build` command (or `make all` to run the unit tests befgore the build).

### Binary

The resulting binary can be found in the `bin/` directory and can be run like: `bin/golang-helloworld`.
It supports a command line flag, which specifies which port the server should listen on: `bin/golang-helloworld --port 5000` (the default value is `8080`). This can be overridden by the `PORT` environment variable.

### Container

The app can be also run as a container, for example with Docker.
The latest version of the image can be found at: `ghcr.io/krisztiansala/golang-helloworld:main`.
To start a container, run: `docker run ghcr.io/krisztiansala/golang-helloworld:main -p 8080:8080`.

### Kubernetes

There are two ways to deploy the application to an existing Kubernetes cluster.
`make deploy` will create the Kubernetes objects defined in the `k8s` folder, while `make helm_deploy` will install the Helm chart from the `helm` directory with the default values.
The created resources can be removed with the `make delete` and `make helm_uninstall` commands respectively.

The created service is of type ClusterIP, so not available publicly. That's why the `make portforward` command has to be run if we wish to access the application locally.

#### Create local cluster

To create a local Kubernetes cluster and deploy the application on it with one command, run `make tf_local_apply`. This will spin up a new k3d cluster (needs Docker) and installs the Helm chart on it.
Destroy the above created cluster with `make tf_local_destroy`.

## Infrastructure

### GKE autopilot cluster

To create a remote Kubernetes cluster for running the application on GCP, use `make tf_remote_apply` command. Before running it though, make sure to change the project ID in the `terraform/gcp/terraform.auto.tfvars` file.
We use Autopilot to ease the management tasks of the cluster - this way no node groups have to be defined. Also, you will pay for only the resources that the deployed pods are using and not for the whole node.

### Terraform Cloud

Using the above described method, the Terraform state will be stored locally. This is not a good practice, so that's why we can use Terraform Cloud (free for up to 5 team members) to store the state remotely, apply remotely, handle concurrency control, etc.

A new workspace will have to be created where the `GOOGLE_CREDENTIALS` variable set to a service account JSON key. The service account needs to have the necessary permissions in the project to create and destroy the managed resources. Feel free to make a new role for it with fine grained permission control, but for the purpose of the example, I just gave it owner on the project.

Integrate the workspace with the github repository and specify the path to the `terraform` directory `gcp` folder. With this, for each change in those files and new run will be created on Terraform Cloud and automatically (or manually) applied (based on preference). 

### Github Actions

There are two workflows in this repository. The deployment workflow will run on all source code changes pushed to the main branch. It will unit test the application, build a docker image from it and deploy the resulted image to the GKE cluster. 

The deployment is done through a service account (also created through Terraform), which has admin access over the cluster (feel free to fine tune this as well).

The other workflow will check the formatting and validity of the Terraform code and run a plan on Terraform cloud (for this an API token needs to be set as a secret for the Action). Apllying the infrastructure changes do not happen in this pipeline, instead it has to be manually approved in Terraform Cloud.