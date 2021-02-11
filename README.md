# Docker Get Started

Check out the exercises in the *ReadMe* for both the *python* and *golang* folders.
There you'll learn how to create a docker file and use it to build and run an application and make some configuration (like which ports to use and specifying environment variables).

## Putting it all together
The second part of the exercise is to put it all together, meaning creating a docker compose file that will:

1. Include the golang backend
   1. Configure it so that you have a custom value returned
2. Include python frontend
   1. Configure it so that you point to the golang backend and show that value
   2. Make sure you can access on your browser by specifying a custom port

Running `docker-compose build` should build both frontend and backend.
Running `docker-compose up` should start both frontend and backend and allow you to browse to the frontend and display the value fetched from the backend.