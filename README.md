# fetch
Fetch is a CLI for downloading files written in Go that enables user to get fast downloads by
utilizing multiple threads and downloading file chunks in parallel

```
Usage:
  fetch <URL> <filename> [flags]

Flags:
      --config string   config file (default is $HOME/.fetch.yaml)
  -h, --help            help for fetch
      --path string     Specify Download Location of the file
      --seq             Download the file sequentially instead of parallel downloading
      --threads int     Specify Number of threads to be used (default 20)
      --verbose         Specify Verbosity of the output
```
