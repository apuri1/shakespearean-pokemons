# Shakespearean Pokemons

## Commands

To build the app outside of a docker container:

```
$ make
```

To run the app outside of a dockeer container:

```
$ cd bin
$ ./app
```

To build docker image containing the app:

```
$ sudo docker build -t shakes .
```

To run docker container containing the app:

```
$ sudo docker run -p 5000:5000 shakes
```

Use curl, postman or http://localhost:5000 in a web browser to send requests to excercise the app

## TODOs:

- Implement front-facing nginx reverse proxy to terminate TLS requests and route those requests to the app (can also run in docker)
- Implement more metrics, measuring latency etc
- Implement more unit tests
- Review the pokeapi documentation to refine how and what pokeapi calls to make
- Implement a caching mechanism to reduce pokeapi and shakespeare translation requests- would also get around issue of request throttling.
- Implement a backend, possible mysql db, to store data.
- Make code more modular, moving any repeating code (HTTP client requests) into separate functions or packages
- Update Makefile to also build the docker image
- Implement a Jenkinsfile or .gitlab-ci.yml - can be used to trigger a build pipeline for the repo upon a git commit, and push the image to a docker repository or ECR
- If available, add sonarcube config to auto-inspect code quality etc.
- Possibly add a config.yml file for greater flexiblity e.g. port numbers, target endpoints.
- Implement a kubernetes deployment spec, to deploy the app in a kubernetes cluster. If config.yml added, add a configmap spec.
- If not deploying via Helm, implement terraform instead of yml specs.
- Implement scripts or use something like Locust load testing, to test the performance of the deployment - does it scale?
