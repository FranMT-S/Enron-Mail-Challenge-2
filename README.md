# Enron Mail

This application allows you to index the Enron Corporation email dataset, which can then be viewed on a web page.

# Requirements
* Golang >=1.21
* Zincsearch
* Node.js >=22.0.0

* Enron Mail Dataset

### Enron Mail Dataset

If you're using Linux, you can download it from here:
```bash
curl -L http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz -o enron_mail_20110402.tgz && tar -xf enron_mail_20110402.tgz
```

## Zincsearch

The database used is [ZincSearch](https://github.com/zincsearch/zincsearch). You can download it from its official page or run the container found in the IndexerV2 directory. First, you need to create a network in Docker:

```bash
docker network create mails-network

docker-compose up
```

## IndexerV2

This program is responsible for indexing emails, to run it you must use.
```bash
go run main.go "mails directory"
```

It also accepts some configuration flags

| Option | Type | Description |
|----------|----------|-------------|
| `-batch` | `int` | Batch size to insert emails into the database. Min: 100, Max: 2000, Default: 1000. |
| `-prof` | `bool` | Enables trace profiling. It is not recommended to use it together with CPU and memory profiling for best results. Default: false. |
| `-trace` | `bool` | Enables memory and CPU profiling. It is not recommended to use it together with trace profiling for best results. Default: false. |
| `-wm` | `int` | 'Workers mails': Number of processes to index emails. Default: number of CPUs, Min: 1, Max: 30, Default: 20. |
| `-wu` | `int` | 'Workers upload': Number of processes to upload emails. Min: 1, Max: 8, Default: 4. |

**Example**
```batch
go run main -wm 4 -wu 1 mailsdir
```
#### Setting Environment Variables

This file documents the environment variables used in the application.

## Database Variables

| Variable | Description | example |
|-----------------------|--------------------------------------------------|-------------------|
| `DATABASE_USER` | Database user. | |
| `DATABASE_PASSWORD` | Database user password. | |
| `DATABASE_HOST` | URL or address of the database server. | `http://localhost` |
| `DATABASE_PORT` | Port on which the database is running. | `4080` |
| `DATABASE_NAME` | Name of the database used. | `mails` |
| `APP_ENV` | Environment in which the application is running. Can be `development` or `production`. | `development` |

## Api

This is the api that will be in charge of getting the emails from the database to run it use *Air* if you have it installed, make sure to fill in the environment variables.

```
run go main.go
```

Or

```
air
```

## App (Frontend)

This is the web application that will get the data from the api to initialize it you must use

```
npm run dev
```
