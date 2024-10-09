# Mon

Система мониторинга

## Установка и запуск

``` bash
git clone https://github.com/farukshin/mon.git
cd mon
go build .
./mon start
```

``` bash
curl localhost:1616
```

## Добавление сенсора

Через командную строку

``` bash
./mon sensors list
./mon sensors add --type=httpcode --target=google.com
# > 27e92b12-0933-4b82-b2b9-96c1b64745a2
```

Через API

``` bash
MON_SRV="localhost:1616" # сервер и порт с запущенной системой мониторинга
curl -X POST $MON_SRV/api/sensors/add \
    -d '{"kind":"httpcode", "target":"google.com", "time":60}'
# > {ok:true, uid:"27e92b12-0933-4b82-b2b9-96c1b64745a2"}
```

## Изменение сенсора

Через командную строку

``` bash
./mon sensors list
./mon sensors edit --uid=27e92b12-0933-4b82-b2b9-96c1b64745a2 --target=farukshin.com
# > 27e92b12-0933-4b82-b2b9-96c1b64745a2 изменено значение target, было google.com, стало farukshin.com
```
Через API

``` bash
MON_SRV="localhost:1616"
curl -X POST $MON_SRV/api/sensors/edit \
    -d '{"uid":"7e92b12-0933-4b82-b2b9-96c1b64745a2", "target":"farukshin.com"}'
# > {ok:true, uid:"27e92b12-0933-4b82-b2b9-96c1b64745a2"}
```

## Удаление сенсора

Через командную строку

``` bash
./mon sensors list
./mon sensors delete --uid=27e92b12-0933-4b82-b2b9-96c1b64745a2
# > success
```

Через API

``` bash
MON_SRV="localhost:1616"
curl -X POST $MON_SRV/api/sensors/delete \
    -d '{"uid":"7e92b12-0933-4b82-b2b9-96c1b64745a2"}'
# > {ok:true}
```