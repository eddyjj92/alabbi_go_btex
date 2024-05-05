CREATE TABLE process (
  id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
  file varchar (255) NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL
);
