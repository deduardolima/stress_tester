# Stress-Tester

## Descrição

Este um sistema CLI em Go para realizar testes de carga em um serviço web. Você deverá fornecer uma URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

## Funcionalidades

- Realiza testes de carga em URLs especificadas.
- Permite configurar o número total de requests e a quantidade de chamadas simultâneas.
- Gera um relatório detalhado com o tempo total, requests bem-sucedidos e distribuição de status HTTP.


## Estrutura do Projeto

- Stress-Tester/
  - cmd/
    - root.go
    - stress_tester/
      - main.go
  - internal/
    - entity/ 
      - report.go 
    - usecase/
      - run_stress_tester.go
    - service/
      - http_client.go
  - pkg/
    - logger.go
  - docker-compose.yml
  - Dockerfile
  - go.mod
  - go.sum
  - README.md


## Pré-requisitos

- Docker
- Docker Compose

## Configuração
```
git clone https://github.com/deduardolima/stress-tester.git
cd stress-tester

```

## Instalação e Execução com Docker
Construa e inicie os containers:
```
docker-compose up --build -d
```

isso irá construir a imagem do aplicativo e iniciar o serviço definido no docker-compose.yml, 
e ja iniciará um teste com  1000 requisições para https://www.globo.com sendo 10 simultaneamente.


## Execucão 

Agora é só utilizar!
- Escolha sua URL de teste e inclua no comando abaixo:

```
docker run stress-tester-stress-tester --url=https://teste.com --requests=1000 --concurrency=10

```

Fique à vontade para alterar as variáveis `url`, `requests` e `concurrency`.


## Exemplo de Saída

```
 Total Time: 12.00 seconds 
 Total Requests: 1000
 Successful Requests: 990
 Status Distribution: map[200:990 429:10]
```

## Créditos

Este projeto foi criado por [Diego Eduardo](http://github.com/deduardolima)







