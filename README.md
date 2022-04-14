# goenvconfig

Inspired by Elixir env configs (prod, test, dev) and Ruby .env

## Usage

- include `.env` , `.env.test` and `.env.development` in your project root folder

```go
package main

import (
    "os"

    "github.com/grandeto/goenvconfig"
)

func init() {
    goenvconfig.LoadConfig()
}

func main() {
    dbDsn := os.Getenv("DB_DSN")
}
```

- build your go app by including in `-tags` prod, test or dev depending on your environment

```bash
go build -tags prod .
```
