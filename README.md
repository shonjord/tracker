Computer Tracker
===============

<img width="250" align="center" alt="Flash" src="https://www.pngkey.com/png/full/55-553923_banner-free-stock-free-computers-cliparts-computer-clipart.png" />

computer tracker is a microservice written in [go](https://golang.org/) responsible to manage the assignment ownership of computers from employees.

## Table of Contents

* [Maintainers](#maintainers)
* [Architecture](#architecture)
* [Getting started](#getting-started)
    * [Prerequisites](#prerequisites)
    * [Development](#development)
    * [Testing](#testing)
* [Points Missing](#points-missing)

## Maintainers

* [Albert Agelviz] aagelviz@gmail.com

[[table of contents]](#table-of-contents)

## Architecture
<img width="600" align="center" alt="Flash" src="https://herbertograca.files.wordpress.com/2018/11/100-explicit-architecture-svg.png?w=1200" />

The architecture chosen for this service is "Hexagonal Architecture" or "Ports & Adapters", the idea behind this architecture is scalability, readability, maintainability, compatible in its entirety with DDD (Domain Driven Design), the layers of this architecture are:

- Application
  - Responsible to orchestrate and process the intents of the domain (use cases)
- Domain
  - Layer where all entities, intents are encapsulated and isolated from the outside
- Infrastructure
  - Here lies all the external communications (Database, HTTP clients, Brokers)
- Presentation
  - This layer is the portion of "Ports" in "Ports & Adapters", here lies all inputs (HTTP, HTTP2, TCP, brokers, etc), the idea is that this layer obtains the inputs, encapsulate the input into a "domain command" or a "domain query" and uses the application layer to orchestrate and process the intent, then it knows what to response (JSON, XML, PDF, HTML, etc) 

Other two patterns have been applied along with this architecture
- CQRS
  - Essentially the idea behind this is to separate READS from WRITES, having them separated allows to define better the intentions of our Domain (AssignNewComputer, UnassignComputer, GetComputerByUUID, etc)
- ADR (Action-Domain-Responder)
  - An architectural pattern to replace MVC, the idea behind this architectural pattern is to have the presentation layer define with multiple components whose follow SRP (Single Responsibility Principle), allowing to connect the domain with its intended response. 

[[table of contents]](#table-of-contents)

## Getting started

### Prerequisites

**Required GO Path**
- [GoPath](https://github.com/golang/go/wiki/SettingGOPATH)
    * This PATH is mandatory to be able to work with this service

**Required Tools**
- [Docker](https://docs.docker.com/docker-for-mac/install/)
    * Also, available via HomeBrew, `brew install docker docker-compose && brew cask install docker`

[[table of contents]](#table-of-contents)

### Development

Clone this repository:
```bash
git clone git@github.com:shonjord/tracker.git
```

install a daemon in case this package is missing.
````bash
make daemon
````

Pull and start containers by executing:
```bash
make container-up
```

While developing is very common and useful to tail the logs, for this you can execute the following command on your CLI:
```bash
docker-compose logs -f {container}
```

Keep also an eye of the logs of the administration notification service to watch the notifications when an employee has been assigned with 3 or more computers.
```bash
docker-compose logs -f admin-notification
```

[[table of contents]](#table-of-contents)

### Testing

#### Unit Tests

To run unit tests, execute the following make target:
```bash
make test-unit
```
[[table of contents]](#table-of-contents)

## Points Missing

Due to lack of time, there are a couple of things that could've been added to improve.
- Integration tests, to test the end to end cycle of the endpoints.
- Any Rest API documentation (RAML, OpenAPI).
- Because an API documentation was not possible to add, a postman collection has been shared inside `./resources/postman`
- Adding endpoints to manage CRUD operations of the employees (right now these are added through migrations)

[[table of contents]](#table-of-contents)
