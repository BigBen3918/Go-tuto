# tokenbased-authentication

Implement Token-based Authentication with Golang and MySQL Server 👏👏👏

## Install this app

1. Create the database in local
```
user: "root"
password: ""
database name: "goblog"
```

2. Clone this repository
```
git clone https://github.com/hashi7412/tokenbased-authentication.git <dir_name>
```

3. Download Golang packages that are used for this app
```
cd <dir_name>

go get github.com/go-sql-driver/mysql

go get golang.org/x/crypto/bcrypt
```

4. Run this app
```
go run ./
```

5. Test this app
SSH to your server on another terminal

Add an user to database
```
curl -X POST http://localhost:8081/registrations -H "Content-Type: application/x-www-form-urlencoded" -d "username=john_doe&password=EXAMPLE_PASSWORD"
```

Get a time-based token using user's credential in request of `/authentications`
```
curl -u john_doe:EXAMPLE_PASSWORD http://localhost:8081/authentications
```

Query any resource that allows authentication using the time-based token: Copy the value of `auth_token` and execute the `curl` command below and include your token in an `Authorization` header proceded by the term `Bearer`
```
curl -H "Authorization: Bearer <auth_token>=" http://localhost:8081/test
```

Attempt authenticating to the application using an invalid token (ex: `fakerandomtoken`)
```
curl -H "Authorization: Bearer fakerandomtoken" http://localhost:8081/test
```

Attempt requesting a token without a valid user account
```
curl -u john_doe:WRONG_PASSWORD http://localhost:8081/authentications
```


## Guide this repository

This repository is for authentication implementation based token with Golang using MySQL as a database

### [main.go](https://github.com/hashi7412/tokenbased-authentication/blob/main/main.go)

- main()
In main function which is executed first when the app is runned, implemented a handler function for multiple URL paths that provide functionalities.

For example
```
http.HandleFunc(<path>, <handler>)
```

- registrationsHandler()

The `registrationsHandler` function retrieves submitted `username` and `password` for any users you're adding to your system and directs the same to a `registerUser` function in a `registrations.go` file which you'll create next.

- authenticationsHandler()

Then, the `authenticationsHandler` extracts log in credentials(`username` and `password`) using the statement `req.BasicAuth()`. Then, it passes these details to a `generateToken` function under an `authentication.go` file, which you'll create later. In case the credentials match a valid account on the `system_users` table, you're issuing the user with a token.

- testResourceHandler()

Next, you have the `testResourceHandler` function. Under this function, you're retrieving the time-based token from the `Authorization` header submitted by the client's request. Then, you're passing it to a `validateToken` function under the `authentication.go` file to check if the token is valid. You're then greeting any authenticated user with a welcome message.

### [registrations.go](https://github.com/hashi7412/tokenbased-authentication/blob/main/registrations.go)

The above file has a single `registerUser` function that inserts data into your `goblog` database in the `system_users` table. You're using the statement `hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)` to hash the plain-text passwords for security purposes. The function returns a `Success` message once you've created a user into the database.

You've imported the `database/sql`, `github.com/go-sql-driver/mysql`, and `golang.org/x/crypto/bcrypt` packages to implement MySQL database and password hashing functions.

### [authentications.go](https://github.com/hashi7412/tokenbased-authentication/blob/main/authentications.go)

In the `generateToken` function, you're accepting a `username` and a `password`. Then, you're running a `SELECT` statement against the `system_users` table to check if a record exists with that username. You're then using the statement `if err == sql.ErrNoRows {}` to determine if there is a matching row for the user. If the user doesn't exist, you're throwing an `Invalid username or password`. error. However, if there is a matching record, you're using the statement `bcrypt.CompareHashAndPassword([]byte(accountPassword), []byte(password))` to determine if the account's password and the supplied password match.

Next, you're using `randomToken := make([]byte, 32)` and `_, err = rand.Read(randomToken)` statements to generate a random token for the user. You're later encoding the token to `base64` using the statement `base64.URLEncoding.EncodeToString(...)`. then, you're permanently saving the token to the authentication_tokens table.

In the `validateToken` function, you're checking the provided token on the `authentication_tokens` table to see if there is a match. If the token is valid, you're returning detailed information about the token, including the matching user's details and token values. Otherwise, you're throwing an error to the calling function.

You're using the statement if `expiryTime.Before(currentTime) {...}` to check if the token has expired.

### [dbconn.go](https://github.com/hashi7412/tokenbased-authentication/blob/main/dbconn.go)

This file has `dbConn()` function to connect database. If it has error, it will stop this program immediately with `panic` finction

## Conclusion

In this repository, we've implemented token-based authentication with Golang and MySQL

Here is some repositories for your guide:
[Hands-on Go](https://github.com/hashi7412/handson-go)
[CRUD with MySQL](https://github.com/hashi7412/crud-with-mysql)

Thank you for looking at this repository. 👋
