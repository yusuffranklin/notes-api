resource "google_sql_database_instance" "sandbox" {
  name             = var.db_instance_name
  region           = "asia-southeast2"
  database_version = "POSTGRES_17"

  settings {
    tier = "db-f1-micro"
    ip_configuration {
      ipv4_enabled    = true
      authorized_networks {
        name  = "allow-all"
        value = "0.0.0.0/0"
      }
    }
    location_preference {
      zone = "asia-southeast2-a"
    }
  }

  deletion_protection = false
}

resource "google_sql_database" "default" {
  name     = "notes_db"
  instance = google_sql_database_instance.sandbox.name
}

resource "google_sql_user" "default" {
  name     = var.db_user
  instance = google_sql_database_instance.sandbox.name
  password = var.db_password
}