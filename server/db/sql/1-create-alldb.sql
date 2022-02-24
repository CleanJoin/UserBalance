CREATE SCHEMA avito AUTHORIZATION avito;

CREATE SEQUENCE avito.transaction_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;



ALTER SEQUENCE avito.transaction_id_seq OWNER TO avito;
GRANT ALL ON SEQUENCE avito.transaction_id_seq TO avito;



CREATE SEQUENCE avito.users_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;

-- Permissions

ALTER SEQUENCE avito.users_id_seq OWNER TO avito;
GRANT ALL ON SEQUENCE avito.users_id_seq TO avito;


CREATE TABLE avito."transaction" (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	useridfrom int4 NOT NULL DEFAULT 0,
	"money" float8 NOT NULL DEFAULT 0,
	transaction_time timestamp NOT NULL DEFAULT now(),
	useridto int4 NOT NULL DEFAULT 0
);
CREATE UNIQUE INDEX transaction_id_idx ON avito.transaction USING btree (id);

ALTER TABLE avito."transaction" OWNER TO avito;
GRANT ALL ON TABLE avito."transaction" TO avito;



CREATE TABLE avito.users (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	username varchar NOT NULL,
	"password" varchar NOT NULL,
	"money" float8 NOT NULL DEFAULT 0
);
CREATE UNIQUE INDEX users_id_idx ON avito.users USING btree (id);


ALTER TABLE avito.users OWNER TO avito;
GRANT ALL ON TABLE avito.users TO avito;




GRANT ALL ON SCHEMA avito TO avito;

INSERT INTO avito.users
(username, "password", "money")
VALUES('Andrey', '8e756c9f2b15da6a63f84852fc39667617523133', 0.0);
INSERT INTO avito.users
(username, "password", "money")
VALUES('Anton', '8e756c9f2b15da6a63f84852fc39667617523134', 0.0);
