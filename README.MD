# inner-shadowsocks(fork) for russian language #

![img](https://i.v2ex.co/5iP665B4.png)

[Русский](#Русский) [Особенности](#Особенности) [Использование](#Использование)

### Русский

Это простое приложение как `ss-local` на Golang. Он предназначен для работы на сервере и подключение его через "shadowsocks", тогда вы можете просто установить `SOCKS5` прокси в Chroimium, Telegram и т. д. без запуска клиента на вашем ноутбуке или сотовом телефоне.

Существует версия этого (Форк!!)приложения на python: [shadowsocks-with-socks-auth](https://github.com/ihciah/shadowsocks-with-socks-auth), которая представляет собой небольшую модифицированную версию оригинального shadowsocks с добавлением ` Поддержка SOCKS5`. Эта версия Go более эффективна и поддерживает баланс нескольких серверов.

-----

#### Особенности:

- `SOCKS5` аутентификация
- multi-server для баланса нагрузки

#### Установка:

Путь №1: Загрузить бинарный файл [Releases](https://github.com/Stezkoy/inner-shadowsocks/releases).Установка  (Go)Golang не требуется.

Путь №2:

```shell
go get -u -v github.com/Stezkoy/inner-shadowsocks/inner-ss
```



#### Использование:

```shell
inner-ss -c config.json
```

```Systemd
systemctl enable inner-ss@config
```

- -c: путь к файлу конфигурации(по умолчанию: `config.json`)
- -v: подробный режим

Настройте файл config.json:


```json
{
  "listen": "127.0.0.1",
  "port": 9980,
  "auth": true,
  "username": "Stezkoy",
  "password": "123456",
  "servers": ["ss://aes-256-cfb:pass1@ip1:port1", "ss://RC4-MD5:pass2@ip2:port2"],
  "maxfail": 3,
  "recovertime": 600,
  "starttimeout": 5,
  "remotetimeout": 60,
  "insidetimeout": 60,
  "whitelistenable": false,
  "whitelistdomains": [".github.com"],
  "whitelistips": ["::/0", "0.0.0.0/0"]
}
```


`auth`, `username`, `password`, `maxfail`, `recovertime`, `remotetimeout`, `insidetimeout`, `starttimeout` являются необязательными, значение по умолчанию которых `false`, `""`, `""`, `10`, `600`, `60`, `60`, `8` .

Также,  `whitelistenable` - `false` по умолчанию. Если вы хотите ограничить прокси только некоторым диапазоном, вы можете использовать `whitelistdomains` and `whitelistips` .

IP должен соответствовать формату CIDR, например `100.110.120.130/32`. `::/0` означает, что все IPv6, `0.0.0.0/0` означает, что все IPv4. Домены будут сопоставляться с суффиксом. Любые домены, суффикс которых совпадает с `whitelistdomains`, могут пройти. Например, `www.github.com` будет соответствовать` .github.com`.

Если вы установите для `auth` значение` true`, вы должны указать `username` и` password`. 

Вы должны перечислить свои серверы по порядку. Если один из них не может быть подключен в течение "maxfail" раз, он будет помечен как битый сервер до "recovertime" секунд спустя. Пожалуйста, обратите внимание, что вы должны убедиться, что пароль и метод шифрования верны. Поскольку мы не можем отличить вредоносные запросы пользователей и неверный пароль серверов shadowsocks, здесь мы оцениваем состояние сервера, если можем подключать его.

`remotetimeout` - это тайм-аут, соединяющий удаленные серверы, а `insidetimeout` - это клиент. `starttimeout` - это время ожидания первого сеанса связи. Короткое `starttimeout` может уменьшить время ожидания ложного пакета, длинные `remotetimeout` и `insidetimeout` могут позволить длинное соединение, например, уменьшить повторное соединение Telegram.

Вы можете использовать `systemd` для запуска демона. Пример настройки находится внизу.

Systemd Service 
```
[Unit]
Description=Inner-SS Client Service
After=network.target

[Service]
Type=simple
User=nobody
ExecStart=/bin/inner-ss -v -c /etc/inner-ss/%i.json

[Install]
WantedBy=multi-user.target
```
