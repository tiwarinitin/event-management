# Event-Management application
This is REST API written in GO for event management. Containing several routes including registering,signup,login, etc actions.

We have used Sqlite db for local db usage. Dockerize this application will be updated in sometime.

## How to run application
> go run .

### API Routes

You can reference *api-test* folder that contains all the api urls and body requires for testing this application.
 
Some routes are mentioned below:

```
GET    /events
GET    /events/id
POST   /events
PUT    /events/id
DELETE /events
DELETE /events/id
POST   /events/signup
POST   /events/login
POST   /events/id/register
DELETE /events/id/register
```




