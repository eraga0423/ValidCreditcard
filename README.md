Управление номерами кредитных карт
Программа для работы с номерами кредитных карт: проверка корректности, генерация, получение информации и выпуск.

Установка
Убедитесь, что установлен Go.

Склонируйте репозиторий.

Соберите программу:

bash
Копировать код
go build -o creditcard .
Использование
Валидация
Проверка номеров карт:

bash
Копировать код
./creditcard validate "4400430180300003"
# Вывод: OK

echo "4400430180300003" | ./creditcard validate --stdin
# Вывод: OK
Генерация
Генерация номеров по шаблону:

bash
Копировать код
./creditcard generate "440043018030****"
# Вывод: 4400430180300003, 4400430180300011...

./creditcard generate --pick "440043018030****"
# Вывод: Случайный номер карты
Информация
Получение информации о карте:

bash
Копировать код
./creditcard information --brands=brands.txt --issuers=issuers.txt "4400430180300003"
# Вывод: Информация о карте
Выпуск
Выпуск номера карты:

bash
Копировать код
./creditcard issue --brands=brands.txt --issuers=issuers.txt --brand=VISA --issuer="Kaspi Gold"
# Вывод: Новый номер карты