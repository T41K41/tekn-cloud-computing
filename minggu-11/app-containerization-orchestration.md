# Application Containerization and Microservice Orchestration

Laporan beserta gambar dibawah ini adalah hasil praktikum melalui [Application Containerization and Microservice Orchestration](https://training.play-with-docker.com/microservice-orchestration/), sehingga untuk materi dan penjelasan lebih detailnya dapat diakses melalui web tersebut.

### Steps

- [Stage Setup](app-containerization-orchestration.md#stage-setup)
- [Step 0: Basic Link Extractor Script](app-containerization-orchestration.md#step-0-basic-link-extractor-script)
- [Step 1: Containerized Link Extractor Script](app-containerization-orchestration.md#step-1-containerized-link-extractor-script)
- [Step 2: Link Extractor Module with Full URI and Anchor Text](app-containerization-orchestration.md#step-2-link-extractor-module-with-full-uri-and-anchor-text)
- [Step 3: Link Extractor API Service](app-containerization-orchestration.md#step-3-link-extractor-api-service)
- [Step 4: Link Extractor API and Web Front End Services](app-containerization-orchestration.md#step-4-link-extractor-api-and-web-front-end-services)
- [Step 5: Redis Service for Caching](app-containerization-orchestration.md#step-5-redis-service-for-caching)
- [Step 6: Swap Python API Service with Ruby](app-containerization-orchestration.md#step-6-swap-python-api-service-with-ruby)
- [Conclusions](app-containerization-orchestration.md#conclusions)

## Stage Setup

Let’s get started by first cloning the demo code repository, changing the working directory, and checking the ```demo``` branch out.

<div><img src="gambar/ss1.png"></div>
<div><img src="gambar/ss2.png"></div>

## Step 0: Basic Link Extractor Script

Checkout the ```step0``` branch and list files in it. The ```linkextractor.py``` file is the interesting one here, so let’s look at its contents:

<div><img src="gambar/ss3.png"></div>

However, this seemingly simple script might not be the easiest one to run on a machine that does not meet its requirements. The ```README.md``` file suggests how to run it, so let’s give it a try. When we tried to execute it as a script, we got the ```Permission denied``` error. Let’s check the current permissions on this file. This current permission ```-rw-r--r--``` indicates that the script is not set to be executable. We can either change it by running ```chmod a+x linkextractor.py``` or run it as a Python program instead of a self-executing script as illustrated below:

<div><img src="gambar/ss4.png"></div>

Here we got the first ```ImportError``` message because we are missing the third-party package needed by the script. We can install that Python package (and potentially other missing packages) using one of the many techniques to make it work, but it is too much work for such a simple script, which might not be obvious for those who are not familiar with Python’s ecosystem.

This is where application containerization tools like Docker come in handy. In the next step we will try to containerize this script and make it easier to execute.

## Step 1: Containerized Link Extractor Script

Checkout the ```step1``` branch and list files in it. We have added one new file (i.e., ```Dockerfile```) in this step. Let’s look into its contents:

<div><img src="gambar/ss5.png"></div>

Using this ```Dockerfile``` we can prepare a Docker image for this script. So far, we have just described how we want our Docker image to be like, but didn’t really build one. So let’s do just that:

<div><img src="gambar/ss6.png"></div>

We have created a Docker image named ```linkextractor:step1``` based on the ```Dockerfile``` illustrated above. If the build was successful, we should be able to see it in the list of image:

<div><img src="gambar/ss7.png"></div>

This image should have all the necessary ingredients packaged in it to run the script anywhere on a machine that supports Docker. Now, let’s run a one-off container with this image and extract links from some live web pages. This outputs a single link that is present in the simple ```example.com``` web page. Let’s try it on a web page with more links in it:

<div><img src="gambar/ss8.png"></div>

In the next step we will make these changes and some other improvements to the script.

## Step 2: Link Extractor Module with Full URI and Anchor Text

Checkout the step2 branch and list files in it.

<div><img src="gambar/ss9.png"></div>

In this step the ```linkextractor.py``` script is updated with the following functional changes. Let’s have a look at the updated script.

<div><img src="gambar/ss10.png"></div>

Let’s have a look at the updated script. We have used a new tag ```linkextractor:step2``` for this image so that we don’t overwrite the image from the ```step1``` to illustrate that they can co-exist and containers can be run using either of these images.

<div><img src="gambar/ss11.png"></div>

Running a one-off container using the ```linkextractor:step2``` image should now yield an improved output:

<div><img src="gambar/ss12.png"></div>

Running a container using the previous image linkextractor:step1 should still result in the old output:

<div><img src="gambar/ss13.png"></div>

So far, we have learned how to containerize a script with its necessary dependencies to make it more portable. We have also learned how to make changes in the application and build different variants of Docker images that can co-exist. In the next step we will build a web service that will utilize this script and will make the service run inside a Docker container.

## Step 3: Link Extractor API Service

Checkout the ```step3``` branch and list files in it. Let’s first look at the ```Dockerfile``` for changes:

<div><img src="gambar/ss14.png"></div>

The ```linkextractor.py``` module remains unchanged in this step, so let’s look into the newly added ```main.py``` file:

<div><img src="gambar/ss15.png"></div>

It’s time to build a new image with these changes in place:

<div><img src="gambar/ss16.png"></div>
<div><img src="gambar/ss17.png"></div>

Then run the container in detached mode (```-d``` flag) so that the terminal is available for other commands while the container is still running. Note that we are mapping the port ```5000``` of the container with the ```5000``` of the host (using ```-p 5000:5000``` argument) to make it accessible from the host. We are also assigning a name (```--name=linkextractor```) to the container to make it easier to see logs and kill or remove the container.

<div><img src="gambar/ss18.png"></div>

We can now make an HTTP request in the form ```/api/<url>``` to talk to this server and fetch the response containing extracted links. Since the container is running in detached mode, so we can’t see what’s happening inside, but we can see logs using the name ```linkextractor``` we assigned to our container:

<div><img src="gambar/ss19.png"></div>

We can see the messages logged when the server came up, and an entry of the request log when we ran the curl command. Now we can kill and remove this container:

<div><img src="gambar/ss20.png"></div>

In this step we have successfully ran an API service listening on port ```5000```. This is great, but APIs and JSON responses are for machines, so in the next step we will run a web service with a human-friendly web interface in addition to this API service.

## Step 4: Link Extractor API and Web Front End Services

Checkout the ```step4``` branch and list files in it.

<div><img src="gambar/ss21.png"></div>

In this step we are planning to run two separate containers, one for the API and the other for the web interface. The latter needs a way to talk to the API server. For the two containers to be able to talk to each other, we can either map their ports on the host machine and use that for request routing or we can place the containers in a single private network and access directly. Docker has excellent support for networking and provides helpful commands for dealing with networks. Additionally, in a Docker network containers identify themselves using their names as hostnames to avoid hunting for their IP addresses in the private network. However, we are not going to do any of this manually, instead we will be using Docker Compose to automate many of these tasks.

So let’s look at the ```docker-compose.yml``` file we have:

<div><img src="gambar/ss22.png"></div>

Now, let’s have a look at the user-facing ```www/index.php``` file:

<div><img src="gambar/ss23.png"></div>

Let’s bring these services up in detached mode using ```docker-compose``` utility. Checking for the list of running containers confirms that the two services are indeed running:

<div><img src="gambar/ss24.png"></div>
<div><img src="gambar/ss25.png"></div>

We should now be able to talk to the API service as before:

<div><img src="gambar/ss26.png"></div>

To access the web interface [click to open the Link Extractor](https://training.play-with-docker.com/). Then fill the form with ```https://training.play-with-docker.com/``` (or any HTML page URL of your choice) and submit to extract links from it.

<div><img src="gambar/ss27.png"></div>

Now, let’s modify the ```www/index.php``` file to replace all occurrences of ```Link Extractor``` with ```Super Link Extractor``` :

<div><img src="gambar/ss28.png"></div>

Reloading the web interface of the application (or [clicking here](https://training.play-with-docker.com/)) should now reflect this change in the title, header, and footer. 

<div><img src="gambar/ss29.png"></div>

Let’s revert these changes now to clean the Git tracking:

<div><img src="gambar/ss30.png"></div>
<div><img src="gambar/ss31.png"></div>

Before we move on to the next step we need to shut these services down, but Docker Compose can help us take care of it very easily:

<div><img src="gambar/ss32.png"></div>

In the next step we will add one more service to our stack and will build a self-contained custom image for our web interface service.

## Step 5: Redis Service for Caching

Checkout the ```step5``` branch and list files in it.

<div><img src="gambar/ss33.png"></div>

Let’s first inspect the newly added ```Dockerfile``` under the ```./www``` folder:

<div><img src="gambar/ss34.png"></div>

Next, we will look at the API server’s ```api/main.py``` file where we are utilizing the Redis cache:

<div><img src="gambar/ss35.png"></div>

Now, let’s look into the updated ```docker-compose.yml``` file:

<div><img src="gambar/ss36.png"></div>

Let’s boot these services up:

<div><img src="gambar/ss37.png"></div>

Now, that all three services are up, access the web interface by [clicking the Link Extractor](https://training.play-with-docker.com/). There should be no visual difference from the previous step. However, if you extract links from a page with a lot of links, the first time it should take longer, but the successive attempts to the same page should return the response fairly quickly. To check whether or not the Redis service is being utilized, we can use ```docker-compose exec``` followed by the ```redis``` service name and the Redis CLI’s [monitor](https://redis.io/commands/monitor) command:

<div><img src="gambar/ss38.png"></div>

Now, try to extract links from some web pages using the web interface and see the difference in Redis log entries for pages that are scraped the first time and those that are repeated. 

<div><img src="gambar/ss39.png"></div>
<div><img src="gambar/ss40.png"></div>
<div><img src="gambar/ss41.png"></div>

Before continuing further with the tutorial, stop the interactive ```monitor``` stream as a result of the above ```redis-cli``` command by pressing ```Ctrl + C``` keys while the interactive terminal is in focus. Now that we are not mounting the ```/www``` folder inside the container, local changes should not reflect in the running service:

<div><img src="gambar/ss42.png"></div>

Verify that the changes made locally do not reflect in the running service by reloading the web interface and then revert changes:

<div><img src="gambar/ss43.png"></div>
<div><img src="gambar/ss44.png"></div>
<div><img src="gambar/ss45.png"></div>

Now, shut these services down and get ready for the next step:

<div><img src="gambar/ss46.png"></div>

We have successfully orchestrated three microservices to compose our Link Extractor application. We now have an application stack that represents the architecture illustrated in the figure shown in the introduction of this tutorial. In the next step we will explore how easy it is to swap components from an application with the microservice architecture.

## Step 6: Swap Python API Service with Ruby

Checkout the step6 branch and list files in it.

<div><img src="gambar/ss47.png"></div>

Notice that the ./api folder does not contain any Python scripts, instead, it now has a Ruby file and a Gemfile to manage dependencies. Let’s have a quick walk through the changed files:

<div><img src="gambar/ss48.png"></div>

This Ruby file is almost equivalent to what we had in Python before, except, in addition to that it also logs the link extraction requests and corresponding cache events. In a microservice architecture application swapping components with an equivalent one is easy as long as the expectations of consumers of the component are maintained.

<div><img src="gambar/ss49.png"></div>

Above Dockerfile is written for the Ruby script and it is pretty much self-explanatory.

<div><img src="gambar/ss50.png"></div>

The ```docker-compose.yml``` file has a few minor changes in it. The ```api``` service image is now named ```linkextractor-api:step6-ruby```, the port mapping is changed from ```5000``` to ```4567``` (which is the default port for Sinatra server), and the `API_ENDPOINT` environment variable in the ```web``` service is updated accordingly so that the PHP code can talk to it.

With these in place, let’s boot our service stack:

<div><img src="gambar/ss51.png"></div>

We should now be able to access the API (using the updated port number):

<div><img src="gambar/ss52.png"></div>

Now, open the web interface by [clicking the Link Extractor](https://training.play-with-docker.com/) and extract links of a few URLs. Also, try to repeat these attempts for some URLs.

<div><img src="gambar/ss53.png"></div>

We can shut the stack down now. Since we have persisted logs, they should still be available after the services are gone:

<div><img src="gambar/ss54.png"></div>

This illustrates that the caching is functional as the second attempt to the ```http://example.com/``` resulted in a cache ```HIT```.

In this step we explored the possibility of swapping components of an application with microservice architecture with their equivalents without impacting rest of the parts of the stack. We have also explored data persistence using bind mount volumes that persists even after the containers writing to the volume are gone.

So far, we have used ```docker-compose``` utility to orchestrate the application stack, which is good for development environment, but for production environment we use ```docker stack deploy``` command to run the application in a [Docker Swarm Cluster](https://training.play-with-docker.com/swarm-stack-intro).

## Conclusions

We started this tutorial with a simple Python script that scrapes links from a given web page URL. We demonstrated various difficulties in running the script. We then illustrated how easy to run and portable the script becomes onces it is containerized. In the later steps we gradually evolved the script into a multi-service application stack. In the process we explored various concepts of microservice architecture and how Docker tools can be helpful in orchestrating a multi-service stack. Finally, we demonstrated the ease of microservice component swapping and data persistence.
