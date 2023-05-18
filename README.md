 # Back4app Containers - Go Application

 This is a simple Go application that has been configured to run on Back4app Containers. It serves a welcome page when accessed via a web browser.

 ## Getting Started

 These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

 ### Prerequisites

 - You should have [Go](https://golang.org/dl/) installed on your local machine. 
 - Install [Docker](https://www.docker.com/products/docker-desktop) if you want to build and run the Docker container locally.

 ### Installing

 1. Clone this repository:

 ```sh
 git clone https://github.com/templates-back4app/containers-go-app.git
 cd containers-go-app
 ```

 2. Start the development server:

 ```sh
 go run main.go
 ```

 Open [http://localhost:8080](http://localhost:8080) to view it in the browser.

 ### Building a Docker Image

 To build a Docker image of the application, run the following command:

 ```sh
 docker build -t containers-go-app .
 ```

 ### Running the Docker Container

 To run the Docker container, use the following command:

 ```sh
 docker run -p 8080:8080 containers-go-app
 ```

 Now, the app is running at [http://localhost:8080](http://localhost:8080)

 ## Deployment

 The project can be deployed on Back4app Containers. Refer to the [Back4app Documentation](https://www.back4app.com/docs/platform/containers) for more information.

 ## Built With

 - [Go](https://golang.org/)
 - [Docker](https://www.docker.com/)

 ## Contributing

 Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests.

 ## Versioning

 We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/templates-back4app/containers-go-app/tags).

 ## Authors

 - **Back4app** - Initial work - [Back4app](https://github.com/back4app)

 See also the list of [contributors](https://github.com/templates-back4app/containers-go-app/contributors) who participated in this project.

 ## License

 This project is licensed under the MIT License
