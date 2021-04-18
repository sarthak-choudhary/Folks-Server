# users_service
An authentication layer

This is a jwt token based authentication system. It has end-points for signup, login and myProfile along with an endpoint for google oauth.

The user model in the service is as shown below:
``` 
    _id : ObjectID
    firstname : string
    lastname : string
    password : string
    phoneNo : string
    interests : []string
    isComplete : boolean
```
To make a user using signup view.

endpoint : /sign_up     Method: POST

The body of the request should contain the json representation of the user including firstname, lastname, password, phoneNo and interests.
Example :

```
{
    "firstname" : "test",
    "lastname" : "user",
    "email" : "testuser@xyz.com",
    "password" : "password@123",
    "phoneNo" : "93152XXXXX",
    "ineterests" : ["sleeping", "sleeping"]
}
```

In case of success, the response will have id and the jwt-token along with request data for the user with status code of 201.
Example :

```
{
    "_id" : "7823teujywgdj9jhfsdhf",
    "firstname" : "test",
    "lastname" : "user",
    "email" : "testuser@xyz.com",
    "phoneNo" : "93152XXXXX",
    "interests" : ["sleeping, "sleeping"],
    "token" : "sdihgfiq3yr98y340hrfohew93q48u90cn5ur3498urojhetcpo8uy49o8p3tu9p3o8qcu409[u.lkqjhweifuy34cni4uu5"
}
```

In case of faulty request data, there can be an error with an appropriate error message.
Example :
```
{
    "error" : "The error message"
}
```

To login using email and password for an existing user:
endpoint : /login     Method : "POST"

The body of request should contain the valid email id and password.
Example:
```
{
    "email" : "testuser@xyz.com",
    "password" : "password@123"
}
```

The response will be similar to the response from the signup view.

To make a request as an authorized user, the header of the request should contain the jwt-token of that user received from signup or login response.
The Authorization key of the header should conatin the token as JWT followed by a blankspace then the token.

For example, if "hdoiweuyihrncwnriuoy233984ry2hkwehgduyt87eqeew" is the token for the user, then the key-value pair in the request header will be as following:

```
    Authorization : JWT hdoiweuyihrncwnriuoy233984ry2hkwehgduyt87eqeew
```
Apart from signup and login view, there is also a view for google login using id token from google oauth playground.
For retrieving the id token from google oauth playground follow these steps:
1. Go to (Google Oauth Playground)[https://developers.google.com/oauthplayground/].
2. In the section for selecting and authorizing APIs, enter these two scopes in the Authorize APIS field as it is ```https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile``` and click on authorize apis.
3. Login with a valid google account.
4. Click on exchange authorization for tokens, you will get a bunch of JSON data on the right side, copy the id_token from there.
5. You have to use this id token for google oauth endpoint.

endpoint:/google_login                  Method: POST

The body should contain the id token from google oauth playground.

Example: 
```
{
    "id_token" : "this sholud be replaced by a valid id token received from google oauth playground before it expires"
}
```
Response of this will same as signup and login view, if the user is new to the database it will create a new user otherwise will fetch the existing user.

For viewing the profile of the logged user, there is a My Profile endpoint.

enpoint: /my_profile        Method: POST

The request header should conatain the JWT token for the user where as body of the request should be empty.
Response will contain the id, name, email, phone number, interests of the current user.

Example: 
```
{
    "_id": "5f3b68e7c2a960131d905edd",
    "firstName": "test",
    "lastName": "user",
    "email": "testuser@xyz.com",
    "phoneNo": "9999999999",
    "interests": [
        "sleeping",
        "sleeping"
    ]
}
```
