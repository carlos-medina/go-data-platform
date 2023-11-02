DROP TABLE IF EXISTS records;
CREATE TABLE records (
  id         INT AUTO_INCREMENT NOT NULL,
  user_id    INT NOT NULL,
  data_id    INT NOT NULL,
  version    INT NOT NULL,
  content    VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);