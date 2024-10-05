# eacclient

A collection of eacnet clients and tools

# Usage

```
eacnet clients and tools

Usage:
  eacclient [command]

Available Commands:
  common      Common service client
  help        Help about any command
  infinitas   Infinitas service client
  inspector   Start the traffic inspector
  konasute    Konasute service client
  send        Send a request to an eacnet server using a property file

Flags:
  -g, --game string        Konasute game ID
  -h, --help               help for eacclient
  -i, --iid string         Infinitas ID
  -p, --protocol string    "konasute" or "infinitas" (default "konasute")
      --timeout duration   Timeout (default 20s)
  -t, --token string       Authentication token
  -u, --url string         Service URL
  -v, --version string     Software version

Use "eacclient [command] --help" for more information about a command.
```

Documentation can be found under the [docs](docs) directory.