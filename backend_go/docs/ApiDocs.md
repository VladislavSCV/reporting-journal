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
  {
    "users": [
        {
            "id": 3,
            "first_name": "Павел",
            "middle_name": "Алексеевич",
            "last_name": "Морозов",
            "role_id": 3,
            "group_id": 1,
            "login": "student1",
            "password": "hashed_password3",
            "salt": "random_salt3",
            "token": "",
            "status": null,
            "role": "Админ",
            "group": "Группа 101"
        },
        {
            "id": 4,
            "first_name": "Анна",
            "middle_name": "Игоревна",
            "last_name": "Смирнова",
            "role_id": 3,
            "group_id": 2,
            "login": "student2",
            "password": "hashed_password4",
            "salt": "random_salt4",
            "token": "",
            "status": null,
            "role": "Админ",
            "group": "Группа 102"
        },
        {
            "id": 2,
            "first_name": "Ольга",
            "middle_name": "Сергеевна",
            "last_name": "Кузнецова",
            "role_id": 2,
            "group_id": null,
            "login": "teacher1",
            "password": "hashed_password2",
            "salt": "random_salt2",
            "token": "",
            "status": null,
            "role": "Преподаватель",
            "group": "Не указана группа"
        },
        {
            "id": 1,
            "first_name": "Иван",
            "middle_name": "Иванович",
            "last_name": "Петров",
            "role_id": 1,
            "group_id": null,
            "login": "admin",
            "password": "hashed_password1",
            "salt": "random_salt1",
            "token": "",
            "status": null,
            "role": "Студент",
            "group": "Не указана группа"
        }
    ]
}
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
  "message": "Group deleted successfully"
}
```
- **Ошибка**:
    - **Код**: `404 Not Found`  
      **Причина**: Группа не найдена.
    - **Код**: `500 Internal Server Error`
}

---

#### 1. Роли

##### 1.1 Получение списка всех ролей
- **Метод**: `GET`
- **URL**: `/api/role/`
- **Описание**: Возвращает список всех ролей.

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
  "roles": [
    {
      "id": 1,
      "value": "Студент"
    },
    {
      "id": 2,
      "value": "Преподаватель"
    },
    {
      "id": 3,
      "value": "Админ"
    }
  ]
}
```

- **Ошибка**:
    - **Код**: `500 Internal Server Error`

---

##### 1.2 Получение роли по ID
- **Метод**: `GET`
- **URL**: `/api/role/:id`
- **Описание**: Возвращает данные роли по идентификатору.

##### Параметры:
- **id**: Идентификатор роли.

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
  "id": 2,
  "value": "Преподаватель"
}
```

- **Ошибка**:
    - **Код**: `404 Not Found`  
      **Причина**: Роль не найдена.
    - **Код**: `500 Internal Server Error`

---

##### 1.3 Создание роли
- **Метод**: `POST`
- **URL**: `/api/role/`
- **Описание**: Создает новую роль.

##### Тело запроса:
```json
{
  "name": "Student"
}
```

##### Ответы:
- **Успех**:  
  **Код**: `201 Created`  
  **Тело ответа**:
```json
{
"role": {
    "id": 1,
    "value": "test"
    }
}
```

- **Ошибка**:
    - **Код**: `400 Bad Request`  
      **Причина**: Неверные данные.
    - **Код**: `500 Internal Server Error`

---

##### 1.4 Удаление роли
- **Метод**: `DELETE`
- **URL**: `/api/role/:id`
- **Описание**: Удаляет роль по идентификатору.

##### Параметры:
- **id**: Идентификатор роли.

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
      **Причина**: Роль не найдена.
    - **Код**: `500 Internal Server Error`

---

#### 2. Расписание

##### 2.1 Получение расписания
- **Метод**: `GET`
- **URL**: `/api/schedule/:id`
- **Описание**: Возвращает расписание по идентификатору.

##### Параметры:
- **id**: Идентификатор расписания.

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
  "schedule": [
    {
      "ScheduleID": 3,
      "GroupName": "Группа 102",
      "DayOfWeek": 5,
      "SubjectName": "Программирование",
      "TeacherName": "Ольга Сергеевна Кузнецова",
      "Location": "Кабинет 303"
    }
  ]
}
```

- **Ошибка**:
    - **Код**: `404 Not Found`  
      **Причина**: Расписание не найдено.
    - **Код**: `500 Internal Server Error`

---

##### 2.2 Создание расписания
- **Метод**: `POST`
- **URL**: `/api/schedule/`
- **Описание**: Создает новое расписание.

##### Тело запроса:
```json
{
  "group_id": 2,
  "day_of_week": 4,
  "subject_id": 1,
  "teacher_id":1,
  "location": "Room 101"
}
```

##### Ответы:
- **Успех**:  
  **Код**: `201 Created`  
  **Тело ответа**:
```json
{
  "message": "Schedule created successfully"
}
```

- **Ошибка**:
    - **Код**: `400 Bad Request`  
      **Причина**: Неверные данные.
    - **Код**: `500 Internal Server Error`

