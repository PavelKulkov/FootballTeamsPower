# FootballTeamsPower
Разработать сервис, определяющий положение сил двух сборных для ЧМ2018.
Сервису передаётся названия двух команд, в ответе должно приходить соотношение сил.
Сила команды - сумма сил игроков в этой команде.
Сила игрока считается количество матчей выигранных за свою команду умноженная на разницу в счёте.
Если игрок не играл в матче, этот матч не идёт в статистику игроку.

Из того что хотел сделать, но не успел или пока не понял:
  - Полное покрытие тестами
  - Расчет сил игроков и команд во время парсинга данных, а не при запросе
  - Обойти бан transferMarkt, парсинг получился бы намного быстрее
  - DI
  - Использование парсера через интерфейс на случай если будет еще парсер
