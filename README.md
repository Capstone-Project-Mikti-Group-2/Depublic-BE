
![Logo](https://yt3.googleusercontent.com/jDWEWvp_TO-Nnkow5uJYHQ03vAFVQ20K-WQ8nPOkgIr6aVBdLZB6DRJX0Xlip2hjp1HahjMZbw=s900-c-k-c0x00ffffff-no-rj)

# **API DOCUMENTATION DEPUBLIC-APP**

## BASE-URL
http://localhost:8080/api/v1

## **API-DOCUMENTATION**
- [Authentication and Authorization](#Authentication-and-Authorization)
    - [Register](#register)
    - [Login](#login)
    - [Logout](#logout)
- [Users](#users)
    - [Create User](#create-user)
    - [Update User](#update-user)
    - [Delete User](#delete-user)
    - [Get All Users](#get-alluser)
    - [Get User By Id](#get-userbyid)
    - [Get User By Email](#get-userbyemail)
    - [Get User By Number](#get-userbynumber)
    - [Get User By Name](#get-userbyname)
- [Profile](#profile)
    - [Create Profile](#create-profile)
    - [Update Profile](#update-profile)
    - [Delete Profile](#delete-profile)
    - [Delete Account](#delete-account)
- [Event](#event)
    - [Create Event](#create-event)
    - [Update Event](#update-event)
    - [Delete Event](#delete-event)
    - [Get All Event](#get-event)
    - [Get Event By Id](#get-eventbyid)
    - [Filter By Search Event](#filter-searchevent)
    - [Filter By Price Event](#filter-priceevent)
    - [Filter By Location Event](#filter-locationevent)
    - [Filter By Available Event](#filter-availableevent)
    - [Filter By Date Event](#filter-dateevent)
    - [Sort By Cheapest Event](#sort-cheapestevent)
    - [Sort By Expensive Event](#sort-expensiveevent)
    - [Sort By Newest Event](#sort-newestevent)
- [Ticket](#ticket)
    - [Create Ticket](#create-ticket)
    - [Get All Ticket](#get-allticket)
    - [Get Ticket By User Id](#get-ticketbyuserid)
- [Transaction Topup](#transaction-topup)
    - [Create Transaction](#create-transaction)
    - [Get Transaction History](#get-transactionhistory)
    - [Input Saldo Without Midtrans](#inputsaldo-withoutmidtrans)
- [Notification](#notification)
    - [Create Notification](#create-notification)
    - [Get Notification](#get-notification)
    - [User Get Notification](#user-getnotification)

# **API REFERENCE**

## Authentication and Authorization

### Register
- Endpoint :
    - /register
- Method :
    - POST
- BODY :
```json
{
    "name" : "string, no whitespace, required",
    "email": "string, email, required",
    "password": "string, min:6, required",
    "number": "int, required",
    "role" : "onof= Administrator & User"
}

```
- RESPONSE :
```json
{
    "email": "Admin1@gmail.com",
    "nama": "Admin1",
    "number": "011111111111",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywibmFtZSI6IkFkbWluMSIsImVtYWlsIjoiQWRtaW4xQGdtYWlsLmNvbSIsInJvbGUiOiJBZG1pbmlzdHJhdG9yIiwiZXhwIjoxNzAyMjAzMTE5fQ.wYHU5KCeuhgvDRVH-PZRkpPEi-rK-MW51KnbuN9HqmQ"
}
```
### Login
- Endpoint :
    - /login
- Method :
    - POST
- BODY :
```json
{
    "email": "Admin1@gmail.com",
    "password": "Admin111"
}
```
- RESPONSE
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6IkFkbWluMSIsImVtYWlsIjoiQWRtaW4xQGdtYWlsLmNvbSIsInJvbGUiOiJBZG1pbmlzdHJhdG9yIiwiZXhwIjoxNzAyMjAzMzYwfQ.Ku1s-G4bBZYVssa0DDYp1TBmSqlAAih6Nd3f1CGlkHM",
    "user": {
        "id": 1,
        "username": "Admin1",
        "email": "Admin1@gmail.com",
        "number": "011111111111",
        "role": "Administrator",
        "saldo": 9700000,
        "profile": null,
        "createdAt": "2023-12-09T22:15:21.743238+07:00",
        "updatedAt": "2023-12-10T15:38:32.230923+07:00",
        "deletedAt": "0001-01-01T07:00:00+07:00"
    }
}
```

### Logout
- Endpoint :
    - /users/logout
- Method :
    - POST
- Response :
```json 
{
    "message": "success logout"
}
```
## Users

### Create User
- Endpoint :
    - /users
- Method :
    - POST
- BODY :
```json
{
    "name": "Budi Santoso",
    "email": "budi.santoso@email.com",
    "password": "userpassword2",
    "number": "082345678901",
    "role": "User"
}

```
- RESPONSE :
```json
{
    "created_at": "2023-12-10T17:47:47.7057222+07:00",
    "message": "user created successfully",
    "user": {
        "id": 5,
        "username": "Budi Santoso",
        "email": "budi.santoso@email.com",
        "number": "082345678901",
        "role": "User",
        "saldo": 0,
        "profile": null,
        "createdAt": "2023-12-10T17:47:47.7057222+07:00",
        "updatedAt": "2023-12-10T17:47:47.7057222+07:00",
        "deletedAt": "0001-01-01T00:00:00Z"
    }
}
```

### Update User
- Endpoint :
    - /users/:id
- Method :
    - PUT
- BODY :
- Param
```param
{
    "id":5
}
```
- Json

```json
{
    "name": "Budi Rahayu",
    "email": "budi.rahayu@email.com",
    "password": "userpassword3",
    "number": "083456789012",
    "role": "User"
}

```
- RESPONSE :
```json
{
    "message": "success update user",
    "updated_at": "2023-12-10T17:51:10.259763+07:00",
    "user": {
        "id": 5,
        "username": "Budi Rahayu",
        "email": "budi.rahayu@email.com",
        "number": "083456789012",
        "role": "User",
        "saldo": 0,
        "profile": null,
        "createdAt": "0001-01-01T00:00:00Z",
        "updatedAt": "2023-12-10T17:51:10.259763+07:00",
        "deletedAt": "0001-01-01T00:00:00Z"
    }
}
```

### Delete User
- Endpoint :
    - /users/:id
- Method :
    - DELETE
- BODY :
- Param
```param
{
    "id":5
}
```

- RESPONSE :
```json
{
    "deleted": "2023-12-10T17:57:24.8747677+07:00",
    "id": 5,
    "message": "success delete user"
}
```

### Get All Users
- Endpoint :
    - /users
- Method :
    - GET
- RESPONSE :
```json
{
    "data": [
        {
            "created_at": "2023-12-09T22:15:46.274855+07:00",
            "email": "usersome@email.com",
            "id": 2,
            "name": "someUser",
            "number": "000000000",
            "updated_at": "2023-12-09T22:15:46.274855+07:00"
        },
        {
            "created_at": "2023-12-10T16:51:59.224289+07:00",
            "email": "Admin1@gmail.com",
            "id": 3,
            "name": "Admin1",
            "number": "011111111111",
            "updated_at": "2023-12-10T16:51:59.224289+07:00"
        }
    ]
}
```

### Get User By Id
- Endpoint :
    - /users/:id
- Method :
    - GET
- BODY :
- Param
```param
{
    "id":2
}
```
- RESPONSE :
```json
{
    "data": {
        "created_at": "2023-12-09T22:15:46.274855+07:00",
        "email": "usersome@email.com",
        "id": 2,
        "name": "someUser",
        "number": "000000000",
        "updated_at": "2023-12-09T22:15:46.274855+07:00"
    }
}
```
### Get User By Email
- Endpoint :
    - /users/email/:email
- Method :
    - GET
- BODY :
- Param
```param
{
    "email":"usersome@email.com"
}
```
- RESPONSE :
```json
{
    "data": {
        "created_at": "2023-12-09T22:15:46.274855+07:00",
        "email": "usersome@email.com",
        "id": 2,
        "name": "someUser",
        "number": "000000000",
        "updated_at": "2023-12-09T22:15:46.274855+07:00"
    }
}
```

### Get User By Number
- Endpoint :
    - /users/number/:number
- Method :
    - GET
- BODY :
- Param
```param
{
    "number":"000000"
}
```
- RESPONSE :
```json
{
    "data": {
        "created_at": "2023-12-09T22:15:46.274855+07:00",
        "email": "usersome@email.com",
        "id": 2,
        "name": "someUser",
        "number": "000000000",
        "updated_at": "2023-12-09T22:15:46.274855+07:00"
    }
}
```

### Get User By Name
- Endpoint :
    - /users/name/:name
- Method :
    - GET
- BODY :
- Param
```param
{
    "name":"someUser"
}
```
- RESPONSE :
```json
{
    "data": {
        "created_at": "2023-12-09T22:15:46.274855+07:00",
        "email": "usersome@email.com",
        "id": 2,
        "name": "someUser",
        "number": "000000000",
        "updated_at": "2023-12-09T22:15:46.274855+07:00"
    }
}
```

## Profile
### Create Profile
- Endpoint :
    - /profile
- Method :
    - POST
- BODY :
- Param
```param
{
    user_id : (from claims JWT)
}
```
- Json
```json
{
	"address": "Jl. Sudirman No. 123, Jakarta Selatan",
	"avatar": "c3RyaW5nYXJlZGF2YXRhZGF0YQ=="
}

```
- RESPONSE :
```json
{
    "created_at": "2023-12-10T22:03:11.5329427+07:00",
    "data": {
        "id": 1,
        "user_id": 1,
        "address": "Jl. Sudirman No. 123, Jakarta Selatan",
        "avatar": "c3RyaW5nYXJlZGF2YXRhZGF0YQ==",
        "createdAt": "2023-12-10T22:03:11.5329427+07:00",
        "updatedAt": "2023-12-10T22:03:11.5329427+07:00",
        "deletedAt": "0001-01-01T00:00:00Z"
    },
    "message": "profile created successfully"
}
```
### Update Profile
- Endpoint :
    - /profile
- Method :
    - PUT
- BODY :
- Param
```param
{
    user_id : (from claims JWT)
}
```
- Json
```json
{
	"address": "Jl. MH Thamrin No. 456, Jakarta Pusat",
	"avatar": "c3RyaW5nYXJlZGF2YXRhZGF0YQ=="
}
```
- RESPONSE :
```json
{
    "message": "profile updated successfully",
    "updated_at": "2023-12-10T22:05:04.7173924+07:00"
}
```
### Update Self Profile User
- Endpoint :
    - /users/profile
- Method :
    - PUT
- BODY :
- Param
```param
{
    user_id : (from claims JWT)
}
```
- Json
```json
{
    "name": "Dewi Novianti", 
    "email": "dewi.novianti@email.com",
    "password": "userpassword5",
    "number": "085678901234"
}

```
- RESPONSE :
```json
{
    "data": {
        "id": 9,
        "username": "Dewi Novianti",
        "email": "dewi.novianti@email.com",
        "number": "085678901234",
        "role": "",
        "saldo": 0,
        "profile": null,
        "createdAt": "0001-01-01T00:00:00Z",
        "updatedAt": "2023-12-10T22:24:29.1948402+07:00",
        "deletedAt": "0001-01-01T00:00:00Z"
    },
    "message": "success update user"
}

```
### Delete Profile
- Endpoint :
    - /profile
- Method :
    - DELETE
- BODY :
- Param
```param
{
    user_id : 1
}
```
- RESPONSE :
```json
{
    "deleted": "2023-12-10T22:27:38.4895179+07:00",
    "id": 9,
    "message": "success delete profile"
}
```
### Delete Account
- Endpoint :
    - /users/profile
- Method :
    - PUT
- BODY :
- Param
```param
{
    user_id : (from claims JWT)
}
```
- RESPONSE :
```json
{
    "id": 10,
    "message": "success delete account"
}

```


## EVENT

### Create Event
- Endpoint :
    - /events
- Method :
    - POST
- BODY :
```json
{
    "name": "Coldplay Jakarta",
    "description": "Coldplay akan konser di jakarta untuk merayakan greenpace dunia",
    "location": "Jakarta - Gelora Bung Karno",
    "price": 300000,
    "quantity": 5000,
    "image": null,
    "start_date": "2023-11-05",
    "end_date": "2023-11-08",
    "available": true
}

```
- RESPONSE :
```json
{
    "created_at": "2023-12-10T22:59:13.2737405+07:00",
    "data": {
        "id": 5,
        "name": "Coldplay Jakarta",
        "description": "Coldplay akan konser di jakarta untuk merayakan greenpace dunia",
        "location": "Jakarta - Gelora Bung Karno",
        "image": null,
        "price": 300000,
        "quantity": 5000,
        "available": true,
        "start_date": "2023-11-05T00:00:00Z",
        "end_date": "2023-11-08T00:00:00Z",
        "createdAt": "2023-12-10T22:59:13.2737405+07:00",
        "updatedAt": "2023-12-10T22:59:13.3610989+07:00",
        "deletedAt": "0001-01-01T00:00:00Z"
    },
    "message": "event created successfully"
}
```

### Update Event
- Endpoint :
    - /events/:id
- Method :
    - PUT
- BODY :
- Param :
```param
{
    "id" : 3
}
```
- JSON :
```json
{
    "name": "Country Roads Jamboree",
    "description": "Get ready for a country music adventure on the open road",
    "location": "Country Trails Fairgrounds - Nashville",
    "price": 240000,
    "quantity": 6900,
    "image": null,
    "start_date": "2023-02-14",
    "end_date": "2023-02-17",
    "available": true
}
```
- RESPONSE :
```json
{
    "data": {
        "id": 3,
        "name": "Country Roads Jamboree",
        "description": "Get ready for a country music adventure on the open road",
        "location": "Country Trails Fairgrounds - Nashville",
        "image": null,
        "price": 240000,
        "quantity": 6900,
        "available": true,
        "start_date": "2023-02-14T00:00:00Z",
        "end_date": "2023-02-17T00:00:00Z",
        "createdAt": "0001-01-01T00:00:00Z",
        "updatedAt": "2023-12-10T23:18:43.767752+07:00",
        "deletedAt": "0001-01-01T00:00:00Z"
    },
    "message": "event updated successfully",
    "updated_at": "2023-12-10T23:18:43.767752+07:00"
}
```
### Delete Event
- Endpoint :
    - /events/:id
- Method :
    - DELETE
- BODY :
- Param :
```param
{
    "id" : 4
}
```
- RESPONSE :
```json
{
    "deleted": "2023-12-10T23:21:46.2115513+07:00",
    "id": 4,
    "message": "event deleted successfully"
}
```
### Get All Event
- Endpoint :
    - /events
- Method :
    - GET
- RESPONSE :
```json
{
    "data": [
        {
            "id": 5,
            "name": "Coldplay Jakarta",
            "description": "Coldplay akan konser di jakarta untuk merayakan greenpace   dunia",
            "location": "Jakarta - Gelora Bung Karno",
            "image": null,
            "price": 300000,
            "quantity": 5000,
            "available": true,
            "start_date": "2023-11-05T00:00:00Z",
            "end_date": "2023-11-08T00:00:00Z",
            "createdAt": "2023-12-10T22:59:13.27374+07:00",
            "updatedAt": "2023-12-10T22:59:13.361098+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        },
        {
            "id": 3,
            "name": "Country Roads Jamboree",
            "description": "Get ready for a country music adventure on the open road",
            "location": "Country Trails Fairgrounds - Nashville",
            "image": null,
            "price": 240000,
            "quantity": 6900,
            "available": true,
            "start_date": "2023-02-14T00:00:00Z",
            "end_date": "2023-02-17T00:00:00Z",
            "createdAt": "2023-12-09T22:16:08.099738+07:00",
            "updatedAt": "2023-12-10T23:18:43.899515+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        },
        {
            "id": 6,
            "name": "EDM Eclipse Carnival",
            "description": "Dance through the night with pulsating EDM beats",
            "location": "Eclipse Arena - Tokyo",
            "image": null,
            "price": 300000,
            "quantity": 7800,
            "available": true,
            "start_date": "2023-06-18T00:00:00Z",
            "end_date": "2023-06-21T00:00:00Z",
            "createdAt": "2023-12-10T23:23:39.419662+07:00",
            "updatedAt": "2023-12-10T23:23:39.421199+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        }
    ]
}
```
### Get Event By Id
- Endpoint :
    - /events/:id
- Method :
    - GET
- BODY :
- Param :
```param
{
    "id" : 3
}
```
- RESPONSE :
```json
{
    "data": {
        "available": true,
        "description": "Get ready for a country music adventure on the open road",
        "end_date": "2023-02-17T00:00:00Z",
        "id": 3,
        "image": null,
        "location": "Country Trails Fairgrounds - Nashville",
        "name": "Country Roads Jamboree",
        "price": 240000,
        "quantity": 6900,
        "start_date": "2023-02-14T00:00:00Z"
    }
}
```
### Filter By Search Event
- Endpoint :
    - /events/keyword/:keyword
- Method :
    - GET
- BODY :
- Param :
```param
{
    "keyword" : country
}
```
- RESPONSE :
```json
{
    "data": [
        {
            "id": 3,
            "name": "Country Roads Jamboree",
            "description": "Get ready for a country music adventure on the open road",
            "location": "Country Trails Fairgrounds - Nashville",
            "image": null,
            "price": 240000,
            "quantity": 6900,
            "available": true,
            "start_date": "2023-02-14T00:00:00Z",
            "end_date": "2023-02-17T00:00:00Z",
            "createdAt": "2023-12-09T22:16:08.099738+07:00",
            "updatedAt": "2023-12-10T23:18:43.899515+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        }
    ]
}
```
### Filter By Price Event
- Endpoint :
    - /events/price/:min/:max
- Method :
    - GET
- BODY :
- Param :
```param
{
    "min" : 200000
    "max" : 250000
}
```
- RESPONSE :
```json
{
    "data": [
        {
            "id": 3,
            "name": "Country Roads Jamboree",
            "description": "Get ready for a country music adventure on the open road",
            "location": "Country Trails Fairgrounds - Nashville",
            "image": null,
            "price": 240000,
            "quantity": 6900,
            "available": true,
            "start_date": "2023-02-14T00:00:00Z",
            "end_date": "2023-02-17T00:00:00Z",
            "createdAt": "2023-12-09T22:16:08.099738+07:00",
            "updatedAt": "2023-12-10T23:18:43.899515+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        },
        {
            "id": 9,
            "name": "Pop Sensation Extravaganza",
            "description": "An explosive night of pop sensations and dance beats",
            "location": "Pop Dome - Miami",
            "image": null,
            "price": 220000,
            "quantity": 6500,
            "available": true,
            "start_date": "2023-08-28T00:00:00Z",
            "end_date": "2023-08-31T00:00:00Z",
            "createdAt": "2023-12-10T23:24:02.463255+07:00",
            "updatedAt": "2023-12-10T23:24:02.477812+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        }
    ]
}
```
### Filter By Location Event
- Endpoint :
    - /events/location/:location
- Method :
    - GET
- BODY :
- Param :
```param
{
    "location" : Nashville
}
```
- RESPONSE :
```json
{
    "data": [
        {
            "id": 3,
            "name": "Country Roads Jamboree",
            "description": "Get ready for a country music adventure on the open road",
            "location": "Country Trails Fairgrounds - Nashville",
            "image": null,
            "price": 240000,
            "quantity": 6900,
            "available": true,
            "start_date": "2023-02-14T00:00:00Z",
            "end_date": "2023-02-17T00:00:00Z",
            "createdAt": "2023-12-09T22:16:08.099738+07:00",
            "updatedAt": "2023-12-10T23:18:43.899515+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        }
    ]
}
```
### Filter By Available Event
- Endpoint :
    - /events/available/:available
- Method :
    - GET
- BODY :
- Param :
```param
{
    "available" : true
}
```
- RESPONSE :
```json
{
    "data": [
        {
            "id": 5,
            "name": "Coldplay Jakarta",
            "description": "Coldplay akan konser di jakarta untuk merayakan greenpace   dunia",
            "location": "Jakarta - Gelora Bung Karno",
            "image": null,
            "price": 300000,
            "quantity": 5000,
            "available": true,
            "start_date": "2023-11-05T00:00:00Z",
            "end_date": "2023-11-08T00:00:00Z",
            "createdAt": "2023-12-10T22:59:13.27374+07:00",
            "updatedAt": "2023-12-10T22:59:13.361098+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        },
        {
            "id": 3,
            "name": "Country Roads Jamboree",
            "description": "Get ready for a country music adventure on the open road",
            "location": "Country Trails Fairgrounds - Nashville",
            "image": null,
            "price": 240000,
            "quantity": 6900,
            "available": true,
            "start_date": "2023-02-14T00:00:00Z",
            "end_date": "2023-02-17T00:00:00Z",
            "createdAt": "2023-12-09T22:16:08.099738+07:00",
            "updatedAt": "2023-12-10T23:18:43.899515+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        },
        {
            "id": 6,
            "name": "EDM Eclipse Carnival",
            "description": "Dance through the night with pulsating EDM beats",
            "location": "Eclipse Arena - Tokyo",
            "image": null,
            "price": 300000,
            "quantity": 7800,
            "available": true,
            "start_date": "2023-06-18T00:00:00Z",
            "end_date": "2023-06-21T00:00:00Z",
            "createdAt": "2023-12-10T23:23:39.419662+07:00",
            "updatedAt": "2023-12-10T23:23:39.421199+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        }
    ]
}
```
### Filter By Date Event
- Endpoint :
    - /event/date/:start_date/:end_date
- Method :
    - GET
- BODY :
- Param :
```param
{
    "start_date" : 2023-02-10
    "end_date" : 2023-02-17
}
```
- RESPONSE :
```json
{
    "data": [
        {
            "id": 3,
            "name": "Country Roads Jamboree",
            "description": "Get ready for a country music adventure on the open road",
            "location": "Country Trails Fairgrounds - Nashville",
            "image": null,
            "price": 240000,
            "quantity": 6900,
            "available": true,
            "start_date": "2023-02-14T00:00:00Z",
            "end_date": "2023-02-17T00:00:00Z",
            "createdAt": "2023-12-09T22:16:08.099738+07:00",
            "updatedAt": "2023-12-10T23:18:43.899515+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        }
    ]
}
```

### Sort By Cheapest Event
- Endpoint :
    - /event/cheapest
- Method :
    - GET
- BODY :
- Param :
```param
{
    "sort" : termurah
}
```
- RESPONSE :
```json
{
    "data": [
        {
            "id": 11,
            "name": "Electro Fusion Festival",
            "description": "Join the electrifying music celebration with top DJs",
            "location": "Tech Park Amphitheatre - San Francisco",
            "image": null,
            "price": 180000,
            "quantity": 7000,
            "available": true,
            "start_date": "2023-10-12T00:00:00Z",
            "end_date": "2023-10-15T00:00:00Z",
            "createdAt": "2023-12-10T23:24:17.5726+07:00",
            "updatedAt": "2023-12-10T23:24:17.583645+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        },
        {
            "id": 7,
            "name": "Acoustic Serenity Evening",
            "description": "Unwind with soothing acoustic melodies under the stars",
            "location": "Serenity Gardens - Sydney",
            "image": null,
            "price": 190000,
            "quantity": 5800,
            "available": true,
            "start_date": "2023-04-05T00:00:00Z",
            "end_date": "2023-04-08T00:00:00Z",
            "createdAt": "2023-12-10T23:23:47.425543+07:00",
            "updatedAt": "2023-12-10T23:23:47.427416+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        },
        {
            "id": 9,
            "name": "Pop Sensation Extravaganza",
            "description": "An explosive night of pop sensations and dance beats",
            "location": "Pop Dome - Miami",
            "image": null,
            "price": 220000,
            "quantity": 6500,
            "available": true,
            "start_date": "2023-08-28T00:00:00Z",
            "end_date": "2023-08-31T00:00:00Z",
            "createdAt": "2023-12-10T23:24:02.463255+07:00",
            "updatedAt": "2023-12-10T23:24:02.477812+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        },
        {
            "id": 3,
            "name": "Country Roads Jamboree",
            "description": "Get ready for a country music adventure on the open road",
            "location": "Country Trails Fairgrounds - Nashville",
            "image": null,
            "price": 240000,
            "quantity": 6900,
            "available": true,
            "start_date": "2023-02-14T00:00:00Z",
            "end_date": "2023-02-17T00:00:00Z",
            "createdAt": "2023-12-09T22:16:08.099738+07:00",
            "updatedAt": "2023-12-10T23:18:43.899515+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        },
    ]
}
```
### Sort By Expensive Event
- Endpoint :
    - /event/expensive
- Method :
    - GET
- BODY :
- Param :
```param
{
    "sort" : termahal
}
```
- RESPONSE :
```json
{
    "data": [
        {
            "id": 6,
            "name": "EDM Eclipse Carnival",
            "description": "Dance through the night with pulsating EDM beats",
            "location": "Eclipse Arena - Tokyo",
            "image": null,
            "price": 300000,
            "quantity": 7800,
            "available": true,
            "start_date": "2023-06-18T00:00:00Z",
            "end_date": "2023-06-21T00:00:00Z",
            "createdAt": "2023-12-10T23:23:39.419662+07:00",
            "updatedAt": "2023-12-10T23:23:39.421199+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        },
        {
            "id": 5,
            "name": "Coldplay Jakarta",
            "description": "Coldplay akan konser di jakarta untuk merayakan greenpace dunia",
            "location": "Jakarta - Gelora Bung Karno",
            "image": null,
            "price": 300000,
            "quantity": 5000,
            "available": true,
            "start_date": "2023-11-05T00:00:00Z",
            "end_date": "2023-11-08T00:00:00Z",
            "createdAt": "2023-12-10T22:59:13.27374+07:00",
            "updatedAt": "2023-12-10T22:59:13.361098+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        },
        {
            "id": 10,
            "name": "Rock Revival Tour",
            "description": "Experience the revival of classic rock hits",
            "location": "Rock Arena - Los Angeles",
            "image": null,
            "price": 280000,
            "quantity": 7500,
            "available": true,
            "start_date": "2023-11-20T00:00:00Z",
            "end_date": "2023-11-23T00:00:00Z",
            "createdAt": "2023-12-10T23:24:09.995672+07:00",
            "updatedAt": "2023-12-10T23:24:09.998791+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        },
        {
            "id": 8,
            "name": "World Rhythms Fiesta",
            "description": "Celebrate global music diversity with rhythmic beats",
            "location": "Cultural Plaza - London",
            "image": null,
            "price": 260000,
            "quantity": 7200,
            "available": true,
            "start_date": "2023-07-10T00:00:00Z",
            "end_date": "2023-07-13T00:00:00Z",
            "createdAt": "2023-12-10T23:23:56.030234+07:00",
            "updatedAt": "2023-12-10T23:23:56.031752+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        }
    ]
}
```

### Sort By Newest Event
- Endpoint :
    - /event/newest
- Method :
    - GET
- BODY :
- Param :
```param
{
    "sort" : terbaru
}
```
- RESPONSE :
```json
{
    "data": [
        {
            "id": 11,
            "name": "Electro Fusion Festival",
            "description": "Join the electrifying music celebration with top DJs",
            "location": "Tech Park Amphitheatre - San Francisco",
            "image": null,
            "price": 180000,
            "quantity": 7000,
            "available": true,
            "start_date": "2023-10-12T00:00:00Z",
            "end_date": "2023-10-15T00:00:00Z",
            "createdAt": "2023-12-10T23:24:17.5726+07:00",
            "updatedAt": "2023-12-10T23:24:17.583645+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        },
        {
            "id": 10,
            "name": "Rock Revival Tour",
            "description": "Experience the revival of classic rock hits",
            "location": "Rock Arena - Los Angeles",
            "image": null,
            "price": 280000,
            "quantity": 7500,
            "available": true,
            "start_date": "2023-11-20T00:00:00Z",
            "end_date": "2023-11-23T00:00:00Z",
            "createdAt": "2023-12-10T23:24:09.995672+07:00",
            "updatedAt": "2023-12-10T23:24:09.998791+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        },
        {
            "id": 9,
            "name": "Pop Sensation Extravaganza",
            "description": "An explosive night of pop sensations and dance beats",
            "location": "Pop Dome - Miami",
            "image": null,
            "price": 220000,
            "quantity": 6500,
            "available": true,
            "start_date": "2023-08-28T00:00:00Z",
            "end_date": "2023-08-31T00:00:00Z",
            "createdAt": "2023-12-10T23:24:02.463255+07:00",
            "updatedAt": "2023-12-10T23:24:02.477812+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        },
        {
            "id": 8,
            "name": "World Rhythms Fiesta",
            "description": "Celebrate global music diversity with rhythmic beats",
            "location": "Cultural Plaza - London",
            "image": null,
            "price": 260000,
            "quantity": 7200,
            "available": true,
            "start_date": "2023-07-10T00:00:00Z",
            "end_date": "2023-07-13T00:00:00Z",
            "createdAt": "2023-12-10T23:23:56.030234+07:00",
            "updatedAt": "2023-12-10T23:23:56.031752+07:00",
            "deletedAt": "0001-01-01T07:00:00+07:00"
        }
    ]
}
```
## Ticket 
### Create Ticket
- Endpoint :
    - /tickets
- Method :
    - POST
- BODY :
- Param
```param
{
    user_id : (from claims JWT)
}
```
- Json
```json
{
	"event_id": 11,
	"quantity": 1
}

```
- RESPONSE :
```json
{
    "message": "ticket created successfully"
}
```
### Get All Ticket
- Endpoint :
    - /tickets
- Method :
    - GET
```json
{
	"event_id": 11,
	"quantity": 1
}

```
- RESPONSE :
```json
{
    "data": [
        {
            "ticket": {
                "id": 11,
                "name": "Electro Fusion Festival",
                "description": "Join the electrifying music celebration with top DJs",
                "location": "Tech Park Amphitheatre - San Francisco",
                "image": null,
                "price": 180000,
                "quantity": 6998,
                "available": true,
                "start_date": "2023-10-12T00:00:00Z",
                "end_date": "2023-10-15T00:00:00Z",
                "createdAt": "2023-12-10T23:24:17.5726+07:00",
                "updatedAt": "2023-12-11T00:08:23.785044+07:00",
                "deletedAt": "0001-01-01T07:00:00+07:00"
            },
            "user_id": 1
        },
        {
            "ticket": {
                "id": 11,
                "name": "Electro Fusion Festival",
                "description": "Join the electrifying music celebration with top DJs",
                "location": "Tech Park Amphitheatre - San Francisco",
                "image": null,
                "price": 180000,
                "quantity": 6998,
                "available": true,
                "start_date": "2023-10-12T00:00:00Z",
                "end_date": "2023-10-15T00:00:00Z",
                "createdAt": "2023-12-10T23:24:17.5726+07:00",
                "updatedAt": "2023-12-11T00:08:23.785044+07:00",
                "deletedAt": "0001-01-01T07:00:00+07:00"
            },
            "user_id": 1
        }
    ]
}
```

### Get Ticket By User Id
- Endpoint :
    - /tickets
- Method :
    - GET
- Param :
```json
{
	"user_id" : 1
}

```
- RESPONSE :
```json
{
    "data": [
        {
            "ticket": {
                "id": 11,
                "name": "Electro Fusion Festival",
                "description": "Join the electrifying music celebration with top DJs",
                "location": "Tech Park Amphitheatre - San Francisco",
                "image": null,
                "price": 180000,
                "quantity": 6998,
                "available": true,
                "start_date": "2023-10-12T00:00:00Z",
                "end_date": "2023-10-15T00:00:00Z",
                "createdAt": "2023-12-10T23:24:17.5726+07:00",
                "updatedAt": "2023-12-11T00:08:23.785044+07:00",
                "deletedAt": "0001-01-01T07:00:00+07:00"
            },
            "user_id": 1
        },
        {
            "ticket": {
                "id": 11,
                "name": "Electro Fusion Festival",
                "description": "Join the electrifying music celebration with top DJs",
                "location": "Tech Park Amphitheatre - San Francisco",
                "image": null,
                "price": 180000,
                "quantity": 6998,
                "available": true,
                "start_date": "2023-10-12T00:00:00Z",
                "end_date": "2023-10-15T00:00:00Z",
                "createdAt": "2023-12-10T23:24:17.5726+07:00",
                "updatedAt": "2023-12-11T00:08:23.785044+07:00",
                "deletedAt": "0001-01-01T07:00:00+07:00"
            },
            "user_id": 1
        }
    ]
}
```
## Transaction Topup
### Create Transaction
- Endpoint :
    - /transactions
- Method :
    - POST
- BODY :
- Param
```param
{
    user_id : (from claims JWT)
}
```
- Json
```json
{
	"amount": 50000
}

```
- RESPONSE :
```json
{
    "url_pembayaran": "https://app.sandbox.midtrans.com/snap/v3/redirection/73f23ec6-7878-490c-9964-b9728f25c049"
}
```

### Get History Transaction
- Endpoint :
    - /transactions/history
- Method :
    - GET
- BODY :
- Param
```param
{
    user_id : (from claims JWT)
}
```
- RESPONSE :
```json
{
    "data": [
        {
            "amount": 50000,
            "created_at": "2023-12-09T22:28:01.894923+07:00",
            "id": 1,
            "order_id": "topup",
            "status": "unpaid",
            "updated_at": "2023-12-09T22:28:01.894923+07:00",
            "user_id": 1
        },
        {
            "amount": 50000,
            "created_at": "2023-12-09T22:28:17.52897+07:00",
            "id": 2,
            "order_id": "topup123",
            "status": "unpaid",
            "updated_at": "2023-12-09T22:28:17.52897+07:00",
            "user_id": 1
        },
        {
            "amount": 50000,
            "created_at": "2023-12-10T11:34:13.560953+07:00",
            "id": 3,
            "order_id": "topup-9de80ac7-eacc-4aa3-bffd-43049b5faba6",
            "status": "paid",
            "updated_at": "2023-12-10T11:34:40.981557+07:00",
            "user_id": 1
        }
    ]
}
```

### Input Saldo Without Midtrans
- Endpoint :
    - /transactions/input-saldo
- Method :
    - POST
- BODY :
- Param
```param
{
    user_id : (from claims JWT)
}
```
- Json
```json
{
    "saldo":5000000
}

```
- RESPONSE :
```json
{
    "data": {
        "id": 1,
        "username": "Admin1",
        "email": "Admin1@gmail.com",
        "number": "011111111111",
        "role": "Administrator",
        "saldo": 9700000,
        "profile": null,
        "createdAt": "2023-12-09T22:15:21.743238+07:00",
        "updatedAt": "2023-12-10T15:05:17.315768+07:00",
        "deletedAt": "0001-01-01T07:00:00+07:00"
    },
    "message": "success input saldo"
}
```

## Notification
### Create Notification
- Endpoint :
    - /notifications
- Method :
    - POST
- BODY :
```json
{
    "type":"warning",
    "content":"something in content",
    "is_read":false
}

```
- RESPONSE :
```json
{
    "message": "success create notification"
}
```

### Get Notification
- Endpoint :
    - /notifications
- Method :
    - GET
- RESPONSE :
```json
{
    "data": [
        {
            "id": 1,
            "type": "warning",
            "content": "something in content",
            "is_read": true,
            "created_at": "2023-12-10T15:11:20.369244+07:00",
            "updated_at": "2023-12-10T15:11:39.104122+07:00",
            "deleted_at": "0001-01-01T07:00:00+07:00"
        }
    ]
}
```
### User Get Notification
- Endpoint :
    - /users/notifications
- Method :
    - GET
- RESPONSE :
```json
{
    "data": [
        {
            "id": 1,
            "type": "warning",
            "content": "something in content",
            "is_read": false,
            "created_at": "2023-12-10T15:11:20.369244+07:00",
            "updated_at": "2023-12-10T15:11:39.104122+07:00",
            "deleted_at": "0001-01-01T07:00:00+07:00"
        }
    ]
}
```