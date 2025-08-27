variable "yc_token" {
  description = "IAM-токен Yandex Cloud"
  type        = string
}

variable "yc_cloud_id" {
  description = "Cloud ID"
  type        = string
}

variable "yc_folder_id" {
  description = "Folder ID"
  type        = string
}

variable "yc_zone" {
  description = "Зона"
  type        = string
  default     = "ru-central1-d"
}