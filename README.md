# Lighthouse

Lighthouse is an analytics tool to use when JS scripts are not allowed such as emails and READMEs

## Getting Started
To get stared using docker run:

(Linux / MacOs)
```bash
$ mv .env.example .env
$ docker-compose up -d
```

(Powershell)
```
$ Rename-Item .env.example .env
$ docker-compose up -d
```

To include the counter badge in markdown format, include the following line.

If a badge is not required exclude the `badge` parameter from the url

```
![lighthouse badge](http://localhost:8080/apps/YOUR_APP_NAME/collect?badge)
```

To retrieve app statistics from an app

```
$curl http://localhost:8080/apps/YOUR_APP_NAME/stats
```

## Is this production ready?
The visit based implementation has a few drawbacks. Exposed links can be queried endlessly even though the time and remote address is logged. Remote addresses can be altered by a user or if accessed behind proxies.

## License
This is licensed under MIT License Copyright(2020) Hannes Kimara