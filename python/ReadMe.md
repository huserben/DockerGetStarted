# Building a Frontend in Python

In *webapp.py* we're creating a python-flask application that will server some fancy webpage on port 5000 (*http://localhost:5000*).
The application will try to read the environment variable *REST_API* and if it's set, make a get request to that url and display the results.
If the value is not set it will return a static message.

# Exercise
Create a dockerfile to build and run the go backend.
To get inputs on the file, you can check [dockerhub](https://hub.docker.com/_/python).

Note: Name the dockerfile *Dockerfile* (without extension) and put it next to *webapp.py*

In order to build, execute `docker build -t dockergetstarted/frontend .` from within the *python* folder.

After you built it successfully, you can start the container:

`docker run --it --name myfirstdockerfrontend dockergetstarted/frontend`

However this will not yet work, as you must make sure that you expose the ports to the host machine - how can you do that?

After you've managed, you can browse to *https://localhost:5000* and see whether it works.

Bonus Point #1: Can you make it listen on a different port without recompiling?
Bonus Point #2: Can you specify a rest backend where it should fetch data from? (E.g: *https://jsonplaceholder.typicode.com/todos/1*)