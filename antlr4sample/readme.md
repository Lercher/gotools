# ANTLR4 Sample

Assume, the .g4 file is compiled with ANTLR 4.9.2 which we see from
the first line in [sample_parser.go](sample_parser.go):

```go
// Code generated from c:\git\src\github.com\lercher\gotools\antlr4sample\Sample.g4 by ANTLR 4.9.2. DO NOT EDIT.
```

Then you need to reference the compatible runtime via

```sh
$ go get github.com/antlr/antlr4/runtime/Go/antlr@4.9.2
go: added github.com/antlr/antlr4 v0.0.0-20210311221813-5e5b6d35b418
```

## Using VS Code

Install <https://marketplace.visualstudio.com/items?itemName=mike-lischke.vscode-antlr4>. Then use this VSCode settings:

```json
"antlr4.generation": {
    "mode": "external",
    "language": "Go",
    "listeners": true,
    "visitors": false,
    "outputDir": "parser"
},
```
