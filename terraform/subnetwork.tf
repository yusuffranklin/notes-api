data "google_compute_network" "default" {
    name = "default"
    project = var.project_id
}

data "google_compute_subnetwork" "default" {
    name = "default"
    project = var.project_id
}

resource "google_compute_subnetwork" "network_with_secondary_ip_ranges" {
  name          = "default"
  ip_cidr_range = "10.184.0.0/20"
  region        = "asia-southeast2"
  network       = data.google_compute_network.default.id
  secondary_ip_range {
    range_name    = "pods-ip"
    ip_cidr_range = "192.168.0.0/16"
  }

  lifecycle {
    prevent_destroy = true
  }
}

import {
    to = google_compute_subnetwork.network_with_secondary_ip_ranges
    id = "projects/${var.project_id}/regions/asia-southeast2/subnetworks/default"
}