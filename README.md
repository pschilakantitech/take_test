# take-test

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

take-test is the simple web application with microservice architecture that enables the student to take the online test. The application uses the microservice architecture backend with HTML, javascript and CSS for UI and postgres DB for data persistence;  All backend series are  integated with web UI and ready to deploy state.
 
# Insatll on linux

- Insatll latest version Golang as per https://golang.org/doc/install  
- Insatll latest version of Postgress data base. 
```sh
$ sudo apt-get update
$ sudo apt-get install postgresql postgresql-contrib
$ sudo -i -u postgres
 ```
  - Run the DB below script   https://github.com/pschilakantitech/take_test/blob/master/pg_dump/take_test_pgdb.sql
  - Build the code with modififing the DB connection details in 'env' package

```sh
$ git clone git@github.com:pschilakantitech/take_test.git   
$ cd take_test
$ make setup
#Modify DB config details than run build cmd
$ go build
$ ./take_test
 ```
- Now application is runing at https://localhost:3000, cool! take the test
- For easy access, exporting my loacl development env to public url. please use link http://64bdf336.ngrok.io


# work flow
- Using the https://echo.labstack.com for  API framework.
- As all know, the flow starts from the main package main method before that init method will check DB connectivity and hold the DB connection reference.
- In the main method, labstack echo instance is created and assigned the API end points with handler functions.
- frontend assets are written in the 'web_ui' folder and  integrated with APIs
