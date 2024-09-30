<a id="readme-top"></a>
# Payroll Absen Dulu Server Side

Stool Server Side is powered by the [Gin](https://gin-gonic.com/) framework</br> Gin is a web framework written in Golang. It features a Martini-like API, but with performance up to 40 times faster than Martini.

&nbsp;

## Server Side Made With

<b>Please make sure you have already installed and set up all the required components in order to run this app</b></br></br>
[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)][Golang-url] [![Postgres][PostgreSQL]][Postgres-url]</br>

- <b>Go </b> is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency.</br>
- <b>PostgreSQL</b>, also known as Postgres, is a free and open-source relational database management system emphasizing extensibility and SQL compliance. It was originally named POSTGRES, referring to its origins as a successor to the Ingres database developed at the University of California, Berkeley.

[PostgreSQL]: https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white
[Postgres-url]: https://www.postgresql.org/
[Golang-url]: https://go.dev/

&nbsp;

## &#128640; Getting Started

Make sure you already in a project folder, then open the folder address in the terminal and do the command below

- Installation

```go
$ go get -u
$ go build
```

- How to run

```go
$ go run .
```

_It is recommended to install [Git Bash](https://git-scm.com/downloads) first_

&nbsp;

## Hot Reload so you don't switch to terminal ffs
Using Air - Live reload for Go apps
full docs https://github.com/cosmtrek/air

- Installation
```go 
$ go install github.com/cosmtrek/air@latest
```
The simplest usage is run. Make sure you're in folder project.
```
# firstly find `.air.toml` in current directory, if not found, use defaults
air -c .air.toml
```

You can initialize the .air.toml configuration file to the current directory with the default settings running the following command.

```
air init
```

After this, you can just run the air command without additional arguments and it will use the .air.toml file for configuration.

```
air
```

<p align="right">(<a href="#readme-top">back to tsop</a>)</p>