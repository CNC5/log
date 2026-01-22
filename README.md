# log
Structured logger with JSON output

## usage
```
package main

import (
	"github.com/CNC5/log"
)

func main() {
	log.Error(). // spawn new msg chain
        Update("iteration", 4). // add some data
        Update("path", "./").
        UpdateWithJSON("{\"field\": \"value\"}"). // also dump some JSON
        Msg("E") // tell the user something
	// every msg chain MUST end with a Msg or Done
}
```
```output
{"caller":"log.Error","field":"value","iteration":4,"level":"error","msg":"e","path":"./","timestamp":"Thu Jan 22 19:18:43 UTC 2026"}
```