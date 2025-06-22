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

IF OBJECT_ID('proyecto1.careers', 'U') IS NOT NULL
    DROP TABLE proyecto1.careers;
GO

IF OBJECT_ID('proyecto1.aptitude', 'U') IS NOT NULL
    DROP TABLE proyecto1.aptitude;
GO

IF OBJECT_ID('proyecto1.skill', 'U') IS NOT NULL
    DROP TABLE proyecto1.skill;
GO

IF OBJECT_ID('proyecto1.interest', 'U') IS NOT NULL
    DROP TABLE proyecto1.interest;
GO

CREATE TABLE proyecto1.careers (
    CareerId INT IDENTITY(0,1) PRIMARY KEY,
    Faculty VARCHAR(24),
    Career VARCHAR(64),
    CONSTRAINT UQ_Careers_Faculty_Career UNIQUE (Faculty, Career)
);
GO

CREATE TABLE proyecto1.aptitude (
    CareerId INT,
    Aptitude VARCHAR(80),
    FOREIGN KEY (CareerId) REFERENCES proyecto1.careers(CareerId),
    PRIMARY KEY (CareerId, Aptitude)
);
GO

CREATE TABLE proyecto1.skill (
    CareerId INT,
    Skill VARCHAR(80),
    FOREIGN KEY (CareerId) REFERENCES proyecto1.careers(CareerId),
    PRIMARY KEY (CareerId, Skill)
);
GO

CREATE TABLE proyecto1.interest (
    CareerId INT,
    Interest VARCHAR(80),
    FOREIGN KEY (CareerId) REFERENCES proyecto1.careers(CareerId),
    PRIMARY KEY (CareerId, Interest)
);
GO



INSERT INTO proyecto1.careers(Faculty, Career)
VALUES ('ingenieria', 'quimica');

INSERT INTO proyecto1.careers(Faculty, Career)
VALUES ('ingenieria', 'electrica');

INSERT INTO proyecto1.careers(Faculty, Career)
VALUES ('ingenieria', 'mecanica');


-- INSERT INTO proyecto1.career(Faculty, Career)
-- VALUES ('', '');

INSERT INTO proyecto1.aptitude(CareerId, Aptitude)
VALUES (0, 'analisis');

INSERT INTO proyecto1.aptitude(CareerId, Aptitude)
VALUES (0, 'ana');

INSERT INTO proyecto1.aptitude(CareerId, Aptitude)
VALUES (1, 'razonamiento');

INSERT INTO proyecto1.aptitude(CareerId, Aptitude)
VALUES (2, 'resolucion_de_problemas');

-- INSERT INTO proyecto1.aptitude(CareerId, Aptitude)
-- VALUES (, '');

INSERT INTO proyecto1.skill(CareerId, Skill)
VALUES (0, 'laboratorio');

INSERT INTO proyecto1.skill(CareerId, Skill)
VALUES (1, 'circuitos');

INSERT INTO proyecto1.skill(CareerId, Skill)
VALUES (1, 'cir');

INSERT INTO proyecto1.skill(CareerId, Skill)
VALUES (2, 'energia');

-- INSERT INTO proyecto1.skill(CareerId, Skill)
-- VALUES (, '');

INSERT INTO proyecto1.interest(CareerId, Interest)
VALUES (0, 'procesos_industriales');

INSERT INTO proyecto1.interest(CareerId, Interest)
VALUES (1, 'energia');

INSERT INTO proyecto1.interest(CareerId, Interest)
VALUES (2, 'maquinaria');

INSERT INTO proyecto1.interest(CareerId, Interest)
VALUES (2, 'maq');

-- INSERT INTO proyecto1.interest(CareerId, Interest)
-- VALUES (, '');