---

##### 2.3 Удаление расписания
- **Метод**: `DELETE`
- **URL**: `/api/schedule/:id`
- **Описание**: Удаляет расписание по идентификатору.

##### Параметры:
- **id**: Идентификатор расписания.

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
  "message": "Schedule deleted successfully"
}
```

- **Ошибка**:
    - **Код**: `404 Not Found`  
      **Причина**: Расписание не найдено.
    - **Код**: `500 Internal Server Error`

---

#### 3. Предметы

##### 3.1 Получение списка предметов
- **Метод**: `GET`
- **URL**: `/api/subject/`
- **Описание**: Возвращает список всех предметов.

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
[
  {
    "id": 1,
    "name": "Математика"
  },
  {
    "id": 2,
    "name": "Физика"
  },
  {
    "id": 3,
    "name": "Программирование"
  }
]
```

- **Ошибка**:
    - **Код**: `500 Internal Server Error`

---

##### 3.2 Получение предмета по id
- **Метод**: `GET`
- **URL**: `/api/subject/:id`
- **Описание**: Возвращает предмет по id.

##### Параметры:
- **id**: Идентификатор расписания.

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
  "id": 2,
  "name": "Физика"
}
```

- **Ошибка**:
    - **Код**: `500 Internal Server Error`

---

##### 3.3 Создание предмета
- **Метод**: `POST`
- **URL**: `/api/subject/`
- **Описание**: Возвращает список всех предметов.

##### Тело запроса:
```json
{
  "name": "TEST"
}
```

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
  "message": "subject created successfully"
}
```

- **Ошибка**:
    - **Код**: `500 Internal Server Error`

---

##### 3.4 Обновление предмета
- **Метод**: `PUT`
- **URL**: `/api/subject/:id`
- **Описание**: Обновляет предмет.

##### Параметры:
- **id**: Идентификатор расписания.

##### Тело запроса:
```json
{
  "name": "TEwqST"
}
```

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
  "message": "subject updated successfully"
}
```

- **Ошибка**:
    - **Код**: `500 Internal Server Error`

---

##### 3.5 Удаляет предмет
- **Метод**: `DELETE`
- **URL**: `/api/subject/:id`
- **Описание**: Удаляет предмет

##### Параметры:
- **id**: Идентификатор расписания.

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
  "message": "subject deleted successfully"
}
```

- **Ошибка**:
    - **Код**: `500 Internal Server Error`

---


#### 4. Администраторы

##### 4.1 Получение данных для панели администратора
- **Метод**: `GET`
- **URL**: `/api/admin/AdminPanel`
- **Описание**: Возвращает данные для панели администратора.

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
  "users": 120,
  "groups": 10,
  "roles": 3
}
```

- **Ошибка**:
    - **Код**: `500 Internal Server Error`

---

### 5. Преподаватели

#### 5.1 Получение списка групп для выставления оценок
- **Метод**: `GET`
- **URL**: `/api/teacher/groups`
- **Описание**: Возвращает список групп, где преподаватель должен выставить оценки.

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
  "groups": "W3siaWQiOjEsIm5hbWUiOiLQk9GA0YPQv9C/0LAgMTAxIn0seyJpZCI6MiwibmFtZSI6ItCT0YDRg9C/0L/QsCAxMDIifV0="
}
```

- **Ошибка**:
    - **Код**: `401 Unauthorized` — пользователь не авторизован.
    - **Код**: `500 Internal Server Error` — ошибка сервера.

---

#### 5.2 Получение посещаемости студентов
- **Метод**: `GET`
- **URL**: `/api/teacher/studentAttendance/:id`
- **Описание**: Возвращает список студентов с их посещаемостью и оценками по заданной группе.

##### Параметры:
- **URL-параметр**:
    - `id` (integer) — ID группы, для которой нужно получить данные о студентах.

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
  "users": "bnVsbA=="
}
```

- **Ошибка**:
    - **Код**: `400 Bad Request` — неверный ID группы.
    - **Код**: `401 Unauthorized` — пользователь не авторизован.
    - **Код**: `404 Not Found` — группа с указанным ID не найдена.
    - **Код**: `500 Internal Server Error` — ошибка сервера.

---

#### 5.3 Обновление посещаемости студентов
- **Метод**: `POST`
- **URL**: `/api/teacher/studentAttendance`
- **Описание**: Обновляет данные о посещаемости и оценках студентов.

##### Тело запроса:
```json
{
  "student_id": 4,
  "status": "Болеет"
}
```

##### Ответы:
- **Успех**:  
  **Код**: `200 OK`  
  **Тело ответа**:
```json
{
  "message": "Attendance updated successfully"
}
```

- **Ошибка**:
    - **Код**: `400 Bad Request` — некорректные данные в запросе.
    - **Код**: `401 Unauthorized` — пользователь не авторизован.
    - **Код**: `500 Internal Server Error` — ошибка сервера.

---