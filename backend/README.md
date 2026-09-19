# Онлайн-сервис для школы Skadi

## Backend

Бэкенд-часть онлайн-сервиса для школы Skadi.

### Конфигурация

#### Насторойка yaml-конфига

Пример `yaml`-конфига приведён в корне проекта.
После сборки этот файл должен лежать рядом с бинарником.

#### Насторойка переменных окружения

Пример переменных окружения:

```dotenv
# dev
ACCESS_SECRET="example-access-secret"
REFRESH_SECRET="example-refresh-secret"
DB_PASSWORD="test_password"

# test
TEST_DSN="test_user:test_password@tcp(127.0.0.1:3306)/meteo_ssc_ras?parseTime=true&timeout=10s"
```

### Тестовый стенд (для frontend'а)

Пошаговая инструкция, как развернуть тестовый стенд для frontend'а.

#### 1. Клонировать репозиторий и перейти в корень проекта

```shell
git clone https://github.com/echo-tokyo/skadi
cd ./skadi
```

> ! _Далее все действия проводятся из корневой директории проекта_ !

#### 2. Скопировать конфигурационные файлы

В каталоге `./test-stand/frontend` лежат конфиг-файлы для тестового стенда для frontend'а.
Нужно скопировать содержимое файлов:

1. `./test-stand/frontend/backend.env` в файл `./backend/.env`
2. `./test-stand/frontend/mysql.env` в файл `./mysql/.env`
3. `./test-stand/frontend/backend-config.yml` в файл `./backend/config.yml`

Команды для `linux`

```shell
mv ./backend/.env ./backend/.env.backup
mv ./mysql/.env ./mysql/.env.backup
mv ./backend/config.yml ./backend/config.yml.backup

cp ./test-stand/frontend/backend.env ./backend/.env
cp ./test-stand/frontend/mysql.env ./mysql/.env
cp ./test-stand/frontend/backend-config.yml ./backend/config.yml
```

Чтобы востановить старые файлы используйте (linux):

```shell
mv ./backend/.env.backup ./backend/.env
mv ./mysql/.env.backup ./mysql/.env
mv ./backend/config.yml.backup ./backend/config.yml

```

#### 3. Запустить `compose` файл для frontend'а

```shell
docker compose -f dev.front.docker-compose.yml up -d --build
```

Проверить, что все сервисы стартовали, можно командой:

```shell
docker compose -f dev.front.docker-compose.yml ps
```

#### 4. Провести миграции БД

```shell
docker exec -it skadi_backend /bin/sh -c "/app/migrator up"
```

#### 5. Создать первого юзера

Эта команда запускает CLI-менеджер.

```shell
docker exec -it skadi_backend /bin/sh -c "/app/app manager"
```

После его запуска предложится список доступных действий.
Необходимо выбрать `create-admin` для создания нового юзера-админа.
Следуя инструкциям, создайте пользователя.

#### 6. Готово

Backend доступен по [адресу](http://127.0.0.1:8000/api/v1).
Swagger документация доступна по [адресу](http://127.0.0.1:8000/api/v1/docs).

> Чтобы опустить стенд, используйте:

```shell
docker compose -f dev.front.docker-compose.yml down
```

> Чтобы снести БД и создать заново, используйте:
>
> \* После первой команды будет запрошен пароль от БД. Его можно найти в файле `./mysql/.env` относительно корня проекта.

```shell
docker compose -f dev.front.docker-compose.yml exec mysql mysql -u root -p -e "DROP DATABASE skadi; CREATE DATABASE skadi;"
docker exec -it skadi_backend /bin/sh -c "/app/migrator up"
```

### Особенности работы программы

Время везде использует часовой пояс `UTC`. Это единый часовой пояс.
Предполагается перевод времени в локальный(ые) часовой(ые) пояс(а) на клиенте.
