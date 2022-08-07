resource "google_container_cluster" "primary" {
  provider = google

  name     = "primary"
  location = var.region

  network    = var.vpc
  subnetwork = var.subnet_id

  private_cluster_config {
    enable_private_endpoint = false
    enable_private_nodes    = false
  }

  maintenance_policy {
    recurring_window {
      start_time = "2021-06-18T00:00:00Z"
      end_time   = "2050-01-01T04:00:00Z"
      recurrence = "FREQ=WEEKLY;BYDAY=MO,TU,WE,TH"
    }
  }

  enable_autopilot = true

  release_channel {
    channel = "REGULAR"
  }
}
