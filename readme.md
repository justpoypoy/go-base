# golang base api with clean architecture
A boilerplate built with go and clean architecture.
Most of this application built by standard libray.

![image](https://user-images.githubusercontent.com/13291041/102681893-84326980-4208-11eb-8f84-2959e03b89d8.png)

# Get Started
`cp app/.env_example app/.env`

# Architecture
```
.
├── domain
│   └── user.go
├── go.mod
├── go.sum
├── infrastructure
│   ├── env.go
│   ├── logger.go
│   ├── router.go
│   └── sql.go
├── interfaces
│   ├── sql.go
│   ├── user_controller.go
│   └── user_repository.go
├── log
│   ├── access.log
│   └── error.log
├── main.go
├── readme.md
└── usecases
    ├── logger.go
    ├── user_interactor.go
    └── user_repository.go

5 directories, 17 files
```

| Layer                | Directory      |
|----------------------|----------------|
| Frameworks & Drivers | infrastructure |
| Interface            | interfaces     |
| Usecases             | usecases       |
| Entities             | domain         |

# API

| ENDPOINT     | HTTP Method    | Parameters   |
|--------------|----------------|--------------|
| /users       | GET            |              |
| /user/:id    | GET            | id=[int]     |
| /user        | POST           |              |
| /user/:id    | PUT            | id=[int]     |
| /user/:id    | DELETE         | id=[int]     |

# Controller method naming rule

| Controller Method | HTTP Method | Description                                |
|-------------------|-------------|--------------------------------------------|
| Index             | GET         | Display a listing of the resource          |
| Store             | POST        | Store a newly created resource in storage  |
| Show              | GET         | Display the specified resource             |
| Update            | PUT/PATCH   | Update the specified resource in storage   |
| Destroy           | DELETE      | Remove the specified resource from storage |

# Repository method naming rule

| Repository Method | Description                                          |
|-------------------|------------------------------------------------------|
| FindByXX          | Returns the entity identified by the given XX        |
| FindAll           | Returns all entities                                 |
| Save              | Saves the given entity                               |
| SaveByXX          | Saves the given entity identified by the given XX    |
| DeleteByXX        | Deletes the entity identified by the given XX        |
| Count             | Returns the number of entities                       |
| ExistsBy          | Indicates whether an entity with the given ID exists |

cf. [Spring Data JPA - Reference Documentation](https://docs.spring.io/spring-data/data-jpa/docs/current/reference/html/#repositories.core-concepts)

# Tests
On going.

# How to read code
| Step | Files              | Directory      |
|------|--------------------|----------------|
| 1    | main.go            | main directory |
| 2    | router.go          | infrastructure |
| 3    | user_controller.go | interfaces     |
| 4    | user_interactor.go | usecases       |
| 5    | user_repository.go | usecases       |
| 6    | user.go            | domain         |



# References
- [Dive to Clean Architecture with Golang](https://dev.to/bmf_san/dive-to-clean-architecture-with-golang-cd4)