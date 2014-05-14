
# gob

A simple go build helper

## install

```
$ go get github.com/Bodhi5/gob
```

## Overview

The folder structure and set up for a go workspace isn't always ideal when working on multiple projects and can be tedious to set up at times.
gob is just a simple project that aids you in creating the workspace on-the-fly when you build your project.
This will allow you to not have to continually set your `$GOPATH` and work with the deeply nested folder structure required, but also use the built in tools to build and run your project.

## Usage

gob uses a custom `$GOPATH` path at `$CWD/gopath` and `GOBIN` path at `$CWD/bin`
allowing for the normal `go build` tool to work with out needing the nested
folder structure usually required.


## Set up a new project to build with gob

Just run the command and you will be prompted for the project `name` and organization `org`
```
$ gob init
```

To build your project simple run the following command from your project dir

## Building your project

```
$ gob build
```

## Credit

Based on the [build](https://github.com/coreos/fleet/blob/master/build) script many of the [CoreOs](https://github.com/coreos) projects use.

## License

The MIT License (MIT)

Copyright (c) 2014 Christian Sullivan &lt;cs@bodhi5.com&gt;

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.