# Microservice template via Golang

### How to run:

```shell
cd src/cmd
cp .env.example .env
```

### Postgres

To add UUID to the postgres, you should run this command in PSQL:

```shell
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

### Directories structure :

```
project
│   .gitignore
│   .gitlab-ci.yml
│   Dockerfile    
│
└───src
│   │
│   └─── cmd
│   │   │   └───   di  (Dependency Injection)
│   │   │   main.go
│   │   │   .env
│   │   │   .env.example
│   │   │   ...
│   │   │   
│   └─── internal
│   │   │   └───   data
│   │   │   └───   domain
│   │   │   └───   presentation
│   │   │
│   └─── pkg
│       │   ...
```

