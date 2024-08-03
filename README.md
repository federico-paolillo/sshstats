# rsyslog to fluentbit to Grafana Loki

Usually, when you have something with SSH on the Internet somebody will try to get access to it and I am curious of which usernames they try to login with.

## Introduction

I have [rsyslog](https://www.rsyslog.com/) that logs messages of authentication attempts under `/var/log/auth/log`. I would like to sift through those logs for failed authentication attempts and record which usernames where used for those failed attempts.

The goal is to make a (live ?) leaderboard of the most used usernames that failed authentication.

For that, I will feed rsyslog logs to [fluentbit](https://fluentbit.io/) for scraping and transformation. Then fluentbit will forward the now structured logs to [Grafana Loki](https://grafana.com/docs/loki/latest/get-started/labels/) hosted on [Grafana Cloud](https://grafana.com/products/cloud/). From Grafana Cloud I will setup a leaderboard query that I will make available through a very thin backend in [Go](https://github.com/gin-gonic/gin).

A even thinnier frontend will take care of showing the leaderboard.

## What's inside

- rsyslog configuration for logging authentication attempts
- fluentbit configuration for scraping, transformaing and sending rsyslog

**Wait, why no Grafana Loki configuration ?**
Because I am using Grafana Cloud. I have absolutely no willpower to configure Grafana Loki

## Requirements

I am using:

- rsyslog v8.2112.0 (aka 2021.12)
- fluentbit v3.1.4
- Grafana Cloud

## How it works

### rsyslog

Usually, on Ubuntu distributions, rsyslog should be already configured to log auth. attempts under `/var/log/auth.log`. If not, check `config/rsyslog/99_authmessages.conf` for an example. Usually you can put that file under `/etc/rsyslog.d/`, restart rsyslog and see your auth. log files.

Remember to configure [logrotate](https://linux.die.net/man/8/logrotate) appropriately. fluentbit input plugin [tail](https://docs.fluentbit.io/manual/pipeline/inputs/tail) is able to [handle logrotate natively](https://docs.fluentbit.io/manual/pipeline/inputs/tail#file-rotation). An example is also available under `config/rsyslog`, you can usually put logrotate configuration under `/etc/logrotate.d`

### fluentbit

I have installed fluentbit from the official APT package, [as described in the official documentation](https://docs.fluentbit.io/manual/installation/linux/ubuntu). The official package runs fluentbit as systemd service. After installation you need to start the service `fluent-bit`.

I have change the unit file `ExecStart`, under `/lib/systemd/system/fluent-bit.service` to use my own YAML configuration file by specifying `-c /etc/fluent-bit/fluent-bit.yml`

fluentbit will run as `root` as indicated by:

```
$ sudo systemctl show -pUser,UID fluent-bit
UID=[not set]
User=
```

You can find in this repository my fluentbit configuration under `config/fluentbit`. I have chosen to use the YAML configuration format because I like it better.

`fluent-bit.yml` file describes a pipeline that tails `/var/log/auth.log` file and parses it using a custom parser. After parsing all entries that do not have an `user` attribute are discarded. Remaining entries are then forwarded to Loki hosted on Grafana Cloud. I am using the [built-in fluentbit Loki](https://docs.fluentbit.io/manual/pipeline/outputs/loki) plugin that has been configured according to [the official documentation](https://docs.fluentbit.io/manual/pipeline/outputs/loki#fluent-bit--grafana-cloud). It is important for Grafana Cloud [to activate TLS](https://docs.fluentbit.io/manual/administration/transport-security#example-enable-tls-on-http-output).

Note that `fluent-bit.yml` file contains [env. variables interpolations](https://docs.fluentbit.io/manual/administration/configuring-fluent-bit/classic-mode/variables) for `GRAFANACLOUD_USER`, `GRAFANACLOUD_PASSD`, `NODENAME`. When using fluentbit as systemd service environment variables should be specified (by default) under `/etc/default/fluent-bit` or `/etc/sysconfig/fluent-bit` as described in the service unit file.`GRAFANACLOUD_*` variables must be set to a valid username and Grafana Cloud token. See [Grafana Cloud documentation](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/authorize-services/). `NODENAME` should be set to the server's name, this will be part of the Loki stream label.

There is a custom `parsers.conf` to convert the rsyslong entries using regex capturing groups. Of particular note is the property `Time_Offset`, this setting is **necessary** because without proper time zones Loki will reject log entries with a messages similar to: `entry for stream <stream> has timestamp too new: <timestamp>`. Ensure you configure `Time_Offset` to match your server time zone.

The parser will consume log entries that match lines like: `Aug  3 13:57:33 <server> sshd[37253]: Failed password for invalid user root from <ip> port 60352 ssh2` and will capture the date of the entry, the username attempted and the ip the request is coming from.

Both `fluent-bit.yml` and `parsers.conf` have to be placed under `/etc/fluent-bit` then fluentbit service unit file `ExecStart` must be changed to consume the configuration appropriately by specifying `-c /etc/fluent-bit/fluent-bit.yml`.

## Loki

Once you are able to get records into Loki you should find entries such as:

```
2024-08-03 13:57:25.000	{"user":"hadoop","ip":"<redacted>","port":47222}
2024-08-03 13:57:31.000	{"user":"gerrit","ip":"<redacted>","port":58582}
2024-08-03 13:57:33.000	{"user":"root","ip":"<redacted>","port":60352}
2024-08-03 13:57:41.000	{"user":"hadoop","ip":"<redacted>","port":42550}
```

You can get a Top 15 list of usernames in the last 24 hours running a LogQL [instant query](https://grafana.com/docs/loki/latest/reference/loki-http-api/#query-logs-at-a-single-point-in-time) like:

```
topk(15,
  sum by(user) (
    count_over_time({job="ssh-authn"} | json [24h])
  )
)
```
