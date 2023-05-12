# Описание работы сервиса

## Запуск сервиса

### Для запуска сервиса необходимо выполнить следующие действия:

*   Установить Postgre SQL и создать базу данных TEST.
*   Запустить локальный сервер на порту, указанном в config.json.
*   Запустить сервис командой go run cmd/main.go.

## Методы API
## Save currency rates for a specific date

### Метод: GET /currency/save/{date}

    Описание:
    Сохраняет курсы валют для указанной даты в локальную базу данных

### Параметры

* date (обязательный): Дата в формате dd.mm.yyyy.

### Ответы:

*    200: Курсы валют успешно сохранены.
*    400: Некорректный запрос.

## Code 
```
func (c *Controller) saveCurrencyDate(w http.ResponseWriter, r *http.Request) {
	date := mux.Vars(r)["date"]

	resp, err := http.Get(fmt.Sprintf(api, date))

	if err != nil {
		log.Printf("Error when making a request to the national bank's API: %v", err)
		http.Error(w, "Error when making a request to the national bank's API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}
	var rates model.Rates

	if err = xml.Unmarshal(body, &rates); err != nil {
		log.Printf("Error parsing response body: %v", err)
		http.Error(w, "Error parsing response body", http.StatusInternalServerError)
		return
	}

	go c.service.Currency.CreateCurrency(context.Background(), rates)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"success": true})

}
```

## Get currency rates for a specific date and code

### Метод: GET /currency/{date}/{code}

    Описание:
    Метод GET /currency/{date}/{code} позволяет получить курсы валют для указанной даты и кода валюты из базы данных TEST

### Параметры

* {date} - обязательный параметр, дата в формате "dd.MM.yyyy", для которой необходимо получить курсы валют.

* {code} - необязательный параметр, код валюты в формате ISO 4217, для которой необходимо получить курсы. Если параметр не указан, возвращаются курсы всех валют.

### Ответ

  В случае успешного выполнения запроса возвращается массив объектов с информацией о курсах валют для указанной даты и кода валюты. Каждый объект содержит следующие поля:

* date - дата, на которую указаны курсы валют.
* fullname - полное название валюты.
* title - код валюты в формате ISO 4217.
* description - курс валюты.

В случае ошибки возвращается JSON объект со статусом "error" и сообщением об ошибке.

### Code

```
func (c *Controller) currencyHandler(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)["date"]
	code := mux.Vars(r)["code"]

	var currency []model.Currency
	var err error

	if code != "" {
		currency, err = c.service.Currency.GetCurrencyByCode(context.Background(), data, code)
	} else {
		currency, err = c.service.Currency.GetCurrency(context.Background(), data)
	}

	if err != nil {
		log.Printf("Error getting currency: %v", err)
		http.Error(w, "Error getting currency", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currency)

}

```

### Пример запроса

### GET /currency/15.04.2021/USD

### Пример ответа:

    Успешный ответ (код 200):


```
[
{
"date": "15.04.2021",
"fullname": "Доллар США",
"title": "USD",
"description": 422.46
}
]
```
    Неуспешный ответ - неверный формат даты (код 400):

```
{
    "status": "error",
    "message": "Invalid date format"
}

```

 Неуспешный ответ - данные не найдены (код 404):

 ```
 {
    "status": "error",
    "message": "Data not found for the given date and code"
 }

 ```
 
 Неуспешный ответ - серверная ошибка (код 500):

 ```
 {
    "status": "error",
    "message": "Internal server error"
 }

 ```