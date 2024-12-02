### 1. Аутентификация

#### 1.1 Регистрация пользователя

- **Метод**: `POST`
- **URL**: `/api/auth/registration`
- **Описание**: Создание новой учетной записи пользователя.

##### Тело запроса:
```json
{
  "first_name": "John",
  "middle_name": "A.",
  "last_name": "Doe",
  "role_id": 1,
  "group_id": 1,
  "login": "johndoe",
  "password": "securepassword123"
}
```

##### Ответы:
- **Успех**:  
  **Код**: `201 Created`  
  **Тело ответа**:

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ...",
    "user": {
    "id": 172,
    "first_name": "John",
    "middle_name": "A.",
    "last_name": "Doe",
    "role_id": 1,
    "group_id": 1,
    "login": "johndoe",
    "password": "",
    "salt": "",
    "token": "",
    "status": null,
    "role": "",
    "group": null
    }
}
```
  
- **Ошибка**:
    - **Код**: `400 Bad Request`  
      **Причина**: Некорректные данные в запросе.
    - **Код**: `409 Conflict`  
      **Причина**: Пользователь с таким `login` уже существует.

---

#### 1.2 Вход в систему

- **Метод**: `POST`
- **URL**: `/api/auth/login`
- **Описание**: Аутентификация пользователя и получение токена.

##### Тело запроса:
```json
{
  "login": "johndoe",
  "password": "securepassword123"
}
```

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNzAsInJvbGVfaWQiOjEsImlzcyI6ImV4YW1wbGVJc3N1ZXIiLCJleHAiOjE3MzM0MTc1MTF9.vhpuatTgrr8IvyWtieYNmDohxJvUN7cClzbf4P4F8_A",
  "user": {
    "id": 170,
    "first_name": "John",
    "middle_name": "A.",
    "last_name": "Doe",
    "role_id": 1,
    "group_id": 1,
    "login": "johndoe",
    "password": "vjM8/Njyu9M9xgcbfcjhMKiaipsqDJ7VVGvpPYGL7LA=",
    "salt": "PP50IBkBk5d5egK5wMOD4g==",
    "token": "",
    "status": null,
    "role": "",
    "group": null
  }
}
  ```
- **Ошибка**:
    - **Код**: `401 Unauthorized`  
      **Причина**: Неверный `login` или `password`.
    - 
---

#### 1.3 Верификация токена

- **Метод**: `POST`
- **URL**: `/api/auth/verify`
- **Описание**: Проверка валидности токена.

##### Тело запроса:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
"id": 170,
"role_id": 1
}
```

- **Ошибка**:
    - **Код**: `401 Unauthorized`  
      **Причина**: Неверный или истекший токен.

---

### 2. Пользователи

#### 2.1 Получение списка пользователей

- **Метод**: `GET`
- **URL**: `/api/user`
- **Описание**: Получение всех пользователей с возможностью фильтрации по роли.
- 
##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
  ```json
  [
    {
    "id": 1,
    "first_name": "StudentFirst1_Group1",
    "middle_name": "Middle1",
    "last_name": "Last1",
    "role_id": 1,
    "group_id": 1,
    "login": "student_1_1",
    "password": "hashed_password_1",
    "salt": "salt_1",
    "token": "",
    "status": null,
    "role": "Student",
    "group": "Group A"
    },
    {
      "id": 1,
      "first_name": "StudentFirst1_Group1",
      "middle_name": "Middle1",
      "last_name": "Last1",
      "role_id": 1,
      "group_id": 1,
      "login": "student_1_1",
      "password": "hashed_password_1",
      "salt": "salt_1",
      "token": "",
      "status": null,
      "role": "Student",
      "group": "Group A"
    }
  ]
  ```
- **Ошибка**:
    - **Код**: `500 Internal Server Error`  
      **Причина**: Проблема с базой данных.

---

#### 2.2 Удаление пользователя

- **Метод**: `DELETE`
- **URL**: `/api/user/:id`
- **Описание**: Удаление пользователя по его идентификатору.

##### Параметры:
- **id**: Идентификатор пользователя.

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
"status": "deleted"
}
```
- **Ошибка**:
    - **Код**: `404 Not Found`  
      **Причина**: Пользователь не найден.
    - **Код**: `500 Internal Server Error`  
      **Причина**: Проблема с базой данных.

---

### 3. Группы

#### 3.1 Получение списка групп

- **Метод**: `GET`
- **URL**: `/api/group`
- **Описание**: Получение списка групп.

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{ 
"groups": [
    {
      "id": 101,
      "name": "Group A"
    },
    {
      "id": 102,
      "name": "Group B"
    }
  ]
} 
```
- **Ошибка**:
    - **Код**: `500 Internal Server Error`  
      **Причина**: Проблема с базой данных.

---

#### 3.2 Создание новой группы

- **Метод**: `POST`
- **URL**: `/api/group`
- **Описание**: Добавление новой группы.

##### Тело запроса:
```json
{
  "name": "Group C"
}
```

##### Ответы:
- **Успех**:  
  **Код**: `201 Created`  
  **Тело ответа**:
  ```json
  {
    "id": 103,
    "name": "Group C"
  }
  ```
- **Ошибка**:
    - **Код**: `400 Bad Request`  
      **Причина**: Неправильные данные.


---

#### 3.3 Удаление группы

- **Метод**: `DELETE`
- **URL**: `/api/group/:id`
- **Описание**: Удаление группы по её идентификатору.

##### Параметры:
- **id**: Идентификатор группы.

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
"status": "deleted"
}
```
- **Ошибка**:
    - **Код**: `404 Not Found`  
      **Причина**: Группа не найдена.
    - **Код**: `500 Internal Server Error`
}