# Building a REST API in Go

In *main.go* we're creating a REST API that will listen on localhost:8888.
The application will try to read the environment variable *RETURN_VALUE* and return this when doing a *GET* request to *http://localhost:8888*.
If the value is not set it will return *Hello from Go*.

# Exercise
Create a dockerfile to build and run the go backend.
To get inputs on the file, you can check [dockerhub](https://hub.docker.com/_/golang).

Note: Name the dockerfile *Dockerfile* (without extension) and put it next to *main.go*

In order to build, execute `docker build -t dockergetstarted/backend .` from within the *golang* folder.

After you built it successfully, you can start the container:

`docker run -it --name myfirstdockerbackend dockergetstarted/backend`

However this will not yet work, as you must make sure that you expose the ports to the host machine - how can you do that?

After you've managed, you can browse to *https://localhost:8888* and see whether it works.

Bonus Point #1: Can you make it listen on a different port without recompiling?
Bonus Point #2: Can you change what the REST API returns without recompiling?
