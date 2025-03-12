![GitHub Release](https://img.shields.io/github/v/release/CETEN-BDE/accounts-backend)
![GitHub License](https://img.shields.io/github/license/CETEN-BDE/accounts-backend)
[![Maintainability](https://api.codeclimate.com/v1/badges/b3ed9e002638692d956e/maintainability)](https://codeclimate.com/github/CETEN-BDE/accounts-backend/maintainability)


## Getting Started

To get a local copy up and running follow these steps.

### Prerequisites

* Install Go
```
  https://go.dev/doc/install
```
* Install Go-Task
```
https://taskfile.dev/installation/
```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/CETEN-BDE/accounts.git
   ```
2. Setup dependencies
   ```sh
   task setup
   ```
   
<!-- USAGE EXAMPLES -->
## Usage

To start the api

```sh
go run cmd/accounts/main.go
```

### OpenApi

The request handler and type for the api are auto-generated with oapi-codegen
Use:
```sh
task regen
```
To update the file accounts.gen.go


_For more examples, please refer to the [Documentation](https://github.com/CETEN-BDE/Go-Backend-Template/wiki)_

<!-- CONTRIBUTING -->
## Contributing

See the organization's [contributing guidelines](https://github.com/CETEN-BDE/.github/CONTRIBUTING.md)
