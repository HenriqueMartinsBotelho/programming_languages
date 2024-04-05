# Preparando o ambiente

1. `pip install -r requirements.txt`

# Como rodar os exercícios

1. `docker compose up -d` para iniciar o banco de dados.
2. `docker exec -it sql_exercises bash` para entrar no container.
3. `psql -U root exercises` para se conectar ao banco.

## Comandos úteis

`/l` lista todos os bancos
`\du` lista todos os usuários

---

Usando o [Material de Havard](https://cs50.harvard.edu/sql/2023/weeks/0/) para treinar SQL

Use o comando sqlite3 long.db para abrir o banco de dados.
