# ✂️ URL Shortener на Go

Простое и быстрое приложение для сокращения URL, написанное на Go с использованием фреймворка Gin, Redis для хранения данных и Testify для тестирования.

## 🚀 Возможности

- Сокращение длинных URL
- Перенаправление по сокращённой ссылке
- Хранение ссылок в Redis
- Генерация случайных хэшей
- API-интерфейс
- Тесты с использованием Testify

## 🛠️ Стек технологий

- **Язык:** Go
- **Веб-фреймворк:** [Gin](https://github.com/gin-gonic/gin)
- **БД:** [Redis](https://redis.io/)
- **Тестирование:** [Testify](https://github.com/stretchr/testify)

## 📦 Установка и запуск

### 1. Клонировать репозиторий

```bash
git clone https://github.com/miraklik/url-shortener.git
cd url-shortener
```

### 2. Установка зависимостей

```bash
go mod tidy
```
### 3. Запустить Redis

```bash
docker run -p 6379:6379 redis
```

### 4. Запуск сервера

```bash
go run main.go
```

