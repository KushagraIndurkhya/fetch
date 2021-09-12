# fetch
Fetch is a CLI for downloading files written in Go that enables user to get fast downloads by
utilizing multiple threads and downloading file chunks in parallel

```
Usage:
  fetch [flags]
  fetch [command]

Available Commands:
  clean       Delete Downloaded files
  completion  generate the autocompletion script for the specified shell
  download    Download from url
  help        Help about any command
  history     Fetch your download history

Flags:
  -h, --help   help for fetch
```
## download
```
Usage:
  fetch download <URL> <filename> [flags]

Flags:
      --config string   config file (default is $HOME/.fetch.yaml)
  -h, --help            help for fetch
      --path string     Specify Download Location of the file
      --seq             Download the file sequentially instead of parallel downloading
      --threads int     Specify Number of threads to be used (default 20)
      --verbose         Specify Verbosity of the output
```

```
	Example Usage: 
	fetch download https://saimei.ftp.acc.umu.se/debian-cd/current/amd64/iso-cd/debian-11.0.0-amd64-netinst.iso debian.iso --path="~/iso/debian.iso" --threads=20 --verbose 
```

## help
Get details of the files downloaded using fetch

```
Usage:
  fetch history [flags]
  fetch history [command]

Available Commands:
  clean       Clear your downloading history

Flags:
  -h, --help       help for history
      --list int   Specify Number of Rows in result (default 10)

```
Example usage: 
```
fetch history --list=5

Shows 5 recent downloads
```

## Clean
Clean Up all the files downloaded using fetch cli

Example usage: 
```
fetch clean
```
