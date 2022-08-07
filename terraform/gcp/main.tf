module "gke-autopilot" {
  source                   = "./modules/compute/gke-autopilot"
  region                   = var.region
  vpc                      = google_compute_network.vpc.name
  subnet_id                = google_compute_subnetwork.subnet.id
}

resource "google_service_account" "github-actions" {
  account_id   = "github-actions"
  display_name = "Service account for deploying to the GKE cluster from Github Actions"
}

resource "google_project_iam_member" "github-actions-gke" {
  project = var.project_id
  role    = "roles/container.admin"
  member  = "serviceAccount:${google_service_account.github-actions.email}"
}
