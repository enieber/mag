# Mag

Project to helper connect ecomerce with proxmox by api

![image](https://github.com/enieber/mag/assets/7907068/d0f2b56d-684b-4fc1-9972-aef936a1fd0c)


## How to run

- install go
- create database sqlite in folder of project with name test.db
- run project with `go run main.go`
- run documentation with swag install local `swag init`
- access [swagger of api](http://localhost:8080/swagger/v1/index.html)

## Roadmap

- [x] create user
- [x] create product
- [x] buy product with user
- [x] payment update
- [x] create resource when payment ok
- [ ] connect with proxmox
- [ ] create vm when create resourece
- [ ] add ssh public key to vm user


## Licence  AGPL-3.0
