# Improve Application Resiliency With Chaos Testing

## (1) Tutorial Intro

### (1.1) Introduction

Testing is essential to ensure that your software application performs as expected. No doubt about it. But even if your application runs smoothly in a testing environment, unexpected incidents can occur in production, such as a crashed database or a lost connection to a third-party service. To ensure your application can handle such incidents and minimize their impact, you need to incorporate chaos testing into your delivery process. In this article, we’ll explore how to implement chaos testing using k6 for two backend services, to help you ensure the resilience of your application.

#### What is chaos testing?
Chaos testing is a software testing methodology that involves intentionally injecting faults and disruptions into your application to identify its weaknesses. It is often used with other testing methodologies, such as functional, load, and performance, to ensure your application’s overall quality and reliability.

#### Why implement chaos testing?
Chaos testing ensures that your application can withstand unexpected failures and disruptions in a production environment. By simulating different scenarios and failure modes, you identify potential issues in your app and develop solutions to resolve them.

#### Chaos xk6-disruptor

Currently, the xk6-disruptor only supports chaos testing for an application running in Kubernetes. The xk6-disruptor runs in your local machine and has two main components:

The xk6-disruptor extension works as a Javascript API to inject faults into your app through Kubernetes components, whether through Kubernetes pods (using PodDisruptor) or Kubernetes services (using ServiceDisruptor).
The xk6-disruptor-agent is installed into the Kubernetes by the xk6-disruptor. Then the xk6-disruptor-agent will inject faults into the Kubernetes component (Kubernetes pod or service) according to instructions from the xk6-disruptor extension.
So, put it simply, the extension exposes the Javascript API, and the agent will inject the faults.

##### Pod disruptor

The xk6-disruptor extension uses Pod disruptor to inject faults to the app inside the Kubernetes cluster in three steps:

The pod disruptor chooses the application pod that it wants to inject faults into
The pod disruptor attaches the disruptor agent to the application pod
The disruptor agent injects faults into the application container
The below diagram illustrates how the pod disruptor works:

![disruptor](./media/disruptor.png)

##### Service disruptor

The xk6-disruptor extension uses the service disruptor as a wrapper around the pod disruptor. Basically, the service disruptor works in three steps:

The service disruptor chooses the Kubernetes service pod that it wants to inject faults into
The service disruptor creates a pod disruptor to interact with the Kubernetes service pod through the pod disruptor
The pod disruptor attaches the disruptor agent to the Kubernetes service pod, then the disruptor agent injects faults into the application container.

#### Chaos testing example

To better understand how to apply chaos testing for a software application, let’s build two simple backend services, as detailed below:

**Service 1**: a Go service that creates an API to display the current timestamp. Let’s call this API a get time API.
**Service 2**: a Go service that implements an API consuming the get time API. This API will respond “Successfully consuming message” if the response status code of the get timestamp API is 200. Let’s call this API a consuming API

You will build these services using the following three steps:

- Clone the GitHub repository
- Build the Docker image for each service using Dockerfile
- Deploy each service to the Kubernetes cluster

## (1.2) Prerequisites

To follow along with the steps in the article, you need to fulfill the following prerequisites:

* Git installed
* Docker engine in your machine to build the service images.
* The kubectl command line tool to interact with the Kubernetes cluster.
* Have an account on Docker hub to store the service images.

## (1.3) Setting up a Kubernetes cluster

## (2) Deployments

### (2.1) Deploy provider API service

### (2.2) Deploy consumer API service

## (3) Chaos Testing

### (3.1) Installing k6 and xk6-disruptor

### (3.2) Implementing chaos testing scenarios

### (3.3) Root cause analysis

### (3.4) Solution

## (4) Conclusion