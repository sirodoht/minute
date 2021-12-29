# minute

Self-hosted uptime monitor.

## Usage

```sh
$ go build
$ ./minute sites.txt
```

## Configuration

Configuration is done through the `sites.txt` file.

```
<SMTP server URL>
<SMTP username>
<SMTP password>
<From email>
<To email>
<Website 1>
<Website 2>
...
<Website n>
```

Check [`example-sites.txt`](example-sites.txt) for an example.

## License

MIT
