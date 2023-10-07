# Simplified Go interface for HTTP Golang Filter on Envoy Proxy

> A simplified interface for building plugin on Envoy Proxy with [HTTP Golang Filter](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/golang_filter).

## Features

* Porting `net/http` interface experience to extend Envoy Proxy behavior with [HTTP Golang Filter](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/golang_filter).

## Development Guide

### Prerequisites

* Go 1.20 or later. Follow [Golang installation guideline](https://golang.org/doc/install).

### Setup

* Install Git.

* Install Go 1.20.

* Clone the project.

    ```bash
    $ git clone -b plugin git@github.com:ardkabs/go-envoy.git
    ```

* Create a meaningful branch

    ```bash
    $ git checkout -b <your-meaningful-branch>
    ```

* Test your changes.

    ```bash
    $ make test
    ```

* We highly recommend instead of only run test, please also do audit which include formatting, linting, vetting, and testing.

    ```bash
    $ make audit
    ```

* Add, commit, and push changes to repository

    ```bash
    $ git add .
    $ git commit -s -m "<conventional commit style>"
    $ git push origin <your-meaningful-branch>
    ```

    For writing commit message, please use [conventionalcommits](https://www.conventionalcommits.org/en/v1.0.0/) as a reference.

* Create a Pull Request (PR). In your PR's description, please explain the goal of the PR and its changes.

### Testing

#### Unit Test

```bash
$ make test
```

### Try It

To try this interface in action, heads to [example](./example) directory.
