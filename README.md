# yak
helper for build and manage microservices

# yak

установка: склонировать проект, выполнить go install в корневой директории проекта

запуск генератора для консольной утилиты

`yak config-helper-gen -i test.toml -o src/conf/Helper.go`

-i|--input путь к файлу конфига *.toml

-o|--output путь для сохранения хелпера *.go

--pg10 true|false опциональный флаг для использования в импортах pg-go 10 версии

## Загрузка конфига и переопределение параметров через env:

при пустом 2м поле идёт в common.toml, иначе по заданным параметрам

ReadConfigFilesV3("./etc/", "", conf.MODULE, conf.Validate)

ReadConfigFilesV3("./etc/", "common2.toml", conf.MODULE, conf.Validate)

для переопределение переменных через env:

os.Setenv("TEST_PJ_INTERFACES_APIPORT", "12349876")

префикс указывается в конфиге в поле module, разделитель _

## Пример config файла

```toml
module = "test-pj"          # название модуля
[pg.default]                # pg - директива для определения коннекта к базе данных
host = "127.0.0.1:1234"     # поля ниже являются обязательными и не должны быть пустыми
user = "user"
pass = "password"
database = "my-db"
[pg.custom]                 # можно определить настройки для подключения к нескольким
host = "127.0.0.1:1234"     # по принципу [pg._NAME_]
user = "user"
pass = "password"
database = "my-db"


[mq.default]
host = "127.0.0.1"          # должен быть специальный метод для получения url rabbit:
port = 8765                 # conf.Mq().Default().Url() который вернет строку вида:
user = "user"               # ("amqp://%s:%s@%s:%s/%s", user, pass, host, port, vhost)
pass = "password"           # vhost опционален и может быть пустым, как и порт, если порт
vhost = "mts-money"         # не указан или zero-value берем стандартный 5672
[mq.default.queues]
proc-queue = "procQueue"    # список имен очередей: получаем conf.Mq().Default().Queues().ProcQueue()
[mq.default.exchanges]
fanout-exch = "fanoutExch"  # аналогично очередям

[ftp.default]                 # для ftp определены параметры
host = "127.0.0.1"          # обязательное поле
port = 666                   # остальные не обязательные - если port не указан или zero-value - берем стандартный 21
user = "user"
pass = "pass"
[ftp.default.params]
pull = "in"                 # просто стандартизируем какие-то параметры для ftp
push = "out"

[services]                    # директива где мы определяем grpc модули, с которыми взаимодействует модуль
serv-admitad = "admitad:80"   # по формату host:port , port можно опустить, тогда 80 по-умолчанию
                              # прошу называть директивы как сервисы в проекте serv-admitad - соответствует названию в репозитории  

[interfaces]
grpc = 80                   # стандартизируем порты, которые открывает сервис
json-rpc = 8080             # думаю, для первых 3х нужно стандартизировать имена,
ping = 9999                 # далее называем на свой вкус

[deps]                        # список зависимостей, с которыми взаимодействует сервис
notadmitad = "http://not-admitad.by"  # можно задавать так
[deps.admitad]
url = "http://admitad.by"   # а можно так, но с обязательным параметром URL
key = "wgwrh34yrhwerh"      # ниже могут быть дополнительные параметры для зависимости, ключи, данные авторизации и т.д.

[vars]
abc = 123                   # договоримся, что всякие дополнительные параметры для сервиса можно группировать в vars

[customDiv]
foo = "bar"                 # но можно создать свою группу для конфигурации какого-то модуля бизнеслогики или инфраструктурного слоя.
```

