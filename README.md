# Install

Create the file config.go
This file is used to configure the server

##### Example config

    package main

    const cfPerPage = 50
    const cfZoneID = "0f48e47b9abd78832f20144b417e671"
    const cfApiKey = "e61efc9086509aa52181f10d242ff98163813"
    const cfEndpoint = "https://api.cloudflare.com/client/v4/zones"
    const cfEmail = "email@example.org"

    const userName = "username"
    const userPass = "password"
    const httpRealm = "example"

    const ttl int64 = 300
    const port = 2130

##### Build server

    ./build.sh

##### Run server
    ./cf-proxy
    
The server will listen on localhost.
It is highly recomended to run nginx or similar in front to provide https.

# Authentication

### Basic
    curl -u "username:password" 

### Headers
    curl -H "Api-User: username" -H "Api-Pass: password" 

# Actions

## add
GET /add/{name}/{type}/{content}

###### Example
/add/foo.example.org/a/41.23.21.19

###### Response

    {
        "result": {
            "id": "ec7cc01e10e674bc152e8c870c6c1296",
            "type": "A",
            "name": "foo.example.org",
            "content": "41.23.21.19",
            "modified_on": "2017-05-01T17:46:56.886666Z",
            "created_on": "2017-05-01T17:46:56.886666Z",
            "ttl": 300
        },
        "success": true
    }

## set
GET /set/{id}/{content}

###### Example
/set/ec7cc01e10e674bc152e8c870c6c1296/80.32.12.10

###### Response

    {
        "result": {
            "id": "ec7cc01e10e674bc152e8c870c6c1296",
            "type": "A",
            "name": "foo.example.org",
            "content": "80.32.12.10",
            "modified_on": "2017-05-01T17:46:56.886666Z",
            "created_on": "2017-05-01T17:46:56.886666Z",
            "ttl": 300
        },
        "success": true
    }

### get
GET /get/{id}

###### Example
/get/ec7cc01e10e674bc152e8c870c6c1296

###### Response

    {
        "result": {
            "id": "ec7cc01e10e674bc152e8c870c6c1296",
            "type": "A",
            "name": "foo.example.org",
            "content": "80.32.12.10",
            "modified_on": "2017-05-01T17:46:56.886666Z",
            "created_on": "2017-05-01T17:46:56.886666Z",
            "ttl": 300
        },
        "success": true
    }

### delete
GET /delete/{id}

###### Example
GET /delete/ec7cc01e10e674bc152e8c870c6c1296

###### Response

    {
        "result": {
            "id": "ec7cc01e10e674bc152e8c870c6c1296",
            "ttl": 0
        },
        "success": true
    }

### list
GET /list

###### Response

    [
        {
            "id": "8679e713e4b6ebacc517cf9ad6cc34f9",
            "type": "A",
            "name": "example.org",
            "content": "91.32.12.12",
            "modified_on": "2017-05-01T17:46:39.539041Z",
            "created_on": "2017-05-01T17:46:39.539041Z",
            "ttl": 300
        },
        {
            "id": "05b1b55b8b131bac6deb0790d97c4a56",
            "type": "A",
            "name": "www.example.org",
            "content": "78.12.32.10",
            "modified_on": "2017-05-01T17:46:56.886666Z",
            "created_on": "2017-05-01T17:46:56.886666Z",
            "ttl": 300
        },
        {
            "id": "3ecc41bc5d046ddb534cbf6415af294b",
            "type": "A",
            "name": "mail.example.org",
            "content": "81.54.81.23",
            "modified_on": "2017-06-26T11:08:21.615539Z",
            "created_on": "2017-06-26T11:08:21.615539Z",
            "ttl": 300
        }
    ]
    