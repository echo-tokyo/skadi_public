# Skadi

Веб-платформа для автоматизации учебного процесса IT-Школы «СКАДИ», обеспечивающей удобное взаимодействие администраторов, преподавателей и учеников при выдаче, выполнении и проверке заданий.

## Конфигурация

Основные настройки сервера лежат в файле [config.yml](./backend/config.yml)

Секреты хранятся в переменных окружения:

```dotenv
ACCESS_SECRET="example-access-secret"
REFRESH_SECRET="example-refresh-secret"
DB_PASSWORD="test_password"

MYSQL_ROOT_PASSWORD="p@ssW0rd"
MYSQL_DATABASE="skadi"
```

### Деплой

#### 1. Запустить `compose` файл

```shell
docker compose -f docker-compose.yml up -d --build
```

#### 2. Проверить запущенные сервисы

```shell
docker compose -f docker-compose.yml ps
```

#### 3. Провести миграции БД

```shell
docker exec -it skadi_backend /bin/sh -c "/app/migrator up"
```

#### 4. Создать админа

Эта команда запускает CLI-менеджер.

```shell
docker exec -it skadi_backend /bin/sh -c "/app/app manager"
```

После его запуска предложится список доступных действий.
Необходимо выбрать `create-admin` для создания нового юзера-админа.
Следуя инструкциям, создайте пользователя.

#### 5. Готово

Чтобы опустить сервисы, используйте:

```shell
docker compose -f docker-compose.yml down
```
