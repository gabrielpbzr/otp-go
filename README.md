# OTP-Go

Uma aplicação web para geração de senhas de uso único do tipo TOTP (Time-based One-Time Password). Projeto de estudo, inspirado na [RFC 6238](https://www.rfc-editor.org/rfc/rfc6238).

Escrito em Go 1.19 usando somente a biblioteca padrão.


> AVISO: NÃO USE ESTE PROJETO EM PRODUÇÃO!

## Como usar

1. Baixe este repositório.
2. Rode o comando `make` para limpar, construir e testar o projeto.
3. Rode este projeto executando o arquivo `bin/otp`.
4. O projeto subirá na porta TCP 3000.
5. Envie um HTTP POST para o servidor com o parâmetro *secret*, contendo a sua chave secreta codificada em base32. Exemplo: Considerando como segredo a string "123456": 
```shell
curl -v http://localhost:3000/ -d secret\=GEZDGNBVGY======
``` 
O aplicativo responderá com uma senha de 6 dígitos, válida por 30 segundos.

6. Para codificar o seu segredo em base32, envie uma requisição para o endpoint `/encode`, com o parâmetro `secret` contendo a sua chave secreta.
```
curl -v http://localhost:3000/encode?secret=123456
``` 
