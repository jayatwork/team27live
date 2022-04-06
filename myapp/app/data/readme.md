### Database details:

    postgres-# \l
                             List of databases
    Name    | Owner  | Encoding | Collate |    Ctype    | Access privileges
    -----------+--------+----------+---------+-------------+-------------------
    postgres  | ****** | UTF8     | C       | en_US.UTF-8 | 

### To for dbadmin:

    export DATABASE_URL="postgres://username:password@localhost:5432/database_name"

    postgres=# \dt
          List of relations

    Schema |   Name    | Type  | Owner  
    --------+-----------+-------+--------
    public | personnel | table | ******
    (1 row)

    postgres=# CREATE TABLE trading (
    id SERIAL PRIMARY KEY,
    implementation INT,
    account TEXT,
    model TEXT, status TEXT UNIQUE NOT NULL
    );

    CREATE TABLE

    postgres=# \dt
    List of relations
    Schema |   Name    | Type  | Owner  
    --------+-----------+-------+--------
    public | personnel | table | ******
    public | trading   | table | ******
    (2 rows)

 	Alpaca Trading API 	GET /v2/account 	Status 	DATE: 
Crytotrading with Golang 	Request 	Status 	Date: 