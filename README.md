Rest api server - управление балансом пользователей:

Swagger:
http://localhost:8000/swagger/index.html

База данных postgres (доступ через .env)
Два пользователя созданы:
Andey пароль Andey
INSERT INTO avito.users
(username, "password", "money")
VALUES('Andrey', '8e756c9f2b15da6a63f84852fc39667617523133', 0.0);
INSERT INTO avito.users
(username, "password", "money")
VALUES('Anton', '8e756c9f2b15da6a63f84852fc39667617523134', 0.0);

Зачисление средств,  (Принимает id пользователя и сколько средств зачислить.)+ (POST addMoneyHandler)

Списание средств, (Принимает id пользователя и сколько средств списать.)+  (POST reduceUserHandler)

Перевод средств от пользователя к пользователю(Принимает id пользователя с которого нужно списать средства, id пользователя которому должны зачислить средства, а также сумму.)+ (POST transferMoneyHandler)

Получения баланса пользователя (Принимает id пользователя. Баланс всегда в рублях.) + (POST getMoneyUserHandler)

Show the status of server (GET heathHandler)

Запуск приложения с помощью: docker-compose up -build
Проверить запросы можно через swagger

Доп. задание 1: (API)

Необходима возможность вывода баланса пользователя в отличной от рубля валюте.
curl -d '{"userid": 1}' -H "Content-Type: application/json" -X POST http://localhost:8000/api/money?currency=USD

Курсу валют беру отсюда https://freecurrencyapi.net/.


Доп. задание 2:(параметр старницы и сортировку по сумму и дате)

curl -d '{"userid": 1}' -H "Content-Type: application/json" -X POST http://localhost:8000/api/getmovemoney?page=1&filtermoney=asc&filtertime=desc

