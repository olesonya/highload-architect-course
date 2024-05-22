# highload-architect-course.homework.01

## описание api

Релизован вариант api на основе контракта [user.proto](./api/grpc/user/v1/user.proto)

## запуск сервиса и базы данных

Контейнер с сервисом собирается при запуске docker-compose.yml

```bash
docker-compose up -d ; docker-compose logs -f
```

## использование grpc api

Утилита для проверки grpc api: [grpcurl](https://github.com/fullstorydev/grpcurl)

Примеры запросов к api:

```bash
# регистрация пользователя
grpcurl -plaintext -d '{"user_info":{"first_name":"Ivan","second_name":"Ivanov","birthdate":"22.02.08","biography":"just cool mate","city":"Moscow"},"user_pass":"12345"}' 127.0.0.1:50051 user.v1.UserService.Register
{
  "userId": "391609da-17bb-11ef-853b-1c1bb59facef"
}

# получение данных о пользователе
grpcurl -plaintext -d '{"user_id":"391609da-17bb-11ef-853b-1c1bb59facef"}' 127.0.0.1:50051 user.v1.UserService.Get
{
  "user": {
    "userId": "391609da-17bb-11ef-853b-1c1bb59facef",
    "userInfo": {
      "firstName": "Ivan",
      "secondName": "Ivanov",
      "birthdate": "22.02.08",
      "biography": "just cool mate",
      "city": "Moscow"
    }
  }
}

# залогинивание и получение токена
grpcurl -plaintext -d '{"user_id":"391609da-17bb-11ef-853b-1c1bb59facef","user_pass":"12345"}' 127.0.0.1:50051 user.v1.UserService.Login
{
  "token": "5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5"
}
```
