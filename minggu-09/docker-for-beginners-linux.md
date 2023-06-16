# Docker for Beginners - Linux

Laporan beserta gambar dibawah ini adalah hasil praktikum melalui [Docker for Beginners - Linux](https://training.play-with-docker.com/beginner-linux/), sehingga untuk materi dan penjelasan lebih detailnya dapat diakses melalui web tersebut.

### Task

- Task 0: Prerequisites 
- Task 1: Run some simple Docker containers 
- Task 2: Package and run a custom app using Docker 
- Task 3: Modify a Running Website

## Task 0: Prerequisites

You will need all of the following to complete this lab: 
- A clone of the lab’s GitHub repo. 
- A DockerID.

### Clone the Lab’s GitHub Repo

Use the following command to clone the lab’s repo from GitHub (you can click the command or manually type it). This will make a copy of the lab’s repo in a new sub-directory called ```linux_tweet_app```.

<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(3).png"></div><br>

### Make sure you have a DockerID

If you do not have a DockerID (a free login used to access Docker Hub), please visit [Docker Hub](https://hub.docker.com/) and register for one. You will need this for later steps.
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(2).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(1).png"></div><br>

## Task 1: Run some simple Docker containers 

There are different ways to use containers. These include: 
1. To run a single task: This could be a shell script or a custom app. 
2. Interactively: This connects you to the container similar to the way you SSH into a remote server. 
3. In the background: For long-running services like websites and databases. 

In this section you’ll try each of those options and see how Docker manages the workload.

### Run a single task in an Alpine Linux container

<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(5).png"></div><br>

### Run an interactive Ubuntu container

<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(6).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(7).png"></div><br>

### Run a background MySQL container

<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(8).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(9).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(10).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(11).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(12).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(13).png"></div><br>

## Task 2: Package and run a custom app using Docker 

In this step you’ll learn how to package your own apps as Docker images using a [Dockerfile](https://docs.docker.com/engine/reference/builder/). 

The Dockerfile syntax is straightforward. In this task, we’re going to create a simple NGINX website from a Dockerfile.

### Build a simple website image

<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(14).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(15).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(16).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(17).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(18).png"></div><br>

## Task 3: Modify a Running Website

### Start our web app with a bind mount

<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(19).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(20).png"></div><br>

### Modify the running website

<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(21).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(22).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(23).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(24).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(25).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(26).png"></div><br>

### Update the image

<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(26).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(27).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(28).png"></div><br>

### Test the new version

<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(29).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(30).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(31).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(32).png"></div><br>

### Push your images to Docker Hub

<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(33).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(34).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(35).png"></div><br>
<div><img src="https://github.com/T41K41/tekn-cloud-computing/blob/174a7c246921c6099f0019166bfd1ce9249fa9fe/minggu-09/gambar/gambar%20(36).png"></div><br>
