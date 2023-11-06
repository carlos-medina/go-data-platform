DROP TABLE IF EXISTS records;
CREATE TABLE records (
  data_id    INT NOT NULL,
  user_id    INT NOT NULL,
  version    INT NOT NULL,
  content    VARCHAR(255) NOT NULL,
  PRIMARY KEY (`data_id`)
);