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


INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Ingenieria', 'Quimica');
INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Ingenieria', 'Electrica');
INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Ingenieria', 'Mecanica');
INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Ingenieria', 'Sistemas');
INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Ingenieria', 'Civil');

INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Medicina', 'Medicina General');
INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Medicina', 'Nutricion');
INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Medicina', 'Enfermeria');

INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Humanidades', 'Filosofia');
INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Humanidades', 'Psicologia');
INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Humanidades', 'Sociologia');

INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Economica', 'Administracion');
INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Economica', 'Economia');
INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Economica', 'Contaduria');

INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Derecho', 'Derecho');

INSERT INTO proyecto1.careers(Faculty, Career) VALUES ('Arquitectura', 'Arquitectura');



-- Line 0
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (0, 'analisis');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (0, 'laboratorio');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (0, 'procesos industriales');

-- Line 1
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (1, 'razonamiento');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (1, 'circuitos');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (1, 'energia');

-- Line 2
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (2, 'resolucion de problemas');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (2, 'diseno mecanico');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (2, 'maquinaria');

-- Line 3
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (3, 'logica');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (3, 'programacion');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (3, 'tecnologia');

-- Line 4
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (4, 'matematica');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (4, 'dibujo');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (4, 'construccion');

-- Line 5
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (5, 'biologia');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (5, 'empatia');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (5, 'salud');

-- Line 6
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (6, 'biologia');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (6, 'comunicacion');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (6, 'bienestar');

-- Line 7
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (7, 'empatia');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (7, 'atencion');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (7, 'salud publica');

-- Line 8
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (8, 'analisis');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (8, 'redaccion');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (8, 'lectura');

-- Line 9
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (9, 'observacion');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (9, 'escucha activa');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (9, 'comportamiento humano');

-- Line 10
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (10, 'pensamiento critico');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (10, 'investigacion');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (10, 'justicia social');

-- Line 11
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (11, 'liderazgo');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (11, 'orgnaizacion');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (11, 'negocios');

-- Line 12
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (12, 'analisis cuantitativo');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (12, 'investigacion');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (12, 'mercados');

-- Line 13
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (13, 'precision');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (13, 'numeros');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (13, 'gestion financiera');

-- Line 14
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (14, 'argumentacion');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (14, 'lectura comprensiva');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (14, 'justicia');

-- Line 15
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (15, 'creatividad');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (15, 'diseno espacial');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (15, 'urbanizacion');


-- Line 0
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (0, 'pensamiento analitico');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (0, 'simulacion');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (0, 'medio ambiente');

-- Line 1
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (1, 'creatividad tecnica');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (1, 'automatizacion');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (1, 'energias renovables');

-- Line 2
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (2, 'perseverancia');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (2, 'manufactura');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (2, 'innovacion');

-- Line 3
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (3, 'adaptabilidad');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (3, 'ciberseguridad');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (3, 'inteligencia artificial');

-- Line 4
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (4, 'organizacion espacial');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (4, 'topografia');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (4, 'infraestructura');

-- Line 5
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (5, 'empatia clinica');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (5, 'diagnostico');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (5, 'bienestar comun');

-- Line 6
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (6, 'curiosidad cientifica');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (6, 'planificacion alimentaria');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (6, 'habitos saludables');

-- Line 7
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (7, 'compromiso');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (7, 'manejo de pacientes');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (7, 'salud mental');

-- Line 8
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (8, 'pensamiento abstracto');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (8, 'debate');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (8, 'etica');

-- Line 9
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (9, 'intuicion');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (9, 'evaluacion psicologica');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (9, 'relaciones humanas');

-- Line 10
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (10, 'sensibilidad social');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (10, 'recoleccion de datos');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (10, 'equidad');

-- Line 11
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (11, 'iniciativa');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (11, 'gestion de proyectos');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (11, 'emprendimiento');

-- Line 12
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (12, 'pensamiento logico');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (12, 'modelado economico');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (12, 'politica publica');

-- Line 13
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (13, 'responsabilidad');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (13, 'auditoria');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (13, 'cumplimiento fiscal');

-- Line 14
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (14, 'analisis juridico');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (14, 'argumentacion legal');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (14, 'derechos humanos');

-- Line 15
INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (15, 'sensibilidad estetica');
INSERT INTO proyecto1.skill(CareerId, Skill) VALUES (15, 'maquetado');
INSERT INTO proyecto1.interest(CareerId, Interest) VALUES (15, 'sostenibilidad');