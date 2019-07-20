# APINIT-GO

APINIT-GO is a setup of an apiREST using Golang with a mongo database and Docker with an authentification service using JWT.

## Installation

Use [curl](https://curl.haxx.se/) to install [Docker Compose](https://docs.docker.com/compose/install/#install-compose).

```bash
sudo curl -L "https://github.com/docker/compose/releases/download/1.23.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```
Then apply executable permissions to the binary:

```bash
sudo chmod +x /usr/local/bin/docker-compose
```

## Usage

To start the container
```bash
sudo docker-compose up --build
```

## AUTHENTIFICATION

* __POST__: 127.0.0.1:8080/api/auth/login

```json
{
    "username": "sankamille",
    "password": "password123"
}
```

* __POST__: 127.0.0.1:8080/api/auth/register

```json
{
    "username": "sankamille",
    "email": "luc.brulet@epitech.eu",
    "password": "password123"
}
```