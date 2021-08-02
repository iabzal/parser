# Mailer microservice

Микросервис для Email рассылки. Записывает EMail сообщения в БД, проверяет запланированные для отправки сообщения 
каждые **5** минут и отправляет. Проверяет статус отправки каждые **5** минут и изменяет его в БД.

**Для Amway** проверяет запланированные для отправки сообщения каждый день в **20:06** и если есть 
доставленные посылки, то отправляет сообщение с вложением файла.

Логи: **error.log**

Порт: **7001**

## Web server endpoints
## Create Mail

Создания Email сообщения для последующей отправки.

**URL** : `/email`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "subject": "[subject in plain text]:string",
    "text": "[message text in plain text]:string",
    "recipient": "[valid email]:array"
}
```

**Data example**

```json
{
    "subject": "Пример темы",
    "text": "Пример текста",
    "recipient": ["test@gmail.com", "test@yandex.ru"]
}
```

#### Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "clientMessageId": "f2f4584d-449f-42c5-8539-6aa5824b3a65"
}
```

#### Error Response

**Condition** : Если неправильно заполнены поля

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
    "message": "Error:Field validation for 'Subject' failed on the 'required' tag"
}
```

## Get sms info

Получение данных о сообщении (статус доставки).

**URL** : `/email/:client_message_id`

**Method** : `GET`

**Auth required** : NO

#### Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "id": 1,
    "subject": "Пример темы",
    "text": "Пример текста",
    "recipient": "test@gmail.com, test@yandex.ru",
    "filepath": "",
    "client_message_id": "56358329-cdc6-459f-b454-16878990252a",
    "status": "delivered",
    "created_at": "2019-12-05T17:00:00Z",
    "updated_at": "2019-12-05T17:05:45Z",
    "planned_sending_at": "2019-12-05T17:05:00Z"
}
```

#### Error Response
**Condition** : Если сообщение не найдено

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
    "message": "сообщение не найдено"
}
```
