# avito-test-go

Файл с описанием базы данных лежит в `scripts/database.sql`

**Метод получения текущего баланса пользователя**

GET `/api/balance`

Тело запроса:
```
{
  "user_id":1         // id пользователя, чей баланс нужно получить
}
```

**Метод начисления/списания средств**

POST `/api/balance`

Тело запроса для пополнения:
```
{
  "user_id":1,        // id пользователя, которому нужно зачислить/списать средства
  "amount":10,        // количество средств для пополнения/списания
  "type":"add"        // "add" - пополнение, "subtract" - списание
}
```

 
**Метод перевода средств от пользователя к пользователю**

POST `/api/p2p`

Тело запроса:
```
{
  "from_user_id":1,   // id пользователя, который переводит средства
  "to_user_id":2,     // id пользователя, которому переводят средства
  "amount":10         // количество средств для перевода
}
```
