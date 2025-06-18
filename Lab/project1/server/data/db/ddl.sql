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

CREATE TABLE proyecto1.careers_knowledge (
    Faculty VARCHAR(24),
    Career VARCHAR(64),
    Aptitude VARCHAR(80),
    Skill VARCHAR(80),
    Interest VARCHAR(80),
    PRIMARY KEY(Faculty, Career)
);
GO

INSERT INTO proyecto1.careers_knowledge(Faculty, Career, Aptitude, Skill, Interest)
VALUES ('ingenieria', 'quimica', 'analisis', 'laboratorio', 'procesos_industriales');

INSERT INTO proyecto1.careers_knowledge(Faculty, Career, Aptitude, Skill, Interest)
VALUES ('ingenieria', 'electricta', 'razonamiento', 'circuitos', 'energia');

INSERT INTO proyecto1.careers_knowledge(Faculty, Career, Aptitude, Skill, Interest)
VALUES ('ingenieria', 'mecanica', 'resolucion_de_problemas', 'diseno_mecanico', 'maquinaria');

-- INSERT INTO proyecto1.careers_knowledge(Faculty, Career, Aptitude, Skill, Interest)
-- VALUES ('', '', '', '', '');

select * from proyecto1.careers_knowledge