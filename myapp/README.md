---
# Trading in Golang

### Prerequisites

- REVEL Framework [Revel](https://github.com/revel/revel)
- Above revel is a  high-productivity web framework for the [Go language](http://www.golang.org/).
- Goland [Goland](https://www.jetbrains.com/go/) or preferred IDE like vscode atom
- Postgresql [Opensource Database](https://www.postgresql.org/)

### Start Postgresql (Macos) and 

    brew services start postgresql 

### DB tables interaction logic

    sqlInsert:= `INSERT INTO personnel (age, email, first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id`

    id := 0
    err = db.QueryRow(sqlInsert, (26,"rommel@500rockets.io","rommel","galisanao")).Scan(&id)
    if err != nil {
        panic(err)
    }
    fmt.Println("New record ID is:", id)

### To Start the web server:

    revel run -a myapp

### Go to http://localhost:9002/ or check page status

    curl -I --http2 -s localhost:9002 | grep HTTP

### To TEST application running those under /tests:

    revel test -a myapp

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


