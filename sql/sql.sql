-- Comandos para criar tabela no terminal,
-- Criando uma tabela devbook,caso ela não existe 
CREATE DATABASE IF NOT EXISTS devbook;

-- Apartir de agora todos os comandos que for digitados serão dentro do devbook
USE devbook;

-- Se existir uma tabela chamada 'usuarios', ela será removida do banco de dados.
DROP TABLE IF EXISTS usuarios;

--Criando uma nova tabela usuarios, com novos dados
CREATE TABLE usuarios(
    ID int auto_increment primary key,
    Nome varchar(50) not null,
    Nick varchar(50) unique,
    Email varchar(50) not null,
    senha varchar(20) not null
    CriadoEm timestamp default current_timestamp() -- current_timestamp()-> Estou falando que esse campo Criado em irá receber uma data atual
)ENGINE=INNODB;

-- ___________________________________________________________________________________________

CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    ID int auto_increment primary key,
    Nome varchar(50) not null,
    Nick varchar(50) unique,
    Email varchar(50) unique,
    senha varchar(20) unique,
    CriadoEm timestamp default current_timestamp() 
)ENGINE=INNODB;