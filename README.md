# 🌤️ WeatherAPIService

[![Go Version](https://img.shields.io/badge/Go-1.25.1%2B-blue.svg)](https://golang.org)
[![HTTP Router](https://img.shields.io/badge/HTTP-Standard%20Library-green.svg)](https://pkg.go.dev/net/http)

---

## 📖 Оглавление

- [🌤️ WeatherAPIService](#️-weatherapiservice)
  - [📖 Оглавление](#-оглавление)
  - [🚀 О проекте](#-о-проекте)
  - [✨ Функциональность](#-функциональность)
  - [🛠 Технологии](#-технологии)
  - [⚡️ Быстрый старт](#️-быстрый-старт)
    - [Предварительные требования](#предварительные-требования)
    - [Установка и запуск](#установка-и-запуск)
  - [📡 API Endpoints](#-api-endpoints)
    - [Получение погоды](#получение-погоды)
  - [🔧 Конфигурация](#-конфигурация)

---

## 🚀 О проекте

WeatherAPIService — это RESTful сервис на Go, предоставляющий актуальную информацию о погоде для различных городов через интеграцию с OpenWeatherMap API.

## ✨ Функциональность

* **Получение текущей погоды** по названию города
* **Локализация** описаний погоды на русский язык
* **Простой REST API** с минимальными зависимостями

## 🛠 Технологии

* **Язык программирования:** [Go (Golang)](https://golang.org/)
* **HTTP фреймворк:** Стандартная библиотека `net/http`
* **Внешнее API:** [OpenWeatherMap](https://openweathermap.org/)
* **Конфигурация:** `.env` файлы

## ⚡️ Быстрый старт

### Предварительные требования

- Go 1.25.1 или выше
- API ключ от [OpenWeatherMap](https://openweathermap.org/api)

### Установка и запуск

1. **Клонируйте репозиторий:**
```bash
git clone https://github.com/ZeroZeroZerooZeroo/WeatherAPIService.git
cd WeatherAPIService
```

2. **Установите зависимости:**
```bash
go mod tidy
```

3. **Настройте окружение:**
Создайте файл `.env` в корне проекта:
```env
SERVER_PORT=8080
API_KEY=your_openweathermap_api_key_here
API_URL=https://api.openweathermap.org
```

4. **Запустите приложение:**
```bash
go run main.go
```

Сервис будет доступен по адресу: `http://localhost:8080`

## 📡 API Endpoints

### Получение погоды

**Запрос:**
```bash
curl "http://localhost:8080/weather?city=Москва"
```

**Ответ:**
```json
{
    "city": "Москва",
    "temperature": 15.5,
    "description": "легкая облачность"
}
```

## 🔧 Конфигурация

| Переменная | По умолчанию | Описание |
|------------|--------------|----------|
| `SERVER_PORT` | `8080` | Порт HTTP сервера |
| `API_KEY` | - | API ключ OpenWeatherMap |
| `API_URL` | `https://api.openweathermap.org` | Базовый URL API |

---

*Для получения API ключа зарегистрируйтесь на [OpenWeatherMap](https://openweathermap.org/api)*