DROP DATABASE IF EXISTS gowebprojectdb;
CREATE DATABASE gowebprojectdb;
USE gowebprojectdb;

-- SET SQL_SAFE_UPDATES = 0;

CREATE TABLE `gowebprojectdb`.`gamelist` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `descricao` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO gamelist (descricao) VALUES ("jogos de aventura");
INSERT INTO gamelist (descricao) VALUES ("jogos de acao");
INSERT INTO gamelist (descricao) VALUES ("jogos de luta");

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

INSERT INTO game (Titulo, Ano, Genero, gamelist) VALUES ("Street Fighter", 1999, "Luta", 3);
INSERT INTO game (Titulo, Ano, Genero, gamelist) VALUES ("Red Dead", 2022, "Faroeste", 2);
INSERT INTO game (Titulo, Ano, Genero, gamelist) VALUES ("LOL", 2009, "Estrategia", 1);
INSERT INTO game (Titulo, Ano, Genero, gamelist) VALUES ("CSGO", 2013, "FPS", 3);
INSERT INTO game (Titulo, Ano, Genero, gamelist) VALUES ("Valorant", 2018, "FPS", 2);
INSERT INTO game (Titulo, Ano, Genero, gamelist) VALUES ("BloonsTD", 2007, "Estrategia", 1);

SELECT * FROM game;
SELECT * FROM gamelist;



