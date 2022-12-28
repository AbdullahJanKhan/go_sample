# Rest Server

This layer is the first responders to an incomming request.

We will contain our middleware and routes in this folder.

## Middleware

Middleware is a small piece of code that performs initial checks on an incomming request, e.g. the most common is the JWT verification i.e. authentication (verifing the user sending in the request is of our application or is authentic), another is authorization; if we have different types of users we may restrict some users types to some specific routes only.

Middlewares are also used for multiple other usecases that one may have.

We will use our middleware here to handle the inital request and set it to the incoming context.

`PS: Here we will be working with Go Lang GIN framework, feel free to change it.`
