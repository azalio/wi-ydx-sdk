# Yandex Cloud Lockbox Secret Fetcher

Go utility для безопасного получения секретов из Yandex Cloud Lockbox через Instance Service Account

## Особенности
- Получение секретов через метаданные ВМ Yandex Cloud
- Поддержка указания версии секрета
- Автоматическая аутентификация через сервисный аккаунт ВМ
- Человеко-читаемый вывод секретов

## Предварительные требования
- Аккаунт Yandex Cloud
- ВМ в Yandex Cloud с привязанным сервисным аккаунтом
- У сервисного аккаунта должна быть роль `lockbox.payloadViewer`

## Установка
```bash
# Сборка бинарного файла
go build -o lockbox-fetcher main.go
```

## Использование
```bash
# Базовое использование (последняя версия секрета)
./lockbox-fetcher --secret-id=<your-secret-id>

# С указанием версии
./lockbox-fetcher --secret-id=<your-secret-id> --version-id=<version-id>
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
https://yandex.cloud/ru/docs/compute/operations/vm-create/create-with-lockbox-secret

## Безопасность
- Секреты никогда не сохраняются на диск
- Аутентификация выполняется через защищенный сервис метаданных
- Рекомендуется использовать на выделенных ВМ с ограниченным доступом

## Разработка
Требования:
- Go 1.20+
- Yandex Cloud SDK для Go

Установка зависимостей:
```bash
go mod tidy
```

## Лицензия
Apache 2.0
