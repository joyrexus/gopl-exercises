# 4.2

Write a program that prints the SHA256 hash of its standard input by default
but supports a command-line flag to print the SHA384 or SHA512 hash instead.


## Examples

```
go run main.go < README.md
```

```
go run main.go -sha 384 < README.md
```

```
go run main.go -sha 512 < README.md
```

```
% echo -n "x" | go run main.go
2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
```
