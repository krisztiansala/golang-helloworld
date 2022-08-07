variable "project_id" {
  default     = ""
  description = "project id"
}

variable "vpc_name" {
  default     = ""
  description = "VPC name"
}

variable "region" {
  default     = ""
  description = "region"
}

variable "gke_username" {
  default     = ""
  description = "gke username"
}

variable "gke_password" {
  default     = ""
  description = "gke password"
}

variable "GOOGLE_CREDENTIALS" {
  type = string
}