# autopprof [![go report card](https://goreportcard.com/badge/github.com/lzakharov/autopprof)](https://goreportcard.com/report/github.com/lzakharov/autopprof) [![PkgGoDev](https://pkg.go.dev/badge/github.com/lzakharov/autopprof)](https://pkg.go.dev/github.com/lzakharov/autopprof)

Tools to automatically collect Go app profiles if something interesting happens to them.

## Usage

The main idea of the project is to provide a framework for creating tools for
automatically collecting profiles of Go applications when certain situations
arise. These tools can help with troubleshooting, postmortem analysis and
optimizations. In addition, this repository contains specific implementations of
such instruments.

### Memory Usage Monitoring

Run a periodic memory usage check that will collect a heap profile when a
specified threshold is exceeded. Check the [memory](./pkg/memory) package to see the
all available options.

```go
const threshold = 1 << 30 // 1GB
monitor := memory.NewMonitor(threshold)
go monitor.Run(ctx)
```

## Examples

See [examples](./examples) for more information.

## TODO

- [ ] Add S3 storage support.

## Alternatives

A good alternative would be to use Continuous Profiling tools. For example, there are such wonderful projects as:
- [Parca](https://github.com/parca-dev/parca)
- [Pyroscope](https://github.com/grafana/pyroscope)
- [profefe](https://github.com/profefe/profefe)

However, unlike this project, continuous profiling services usually require much
more resources. The author believes that for many projects such systems may be
redundant.

## Issue tracker

Please report any bugs and enhancement ideas using the GitHub issue tracker:

https://github.com/lzakharov/autopprof/issues

Feel free to also ask questions on the tracker.
