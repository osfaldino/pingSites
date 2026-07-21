# Это проект для параллельного пинга выбранных серверов

## Как Закинуть программу в терминал?

```bash
go build -o pingSites main.go # компилируем бинарник
sudo cp pingSites /usr/local/bin/
sudo cp sites.conf /usr/local/bin/
```

## Как изменить конфиг?
```bash
sudo nano /usr/local/bin/sites.conf
```
каждый URL должен быть на новой строке
