package main

import (
    "fmt"
    "os"
    "bufio"
    "database/sql"
    _"github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)

type User struct {
    Id int
    Name string
    Email string 
    Address string
}

func main() {
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }

    database := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
        os.Getenv("DATABASE_USERNAME"),
        os.Getenv("DATABASE_PASSWORD"),
        os.Getenv("DATABASE_HOST"),
        os.Getenv("DATABASE_PORT"),
        os.Getenv("DATABASE_NAME"),
    )

    db, err := sql.Open("mysql", database)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    f, err := os.Create("rw_data_test1.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    w := bufio.NewWriter(f)
    defer w.Flush()

    fmt.Println("Start....")

    for rows.Next() {
        var user User
        err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Address)
        if err != nil {
            panic(err)
        }

        _, err = fmt.Fprintf(w, "%d|%s|%s|%s\n", user.Id, user.Name, user.Email, user.Address)
        if err != nil {
            panic(err)
        }
    }

    fmt.Println("Done....")
}