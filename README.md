# Modulo go para saludos

# Practicas de aprendizaje GO

## Instalación
Ejecuatr el comando 
```bash
go get -u github.com/migleor/greetings
```
## Uso
Para utilizar el módulo llama la función Hellos pasando un slice de nombres

```go
package main

import (
	"fmt"
	"log"

	"github.com/migleor/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Name 1", "Name 2", "Name 3"}
	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, message := range message {
		fmt.Println(message)
	}

}
```
