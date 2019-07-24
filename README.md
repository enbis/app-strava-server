# app-strava-server

### Authorization Error

Auth process with new list of scopes, since October 2018.

- Step 1 redirect user to Strava's auth page, and allow permissions
  
```
https://www.strava.com/oauth/authorize?
    client_id=YOUR_CLIENT_ID&
    redirect_uri=YOUR_CALLBACK_DOMAIN&
    response_type=code&
    scope=YOUR_SCOPE
```

- Step 2 read code from response, and `POST` the following request to obtain new `access_token` with expiration date

```
https://www.strava.com/oauth/token?
    client_id=YOUR_CLIENT_ID&
    client_secret=YOUR_CLIENT_SECRET&
    code=AUTHORIZATION_CODE_FROM_STRAVA&
    grant_type=authorization_code
```

- Step 3 with that token you will be able to continue launch the requests

### Server

Launch server `go run main.go`

### Test
Run test `go test -run [NameOfTest]`