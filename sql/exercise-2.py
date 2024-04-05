import psycopg2
from faker import Faker
import random

# Configuração da conexão ao banco de dados PostgreSQL
conn_params = {
    "dbname": "exercises",
    "user": "root",
    "password": "supersecret",
    "host": "localhost"
}

# Gerando dados fictícios
fake = Faker()


def create_table():
    # SQL para criar a tabela CITY
    create_table_sql = """
    CREATE TABLE IF NOT EXISTS CITY (
        ID SERIAL PRIMARY KEY,
        NAME VARCHAR(17),
        COUNTRYCODE VARCHAR(3),
        DISTRICT VARCHAR(20),
        POPULATION INTEGER
    );
    """
    conn = None
    try:
        conn = psycopg2.connect(**conn_params)
        cur = conn.cursor()
        cur.execute(create_table_sql)
        conn.commit()  # Confirmar a criação da tabela
        cur.close()
    except (Exception, psycopg2.DatabaseError) as error:
        print(error)
    finally:
        if conn is not None:
            conn.close()


def create_cities(n):
    cities = []
    for _ in range(n):
        # Truncando para garantir o limite de 17 caracteres
        name = fake.city()[:17]
        countrycode = fake.country_code()
        # Truncando para garantir o limite de 20 caracteres
        district = fake.state()[:20]
        population = random.randint(1000, 1000000)
        cities.append((name, countrycode, district, population))
    return cities


def insert_cities(cities):
    query = """
    INSERT INTO CITY (NAME, COUNTRYCODE, DISTRICT, POPULATION)
    VALUES (%s, %s, %s, %s)
    """
    conn = None
    try:
        # Conectando ao banco de dados
        conn = psycopg2.connect(**conn_params)
        cur = conn.cursor()
        # Inserindo cada cidade
        for city in cities:
            cur.execute(query, city)
        conn.commit()
        cur.close()
    except (Exception, psycopg2.DatabaseError) as error:
        print(error)
    finally:
        if conn is not None:
            conn.close()


def show_all_cities():
    query = """
    SELECT * FROM CITY
    """
    conn = None
    try:
        conn = psycopg2.connect(**conn_params)
        cur = conn.cursor()
        cur.execute(query)
        rows = cur.fetchall()
        for row in rows:
            print(row)
        cur.close()
    except (Exception, psycopg2.DatabaseError) as error:
        print(error)
    finally:
        if conn is not None:
            conn.close()


def select_big_cities():
    query = """
    SELECT * FROM CITY WHERE POPULATION > 100000 AND COUNTRYCODE = 'CN' """
    conn = None
    try:
        conn = psycopg2.connect(**conn_params)
        cur = conn.cursor()
        cur.execute(query)
        rows = cur.fetchall()
        for row in rows:
            print(row)
        cur.close()
    except (Exception, psycopg2.DatabaseError) as error:
        print(error)
    finally:
        if conn is not None:
            conn.close()


# contiuar estudado com o obsidian

# Criação da tabela CITY
# create_table()

# Gerar 1000 cidades fictícias
# cities = create_cities(1000)
# Inserir as cidades no banco de dados
# insert_cities(cities)
select_big_cities()
