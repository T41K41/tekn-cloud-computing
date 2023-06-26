# Docker Networking Hands-on Lab

Laporan beserta gambar dibawah ini adalah hasil praktikum melalui [Docker Networking Hands-on Lab](https://training.play-with-docker.com/docker-networking-hol/), sehingga untuk materi dan penjelasan lebih detailnya dapat diakses melalui web tersebut.

### Task

- [Section #1 - Networking Basics](https://github.com/isaanggi/tekn-cloud-computing/blob/main/minggu-10/docker-networking.md#section-1---networking-basics)
- [Section #2 - Bridge Networking](https://github.com/isaanggi/tekn-cloud-computing/blob/main/minggu-10/docker-networking.md#section-2---bridge-networking)
- [Section #3 - Overlay Networking](https://github.com/isaanggi/tekn-cloud-computing/blob/main/minggu-10/docker-networking.md#section-3---overlay-networking)
- [Cleaning Up](https://github.com/isaanggi/tekn-cloud-computing/blob/main/minggu-10/docker-networking.md#cleaning-up)

## Section #1 - Networking Basics

### Step 1: The Docker Network Command
The ```docker network``` command is the main command for configuring and managing container networks. Run the ```docker network``` command from the first terminal.
<div><img src="gambar/SS1.png"></div>
<div><img src="gambar/SS2.png"></div>

### Step 2: List networks
Run a ```docker network ls``` command to view existing container networks on the current Docker host.
<div><img src="gambar/SS3.png"></div><br>

### Step 3: Inspect a network
The ```docker network inspect``` command is used to view network configuration details. Use ```docker network inspect <network>``` to view configuration details of the container networks on your Docker host. The command below shows the details of the network called ```bridge```.
<div><img src="gambar/SS4.png"></div><br>

### Step 4: List network driver plugins
The ```docker info``` command shows a lot of interesting information about a Docker installation. Run the ```docker info``` command and locate the list of network plugins.
<div><img src="gambar/SS5.png"></div><br>

## Section #2 - Bridge Networking

### Step 1: The Basics
Every clean installation of Docker comes with a pre-built network called bridge. Verify this with the docker ```network ls```.
<div><img src="gambar/SS6.png"></div><br>

Install the ```brctl``` command and use it to list the Linux bridges on your Docker host. You can do this by running ```sudo apt-get install bridge-utils```.
<div><img src="gambar/SS7.png"></div><br>

Then, list the bridges on your Docker host, by running ```brctl show```. You can also use the ```ip a``` command to view details of the docker0 bridge.

<div><img src="gambar/SS8.png"></div><br>

### Step 2: Connect a container
The bridge network is the default network for new containers. This means that unless you specify a different network, all new containers will be connected to the bridge network. Create a new container by running ```docker run -dt ubuntu sleep infinity```. This command will create a new container based on the ```ubuntu:latest``` image and will run the ```sleep``` command to keep the container running in the background. You can verify our example container is up by running ```docker ps```.
<div><img src="gambar/SS9.png"></div><br>

As no network was specified on the ```docker run``` command, the container will be added to the bridge network. Run the ```brctl show``` command again. Notice how the docker0 bridge now has an interface connected. This interface connects the docker0 bridge to the new container just created. You can inspect the bridge network again, by running ```docker network inspect bridge```, to see the new container attached to it.
<div><img src="gambar/SS10.png"></div><br>

### Step 3: Test network connectivity
Ping the IP address of the container from the shell prompt of your Docker host by running ```ping -c5 <IPv4 Address>```. Remember to use the IP of the container in your environment.
<div><img src="gambar/SS11.png"></div><br>

The replies above show that the Docker host can ping the container over the bridge network. But, we can also verify the container can connect to the outside world too. Lets log into the container, install the ```ping program```, and then ping ```www.github.com```.

First, we need to get the ID of the container started in the previous step. You can run ``docker ps`` to get that. Next, lets run a shell inside that ubuntu container, by running ```docker exec -it <CONTAINER ID> /bin/bash```. Next, we need to install the ping program. So, lets run ```apt-get update && apt-get install -y iputils-ping```.
<div><img src="gambar/SS12.png"></div><br>

Lets ping ```www.github.com``` by running ```ping -c5 www.github.com```. Finally, lets disconnect our shell from the container, by running ```exit```. We should also stop this container so we clean things up from this test, by running ```docker stop <CONTAINER ID>```.
<div><img src="gambar/SS13.png"></div><br>

### Step 4: Configure NAT for external connectivity
In this step we’ll start a new NGINX container and map port 8080 on the Docker host to port 80 inside of the container. This means that traffic that hits the Docker host on port 8080 will be passed on to port 80 inside the container.

Start a new container based off the official NGINX image by running ```docker run --name web1 -d -p 8080:80 nginx```. Review the container status and port mappings by running ```docker ps```.
<div><img src="gambar/SS14.png"></div><br>

If for some reason you cannot open a session from a web broswer, you can connect from your Docker host using the ```curl 127.0.0.1:8080``` command.
<div><img src="gambar/SS15.png"></div><br>

## Section #3 - Overlay Networking

### Step 1: The Basics
In this step you’ll initialize a new Swarm, join a single worker node, and verify the operations worked. Run ```docker swarm init --advertise-addr $(hostname -i)```. In the first terminal copy the entire ```docker swarm join ...``` command that is displayed as part of the output from your terminal output. Then, paste the copied command into the second terminal. Run a docker node ls to verify that both nodes are part of the Swarm.
<div><img src="gambar/SS16.png"></div><br>

### Step 2: Create an overlay network
Now that you have a Swarm initialized it’s time to create an overlay network. Create a new overlay network called “overnet” by running ```docker network create -d overlay overnet```. Use the ```docker network ls``` command to verify the network was created successfully. Run the same ```docker network ls``` command from the second terminal.
<div><img src="gambar/SS17.png"></div><br>

Notice that the “overnet” network does not appear in the list. This is because Docker only extends overlay networks to hosts when they are needed. This is usually when a host runs a task from a service that is created on the network. We will see this shortly.

Use the ```docker network inspect <network>``` command to view more detailed information about the “overnet” network. You will need to run this command from the first terminal.
<div><img src="gambar/SS18.png"></div><br>

### Step 3: Create a service
Now that we have a Swarm initialized and an overlay network, it’s time to create a service that uses the network.

Execute the following command from the first terminal to create a new service called myservice on the overnet network with two tasks/replicas.

```
docker service create --name myservice \
--network overnet \
--replicas 2 \
ubuntu sleep infinity
```

Verify that the service is created and both replicas are up by running ```docker service ls```.
<div><img src="gambar/SS19.png"></div><br>

The ```2/2``` in the ```REPLICAS``` column shows that both tasks in the service are up and running. Verify that a single task (replica) is running on each of the two nodes in the Swarm by running ```docker service ps myservice```.
<div><img src="gambar/SS20.png"></div><br>

We can also run ```docker network inspect overnet``` on the second terminal to get more detailed information about the “overnet” network and obtain the IP address of the task running on the second terminal.
<div><img src="gambar/SS21.png"></div><br>

### Step 4: Test the network
To complete this step you will need the IP address of the service task. Execute the following commands, ```docker network inspect overnet```. Notice that the IP address listed for the service task (container) running on the first node is different to the IP address for the service task running on the second node. Note also that they are on the same “overnet” network.
<div><img src="gambar/SS22.png"></div><br>

Run a ```docker ps``` command to get the ID of the service task so that you can log in to it in the next step. Log on to the service task. Be sure to use the container ```ID``` from your environment as it will be different from the example shown below. We can do this by running ```docker exec -it <CONTAINER ID> /bin/bash```. Install the ping command and ping the service task running on the both node.
<div><img src="gambar/SS23.png"></div><br>
<div><img src="gambar/SS24.png"></div><br>
The output above shows that both tasks from the myservice service are on the same overlay network spanning both nodes and that they can use this network to communicate.

### Step 5: Test service discovery

Run ```cat /etc/resolv.conf``` form inside of the container. The value that we are interested in is the ```nameserver 127.0.0.11```. This value sends all DNS queries from the container to an embedded DNS resolver running inside the container listening on ```127.0.0.11:53```. All Docker container run an embedded DNS server at this address. Try and ping the “myservice” name from within the container by running ```ping -c5 myservice```.
<div><img src="gambar/SS25.png"></div><br>

Type the ```exit``` command to leave the ```exec``` container session and return to the shell prompt of your Docker host. Inspect the configuration of the “myservice” service by running ```docker service inspect myservice```. Lets verify that the VIP value matches the value returned by the previous ```ping -c5 myservice``` command.
<div><img src="gambar/SS26.png"></div><br>

## Cleaning Up
Execute the ```docker service rm myservice``` command to remove the service called myservice. Execute the ```docker ps``` command to get a list of running containers. You can use the ```docker kill <CONTAINER ID ...>``` command to kill the ubuntu and nginx containers we started at the beginning.
<div><img src="gambar/SS27.png"></div><br>

Finally, lets remove node1 and node2 from the Swarm. We can use the ```docker swarm leave --force``` command to do that.
<div><img src="gambar/SS28.png"></div>
<div><img src="gambar/SS29.png"></div>

Congratulations! You’ve completed this lab!
