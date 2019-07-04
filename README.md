apache-logformat
===================

[![Build Status](https://travis-ci.org/charlysan/apache-logformat.png?branch=master)](https://travis-ci.org/charlysan/apache-logformat)

[![GoDoc](https://godoc.org/github.com/charlysan/apache-logformat?status.svg)](https://godoc.org/github.com/charlysan/apache-logformat)

[![Coverage Status](https://coveralls.io/repos/charlysan/apache-logformat/badge.png?branch=topic%2Fgoveralls)](https://coveralls.io/r/charlysan/apache-logformat?branch=topic%2Fgoveralls)

# SYNOPSYS

```go
import (
  "net/http"
  "os"

  "github.com/charlysan/apache-logformat"
)

func main() {
  var s http.ServeMux
  s.HandleFunc("/", handleIndex)
  s.HandleFunc("/foo", handleFoo)

  http.ListenAndServe(":8080", apachelog.CombinedLog.Wrap(&s, os.Stderr))
}
```

# DESCRIPTION

This is a port of Perl5's [Apache::LogFormat::Compiler](https://metacpan.org/release/Apache-LogFormat-Compiler) to golang
