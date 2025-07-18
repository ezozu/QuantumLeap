// cmd/quantumleap/main.go
package main

import (
"flag"
"log"
"os"

"quantumleap/internal/quantumleap"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := quantumleap.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
