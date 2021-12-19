# slogx
[![Release](https://img.shields.io/github/v/release/IchBinLeoon/slogx?style=flat-square)](https://github.com/IchBinLeoon/slogx/releases)
[![Go](https://img.shields.io/github/go-mod/go-version/IchBinLeoon/slogx?style=flat-square)](https://golang.org/)
[![License](https://img.shields.io/github/license/IchBinLeoon/slogx?style=flat-square)](https://github.com/IchBinLeoon/slogx/blob/main/LICENSE)
[![Code Size](https://img.shields.io/github/languages/code-size/IchBinLeoon/slogx)](https://github.com/IchBinLeoon/slogx/blob/main/slogx.go)

A minimal level based logging library for Go.

- [Installation](#Installation)
- [Example](#Example)
- [Usage](#Usage)
    - [Logger](#Logger)
    - [Log](#Log)
    - [Level](#Level)
    - [Format](#Format)
    - [Output](#Output)
- [Contribute](#Contribute)
- [License](#License)

## Installation
```
go get github.com/IchBinLeoon/slogx
```

## Example
```go
package main

import "github.com/IchBinLeoon/slogx"

func main() {
    logger := slogx.NewLogger("EXAMPLE")

    logger.Error("This is Error!")
    logger.Errorf("This is %s!", "Error")

    logger.Warning("This is Warning!")
    logger.Warningf("This is %s!", "Warning")

    logger.Info("This is Info!")
    logger.Infof("This is %s!", "Info")
}
```
Output:
```
2021-06-08 20:08:19 ERROR main.go:11 EXAMPLE: This is Error!
2021-06-08 20:08:19 ERROR main.go:12 EXAMPLE: This is Error!
2021-06-08 20:08:19 WARNING main.go:14 EXAMPLE: This is Warning!
2021-06-08 20:08:19 WARNING main.go:15 EXAMPLE: This is Warning!
2021-06-08 20:08:19 INFO main.go:17 EXAMPLE: This is Info!
2021-06-08 20:08:19 INFO main.go:18 EXAMPLE: This is Info!
```

## Usage
### Logger
Create a new logger:
```go
logger := slogx.NewLogger("awesome name")
```

Get an existing logger by its name:
```go
logger := slogx.GetLogger("awesome name")
```

### Log
Log a message at Fatal level and exit:
```go
logger.Fatal("This is Fatal!")
logger.Fatalf("This is %s!", "Fatal")
```

Log a message at Error level:
```go
logger.Error("This is Error!")
logger.Errorf("This is %s!", "Error")
```

Log a message at Warning level:
```go
logger.Warning("This is Warning!")
logger.Warningf("This is %s!", "Warning")
```

Log a message at Info level:
```go
logger.Info("This is Info!")
logger.Infof("This is %s!", "Info")
```

Log a message at Debug level:
```go
logger.Debug("This is Debug!")
logger.Debugf("This is %s!", "Debug")
```

Log a message at a specified level:
```go
logger.Log(slogx.ERROR, "This is Error!")
logger.Logf(slogx.INFO, "This is %s!", "Info")
```

### Level
The default logging level is `INFO`.

Set the logging level:
```go
logger.SetLevel(slogx.DEBUG)
```

Set the logging level from a string:
```go
logger.SetLevel(slogx.ParseLevel("DEBUG"))
```

Get the current logging level:
```go
level := logger.GetLevel()
```

### Format
Set a custom format:
```go
err := logger.SetFormat("[${time}] ${file} (${level}): ${message}")
if err != nil {
    // Handle error...
}
```
|Verb|Description|
|----|-----------|
|${time}|The current time|
|${level}|The logging level|
|${file}|The file the log statement is in|
|${line}|The line the log statement is on|
|${name}|The name of the logger|
|${message}|The log message|

The default time format is `2006-01-02 15:04:05`. 

To change the time format:
```go
logger.SetTimeFormat("Jan _2 15:04:05")
```
The time format must be a layout supported by the go time package.

### Output
The default output is `Stdout`.

Set the output:
```go
f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    // Handle error...
}
defer f.Close()

logger.SetOutput(f)
```
The output can be any `io.Writer`.

## Contribute
Contributions are welcome! Feel free to open issues or submit pull requests!

## License
This project is licensed under the MIT License. See the [LICENSE](https://github.com/IchBinLeoon/slogx/blob/main/LICENSE) file for more details.