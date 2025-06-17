USE master
GO

DROP DATABASE IF EXISTS ia1
GO

CREATE DATABASE ia1
GO

USE ia1
GO

DROP SCHEMA IF EXISTS proyecto1
GO

CREATE SCHEMA proyecto1
GO

IF OBJECT_ID('proyecto1.careers_knowledge', 'U') IS NOT NULL
    DROP TABLE proyecto1.careers_knowledge;
GO

CREATE TABLE careers_knowledge (
    Faculty VARCHAR(24),
    Career VARCHAR(64),
    Aptitude VARCHAR(80),
    Skill VARCHAR(80),
    Interest VARCHAR(80),
    PRIMARY KEY(Faculty, Career)
);
GO