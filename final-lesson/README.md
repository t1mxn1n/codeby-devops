Пример работы проекта продемонстрирован в папке `report`. Там же есть ссылка на google drive видео для наглядности.

## Инструкция для локального запуска

```bash
yc init
```

Переменные из конфига необходимо указать в файл `terraform/terraform.tfvars` по примеру в `terraform/terraform.tfvars.example`
```bash
yc list config
```

Настроить зеркало для YC к terraform
```bash
nano ~/.terraformrc
```

```js
provider_installation {
  network_mirror {
    url = "https://terraform-mirror.yandexcloud.net/"
    include = ["registry.terraform.io/*/*"]
  }
  direct {
    exclude = ["registry.terraform.io/*/*"]
  }
}

```

```bash
terraform init

terraform apply
```

После `apply` предоставится команда для `kubectl` с подключением к созданному кластеру.

После необходимо применить конфигурации из `k8s/`
```bash
kubectl apply -k ./k8s/postgres
kubectl apply -k ./k8s/api
kubectl apply -k ./k8s/monitoring --server-side=true
```

> Возможно потребуется авторизация в `ghcr.io`
