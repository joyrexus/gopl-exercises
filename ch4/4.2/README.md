# 4.2

Write a program that prints the SHA256 hash of its standard input by default
but supports command-line flags to print the SHA384 (`-stronger`) or SHA512 (`-strongest) hash instead.


## Examples

```
% echo -n "x" | go run main.go
2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
```

```
% go run main.go < README.md
430ecc20b954ae99d87afa8ce013a2002a8153fadf3b80cfe6a98444a22e46c5
```

```
% go run main.go -stronger < README.md
8e50c94ed99592f78d8d7443a0128555a0c06f147d71f7a77a36d24bfe497e9938ea22e86a4532289834b568c3193713
```

```
% go run main.go -strongest < README.md
a8844ad0b99b791bd51ce108ff7cf7e70e342339c23fa21277eda12d22760a55b0b6da46f43765f344300ef516f1e9c7788387526f31dc262f0ab1cfbc709840
```
