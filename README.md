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
  - [(2) Deployments](#2-deployments)
    - [(2.1) Deploy provider API service](#21-deploy-provider-api-service)
    - [(2.2) Deploy consumer API service](#22-deploy-consumer-api-service)
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

Now you are all set and ready to deploy your services!:hugs:

## (2) Deployments

### (2.1) Deploy provider API service

### (2.2) Deploy consumer API service

## (3) Chaos Testing

### (3.1) Installing k6 and xk6-disruptor

### (3.2) Implementing chaos testing scenarios

### (3.3) Root cause analysis

### (3.4) Solution

## (4) Conclusion