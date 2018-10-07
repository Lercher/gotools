# Gotools

Several simple command line tools, written in Go (golang). Combined in a single git project. Have a look. To build the tool, unless stated otherwise, use go's standard procedure in the tools subdir: go get/build/install.

To see what each thing does, I'm afraid you have to look at the code. The only command that has significant complexity currently, is the web server based Mandelbrot set explorer [mandel](mandel).

## Verbose Build

````bat
cd <tool>
go get -v && go install -v
````

## Silent Build

````bat
cd <tool>
go get && go install
````
