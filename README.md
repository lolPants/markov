# 💬 Markov ![Golang Build](https://github.com/lolPants/markov/workflows/Golang%20Build/badge.svg)
> CLI tool to analyse text and generate like sentences using markov chains.

## 💾 Installation
Markov is built into a single binary. You can build this yourself with Golang tooling or download a prebuilt release from the [Releases](https://github.com/lolPants/markov/releases) page.

Nightly builds are also available as artifacts on the [Actions](https://github.com/lolPants/markov/actions?query=workflow%3A%22Golang+Build%22) page.

## 🚀 Usage
Markov is a CLI tool, and as such you can read the included help with `markov --help`.

```
Available Commands:
  analyse     reads lines from stdin and outputs a model file to stdout
  generate    uses a model from stdin to output new strings to stdout
  help        Help about any command
  version     Print version information

Flags:
  -h, --help      help for markov
  -v, --version   print version

Use "markov [command] --help" for more information about a command.
```

## 📋 Example
> This example assumes you are running a Bash-like shell and have already organised your sample lines into a line separated `lines.txt` file.

```sh
# Store model for repeated use
$ markov analyse < ./lines.txt > ./model.json
$ markov generate < ./model.json

# As a one liner
$ cat ./lines.txt | markov analyse | markov generate
```
