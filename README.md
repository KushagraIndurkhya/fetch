# Fetch
Fetch is a lightweight CLI for fast downloading .It breaks a file in chunks and utilizes multiple goroutines to download these chunks in parallel to disks and later merge them.

## Benchmark

I benchmarked fetch with two popular cli for downloading cURL and wget on an average it performed better with significantly less time taken by fetch.

![Benchmark](https://github.com/KushagraIndurkhya/fetch/blob/master/benchmark/benchmark.png)

Avg Time taken to download [debian.iso](https://saimei.ftp.acc.umu.se/debian-cd/current/amd64/iso-cd/debian-11.0.0-amd64-netinst.iso) (377 Mb):
- fetch (dividing in 20 chunks)-34.23124401569366 seconds
- cURL- 61.38265061378479 seconds
- wget- 69.53341414928437 seconds

Fetch took <b>44.23% less time than curl and 50.77% less time than wget </b>

## Installation
- For linux/unix systems
	- Download the latest binary from [releases](https://github.com/KushagraIndurkhya/fetch/releases)

	- Download the installation script from [here](http://projects.kindurkhya.me/fetch-install.sh) and run this script:

	$ ``` wget http://projects.kindurkhya.me/fetch-install.sh && ./fetch-install.sh && rm fetch-install.sh ```
- For Windows Systems
 	- Download the latest executable from [releases](https://github.com/KushagraIndurkhya/fetch/releases)
 		-add it to path
- For Building from source Clone this repo
	Build using:
	
	$ ``` go build ```
	
	add the compiled binary to path
## Usage
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
	fetch download https://saimei.ftp.acc.umu.se/debian-cd/current/amd64/iso-cd/debian-11.0.0-amd64-netinst.iso debian.iso --path="/home/stark/iso" --threads=20 --verbose 
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
