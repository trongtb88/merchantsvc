## How to run this service at localhost
1. Start your mysql at your localhost machine successfully
2. Git clone this repo
3. **Change file .env which is mapped with your config (DB_USER, DB_PASSWORD, DB_HOST, SERVER_PORT)**
4. **Change DB_HOST value to 127.0.0.1**
5. Note we set Auth_usename, Auth_password at env file to authenticate APIs
6. **Create database merchant_db by yourself**
7.  Go to terminal at root of project
```sh
   go get .    
   go run src/cmd/main.go
```
8. If you change something related to swagger, run this command to update documentation, 
   it will re-generate docs folder at project source code
```sh
   swag init -g ./src/cmd/main.go
   
   It will say something like 
    2021/09/05 08:45:17 Generate swagger docs....
    2021/09/05 08:45:17 Generate general API Info, search dir:./
    2021/09/05 08:45:17 Generating entity.CreateAccountParam
    2021/09/05 08:45:17 Generating entity.CreateMemberParam
    2021/09/05 08:45:17 Generating entity.UpdateMerchantAccount
    2021/09/05 08:45:17 Generating entity.UpdateMerchantMember
```
9. FE can access this page to see API documentation
   http://localhost:8089/swagger/index.html#
   ![API documentation swagger](https://i.im.ge/2021/09/05/QSe4Tz.png)

## How to run this service at docker
1. Run DOCKER DEAMON at your machine successfully
2. Make sure don't have any image mysql is running at port 3306, otherwise you will have 1 error
3. Git clone this repo
4. Comment all values for localhost like
```
    # for localhost
    #DB_DRIVER=mysql
    #DB_USER=root
    #DB_PASSWORD=1234567890
    #DB_PORT=3306
    #DB_HOST=127.0.0.1 # For running the app without docker
    #DB_NAME=merchant_db
```
5. UNcomment all values for docker
```
    #for docker
    DB_DRIVER=mysql
    DB_USER=merchant_user
    DB_PASSWORD=merchant_password
    DB_PORT=3306
    DB_HOST=merchant-mysql
    DB_NAME=merchant_db
```
5.  Go to terminal at root of project
```sh
   chmod 755 start.sh
   ./start.sh
```

6. If have some logs at console like, server started and worked successfully

```sh
merchant_app      | We are connected to the mysql database
merchant_app      | 2021/09/05 02:22:12 /app/src/cmd/db/db.go:28
merchant_app      | [0.276ms] [rows:-] SELECT DATABASE()
merchant_app      | 
merchant_app      | 2021/09/05 02:22:12 /app/src/cmd/db/db.go:28
merchant_app      | [1.201ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'merchant_db' AND table_name = 'merchant_accounts' AND table_type = 'BASE TABLE'
merchant_app      | 
merchant_app      | 2021/09/05 02:22:12 /app/src/cmd/db/db.go:28
merchant_app      | [8.595ms] [rows:0] CREATE TABLE `merchant_accounts` (`id` bigint AUTO_INCREMENT,`merchant_code` varchar(50) UNIQUE,`merchant_name` varchar(200),`merchant_status` tinyint(1),`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,PRIMARY KEY (`id`),INDEX idx_merchant_accounts_merchant_code (`merchant_code`))
merchant_app      | 
merchant_app      | 2021/09/05 02:22:12 /app/src/cmd/db/db.go:29
merchant_app      | [0.238ms] [rows:-] SELECT DATABASE()
merchant_app      | 
merchant_app      | 2021/09/05 02:22:12 /app/src/cmd/db/db.go:29
merchant_app      | [0.633ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'merchant_db' AND table_name = 'merchant_members' AND table_type = 'BASE TABLE'
merchant_app      | 
merchant_app      | 2021/09/05 02:22:12 /app/src/cmd/db/db.go:29
merchant_app      | [9.542ms] [rows:0] CREATE TABLE `merchant_members` (`id` bigint AUTO_INCREMENT,`merchant_id` bigint(20),`member_name` varchar(200),`member_email` varchar(100) UNIQUE,`member_status` tinyint(1),`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,PRIMARY KEY (`id`),INDEX idx_merchant_members_member_email (`member_email`))
merchant_app      | 2021/09/05 02:22:12 Starting server at port:  8089

```

## How to test this service
### Tesing by command
At root folder of project, run all ingtegration tests
**You must using localhost to run test**
Uncomment all values for localhost at .env file
```
go test ./src/handler/rest
Return like below means pass all tests

 go test ./src/handler/rest 
ok      github.com/trongtb88/merchantsvc/src/handler/rest       3.633s

```
### Testing using swagger
1. FE can access this page to see API documentation
   http://localhost:8089/swagger/index.html#
   ![API documentation swagger](https://i.im.ge/2021/09/05/QSe4Tz.png)

2. **At swagger page, Click button Authorize, input Auth_usename, Auth_password in .env file and click Authorize**
3. Call APIs at this page, and if return 200, that means API worked fine.

### Testing using Posman
1. Import like below, note we are using Basic Authorization, you can using value define in .env file
2. Token = Encode_Base64(Auth.username:Auth.password)
```
    For example :  Encode_Base64(abc:123) = YWJjOjEyMw==
```
Import
```
curl --location --request POST 'http://localhost:8089/v1/accounts' \
--header 'Authorization: Basic YWJjOjEyMw==' \
--header 'Content-Type: application/json' \
--data-raw '{
    "merchant_code" : "abc4",
    "merchant_name" : "merchant1"
}'
}'
```
3. Request success
```
{
    "metadata": {
        "path": "/v1/accounts",
        "status_code": 201,
        "status": "Created",
        "error": {
            "code": "OK",
            "message": "Success"
        },
        "timestamp": "2021-09-05T02:33:44Z"
    },
    "data": {
        "id": 1,
        "merchant_code": "abc4",
        "merchant_name": "merchant1",
        "merchant_status": "Active",
        "CreatedAt": "2021-09-05T02:33:44.192Z",
        "UpdatedAt": "2021-09-05T02:33:44.192Z"
    }
}
```
4. Request failed because wrong authentication
```
{
    "metadata": {
        "path": "/v1/accounts",
        "status_code": 401,
        "status": "Unauthorized",
        "error": {
            "code": "Unauthorized",
            "message": "Unauthorized"
        },
        "timestamp": "2021-09-05T02:30:30Z"
    }
}
```
## Improve
#### Develop endpoint to Authorize when click button Authorize at Swagger
#### using JWT with expired token
#### Build more real authentication flow
#### Build dynamic sql before put to gorm or sql
#### Logger
#### Add request_id or correlation_id to trace request
#### Add some telemetry to monitors
#### Build CI Pipeline to deploy using Jenkin


