# points-web-service

Go Lang points web service

# Programming concepts

-   RESTFUL application
    -   GET, POST, PUT, DELETE implemented in Go Lang
    -   JavaScript Fetch()
        -   PUT handled with text response
        -   DELETE handled with JSON response
-   Testing
    -   Test Suite - (testStatementsTransactions)
    -   Heartbeat - (TestScenario) Ensures confidence that the programming is functioning by concurrently using a ticker to get feedback
    -   Test Helper - (tcSetup) used to setup similar requirements
    -   Golden File
    -   Configuration options
-   Go Lang concepts/patterns
    -   Functional Options (NewApps)
    -   Html Templates
	
_There is an assumption Go Language is installed properly_  
_follow instructions at https://golang.org/ to install go_  
_once Go has been installed, then follow the install instructions_

# Install

to download source code and create binary (type in following command in console)

```
go get github.com/pwarmuz/points-exercise
```

> it is important to understand where the source code is located  
> this depends on where your gopath is configured, The general path would be as follows  
> for source code: {root of wherever gopath configuration}/src/github.com/pwarmuz/points-exercise  
> for binary: {root of wherever gopath configuration}/bin/points-exercise(<-Executable>)  
> There is an option to recompile it from source under the run heading below if there are any running issues

# Docker / Kubernetes
- There is no downloadable repo for this. The intent is to simulate a private project that is built and deployed.
- If running minikube then ensure ```eval $(minikube docker-env)``` is set to target the kubernetes environment
- This dockerfile requires buildkit 1, use ```export DOCKER_BUILDKIT=1```
- To run the unit test use ```docker build -t pwarmuz/points . --target unit-test```
- To run the lint use ```docker build -t pwarmuz/points . --target lint```
- To run locally use ```docker build -t pwarmuz/points . --target localhost```
- To run kubernetes deployment use ```kubectl apply -f points-web.yml```
- If using kubenetes, use port-forward to locally test application with ```kubectl port-forward deployment/points 8080:8080```
- assuming 8080 were the port targeted

- # removing
- To remove the docker images use ```docker rmi pwarmuz/points```
- If you want to remove dangling images use ```docker image prune --filter "dangling=true"```
- To remove kubernetes deployment use ```kubectl delete -f points-web.yml```

# Run

## Notes

-   source code uses Port 8080 for http serving
-   to change this, change `const PORT string = ":8080"` to `const PORT string = ":XXXX"` within `main.go`, where XXXX is the port number, making sure to leave it as a string with `:` prepended

## **Quick run**

_In console as a command_  
:exclamation: _make sure to change the current directory to the source code folder before executing to fulfil the html dependencies, as shown in the command list_

```
cd {root of wherever gopath configuration}/src/github.com/pwarmuz/points-exercise
points-exercise
```

:exclamation: _if you get any html dependencies issues, and you've selected the source code folder in the console, then move the executable to the points-exercise source code folder_

## _Re-compile and run (if all else fails due to configuration or your source code changes)_

_In console as a command, with source code location as directory_

```
cd {root of wherever gopath configuration}/src/github.com/pwarmuz/points-exercise
go build && go run .
```

# Interacting with server

-   use any browser
-   type `http://localhost:8080/` where 8080 is the port number, if you decided to change this then adjust accordingly
-   to play out the scenario in the exercise, type in 5000 within the "Deduct field" and submit

## Valid URLS

-   `http://localhost:8080/`
-   `http://localhost:8080/balance`

## To add points

use the Add field, insert required fields and submit

## To deduct points

use the Deduct field, insert points to deduct and submit

## To add user/ change current user

use the Add User/ Change User field, insert a username and submit

## To delete a user

use the Delete User field, insert username and submit

# Testing

_In console as a command, with source code location as directory_

```
cd {root of wherever gopath configuration}/src/github.com/pwarmuz/points-exercise
go test -cover
```

[response] PASS coverage: 48.5% of statements
