# Slog (WORK IN PROGRESS)

[![PkgGoDev](https://pkg.go.dev/badge/gitHub.com/luigiBaldanza/slog)](https://pkg.go.dev/gitHub.com/luigiBaldanza/slog) [![Go Report Card](https://goreportcard.com/badge/github.com/luigiBaldanza/slog)](https://goreportcard.com/report/github.com/luigiBaldanza/slog)

## TODO
- Review all names (variables, funcs, types)
- O código deve falar por si só, menos comentários e mais legibilidade
- Funcoes descreverem exatamente oq vão fazer
- Principio de responsabilidade de código
- 5 stages of devops evolution
- Fazer PR para o testify com possivel melhoria ao comparar valores
- The art of readable code

Slog is a logger for Go based on the standard library logger.

```bash
[INFO ] 2020/09/25 23:59:59 main.go:6 This is an information log
[ERROR] 2020/09/25 23:59:59 main.go:7 This is an error log
[WARN ] 2020/09/25 23:59:59 main.go:8 This is a warning log
[PANIC] 2020/09/25 23:59:59 main.go:9 This is a panic log
[DEBUG] 2020/09/25 23:59:59 main.go:10 This is a debug log
```
### Installing
	go get github.com/luigiBaldanza/slog


### Info Log
Info log serve para apresentar informações sobre o fluxo de informação no terminal.
```go
package main

import "github.com/luigiBaldanza/slog"

func main() {
    slog.Info("This is an information log")
}
```
```bash
[INFO ] 2020/09/25 23:59:59 main.go:6 This is an information log
```

### Error Log
Info log serve para apresentar informações sobre o fluxo de informação no terminal.
```go
package main

import "github.com/luigiBaldanza/slog"

func main() {
    slog.Error("This is an error log")
}
```
	[ERROR] 2020/09/25 23:59:59 main.go:6 This is an error log


### Warn Log
Info log serve para apresentar informações sobre o fluxo de informação no terminal.
```go
package main

import "github.com/luigiBaldanza/slog"

func main() {
    slog.Warn("This is a warning log")
}
```
	[WARN ] 2020/09/25 23:59:59 main.go:6 This is a warning log


### Panic Log
Info log serve para apresentar informações sobre o fluxo de informação no terminal.
```go
package main

import "github.com/luigiBaldanza/slog"

func main() {
    slog.Warn("This is a panic log")
}
```
	[PANIC] 2020/09/25 23:59:59 main.go:6 This is a panic log

### Debug Log
Info log serve para apresentar informações sobre o fluxo de informação no terminal.
```go
package main

import "github.com/luigiBaldanza/slog"

func main() {
    slog.Warn("This is a debug log")
}
```
	[PANIC] 2020/09/25 23:59:59 main.go:6 This is a debug log

Existem logs específicos para debug, onde o mesmo só é ativado caso a variável de ambiente esteja ativada. TRADUZIRRRRRRRRRR
	
	export SLOG_DEBUG=true

### Custom Log
Info log serve para apresentar informações sobre o fluxo de informação no terminal.
```go
package main

import "github.com/luigiBaldanza/slog"

func main() {
    slog.Warn("This is a debug log")
}
```

### Colorize Log
LOG COM COLORAÇÃO

### Humanize Log
	export SLOG_HUMANIZE=true

### Stack Log
	export SLOG_STACK_ALL=true
	export SLOG_STACK_DEBUG=true
	export SLOG_STACK_ERROR=true
	export SLOG_STACK_INFO=true
	export SLOG_STACK_WARN=true
	export SLOG_STACK_PANIC=true

### File Log

#### Error File Log

#### Info File Log

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)
