# inner-shadowsocks for russian language #

![img](https://i.v2ex.co/5iP665B4.png)

[Русский](#russia) 

### Russia

This is a simple app like `ss-local` in Golang. It is designed for running on a server inside China mainland and connecting a server outside through `shadowsocks`, then you can simply set `SOCKS5` proxy in your Chrome, Telegram etc. without running a client on your laptop or cell phones.

There is a python version of this app: [shadowsocks-with-socks-auth](https://github.com/ihciah/shadowsocks-with-socks-auth), which is just a small modified version of original shadowsocks adding `SOCKS5` auth support. This go version is more efficient and support multi-server balance.

-----

#### Features:

- `SOCKS5` auth
- multi-server for load balance

#### Install:

Way 1: Download binary from [Releases](https://github.com/stezkoy/inner-shadowsocks/releases). Golang environment is not required.

Way 2:

```shell
go get -u -v github.com/stezkoy/inner-shadowsocks/inner-ss
```

Due to network issue in China mainland, in most situation, `go get` cannot be executed correctly.

#### Usage:

```shell
inner-ss -c config.json

```Systemd
systemctl enable inner-ss@config

```
- -c: configure file path(default: `config.json`)
- -v: verbose mode

Configure file:

A json string like

```json
{
  "listen": "0.0.0.0",
  "port": 23333,
  "auth": false,
  "username": "ihc",
  "password": "iah",
  "servers": ["ss://RC4-MD5:pass1@ip1:port1", "ss://RC4-MD5:pass2@ip2:port2"],
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

The `auth`, `username`, `password`, `maxfail`, `recovertime`, `remotetimeout`, `insidetimeout`, `starttimeout` are optional whose default value are `false`, `""`, `""`, `10`, `600`, `60`, `60`, `8` .

Also, `whitelistenable` is `false` by default. If you want to limit proxy only to some range, you can use `whitelistdomains` and `whitelistips` .

IP should follow CIDR format such as `100.110.120.130/32`. `::/0` means all IPv6, `0.0.0.0/0` means all IPv4. Domains will be matched by the suffix. Any domains has suffix matched in `whitelistdomains` can pass. For example, `www.github.com` will match `.github.com`.

If you set `auth` to `true`, you should provide `username` and `password`. 

You should list your servers in order. If one of them cannot be connected for `maxfail` times, it will be marked as downed-server until `recovertime` seconds later. Please pay attention that you must ensure the password and encryption method is correct. Because we cannot distinguish users' malicious requests and the wrong password of shadowsocks servers, here we judge server status through if we can connect it.

`remotetimeout` is the timeout connecting remote servers, and `insidetimeout` is the one to clients. `starttimeout`  is the timeout of the first communication. A short `starttimeout` can decrease the waiting time of the false packet, long `remotetimeout` and `insidetimeout` can allow long connection, for example to decrease the reconnecting of Telegram.

You can use `systemd` or `supervisor` to run the daemon. The example configure of supervisor is at the bottom.

Systemd Service 
