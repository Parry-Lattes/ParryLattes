INSERT INTO Pessoa (idLattes,Nome,Nacionalidade) VALUES ("0","root",".");
INSERT INTO Academico (idPessoa,Nivel,AnoTitulacao) VALUES (1,".","2025");
INSERT INTO Docente (idAcademico,SIAPE,Titulo) VALUES (1,"0","");
INSERT INTO Coordenador (idDocente) VALUES (1);
INSERT INTO Login (idCoordenador,Email,Senha) VALUES (1,"root","1234");
