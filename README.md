[![Build Status](https://travis-ci.com/tamada/qrg.svg?branch=master)](https://travis-ci.com/tamada/qrg)
[![codebeat badge](https://codebeat.co/badges/9aea5795-9f10-4dc2-b63b-d4e12f3aed3f)](https://codebeat.co/projects/github-com-tamada-qrg-master)
[![License](https://img.shields.io/badge/License-WTFPL-blue.svg)](https://github.com/tamada/qrg/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-1.0.0-yellowgreen.svg)](https://github.com/tamada/qrg/releases/tag/v1.0.0)


# qrg

QR code generator.

## Install

### Homebrew

```
$ brew tap tamada/brew
$ brew install qrg
```


### Go lang

```
$ go get github.com/tamada/qrg
```


## Usage

```
$ qrg --help
qrg [OPTIONS] <CONTENT>
OPTIONS
    -d, --dest <DEST>      specifies destination.
                           if this option did not specified,
                           it generates to 'qrcode.png'.
    -l, --level <LEVEL>    specifies error correction level.
                           available values: L, M, Q, H.  default is M.
                           L: 7%, M: 15%, Q: 25%, H: 30%.
    -s, --size <SIZE>      specifies pixel size of QR code. default is 256.
    -h, --help             print this message.
    -v, --version          print version.
CONTENT
    specifies content of QR code.
```


## License

[WTFPL](https://github.com/tamada/blogthumbs/blob/master/LICENSE)
