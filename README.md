## Sobre o projeto

Nesse desafio vocë deve criar um programa de linha de comando (cli) que lê um arquivo csv de filmes e popula um banco de dados pensando em performance e esperando que o arquivo pode crescer muito. 

## Requisitos

- Cada linha do csv contém colunas que devem ser salvas em colunas/campos separados no banco de dados:
- ID - número inteiro que identifica o filme encontrado na 1a coluna do csv
- title - título do filme encontrado na 2a coluna do csv. Com o valor "Jumanji (1995)" o title é Jumanji
- Year - ano do filme encontrado na 2a coluna do csv. Com o valor "Jumanji (1995)" o year é 1995
- Genres - múltiplos valores com os gêneros do filme separados por |. Encontrado na 3a coluna.
- O script deve pensar em performance e tirar proveito de concorrência/paralelismo para popular o banco de dados.

O arquivo está dentro da pasta "assets". 

## Tecnologias usadas
- Golang 
- Postgres
- Goroutines
- Canais 

### Coisas para verificar 

- Conexao com o banco de dados
- Carregar variaveis de ambiente
- Estruturacao do projeto