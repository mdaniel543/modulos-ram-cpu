CREATE DATABASE modules;

CREATE TABLE process(
    pid INTEGER PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    user VARCHAR(100) NOT NULL,
    status VARCHAR(100) NOT NULL,
    percentage_ram INTEGER NOT NULL,
    pid_padre INTEGER,
    FOREIGN KEY (pid_padre) REFERENCES process(pid)
);

CREATE TABLE ram(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    total INTEGER NOT NULL,
    free INTEGER NOT NULL,
    used INTEGER NOT NULL,
    percentage INTEGER NOT NULL
);
