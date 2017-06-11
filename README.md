# TEWW
Golang microservice architecture sample. The sample includes 4 services with trivial names: client, auth, backend and db.

Client - web entrypoint for user.
<br/>
Auth - JWT authentication implementation.
<br/>
Backend - server side worker for processing users requests.
<br/>
DB - storage for user data.

*Ports*
* *client :3000*
* *auth :3001*
* *backend :3002*
* *db (redis) :6379*

*TODO*

* *connect auth to db*
* *add loader*
* *add autorization functionality*
* *add logs*
* *add tests*
