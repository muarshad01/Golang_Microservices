## 18. What we'll cover in this section
* User try to Authenticate through the Broker
* Persistence (Postgres db)
  
***

## 19. Setting up a stub Authentication service
```go
$ cd authentication-service
$ go mod init
```

```go
$ go get github.com/go-chi/chi/v5/middleware
$ go get github.com/go-chi/cors
```

***

## 20. Creating and connecting to Postgres from the Authentication service
```go
$ go get github.com/jackc/pgconn
$ go get github.com/jackc/pgx/v4
$ go get github.com/jackc/pgx/stdlib
```
***

## 21. A note about PostgreSQL
```dockerfile
postgres:
    image: 'postgres:14.2'
    ports:
      - "5432:5432"
    deploy:
      mode: replicated
      replicas: 1
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
      POSTGRES_DB: users
    volumes:
      - ./db-data/postgres/:/var/lib/postgresql/data/
```

```
$ mkdir ./project/db-data/postgres
```
***

## 22. Updating our `docker-compose.yml` for Postgres and the Authentication service

```
FROM alpine:latest

RUN mkdir /app

COPY authApp /app

CMD ["/app/authApp"]
```

```go
$ cd <...>/project
$ make up_build
$ make down
$ make up
```

***

## 23. Populating the Postgres database
```sql


--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer DEFAULT nextval('public.user_id_seq'::regclass) NOT NULL,
    email character varying(255),
    first_name character varying(255),
    last_name character varying(255),
    password character varying(60),
    user_active integer DEFAULT 0,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_id_seq', 1, true);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


INSERT INTO "public"."users"("email","first_name","last_name","password","user_active","created_at","updated_at")
VALUES
(E'admin@example.com',E'Admin',E'User',E'$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe',1,E'2022-03-14 00:00:00',E'2022-03-14 00:00:00');
```

### Beekeeper Studio
* https://www.beekeeperstudio.io/
 
***

## 24. Adding a route and handler to accept JSON
* https://github.com/tsawler/toolbox
```
$ go get -u github.com/tsawler/toolbox
```

***

## 25. Update the Broker for a standard JSON format, and conect to our Auth service

***

## 26. Updating the front end to authenticate thorough the Broker and trying things out
```go
$ make up_build
$ make start
```

***

## JSON Toolbox
* [BeekeeperStudio Client for Postgres](https://www.beekeeperstudio.io/)
* [JSON Toolboc](https://github.com/tsawler/toolbox)

***

## Postgres
```
$ go get github.com/jackc/pgconn
$ go get github.com/jackc/pgx/v4
$ go get github.com/jackc/pgx/v4/stdlib
```

***
