### Database details:

    postgres-# \l
                             List of databases
    Name    | Owner  | Encoding | Collate |    Ctype    | Access privileges
    -----------+--------+----------+---------+-------------+-------------------
    postgres  | 509318 | UTF8     | C       | en_US.UTF-8 | 

### To for dbadmin:

    export DATABASE_URL="postgres://username:password@localhost:5432/database_name"