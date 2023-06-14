# Oohooh-Aahaah Interpreter

[![Go](https://github.com/claudemuller/oohooh-aahaah-go/actions/workflows/go.yml/badge.svg)](https://github.com/claudemuller/oohooh-aahaah-go/actions/workflows/go.yml)

```
              __,__                                  _                       _
     .--.  .-"     "-.  .--.          ___     ___   | |__     ___     ___   | |__
    / .. \/  .-. .-.  \/ .. \        / _ \   / _ \  | '_ \   / _ \   / _ \  | '_ \
   | |  '|  /   Y   \  |'  | |      | (_) | | (_) | | | | | | (_) | | (_) | | | | |
   | \   \  \ 0 | 0 /  /   / |       \___/   \___/  |_| |_|  \___/   \___/  |_| |_|
    \ '- ,\.-"""""""-./, -' /
     ''-' /_   ^ ^   _\ '-''                         _                       _
        |  \._     _./  |             __ _    __ _  | |__     __ _    __ _  | |__
         \   \ '~' /   /             / _` |  / _` | | '_ \   / _` |  / _` | | '_ \
          '._ '-=-' _.'             | (_| | | (_| | | | | | | (_| | | (_| | | | | |
             '-----'                 \__,_|  \__,_| |_| |_|  \__,_|  \__,_| |_| |_|

```
Is an interpreter for the [Monkey language](https://monkeylang.org/) built in Go.

## Usage

One can use the REPL or feed a `.mon` file into the interpreter. After compiling the binary, use the interpreter like so:

```bash
ohah examples/add.mon
```

## Running the REPL

```bash
make repl
```

## Building the interpreter

```bash
make build
```

## Running the tests

```bash
make test
```
