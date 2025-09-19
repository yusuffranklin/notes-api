output "secondary_ip_range" {
    value = google_compute_subnetwork.network_with_secondary_ip_ranges.secondary_ip_range
}

output "gke_id" {
    value = module.gke.cluster_id
}

output "gke_name" {
    value = module.gke.name
}

output "cloudsql_id" {
    value = google_sql_database_instance.sandbox.id
}

output "cloudsql_name" {
    value = google_sql_database_instance.sandbox.name
}