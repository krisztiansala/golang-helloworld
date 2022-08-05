provider "helm" {
  kubernetes {
    config_path    = "~/.kube/config"
    config_context = "k3d-${var.cluster_name}"
  }
}

terraform {
  required_providers {
    k3d = {
      source  = "pvotal-tech/k3d"
      version = "0.0.6"
    }
  }
}

provider "k3d" {}

resource "k3d_cluster" "mycluster" {
  name = var.cluster_name

  kubeconfig {
    update_default_kubeconfig = true
    switch_current_context    = true
  }
}

resource "helm_release" "golang-helloworld" {
  name      = "golang-helloworld"
  chart     = "${path.root}/../../helm"
  namespace = "default"
  depends_on = [
    k3d_cluster.mycluster
  ]
}