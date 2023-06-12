# Echo-server

## Objetivo

Essa aplicação foi desenvolvida didaticamente afim de aprimorar o conhecimento na linguagem Go e Container.
O software inicia um servidor web onde aceita requisição HTTP e retorna para o cliente as seguintes informações: hostname, data da requisição, headers e por fim às variáveis de ambientes. É possível obter dois tipos de retornos, sendo eles: HTML ou JSON. Por padrão as requisições retornam em HTML caso deseje um retorno em JSON basta passar a QueryString `?json=true`.

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