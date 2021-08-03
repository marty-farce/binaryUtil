# Binary Data probe simulator

this Library encodes randomly generated probe data into binary and also decodes it

## Dependencies

* Go 1.16.6 or later

## Implementing this library

``` bash
go get github.com/Teralytic/probe-sim
```

## Usage

* Encode(PacketNum) \*binary.Buffer returns a single binary encoded packet
* Decode(\*binary.Buffer) decodes the binary
