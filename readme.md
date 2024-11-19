# Go Backend Template

Esta es una plantilla para crear un backend en Go. Proporciona una estructura básica y componentes comunes para acelerar el desarrollo de aplicaciones backend.

## Características

- Estructura de proyecto organizada
- Manejo de dependencias con Go Modules
- Ejemplos de controladores y rutas
- Configuración de entorno
- Integración con bases de datos
- Middleware común (autenticación, logging, etc.)

## Requisitos

- Go 1.16 o superior
- Docker (opcional, para desarrollo y despliegue)

## Instalación

1. Clona el repositorio:

```sh
git clone https://github.com/ruxwez/go-backend-template.git
cd go-backend-template
```

2. Instala las dependencias:
```sh
go mod tidy
```

3. Crea un archivo `.env` basado en `.env.example` y configura tus variables de entorno.
```sh
cp .env.example .env
```

4. Ejecuta el servidor:
```sh
go build -o main .
./main
```


## Ejecución con Docker

1. Construye la imagen Docker:
```sh
docker build -t <name-of-your-image> .
```

2. Ejecuta la imagen Docker:
```sh
docker run -p 3000:3000 --env-file .env <name-of-your-image>
```
Es importante tener en cuenta que el puerto debe coincidir con el configurado en el archivo `.env`.

Ademas, es importante tener en cuenta que el nombre del archivo `.env` debe coincidir con el nombre del archivo de entorno que se está utilizando en la línea de comandos.
