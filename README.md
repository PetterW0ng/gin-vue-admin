﻿
English | [简体中文](./README-zh_CN.md)

# Project Guidelines
[Online Documentation](http://doc.henrongyi.top/)

- Web UI Framework：[element-ui](https://github.com/ElemeFE/element)  
- Server Framework：[gin](https://github.com/gin-gonic/gin) 

## 1. Basic Introduction

### 1.1 Project Introduction

[Online Demo](http://qmplus.henrongyi.top/)
> Gin-vue-admin is a full-stack (frontend and backend separation) framework designed for management system. 
> It integrates multiple functions, such as JWT authentication, dynamic routing, dynamic menu, casbin authentication, form generator, code generator, etc. So that you can focus more time on your business Requirements.

### 1.2 Contributing Guide
#### 1.2.1 Issue Guidelines

- Issues are exclusively for bug reports, feature requests and design-related topics. Other questions may be closed directly. If any questions come up when you are using Element, please hit [Gitter](https://gitter.im/element-en/Lobby) for help.

- Before submitting an issue, please check if similar problems have already been issued.

#### 1.2.2 Pull Request Guidelines

- Fork this repository to your own account. Do not create branches here.

- Commit info should be formatted as `[File Name]: Info about commit.` (e.g. `README.md: Fix xxx bug`)

- <font color=red>Make sure PRs are created to `develop` branch instead of `master` branch.</font>

- If your PR fixes a bug, please provide a description about the related bug.

- Merging a PR takes two maintainers: one approves the changes after reviewing, and then the other reviews and merges.

### 1.3 Version list

- master: 2.0 dev code, for prod

- develop: 2.0 dev code, for test

## 2. Getting started
```
- node version > v8.6.0
- golang version >= v1.11
- IDE recommendation: Goland
- After you clone the project, use the scripts in directory db to create your own database.
- We recommend you to apply for your own cloud service in QINIU. Replace the public key, private key, warehouse name and default url address with your own, so as not to mess up the test database.
```

### 2.1 Web

```bash
# clone the project
git clone https://github.com/piexlmax/gin-vue-admin.git

# enter the project directory
cd web

# install dependency
npm install

# develop
npm run serve
```

### 2.2 Server

```bash
# using go.mod

# install go modules
go list (go mod tidy)

# build the server
go build
```

### 2.3 API docs auto-generation using swagger

#### 2.3.1 install swagger 

##### (1) Using VPN or outside mainland China
````
go get -u github.com/swaggo/swag/cmd/swag
````

##### (2) In mainland China 
In mainland China, access to go.org/x is prohibited，we recommend `gopm`
````bash
# install gopm
go get -v -u github.com/gpmgo/gopm

# get swag
gopm get -g -v github.com/swaggo/swag/cmd/swag

# cd GOPATH/src/github.com/swaggo/swag/cmd/swag
go install
````

#### 2.3.2 API docs generation

````
cd server
swag init
````
After executing the above command，`docs` will show in `server/`，then open your browser, jump into `http://localhost:8888/swagger/index.html` to see the swagger APIs.


## 3. Technical selection

- Frontend: using `Element-UI` based on vue，to code the page.
- Backend: using `Gin` to quickly build basic RESTful API. `Gin` is a web framework written in Go (Golang).
- DB: `MySql`(5.6.44)，using `gorm` to implement data manipulation, added support for SQLite databases.
- Cache: using `Redis` to implement the recording of the JWT token of the currently active user and implement the multi-login restriction.
- API: using Swagger to auto generate APIs docs。
- Config: using `fsnotify` and `viper` to implement `yaml` config file。
- Log: using `go-logging` record logs。

## 4. Project Architecture

### 4.1 Architecture Diagram

![Architecture diagram](./docs/gin-vue-admin.png)

### 4.2 Front-end Detailed Design Diagram (Contributor: <a href="https://github.com/baobeisuper">baobeisuper</a>)

![Front-end Detailed Design Diagram](http://qmplusimg.henrongyi.top/naotu.png)

### 4.3 Project Layout

```
    ├─server  	     （backend）
    │  ├─api            （API entrance）
    │  ├─config         （config file）
    │  ├─core  	        （core code）
    │  ├─db             （db scripts）
    │  ├─docs  	        （swagger APIs docs）
    │  ├─global         （global objet）
    │  ├─initialiaze    （initialiazation）
    │  ├─middleware     （middle ware）
    │  ├─model          （model and services）
    │  ├─resource       （resources, such as static pages, templates）
    │  ├─router         （routers）
    │  ├─service         (services)
    │  └─utils	        （common utilities）
    └─web            （frontend）
        ├─public        （deploy templates）
        └─src           （source code）
            ├─api       （frontend APIs）
            ├─assets	（static files）
            ├─components（components）
            ├─router	（frontend routers）
            ├─store     （vuex state management）
            ├─style     （common styles）
            ├─utils     （frontend common utilitie）
            └─view      （pages）

```

## 5. Features

- Authority management: Authority management based on `jwt` and `casbin`. 
- File upload & download: File upload operation based on Qiniu Cloud (In order to make it easier for everyone to test, I have provided various important tokens of my Qiniu test number, and I urge you not to make things a mess).
- Pagination Encapsulation：The frontend uses mixins to encapsulate paging, and the paging method can call mixins
- User management: The system administrator assigns user roles and role permissions.
- Role management: Create the main object of permission control, and then assign different API permissions and menu permissions to the role.
- Menu management: User dynamic menu configuration implementation, assigning different menus to different roles.
- API management: Different users can call different API permissions.
- Configuration management: The configuration file can be modified in the web page (the test environment does not provide this function).
- Rich text editor: Embed MarkDown editor function.
- Conditional search: Add an example of conditional search.
- Restful example: You can see sample APIs in user management module.

```
fontend code file: src\view\superAdmin\api\api.vue 
backend code file: model\dnModel\api.go 
```
- Multi-login restriction: Change `userMultipoint` to true in `system` in `config.yaml` (You need to configure redis and redis parameters yourself. During the test period, please report in time if there is a bug).
- Upload file by chunk：Provides examples of file upload and large file upload by chunk.
- Form Builder：With the help of [@form-generator](https://github.com/JakHuang/form-generator).
- Code generator: Providing backend with basic logic and simple curd code generator.