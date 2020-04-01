BEGIN;

CREATE TABLE users (
   id UUID NOT NULL PRIMARY KEY,
   name VARCHAR (64),
   age INTEGER ,
   updated_at TIMESTAMP DEFAULT now() NOT NULL,
   created_at TIMESTAMP DEFAULT now() NOT NULL
);

CREATE TABLE houses (
   id UUID NOT NULL PRIMARY KEY,
   location VARCHAR (64),
   owner_id UUID NOT NULL,
   updated_at TIMESTAMP DEFAULT now() NOT NULL,
   created_at TIMESTAMP DEFAULT now() NOT NULL,
   FOREIGN KEY (owner_id) REFERENCES users (id)
);

COMMIT;