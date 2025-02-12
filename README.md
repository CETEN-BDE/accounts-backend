![GitHub Release](https://img.shields.io/github/v/release/CETEN-BDE/Log-backend)
![GitHub License](https://img.shields.io/github/license/CETEN-BDE/Log-backend)
![Code Climate maintainability](https://img.shields.io/codeclimate/maintainability/CETEN-BDE/basic-project) <!-- Remove this if you do not want the codeclimate badge-->

# Go Backend Template

This is a template for go backen for ceten project

<!-- TEMPLATE INSTRUCTIONS -->

## Template content

- Default MIT license
- README.md template with badges
- Issue labels
- Issue templates
- Branch protection rules
- api
   - example.go : an exemple of api call handler
   - impl.go : basic file to create echo server
- autogen
   - project.gen.go : auto-generated file via oapi-codegen with project.openapi.yml
- cmd
   - project
      - main.go : The main to start the api
- internal
   - db
      - init.go : init gorm to connect to MariaDB
   - models
      - example.go : example of DB model
- .env.example : example of .env file you need to change it when you change the db name
- .gitignore
- docker-compose.yml : it create the password and db name
- go.mod
- go.sum
- project.openapi.yml : the openapi file to explain all your route
- Taskfile.yml : the task file to simplify the code generation

## How to use this template

1. Click on the "Use this template" button on the top right of the repo's page.
3. Follow the next steps to setup the newly created repository

### Git flow and branch protections

To follow the git flow structure, create a branch named main, and import the rulesets located in the `rulesets` folder. You can delete the folder after importing.

### Update the issue templates

Change the assignees for each issue template in the [`.github/ISSUE_TEMPLATE`](.github/ISSUE_TEMPLATE) folder.

### Update the README

The easiest way to use this README template is with a text editor that supports replacing strings with regexp (eg. vscode).

#### Code Climate maintainability badge

This badge displays the CodeClimate maintainability report for the repository. For this badge to work, you need to [add the created repository in codeclimate](https://codeclimate.com/github/repos/new).

Login with your github account, the CETEN-BDE organization should already allow access for codeclimate. Contact the CETEN-BDE maintainers if you need help for this step.

If the repository needs to remain private, or maintainability is irrelevant, remove the badge from the top of the README.

#### Fill the template

Fill the README template below these instructions. You can replace the following placeholders :

- `PROJECT_TITLE` : The project name, used in the title
- `basic-project` : The name of the repository (the part after the organization's name in the URL)

You need to change :

- **PROJECT_TITLE**.openapi.yml

In the taskfile ->
- oapi-codegen -generate="types,server,strict-server,spec" -package autogen **PROJECT_TITLE**.openapi.yml > autogen/**PROJECT_TITLE**.gen.go
- gomodifytags -all -file autogen/**PROJECT_TITLE**.gen.go -add-tags bson -w > /dev/null

If you already have MariaDB

- you need to create a Database **PROJECT_TITLE**
- change .env to to match your user:password and Database Name

Else

- cmd/**PROJECT_TITLE**/**PROJECT_TITLE**.gen.go
In .docker-compose.yml
- MYSQL_DATABASE: **PROJECT_TITLE**


In .env.example
- **PROJECT_TITLE**_BACKEND_DSN="root:changeme@tcp(127.0.0.1:3306)/**PROJECT_TITLE**?charset=utf8mb4&parseTime=True&loc=Local"

In init.go
- dsn := os.Getenv("**PROJECT_TITLE**_BACKEND_DSN")

In go.mod
- package **basic-project**

In Go file
- "basic-project/autogen"
- "basic-project/internal/models"
- "basic-project/internal/db"
- "basic-project/internal/api"


#### Remove the template instructions

When everything is setup, remove these instructions from the README. This can be done eaysily by replacing the following regexp with an empty string :

```regex
<!-- TEMPLATE INSTRUCTIONS -->(.|\n)*<!-- TEMPLATE INSTRUCTIONS -->
```

<br/>
<br/>

Following this line is the README.md template to fill.

---

<!-- TEMPLATE INSTRUCTIONS -->

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
   git clone https://github.com/CETEN-BDE/basic-project.git
   ```
2. Setup dependencies
   ```sh
   task setup
   ```
   
<!-- USAGE EXAMPLES -->
## Usage

To start the api

```sh
go run cmd/PROJECT_TITLE/main.go
```

### OpenApi

The request handler and type for the api are auto-generated with oapi-codegen
Use:
```sh
task regen
```
To update the file project.gen.go


_For more examples, please refer to the [Documentation](https://github.com/CETEN-BDE/Go-Backend-Template/wiki)_

<!-- CONTRIBUTING -->
## Contributing

See the organization's [contributing guidelines](https://github.com/CETEN-BDE/.github/CONTRIBUTING.md)
