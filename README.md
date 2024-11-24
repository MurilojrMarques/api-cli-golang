Aplicação Command-Line Interface.

O objetivo é retornar IPs e os servidores de domínios da internet utilizando Go.

Para testes
1-Para testar é necessário que você tenha o Go instalado.
2- Em uma pasta destinada ao projeto, utilizando os comandos:
git init
git clone https://github.com/MurilojrMarques/api-cli-golang.git
3- Para testar é necessário estar dentro da pasta api-cli-golang dentro do terminal e rodar as seguintes linhas:

Para ver os IPs:
go run main.go ip --host (domínio que você deseja consultar)

Para ver os servidores:
go run main.go servidores --host (domínio que você deseja consultar)

Exemplo: go run main.go ip --host amazon.com.br
Rodando a linha acima a aplicação deverá devolver os ips da amazon.com.br
