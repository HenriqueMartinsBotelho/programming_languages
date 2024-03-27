CREATE TABLE usuarios(
  id SERIAL PRIMARY KEY,
  nome VARCHAR(100) NOT NULL,
  email VARCHAR(150) UNIQUE NOT NULL,
  senha VARCHAR(50) NOT NULL
);

INSERT INTO usuarios (nome, email, senha) VALUES
('Bromodomo', 'bro@gmail.com', '123456'),

SELECT * FROM usuarios; 
