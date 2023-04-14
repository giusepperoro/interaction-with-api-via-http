create extension if not exists pg_stat_statements;

-- Creation of accounts table
CREATE TABLE IF NOT EXISTS accounts (
                                        client_id BIGINT,
                                        balance int NOT NULL,
                                        PRIMARY KEY(client_id)
    );