# Echo-server

## Objetivo

Essa aplicação foi desenvolvida didaticamente afim de aprimorar o conhecimento na linguagem Go e Docker.
O software inicia um servidor web onde aceitar qualquer requisição HTTP e retornando para o cliente o hostname, data da requisição e os headers além das variáveis de ambientes.

## Recursos necessários

- Golang
- Docker
- cURL

## Como executar à aplicação `echo-server`

Faça download do repositório no seu computador pessoal.
```bash
git clone https://github.com/ojaoferreira/echo-server.git
```

Entre no diretório `echo-server`.
```bash
cd echo-server
```

Crie sua imagem Docker com aplicação `echo-server`.
```bash
docker build . -t echo-server
```

Inicialize à aplicação com o comando abaixo.
```bash
docker run -p 8080:8080 echo-server
```

Com o comando `cURL` vamos verificar se aplicação está rodando.
```bash
curl http://localhost:8080/qualquer/path
```

Utilizando a QueryString `?json=true` o retorno do request será um JSON com o `content-type` igual `application/json`. 
```bash
curl http://localhost:8080/qualquer/path?json=true
```

## Capturas de tela

![Requisição via Google Chrome](/screenshots/chrome.png?raw=true "Requisição via Google Chrome")

![Requisição via Postman](/screenshots/postman.png?raw=true "Requisição via Postman")