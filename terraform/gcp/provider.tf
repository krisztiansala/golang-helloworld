terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.3.0"
    }
  }
  backend "remote" {
    organization = "krisztian-test"

    workspaces {
      name = "golang-helloworld"
    }
  }
  required_version = ">= 1.1.1"
}
provider "google" {
  project = var.project_id
  region  = var.region
  credentials = var.GOOGLE_CREDENTIALS
}
