# Learn Kubernetes Basics

Laporan beserta gambar dibawah ini adalah hasil praktikum melalui [Learn Kubernetes Basics](https://kubernetes.io/docs/tutorials/kubernetes-basics/), sehingga untuk materi dan penjelasan lebih detailnya dapat diakses melalui web tersebut.

### Tutorials

- [Create a Cluster](https://github.com/isaanggi/tekn-cloud-computing/blob/main/minggu-13/learn-kubernetes-basics.md#create-a-cluster)
- [Deploy an App](https://github.com/isaanggi/tekn-cloud-computing/blob/main/minggu-13/learn-kubernetes-basics.md#deploy-an-app)
- [Explore Your App](https://github.com/isaanggi/tekn-cloud-computing/blob/main/minggu-13/learn-kubernetes-basics.md#explore-your-app)
- [Expose Your App Publicly](https://github.com/isaanggi/tekn-cloud-computing/blob/main/minggu-13/learn-kubernetes-basics.md#expose-your-app-publicly)
- [Scale Your App](https://github.com/isaanggi/tekn-cloud-computing/blob/main/minggu-13/learn-kubernetes-basics.md#scale-your-app)
- [Update Your App](https://github.com/isaanggi/tekn-cloud-computing/blob/main/minggu-13/learn-kubernetes-basics.md#update-your-app)

## Create a Cluster

### Using Minikube to Create a Cluster

Kubernetes coordinates a highly available cluster of computers that are connected to work as a single unit. Kubernetes automates the distribution and scheduling of application containers across a cluster in a more efficient way.<br><br>
A Kubernetes cluster consists of two types of resources:

- The Control Plane coordinates the cluster
- Nodes are the workers that run applications

<div><img src="gambar/ss1.jpg"></div>

The Control Plane is responsible for managing the cluster. A node is a VM or a physical computer that serves as a worker machine in a Kubernetes cluster. The nodes communicate with the control plane using the Kubernetes API, which the control plane exposes. End users can also use the Kubernetes API directly to interact with the cluster.

### Interactive Tutorial - Creating a Cluster

#### Cluster up and running

<div><img src="gambar/ss2.jpg"></div>

#### Cluster version

<div><img src="gambar/ss3.jpg"></div>

#### Cluster details

<div><img src="gambar/ss4.jpg"></div>

## Deploy an App

### Using kubectl to Create a Deployment

You can create and manage a Deployment by using the Kubernetes command line interface, Kubectl. Kubectl uses the Kubernetes API to interact with the cluster. In this module, you'll learn the most common Kubectl commands needed to create Deployments that run your applications on a Kubernetes cluster.

<div><img src="gambar/ss5.jpg"></div>

### Interactive Tutorial - Deploying an App

#### kubectl basics

<div><img src="gambar/ss6.jpg"></div>

#### Deploy our app

<div><img src="gambar/ss7.jpg"></div>

#### View our app

<div><img src="gambar/ss8.jpg"></div>
<div><img src="gambar/ss9.jpg"></div>
<div><img src="gambar/ss10.jpg"></div>
<div><img src="gambar/ss11.jpg"></div>

## Explore Your App

### Viewing Pods and Nodes

When you created a Deployment in Module 2, Kubernetes created a Pod to host your application instance. A Pod is a Kubernetes abstraction that represents a group of one or more application containers (such as Docker), and some shared resources for those containers. Those resources include:

- Shared storage, as Volumes
- Networking, as a unique cluster IP address
- Information about how to run each container, such as the container image version or specific ports to use

<div><img src="gambar/ss12.jpg"></div>

A Pod always runs on a Node. A Node is a worker machine in Kubernetes and may be either a virtual or a physical machine, depending on the cluster. Each Node is managed by the control plane. A Node can have multiple pods, and the Kubernetes control plane automatically handles scheduling the pods across the Nodes in the cluster. The control plane's automatic scheduling takes into account the available resources on each Node.
Every Kubernetes Node runs at least:

- Kubelet, a process responsible for communication between the Kubernetes control plane and the Node; it manages the Pods and the containers running on a machine.
- A container runtime (like Docker) responsible for pulling the container image from a registry, unpacking the container, and running the application.

<div><img src="gambar/ss13.jpg"></div>

### Interactive Tutorial - Exploring Your App

#### Step 1 Check application configuration

<div><img src="gambar/ss14.jpg"></div>

#### Step 2 Show the app in the terminal

<div><img src="gambar/ss15.jpg"></div>
<div><img src="gambar/ss16.jpg"></div>

#### Step 3 View the container logs

<div><img src="gambar/ss17.jpg"></div>

#### Step 4 Executing command on the container

<div><img src="gambar/ss18.jpg"></div>
<div><img src="gambar/ss19.jpg"></div>
<div><img src="gambar/ss20.jpg"></div>

## Expose Your App Publicly

### Using a Service to Expose Your App

A Service routes traffic across a set of Pods. Services are the abstraction that allows pods to die and replicate in Kubernetes without impacting your application. Discovery and routing among dependent Pods (such as the frontend and backend components in an application) are handled by Kubernetes Services.

Services match a set of Pods using labels and selectors, a grouping primitive that allows logical operation on objects in Kubernetes. Labels are key/value pairs attached to objects and can be used in any number of ways:

- Designate objects for development, test, and production
- Embed version tags
- Classify an object using tags

<div><img src="gambar/ss21.jpg"></div>

### Interactive Tutorial - Exposing Your App

#### Step 1 Create a new service

<div><img src="gambar/ss22.jpg"></div>
<div><img src="gambar/ss23.jpg"></div>

#### Step 2 Using labels

<div><img src="gambar/ss24.jpg"></div>
<div><img src="gambar/ss25.jpg"></div>
<div><img src="gambar/ss26.jpg"></div>
<div><img src="gambar/ss27.jpg"></div>

#### Step 3 Deleting a service

<div><img src="gambar/ss28.jpg"></div>

## Scale Your App

### Running Multiple Instances of Your App

Scaling is accomplished by changing the number of replicas in a Deployment

<div><img src="gambar/ss29.jpg"></div>
<div><img src="gambar/ss30.jpg"></div>

Scaling out a Deployment will ensure new Pods are created and scheduled to Nodes with available resources. Scaling will increase the number of Pods to the new desired state. Kubernetes also supports autoscaling of Pods, but it is outside of the scope of this tutorial. Scaling to zero is also possible, and it will terminate all Pods of the specified Deployment.

Running multiple instances of an application will require a way to distribute the traffic to all of them. Services have an integrated load-balancer that will distribute network traffic to all Pods of an exposed Deployment. Services will monitor continuously the running Pods using endpoints, to ensure the traffic is sent only to available Pods.

### Interactive Tutorial - Scaling Your App

#### Step 1 Scaling a deployment

<div><img src="gambar/ss31.jpg"></div>
<div><img src="gambar/ss32.jpg"></div>
<div><img src="gambar/ss33.jpg"></div>
<div><img src="gambar/ss34.jpg"></div>

#### Step 2 Load Balancing

<div><img src="gambar/ss35.jpg"></div>
<div><img src="gambar/ss36.jpg"></div>

#### Step 3 Scale Down

<div><img src="gambar/ss37.jpg"></div>

## Update Your App

### Performing a Rolling Update

Similar to application Scaling, if a Deployment is exposed publicly, the Service will load-balance the traffic only to available Pods during the update. An available Pod is an instance that is available to the users of the application.

Rolling updates allow the following actions:

- Promote an application from one environment to another (via container image updates)
- Rollback to previous versions
- Continuous Integration and Continuous Delivery of applications with zero downtime

<div><img src="gambar/ss38.jpg"></div>
<div><img src="gambar/ss39.jpg"></div>
<div><img src="gambar/ss40.jpg"></div>
<div><img src="gambar/ss41.jpg"></div>

### Interactive Tutorial - Updating Your App

#### Step 1 Update the version of the app

<div><img src="gambar/ss42.jpg"></div>
<div><img src="gambar/ss43.jpg"></div>
<div><img src="gambar/ss44.jpg"></div>

#### Step 2 Verify an update

<div><img src="gambar/ss45.jpg"></div>
<div><img src="gambar/ss46.jpg"></div>

#### Step 3 Rollback an update

<div><img src="gambar/ss47.jpg"></div>
<div><img src="gambar/ss48.jpg"></div>
<div><img src="gambar/ss49.jpg"></div>
<div><img src="gambar/ss50.jpg"></div>
