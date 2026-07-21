# Это проект для параллельного пинга выбранных серверов

## Как Закинуть программу в терминал?

```bash
go build -o pingSites main.go # компилируем бинарник
sudo cp pingSites /usr/local/bin/ # кидаем в bin программу
sudo cp sites.conf /usr/local/bin/ # кидаем в bin конфиг
```

## Как изменить конфиг?
```bash
sudo nano /usr/local/bin/sites.conf # изменяем конфиг
```
каждый URL должен быть на новой строке
