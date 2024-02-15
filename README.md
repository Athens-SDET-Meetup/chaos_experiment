# Improve Application Resiliency With Chaos Testing

- [Improve Application Resiliency With Chaos Testing](#improve-application-resiliency-with-chaos-testing)
  - [(1) Tutorial Intro](#1-tutorial-intro)
    - [(1.1) Introduction](#11-introduction)
      - [(1.1.1) What is chaos testing?](#111-what-is-chaos-testing)
      - [(1.1.2) Why implement chaos testing?](#112-why-implement-chaos-testing)
      - [(1.1.3) Chaos xk6-disruptor](#113-chaos-xk6-disruptor)
        - [(1.1.4) Pod disruptor](#114-pod-disruptor)
        - [(1.1.5) Service disruptor](#115-service-disruptor)
      - [(1.1.6) Chaos testing example :smiling\_imp:](#116-chaos-testing-example-smiling_imp)
    - [(1.2) Prerequisites](#12-prerequisites)
    - [(1.3) Setting up a Kubernetes cluster :nerd\_face:](#13-setting-up-a-kubernetes-cluster-nerd_face)
  - [(2) Deployments :rocket:](#2-deployments-rocket)
    - [(2.1) Deploy provider API service](#21-deploy-provider-api-service)
      - [(2.1.1) Going through the files inside the service directory](#211-going-through-the-files-inside-the-service-directory)
      - [(2.1.2) Building the service image and push it to Docker Hub :building\_construction:](#212-building-the-service-image-and-push-it-to-docker-hub-building_construction)
      - [(2.1.3) Deploying the service to the Kubernetes cluster :rocket:](#213-deploying-the-service-to-the-kubernetes-cluster-rocket)
    - [(2.2) Deploy consumer API service :rocket:](#22-deploy-consumer-api-service-rocket)
      - [(2.2.1) Going through the files inside the service directory](#221-going-through-the-files-inside-the-service-directory)
      - [(2.2.2) Building the service image and push it to Docker Hub](#222-building-the-service-image-and-push-it-to-docker-hub)
      - [(2.2.3) Deploying the service to the Kubernetes cluster](#223-deploying-the-service-to-the-kubernetes-cluster)
  - [(3) Chaos Testing](#3-chaos-testing)
    - [(3.1) Installing k6 and xk6-disruptor](#31-installing-k6-and-xk6-disruptor)
    - [(3.2) Implementing chaos testing scenarios](#32-implementing-chaos-testing-scenarios)
    - [(3.3) Root cause analysis](#33-root-cause-analysis)
    - [(3.4) Solution](#34-solution)
  - [(4) Conclusion](#4-conclusion)


## (1) Tutorial Intro

### (1.1) Introduction

Testing is essential to ensure that your software application performs as expected. No doubt about it. But even if your application runs smoothly in a testing environment, unexpected incidents can occur in production, such as a crashed database or a lost connection to a third-party service. To ensure your application can handle such incidents and minimize their impact, you need to incorporate chaos testing into your delivery process. In this article, we’ll explore how to implement chaos testing using k6 for two backend services, to help you ensure the resilience of your application.

#### (1.1.1) What is chaos testing?
Chaos testing is a software testing methodology that involves intentionally injecting faults and disruptions into your application to identify its weaknesses. It is often used with other testing methodologies, such as functional, load, and performance, to ensure your application’s overall quality and reliability.

#### (1.1.2) Why implement chaos testing?
Chaos testing ensures that your application can withstand unexpected failures and disruptions in a production environment. By simulating different scenarios and failure modes, you identify potential issues in your app and develop solutions to resolve them.

#### (1.1.3) Chaos xk6-disruptor

Currently, the xk6-disruptor only supports chaos testing for an application running in Kubernetes. The xk6-disruptor runs in your local machine and has two main components:

The xk6-disruptor extension works as a Javascript API to inject faults into your app through Kubernetes components, whether through Kubernetes pods (using PodDisruptor) or Kubernetes services (using ServiceDisruptor).
The xk6-disruptor-agent is installed into the Kubernetes by the xk6-disruptor. Then the xk6-disruptor-agent will inject faults into the Kubernetes component (Kubernetes pod or service) according to instructions from the xk6-disruptor extension.
So, put it simply, the extension exposes the Javascript API, and the agent will inject the faults.

##### (1.1.4) Pod disruptor

The xk6-disruptor extension uses Pod disruptor to inject faults to the app inside the Kubernetes cluster in three steps:

The pod disruptor chooses the application pod that it wants to inject faults into
The pod disruptor attaches the disruptor agent to the application pod
The disruptor agent injects faults into the application container
The below diagram illustrates how the pod disruptor works:

![disruptor](./media/disruptor.png)

##### (1.1.5) Service disruptor

The xk6-disruptor extension uses the service disruptor as a wrapper around the pod disruptor. Basically, the service disruptor works in three steps:

The service disruptor chooses the Kubernetes service pod that it wants to inject faults into
The service disruptor creates a pod disruptor to interact with the Kubernetes service pod through the pod disruptor
The pod disruptor attaches the disruptor agent to the Kubernetes service pod, then the disruptor agent injects faults into the application container.

#### (1.1.6) Chaos testing example :smiling_imp:

To better understand how to apply chaos testing for a software application, let’s build two simple backend services, as detailed below:

**Service 1**: a Go service that creates an API to display the current timestamp. Let’s call this API a get time API.
**Service 2**: a Go service that implements an API consuming the get time API. This API will respond “Successfully consuming message” if the response status code of the get timestamp API is 200. Let’s call this API a consuming API

You will build these services using the following three steps:

- Clone the GitHub repository
- Build the Docker image for each service using Dockerfile
- Deploy each service to the Kubernetes cluster

### (1.2) Prerequisites

To follow along with the steps in the article, you need to fulfill the following prerequisites:

* [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) installed
* [Docker engine](https://docs.docker.com/engine/install/ubuntu/) in your machine to build the service images.
* The [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/) command line tool to interact with the Kubernetes cluster.
* Have an account on Docker hub to store the service images.

### (1.3) Setting up a Kubernetes cluster :nerd_face:
You can set up a Kubernetes cluster by using cloud provider-supported tools like [Vultr Kubernetes](https://www.vultr.com/kubernetes/), [AWS Elastic Kubernetes Engine](https://aws.amazon.com/eks/), or [Google Kubernetes Engine](https://cloud.google.com/kubernetes-engine). Or you can create your own Kubernetes cluster setup using tools like [Minikub](https://minikube.sigs.k8s.io/docs/) or [kind](https://kind.sigs.k8s.io/). To make things simpler, let’s use Minikube to set up the Kubernetes cluster.

Open a new terminal in your local machine and run the following commands to install Minikube:

```
// First, download the minikube client tool to the current directory
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64

// Make the minikube client to be executable
chmod +x minikube

// Create a new directory in /usr/local/bin 
sudo mkdir -p /usr/local/bin/

// Install minikube into /usr/local/bin directory
sudo install minikube /usr/local/bin/
```

Then start the Minikube service in your current machine by running the below command:
```
minikube start
```
To check whether the Kubernetes cluster is running by using Minikube, run the below command:
```
kubectl get node
``` 
You should see the similar output to that shown below:

```
NAME       STATUS   ROLES           AGE    VERSION
minikube   Ready    control-plane   145m   v1.26.3
```

Now you are all set and ready to deploy your services! :hugs:

## (2) Deployments :rocket:

### (2.1) Deploy provider API service

#### (2.1.1) Going through the files inside the service directory

The current structure of the directory looks like the listing below:
```
.
├── chaos-test.js
├── deployment.yml
├── Dockerfile
├── go.mod
├── go.sum
├── provider.go
└── service.yml
```
1. The **chaos-test.js** file is for defining the chaos testing scenario for your test. Let’s skip this file for now. You will learn about this file in detail in the “Implementing chaos testing scenarios” sub-section.

2. The **go.mod** and **go.sum** files define the service dependencies to be run.

3. The **provider.go** file implements all the logic of the service. You create an API that shows the current timestamp in the response body.

```
package main

import (
  "fmt"
  "Time"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.New()

  r.GET("/", func(c *gin.Context) {
    c.String(200, "Current time is: "+fmt.Sprint(time.Now().Unix()))
  })

  r.Run(":3000")
}
```
The service will use the [Gin](https://gin-gonic.com/) framework to implement the web service and will run on port 3000.

4. The Dockerfile defines all the steps to build the service image so that you will deploy the service to Kubernetes later on using this image.
   
```
FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /main

CMD [ "/main" ]
```

To build the Docker image, Docker will perform the following steps :

- Use the “golang:1.19-alpine” image as a Go run time
- Set the working directory to “/app” path
- Copy go.mod, go.sum file to the work directory
- Install the dependencies using go mod download
- Copy all the go files in the root directory to the “/app” path
- Run the go build -o /main to package the service into an execution file named “main”
- Set command to run the service inside the Docker container using “main” execution file

5. The **deployment.yml** file defines the deployment configuration to create the Kubernetes pods for the service.
   
```
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: provider-chaos
spec:
  replicas: 5
  selector:
    matchLabels:
      name: provider-chaos
  template:
    metadata:
      labels:
        name: provider-chaos
    spec:
      containers:
      - name: application
        image: your_docker_hub_account/provider-chaos:latest
        imagePullPolicy: Always
        envFrom:
        - secretRef:
            name: dockerhub-secret
        ports:
          - containerPort: 3000
```

The Kubernetes cluster will execute the deployment to:

- Create five pods since the number of replicas is “5”.
- Name the pods with the prefix “provider-chaos” since the selector.matchLabels.name value is “provider-chaos”.
- Create the pods using the container image. Modify this line so it uses your Docker Hub account name.
- To access the Docker Hub image, Kubernetes will use the secret variable named dockerhub-secret. As a security best practice, you should not show your Docker Hub credentials in a deployment file. Instead, you need to create a secret variable and then access it via the Kubernetes deployment file. You will create this secret variable later in this section.
- Kubernetes expose the running service container via port “3000”.
You need to replace “your_docker_hub_account” with your actual value at the line “image: your_docker_hub_account/provider-chaos:latest” in the deployment.yml file.

6. The **service.yml** file defines a Kubernetes service as a load balancer so that you can access the demo service from outside the Kubernetes cluster.

```
---
apiVersion: v1
kind: Service
metadata:
  name: provider-chaos-service
spec:
  type: LoadBalancer
  ports:
  - name: HTTP
    port: 3001
    targetPort: 3000
  selector:
    name: provider-chaos
```
The Kubernetes cluster will :

- Create a new service with the name as “provider-chaos-service” and type “LoadBalancer”
- Select all the running pods with the prefix “provider-chaos”
- Access these pods via port “3000”
- Then expose these pods via port “3001”

#### (2.1.2) Building the service image and push it to Docker Hub :building_construction:

In order to deploy the service to the Kubernetes cluster, you need to create the Docker image first. From the current terminal, run the following command to create a new Docker image and replace “your_docker_hub_account” with your actual Docker Hub account value.

```
docker build -t your_docker_hub_account/provider-chaos .
```

You should see a similar output as that shown below:
```
...
 => [3/7] COPY go.mod ./    0.1s
 => [4/7] COPY go.sum ./     0.1s
 => [5/7] RUN go mod download   6.7s
 => [6/7] COPY *.go ./                                              0.1s
 => [7/7] RUN go build -o /main                                            16.2s
 => exporting to image                                          3.1s
 => => exporting layers            3.1s
 => => writing image sha256:b714907f712bf2889cde0840e83fdb4b20d4e167aee3462880d28517afb47c1f      0.0s
 => => naming to docker.io/your_docker_hub_account/provider-chaos
```

Log in to Docker Hub from the current terminal by running:

```
docker login
```

The docker command line will ask for your Docker Hub account name and password. After providing your values, you should see the following output showing that you have logged in to Docker Hub successfully.

```
Authenticating with existing credentials...
WARNING! Your password will be stored unencrypted in /home/donald/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
```

Finally, run the following command to publish the service image to your Docker Hub page and replace “your_docker_hub_account” with your actual value.

```
docker push your_docker_hub_account/provider-chaos
```


You should see the following output:

```
Using default tag: latest
The push refers to repository [docker.io/your_docker_hub_account/provider-chaos]
363f5106a57e: Pushed 
2722c1fdbed8: Pushed 
6c35a8e7f30c: Pushed 
20cb7f1cf854: Pushed 
368e5a7a0116: Pushed 
73b40ab56d66: Layer already exists 
d0bf2a758a06: Layer already exists 
10691ab46c83: Layer already exists 
d270ab11cf6e: Layer already exists 
7cd52847ad77: Layer already exists 
latest: digest: sha256:f5b5fd30150b88ab0b6fec772bcde7d69c7df371181af5c80cd395bae79cfab3 size: 2410
```

Now that you have successfully pushed the service image to Docker Hub. Let’s move on to deploy the service to the Kubernetes cluster. :rocket:

#### (2.1.3) Deploying the service to the Kubernetes cluster :rocket:

You need to create a Kubernetes secret named dockerhub-secret to allow Kubernetes to get access to the Docker Hub images. Run the following command to do it, replace “your_docker_hub_user”, “your_docker_hub_password”, and “your_docker_hub_email” with your actual values.


```
kubectl create secret docker-registry dockerhub-secret \
  --docker-server=docker.io \
  --docker-username=your_docker_hub_user \
  --docker-password=your_docker_hub_password \
  --docker-email=your_docker_hub_email
```

You should see the following output from the console showing that you have successfully created the “dockerhub-secret” secret.

```
secret/dockerhub-secret created
```
Then, you deploy the Kubernetes pods for the application service using deployment.yml file by running the following command:

```
kubectl apply -f deployment.yml
```

You should see the following output:


```
deployment.apps/provider-chaos created
```

By default, Minikube does not come with LoadBalancer support activated. You need to run the following command to set up the load balancer in your local machine. This will allow you to access the API service from outside the Kubernetes cluster using a Kubernetes service component. Open up a new terminal and run the following command:

```
minikube tunnel
```

You should see the following output which shows that you have successfully started the minikube tunnel process:

```
Status:
        machine: minikube
        pid: 79041
        route: 10.96.0.0/12 -> 192.168.49.2
        minikube: Running
        services: []
    errors:
                minikube: no errors
                router: no errors
                loadbalancer emulator: no errors

```

Keep this terminal open so that the tunnel process does not stop.

Finally, you deploy the Kubernetes service so that you can access the application service from outside the Kubernetes cluster.

```
kubectl apply -f service.yml
```

You should see the following output:

```
service/provider-chaos-service created
```

Run the following command to get information about all the running Kubernetes services:
```
kubectl get service
```

You should see a similar output to that shown below:

```
NAME                TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)          AGE
kubernetes          ClusterIP      10.96.0.1       <none>          443/TCP          130m
provider-chaos-service   LoadBalancer   10.105.156.41   10.105.156.41   3001:30057/TCP   104s
```

Now the “EXTERNAL-IP” for “provider-chaos-service” is available at the address is “10.105.156.41”. You can try to make an API request to the API to get the current timestamp by running the following command. Remember to replace “139.180.223.84” with your actual EXTERNAL_IP value.

```
curl --location 'http://10.105.156.41:3001/'
```

You should see a similar output to that shown below:

```
Current time is: 1679546283
```

You have now successfully deployed the get time API to the Kubernetes cluster and are able to access the API from your local machine. Let’s move on to deploying the consuming API service.

### (2.2) Deploy consumer API service :rocket:

#### (2.2.1) Going through the files inside the service directory

The current structure of the directory looks like the following:

```
.
├── deployment.yml
├── Dockerfile
├── go.mod
├── go.sum
├── consumer.go
└── service.yml
```
1. The **go.mod** and **go.sum** files are for defining the dependencies that the service needs to use.

2. The **consumer.go** defines all the logic of this service. This service implements an API with the path as “/consuming”, which makes an API request to the get current timestamp API, then responds with the appropriate message depending on the status code of the get current timestamp API.

```
package main

import (
  "fmt"
  "net/HTTP"
  "os"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.New()

  r.GET("/consuming", func(ctx *gin.Context) {

    response, err := http.Get(fmt.Sprintf("http://%s:%s/", os.Getenv("PROVIDER_URL"), os.Getenv("PROVIDER_PORT")))
    if err != nil {
      fmt.Printf("error making http request: %s\n", err)
      os.Exit(1)
    }

    if response.StatusCode == 200 {
      ctx.String(200, "Successfully consuming message")
    } else {
      ctx.String(500, "Internal server error")
    }

  })

  r.Run(":3000")
}
```
This consuming API will return “200” status code with the message “Successfully consuming message” if the get current timestamp API responds with a 200 status code. If the get current timestamp API returns another status code other than “200”, the consuming API will return a “500” status code with the message “Internal server error”.

The consuming API will get the domain and port of the get current timestamp API from environment variables with the keys PROVIDER_URL and PROVDER_PORT. You will provide these environment variables to the Kubernetes pod using the Kubernetes config map.

3. The **Dockerfile** defines all the steps to build the Docker image for the consuming API service. It looks the same as the Dockerfile for the get current timestamp API service.


```
# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /main

CMD [ "/main" ]
```

The **deployment.yml** file tells Kubernetes to create the pods and the config map for the consuming API service.


```
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-consumer-app
spec:
  replicas: 5
  selector:
    matchLabels:
      name: my-consumer-app
  template:
    metadata:
      labels:
        name: my-consumer-app
    spec:
      containers:
      - name: application
        image: your_docker_hub_account/my-consumer-app:latest
        imagePullPolicy: Always
        envFrom:
        - secretRef:
            name: dockerhub-secret
        ports:
          - containerPort: 3000
        env:
        # Define the environment variable
        - name: PROVIDER_URL 
          value: consumer-chaos-service
        - name: PROVIDER_PORT
          value: "3001"
```

You need to make some modifications to the file to match your actual values:

* Replace the “your_docker_hub_account” in the line “image: your_docker_hub_account/my-consumer-app:latest” with your actual Docker Hub account
The Kubernetes pod for the consuming API service will access the Kubernetes config map values using the environment variables from config map.

5. The **service.yml** file defines the Kubernetes service that will serve as a load balancer so that you can access the consuming API service from outside the Kubernetes cluster.

```
---
apiVersion: v1
kind: Service
metadata:
  name: my-consumer-app-service
spec:
  type: LoadBalancer
  ports:
  - name: HTTP
    port: 3001
    targetPort: 3000
  selector:
    name: my-consumer-app
```
Kubernetes will look for the pod which has a name with the prefix “my-consumer-app” then forward container port “3000” to port “3001” of the Kubernetes service named “my-consumer-app-service”.

#### (2.2.2) Building the service image and push it to Docker Hub 

Run the following commands to build the new image for the consumer app and push it to Docker Hub.

```
docker build -t your_docker_hub_account/my-consumer-app .
docker push your_docker_hub_account/my-consumer-app
```

Now that you have successfully pushed the Docker image to your Docker Hub page. Let’s move on to deploy the consuming API service to Kubernetes cluster.

#### (2.2.3) Deploying the service to the Kubernetes cluster

First, run the following command to deploy the Kubernetes pods and config map for the consuming API service.

```
kubectl apply -f deployment.yml
```

Then, run the following command to deploy the Kubernetes service:

```
kubectl apply -f service.yml
```
Run the following command to check whether the services are up and running.

```
kubectl get service
```
You should be able to see the similar output to that shown below:


```
NAME                      TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)          AGE
kubernetes                ClusterIP      10.96.0.1       <none>          443/TCP          155m
my-consumer-app-service   LoadBalancer   10.110.54.31    10.110.54.31    3001:31258/TCP   23m
consumer-chaos-service         LoadBalancer   10.105.156.41   10.105.156.41   3001:30057/TCP   26m
```

You can make an API request to the API “/consuming” using the command shown below (You need to replace the “your_external_ip” with your actual value for the “EXTERNAL-IP” of your “my-consumer-app-service”):

```
curl --location 'http://your_external_ip:3001/consuming'
```
You should see a similar output to that shown below:

```
Successfully consuming message
```
Now you have successfully deployed the consuming API service to the Kubernetes cluster and are able to access the API from outside the cluster. Let’s continue to the next section to apply chaos testing for the two services.

## (3) Chaos Testing

### (3.1) Installing k6 and xk6-disruptor

### (3.2) Implementing chaos testing scenarios

### (3.3) Root cause analysis

### (3.4) Solution

## (4) Conclusion