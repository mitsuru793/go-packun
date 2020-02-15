# go-packun

This is a cross package manger. You can collect some packages into one file.  
Here is a store file. This has brew and yarn package manager in one file.

```
# ~/.go-packun/store.toml

[[Packages]]
  Manager = "brew"
  Name = "fzf"
  Tags = ["text"]

[[Packages]]
  Manager = "yarn"
  Name = "create-react-app"
  Tags = ["boilerplate"]
```

## Installation

From souse:
```
$ go get github.com/mitsuru793/go-packun
```

## Usage

`go-packun help` lists commands.

### Add

It's interactive and validate package name.

```
$ go-packun add
which package
> brew
package name
> tig
tags(split by space)
> git
```

Package tig will be installed by brew. Its meta information is stored `~/.go-packun/store.toml`.

### Install

Install all packages from `~/.go-packun/store.toml`.

```
$ go-packun install
```
