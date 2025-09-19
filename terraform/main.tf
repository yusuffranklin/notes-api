terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 6.42"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = "asia-southeast2"
  zone    = "asia-southeast2-a"
}