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

1. Falha nas goroutines 

Ao rodar o projeto, está apresentando a seguinte falha:
    ``` panic: runtime error: slice bounds out of range ```
Mesmo aumentando a quantidade de goroutines não está funcionando, foi alterado para que faça a leitura, do arquivo em lotes porém sem sucesso. 

Ao rodar o projeto, apresentou outra falha:
    ``` panic: runtime error: slice bounds out of range [:20] with lenght 19 ```

2. O que está ainda faltando 

Ainda não foi feito a gravação dentro do banco de dados, verificar se com isso não poderia fazer a gravação de linha a linha dentro do banco de dados e depois verificar se a falha não é sanada. 

3. Processamento das informações

Pensar em outro tipo de leitura, as vezes exibir ou gravar em partes, faz a leitura de uma parte grava/exibe, depois vai para o próximo. 