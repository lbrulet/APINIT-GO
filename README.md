[![Go Report Card](https://goreportcard.com/badge/github.com/lbrulet/APINIT-GO)](https://goreportcard.com/report/github.com/lbrulet/APINIT-GO)
  
# APINIT-GO

APINIT-GO is a setup of an apiREST using Golang with a mongo database and Docker with an authentification service using JWT.

* __TODO__:
    - CRUD User
    - Unit Testing with moq & testify

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

## CHOOSE YOUR ENVIRONNMENT !

Use this export to connect the api to your local mongo.
```bash
export ENVIRONMENT=LOCAL
```

Otherwise you can just used the docker compose.

## MAIL SENDER

Don't forget to write your smtp detail into the /configs/local/config.json or /configs/dev/config.json
And please export your mail account with password like below :

```bash
export MAIL_ADDRESS=luc.brulet@epitech.eu
export MAIL_PASSWORD=azertyuiop
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
