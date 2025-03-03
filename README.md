# Yandex Cloud Lockbox Secret Fetcher

Go utility для безопасного получения секретов из Yandex Cloud Lockbox с поддержкой различных методов аутентификации.

## Особенности
- Получение секретов через метаданные ВМ Yandex Cloud
- Поддержка указания версии секрета
- Автоматическая аутентификация через сервисный аккаунт ВМ
- Поддержка Kubernetes через Workload Identity Federation
- Человеко-читаемый вывод секретов

## Предварительные требования

### Для виртуальных машин
- Аккаунт Yandex Cloud
- ВМ в Yandex Cloud с привязанным сервисным аккаунтом
- У сервисного аккаунта должна быть роль `lockbox.payloadViewer`

### Для Kubernetes
- Кластер Kubernetes с настроенной Workload Identity Federation
- Сервисный аккаунт с ролью `lockbox.payloadViewer`
- Настроенный ServiceAccount в Kubernetes

## Установка
```bash
# Сборка бинарного файла
go build -o lockbox-fetcher main.go
```

## Использование

### Виртуальные машины
```bash
# Базовое использование (последняя версия секрета)
./lockbox-fetcher --secret-id=<your-secret-id>

# С указанием версии
./lockbox-fetcher --secret-id=<your-secret-id> --version-id=<version-id>
```

### Kubernetes
Для использования в Kubernetes необходимо настроить Pod с соответствующими параметрами:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: lockbox-fetcher-pod
spec:
  containers:
  - name: lockbox-fetcher
    image: <your-image>
    env:
    - name: SECRET_ID
      value: "<your-secret-id>"
    volumeMounts:
    - name: sa-token
      mountPath: /var/run/secrets/tokens
  serviceAccountName: <your-k8s-sa>
  volumes:
  - name: sa-token
    projected:
      sources:
      - serviceAccountToken:
          path: sa-token
          expirationSeconds: 7200
          audience: <your-audience>
```

Пример вывода:
```
Secret payload:
Version ID: v1
Entries:
  Key: DB_PASSWORD, Value: s3cr3tP@ss
  Key: API_KEY, Value: a1b2c3d4
```

## Настройка прав доступа
- Для ВМ: https://cloud.yandex.ru/ru/docs/compute/operations/vm-create/create-with-lockbox-secret
- Для Kubernetes: https://cloud.yandex.ru/ru/docs/managed-kubernetes/operations/applications/workload-identity

## Безопасность
- Секреты никогда не сохраняются на диск
- Аутентификация выполняется через защищенные механизмы
- Поддержка современных методов аутентификации (OIDC, JWT)
- Рекомендуется использовать на выделенных ВМ или в изолированных Pod с ограниченным доступом

## Разработка
Требования:
- Go 1.20+
- Yandex Cloud SDK для Go

Установка зависимостей:
```bash
go mod tidy
```

## Примеры интеграции
В репозитории представлены примеры использования:
- `main.go` - базовый пример для ВМ
- `wi-token-exchange-simple-pod.yaml` - пример для Kubernetes с Workload Identity

## Лицензия
Apache 2.0
