## Методы аутентификации

В репозитории реализованы два подхода к получению токенов для доступа к секретам:

1. **Через метаданные ВМ** (`main.go`):
   - Автоматическое получение токена через Instance Service Account
   - Использует SDK Yandex Cloud для Go
   - Простая реализация без необходимости ручного обмена токенов
   - Подходит для виртуальных машин в Yandex Cloud

2. **Через прямой обмен токена** (`wi-token-exchange-simple-pod.yaml`):
   - Получение JWT токена из Kubernetes ServiceAccount
   - Обмен JWT токена на IAM токен через API Yandex Cloud
   - Ручная реализация HTTP запросов для обмена токенов
   - Подходит для Kubernetes с настроенной Workload Identity Federation

## Примеры интеграции
В репозитории представлены примеры использования:
- `main.go` - базовый пример для ВМ с использованием метаданных
- `wi-token-exchange-simple-pod.yaml` - пример для Kubernetes с прямым обменом токенов
