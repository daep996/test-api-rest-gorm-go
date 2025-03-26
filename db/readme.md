```bash
sudo -u postgres psql
```

```postgres

GRANT ALL ON SCHEMA public TO your_user;

GRANT ALL ON ALL TABLES IN SCHEMA public TO your_user;

ALTER SCHEMA public OWNER TO your_user;


GRANT CREATE ON SCHEMA public TO nombre_usuario;

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO nombre_usuario;

ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO nombre_usuario;

ALTER DATABASE my_database OWNER TO my_database_user;
```