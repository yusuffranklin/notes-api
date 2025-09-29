# google_client_config and kubernetes provider must be explicitly specified like the following.
data "google_client_config" "default" {}

module "gke" {
    depends_on = [ google_compute_subnetwork.network_with_secondary_ip_ranges ]
  source                     = "terraform-google-modules/kubernetes-engine/google"
  project_id                 = var.project_id
  name                       = "gke-test-1"
  region                     = "asia-southeast2"
  zones                      = ["asia-southeast2-a"]
  network                    = "default"
  subnetwork                 = "default"
  ip_range_pods              = google_compute_subnetwork.network_with_secondary_ip_ranges.secondary_ip_range[0].range_name
  http_load_balancing        = true
  network_policy             = false
  horizontal_pod_autoscaling = false
  filestore_csi_driver       = false
  dns_cache                  = false
  deletion_protection        = false

  node_pools = [
    {
      name                        = "default-node-pool"
      machine_type                = "e2-medium"
      node_locations              = "asia-southeast2-a"
      min_count                   = 3
      max_count                   = 3
      spot                        = false
      disk_size_gb                = 60
      disk_type                   = "pd-standard"
      image_type                  = "COS_CONTAINERD"
      preemptible                 = true
    },
  ]

  node_pools_oauth_scopes = {
    all = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }

  node_pools_labels = {
    all = {}

    default-node-pool = {
      default-node-pool = true
    }
  }

  node_pools_metadata = {
    all = {}

    default-node-pool = {
      node-pool-metadata-custom-value = "my-node-pool"
    }
  }

  node_pools_taints = {
    all = []

    default-node-pool = [
      {
        key    = "default-node-pool"
        value  = true
        effect = "PREFER_NO_SCHEDULE"
      },
    ]
  }

  node_pools_tags = {
    all = []

    default-node-pool = [
      "default-node-pool",
    ]
  }
}