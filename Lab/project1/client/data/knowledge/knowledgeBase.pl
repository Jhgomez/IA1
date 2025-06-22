% Career Fact: career(Faculty, Career, Aptitude, Skill, Interest)

% Engineering
career(ingenieria, quimica, analisis, laboratorio, procesos industriales).
career(ingenieria, electricta, razonamiento, circuitos, energia).
career(ingenieria, mecanica, resolucion de problemas, diseno mecanico, maquinaria).
career(ingenieria, sistemas, logica, programacion, tecnologia).
career(ingenieria, civil, matematica, dibujo, construccion).

% Medicine
career(medicina, medicina general, biologia, empatia, salud).
career(medicina, nutricion, biologia, comunicacion, bienestar).
career(medicina, enfermeria, empatia, atencion, salud publica).

% Social Sciences and Humanities
career(humanidades, filosofia, analisis, redaccion, lectura).
career(humanidades, psicologia, observacion, escucha activa, comportamiento humano).
career(humanidades, sociologia, pensamiento critico, investigacion, justicia social).

% Economic Sciences
career(economica, administracion, liderazgo, orgnaizacion, negocios).
career(economica, economia, analisis cuantitativo, investigacion, mercados).
career(economica, contaduria, precision, numeros, gestion financiera).

% Judgement Sciences and Socials
career(derecho, derecho, argumentacion, lectura comprensiva, justicia).

% Design and Architecture
career(arquitectura, arquitectura, creatividad, diseno espacial, urbanizacion).

% Rules
suggested career(Faculty, Career, Aptitude, Skill, Interest) :- career(Faculty, Career, Aptitude, Skill, Interest).