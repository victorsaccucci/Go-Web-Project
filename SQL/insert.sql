DROP DATABASE IF EXISTS gowebprojectdb;
CREATE DATABASE gowebprojectdb;
USE gowebprojectdb;

--SET SQL_SAFE_UPDATES = 0;

CREATE TABLE `gowebprojectdb`.`gamelist` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `descricao` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`)
);


INSERT INTO gamelist (descricao) VALUES ("Aventura");
INSERT INTO gamelist (descricao) VALUES ("Ação");
INSERT INTO gamelist (descricao) VALUES ("Luta");

CREATE TABLE `gowebprojectdb`.`game` (
  `idgame` INT NOT NULL AUTO_INCREMENT,
  `Titulo` VARCHAR(45) NOT NULL,
  `Ano` INT NOT NULL,
  `Genero` VARCHAR(45) NOT NULL,
  `gamelist` INT NOT NULL,
  PRIMARY KEY (`idgame`),
  CONSTRAINT `gamelist_fk`
    FOREIGN KEY (`gamelist`)
    REFERENCES `gowebprojectdb`.`gamelist` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);

INSERT INTO game (Titulo, Ano, Genero, gamelist) VALUES ("The Witcher 3", 2015, "RPG", 6);
INSERT INTO game (Titulo, Ano, Genero, gamelist) VALUES ("Red Dead Redemption 2", 2018, "Ação/Aventura", 6);
INSERT INTO game (Titulo, Ano, Genero, gamelist) VALUES ("Persona 5", 2016, "RPG", 7);
INSERT INTO game (Titulo, Ano, Genero, gamelist) VALUES ("The Last of Us Part II", 2020, "Ação/Aventura", 7);
INSERT INTO game (Titulo, Ano, Genero, gamelist) VALUES ("Cyberpunk 2077", 2020, "RPG", 6);
INSERT INTO game (Titulo, Ano, Genero, gamelist) VALUES ("The Legend of Zelda: Breath of the Wild", 2017, "Ação/Aventura", 6);
INSERT INTO game (Titulo, Ano, Genero, gamelist) VALUES ("FIFA 22", 2021, "Esportes", 8);
INSERT INTO game (Titulo, Ano, Genero, gamelist) VALUES ("Call of Duty: Warzone", 2020, "FPS", 8);
INSERT INTO game (Titulo, Ano, Genero, gamelist) VALUES ("Minecraft", 2011, "Sandbox", 7);

SELECT * FROM game;
SELECT * FROM gamelist;




