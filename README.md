# Enron Mail

Esta aplicacion permite indexar el dataset de correos de Enron Corporation. para luego poder ser visualizados en una pagina web.

# Requerimientos
* Golang >=1.21
* Zincsearch
* nodejs >=22.0.0


* Enron Mail Data set

### Enron Mail Data set

Si usas linux puedes descargarlo de aqui
```bash
curl -L http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz -o enron_mail_20110402.tgz && tar -xf enron_mail_20110402.tgz
```

## Zincsearch
la base de datos usada es [ZincSearch](https://github.com/zincsearch/zincsearch) puedes descargarlo en su pagina oficial o levantar el contenedor que se encuentra en el directorio IndexerV2, primero se debe crear una red en docker

```bash
docker network create mails-network

docker-compose up
```

## IndexerV2

Este programa se encarga de indexar los correos para ejecutarlo debes usar.
```bash
go run main.go "mails directory"
```

tambien acepta algunos flag de configuraciones

| Opción   | Tipo  | Descripción |
|----------|-------|-------------|
| `-batch` | `int` | Batch size para insertar emails en la base de datos. Mín: 100, Máx: 2000, Default: 1000. |
| `-prof`  | `bool` | Activa trace profiling. No se recomienda usarlo junto con CPU y memory profiling para mejor resultado. Default: false. |
| `-trace` | `bool` | Activa memory y CPU profiling. No se recomienda usarlo junto con trace profiling para mejor resultado. Default: false. |
| `-wm`    | `int` | 'Workers mails': Número de procesos para indexar emails. Default: número de CPU, Mín: 1, Máx: 30, Default: 20. |
| `-wu`    | `int` | 'Workers upload': Número de procesos para subir emails. Mín: 1, Máx: 8, Default: 4. |

**Example**
```batch
go run main -wm 4 -wu 1 mailsdir
```
#### Configuración de Variables de Entorno

Este archivo documenta las variables de entorno utilizadas en la aplicación.

## Variables de Base de Datos

| Variable              | Descripción                                      | example |
|-----------------------|--------------------------------------------------|-------------------|
| `DATABASE_USER`       | Usuario de la base de datos.                     |           |
| `DATABASE_PASSWORD`   | Contraseña del usuario de la base de datos.      |  |
| `DATABASE_HOST`       | URL o dirección del servidor de base de datos.   | `http://localhost` |
| `DATABASE_PORT`       | Puerto en el que se ejecuta la base de datos.    | `4080`           |
| `DATABASE_NAME`       | Nombre de la base de datos utilizada.            | `mails`         |
| `APP_ENV` | Entorno en el que se ejecuta la aplicación. Puede ser `development` o `production`. | `development` |

## Api

Este es el api que se encargara de obtener los correos de la base de datos para correrlo usa o *Air* si lo tienes instalado, asegurate de llenar las variables de entorno.

```
run go main.go
```

O

```
air
```

## App (Frontend)

Esta es la aplicacion web que obtendra los datos desde el api para inicializarlo debes usar 

```
npm run dev
```




