CREATE TABLE processes (
  id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
  file varchar (255) NOT NULL,
  extension varchar (255) NOT NULL,
  folder varchar (255) NOT NULL,
  route varchar (255) NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL
);
