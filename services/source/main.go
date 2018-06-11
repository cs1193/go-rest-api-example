package main

import (
  "os"
)

func main() {
  a := App{}

  a.Initialize(
    os.Getenv("POSTGRES_USERNAME"),
    os.Getenv("POSTGRES_PASSWORD"),
    os.Getenv("POSTGRES_DATABASE_NAME"))

  a.Run(":8000")
}
