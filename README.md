# Simple Converter Program
This is a simple converter program that will convert values between different units of measurement.

## How to use
### get golang
```
$ sudo apt-get install golang
```
### clone repo
```
git clone https://github.com/finbaromahony/converter.git 
```
### Without buildig
```
$ go run converter.go -h
## Conversion program ##
usage: converter [-h|--help] -f|--from (kg|stone|lbs|grams|oz) -t|--to
                 (kg|stone|lbs|grams|oz) -v|--value <float>

                 Convert Units of measurement from one type to another

Arguments:

  -h  --help   Print help information
  -f  --from   Measurement Unit to convert from
  -t  --to     Measurement Unit to convert to
  -v  --value  Value to Convert

$ go run converter --from kg --to stone --value 55
## Conversion program ##
- Convert kg to stone -
55kg = 8.661017442977332 Stone
or
8 Stone; 9.25 lbs
```
### building
```
$ go build

$ ./converter -h
## Conversion program ##
usage: converter [-h|--help] -f|--from (kg|stone|lbs|grams|oz) -t|--to
                 (kg|stone|lbs|grams|oz) -v|--value <float>

                 Convert Units of measurement from one type to another

Arguments:

  -h  --help   Print help information
  -f  --from   Measurement Unit to convert from
  -t  --to     Measurement Unit to convert to
  -v  --value  Value to Convert

$ ./converter --from kg --to stone --value 55
## Conversion program ##
- Convert kg to stone -
55kg = 8.661017442977332 Stone
or
8 Stone; 9.25 lbs

```

## Development Notes
- started with converter.go as a simple entrypoint
- decided on using an argparse and found one developed by clagraff
- this brought up the issues regarding how does go do dependancies, had a look and found glide, but the glide site said that it was no longer maintained and suggested using go modules.
- go modules worked great
    - imported and used argparse in code
    - ran `go mod init converter.go`
    - ran `go run converter.go`
    - and it did the rest itself, impressive stuff
    ```
    ~/dev/golang/converter$ go run converter.go 
    go: finding github.com/clagraff/argparse v1.0.1
    go: downloading github.com/clagraff/argparse v1.0.1
    go: extracting github.com/clagraff/argparse v1.0.1
    go: finding github.com/nsf/termbox-go latest
    go: downloading github.com/nsf/termbox-go v0.0.0-20191229070316-58d4fcbce2a7
    go: extracting github.com/nsf/termbox-go v0.0.0-20191229070316-58d4fcbce2a7
    go: finding github.com/mattn/go-runewidth v0.0.8
    go: downloading github.com/mattn/go-runewidth v0.0.8
    go: extracting github.com/mattn/go-runewidth v0.0.
    ```
- clagraff argparse no longer being developed so moved to using akamensky version