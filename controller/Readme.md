# Controller

In a rest server we have some routes that are accpeted and returned a response.

We have divided our server in such a way in form of folder structure, that we will define all the controllers, i.e. the code base that will just redirect the request to `service layer`, and do final logging of the request received.

When the result is returned from the `service layer` we then construst the a warapper around the result and respond back in a structured from.
