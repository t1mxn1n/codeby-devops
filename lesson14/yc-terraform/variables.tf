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

variable "image_id" {
  description = "ID образа (например ubuntu-2004-lts)"
  type        = string
  default     = "fd86lblkag66uq48h63d" # Ubuntu 20.04 LTS
}

variable "public_key_path" {
  description = "Путь к публичному SSH ключу"
  type        = string
}

variable "private_key_path" {
  description = "Путь к приватному SSH ключу"
  type        = string
}
