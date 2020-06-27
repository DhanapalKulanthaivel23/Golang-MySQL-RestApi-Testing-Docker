# Golang UserInfo RESTful API
    - Restful API to create, update, list, and delete Users, which includes fisrtname, lastname, referencename and country.
    - Listing users based on country.
	- Notify service function enables ServerSentEvent notify other service.
    - Used test framework github.com/stretchr/testify for testing the package.
    - Used logging system like zap/zapcore which is fast level logging.
	- Containerised with docker to build and run.
	
## Directory Structure
```
Golang-MySQL-RestApi-Testing-Docker
    |--base                   		- Application Base
        |--application.go       	- Application Startup base
    |--controller              		- package controllers
        |--createuser.go        	- to handle Create User
        |--listuser.go          	- to handle List User
        |--removeuser.go        	- to handle Remove User
        |--updateuser.go        	- to handle Update User
        |--userinfo.go          	- to handle UserInfo
        |--tests                	- Test Script
            |--base.go          	- Testing Base
            |--createuser_test.go	- User Creation Test scripts
    |--libs                       	- Libraries
        |--utils                 	- Util files
            |--logger            	- Logger Files
                |--logger.go     	- to structured logger
        |--misc.go               	- Miscellaneous Env Details
        |--sqldb.go              	- Database Details
    |--Docker                
    |--main.go  
```	
	
## Local Compile and Run 
```shell
$go test
$go build
$go run main.go
```

## Post User

- http://localhost:8084/users
```json
{
    "Firstname":"firstname",
    "Lastname":"lastname",
    "Referencename":"referencename",
    "Country":"country"
}
```
## Put User

- http://localhost:8084/users/referencename
```json
{
    "Firstname":"firstname",
    "Lastname":"lastname",
    "Country":"country"
}
```
## Delete User

- http://localhost:8084/users/referencename
```json
{
    "Firstname":"firstname",
    "Lastname":"lastname",
    "Country":"country"
}
```
## Get Users

- http://localhost:8084/users/country
```json
{
    "Firstname":"firstname",
    "Lastname":"lastname",
    "Referencename":"referencename",
    "Country":"country"
}
```

## Docker Build & Run 
```shell
$docker build -t Users-Info .

$docker run -p 8080:8086 -it Users-Info
```
