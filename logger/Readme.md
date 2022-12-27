# Logger

Logging required information is esential to monitor requests and responses in a running services.
It helps in tracing, debugging and troubleshooting

Many loggers are available, we choose logger from github.com/sirupsen/logrus to maintain logs in our services.

## Sample Code

```go
package main

import (
  log "github.com/sirupsen/logrus"
)

func main() {
  log.WithFields(log.Fields{
    "animal": "walrus",
    "number": 1,
    "size":   10,
  }).Info("A walrus appears")
}
```
