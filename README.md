# OTP-Go

Uma aplicação web para geração de senhas de uso único do tipo TOTP (Time-based One-Time Password). **Projeto de estudo**, inspirado na [RFC 6238](https://www.rfc-editor.org/rfc/rfc6238).

Escrito em Go 1.19 usando somentea a biblioteca padrão.


> AVISO: NÃO USE ESTE PROJETO EM PRODUÇÃO!

> WARNING: DO NOT USE THIS PROJECT IN PRODUCTION


## Como usar

1. Baixe este repositório.
2. Rode o comando `make` para limpar, construir e testar o projeto.
3. Rode este projeto executando o arquivo `bin/otp`.
4. O projeto subirá na porta TCP 3000.
5. Envie um HTTP POST para o servidor com o parâmetro *secret*, contendo a sua chave secreta codificada em base32
