terraform {
  required_providers {
    yandex = {
      source  = "yandex-cloud/yandex"
      version = "~> 0.117"
    }
  }
  required_version = ">= 1.5.0"
}

provider "yandex" {
  token     = var.yc_token
  cloud_id  = var.yc_cloud_id
  folder_id = var.yc_folder_id
  zone      = var.yc_zone
}

# ----------- СЕТЬ -----------
resource "yandex_vpc_network" "main" {
  name = "demo-network"
}

resource "yandex_vpc_subnet" "subnet_a" {
  name           = "demo-subnet"
  zone           = var.yc_zone
  network_id     = yandex_vpc_network.main.id
  v4_cidr_blocks = ["10.10.0.0/24"]
}

# ----------- СЕРВИСНЫЙ АККАУНТ -----------
resource "yandex_iam_service_account" "k8s_sa" {
  name = "k8s-service-account"
}

resource "yandex_resourcemanager_folder_iam_member" "k8s_sa_roles" {
  folder_id = var.yc_folder_id
  role      = "editor"
  member    = "serviceAccount:${yandex_iam_service_account.k8s_sa.id}"
}

# ----------- КЛАСТЕР -----------
resource "yandex_kubernetes_cluster" "demo_cluster" {
  name        = "demo-cluster"
  description = "Demo cluster created by Terraform"

  network_id = yandex_vpc_network.main.id

  master {
    version = "1.29"
    zonal {
      zone      = var.yc_zone
      subnet_id = yandex_vpc_subnet.subnet_a.id
    }
    public_ip = true
  }

  service_account_id = yandex_iam_service_account.k8s_sa.id
}

# ----------- NODE GROUP -----------
resource "yandex_kubernetes_node_group" "demo_nodes" {
  cluster_id  = yandex_kubernetes_cluster.demo_cluster.id
  name        = "demo-nodes"
  description = "Node group for codeby final lesson"
  version     = "1.29"

  instance_template {
    platform_id = "standard-v3"
    resources {
      cores  = 2
      memory = 2
    }
    boot_disk {
      size = 10
      type = "network-ssd"
    }
    network_interface {
      subnet_ids = [yandex_vpc_subnet.subnet_a.id]
      nat        = true
    }
  }

  scale_policy {
    fixed_scale {
      size = 2
    }
  }

  allocation_policy {
    location {
      zone = var.yc_zone
    }
  }

  service_account_id = yandex_iam_service_account.k8s_sa.id
}
