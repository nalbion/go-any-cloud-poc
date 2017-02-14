This is a "Hello World" app, written in Go to be deployed on either AWS Lambda or Google App Engine.
 
The main application code can be found under `lib/`, and the entry points are `lambda/index.go` and `app-engine/app.go`.
The entry points wrap the Lambda event or `http.ResponseWriter` and `http.Request` in an `InboundRequest` instance and route to the appropriate controller.
Throughout the application code `InboundRequest` is used to read the request params, headers and body, and also to create outbound requests to other web services etc.
 
# AWS Lambda

Running the Go app on Lambda is made possible by [aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim)
and the Makefile and Dockerfile are based on the examples in that repo.
 
To build on your own machine, simply run `make`.  The `dev` target will run `make test lambda/$(PACKAGE).zip` within a Docker container.

To build on a CI server, use the Docker image `nalbion/go-lambda-build` and run `make test lambda/$(PACKAGE).zip`.


# Google App Engine

There is an App-Engine specific `Makefile` in the `app-engine` directory, to build for App Engine move it into the project root directory.
