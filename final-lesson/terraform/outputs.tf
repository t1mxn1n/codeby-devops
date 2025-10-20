output "cluster_id" {
  value = yandex_kubernetes_cluster.demo_cluster.id
}

output "kubeconfig_command" {
  value = "yc managed-kubernetes cluster get-credentials ${yandex_kubernetes_cluster.demo_cluster.name} --external --force"
}
