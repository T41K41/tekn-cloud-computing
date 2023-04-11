# Mininet Walkthrough on Ubuntu

### Step 1: Install

Login, install net-tools, git clone and install Mininet on your Ubuntu system.

```bash
osboxes@isaanggi:~$ sudo apt-get install net-tools
osboxes@isaanggi:~$ git clone https://github.com/mininet/mininet
osboxes@isaanggi:~$ mininet/util/install.sh -w
osboxes@isaanggi:~$ mininet/util/install.sh -a
```

Wait until the installation is complete.

### Step 2: Minimal Topology

Start a minimal topology and enter the CLI:<br>
![1](https://github.com/T41K41/tekn-cloud-computing/blob/935ef3e9400c5eeaec1c2cc94b691738dbf00640/minggu-04/image/mininet-1.png)<br>
The default topology is the minimal topology, which includes one OpenFlow kernel switch connected to two hosts, plus the OpenFlow reference controller.

### Step 3: Interact with Hosts and Switches

Display Mininet CLI commands:<br>
![2](https://github.com/T41K41/tekn-cloud-computing/blob/935ef3e9400c5eeaec1c2cc94b691738dbf00640/minggu-04/image/mininet-2.png)<br>
Display nodes and links:<br>
![3](https://github.com/T41K41/tekn-cloud-computing/blob/935ef3e9400c5eeaec1c2cc94b691738dbf00640/minggu-04/image/mininet-3.png)<br>
Dump information about all nodes:<br>
![4](https://github.com/T41K41/tekn-cloud-computing/blob/935ef3e9400c5eeaec1c2cc94b691738dbf00640/minggu-04/image/mininet-4.png)<br>
Run a command on a host process:<br>
![5](https://github.com/T41K41/tekn-cloud-computing/blob/935ef3e9400c5eeaec1c2cc94b691738dbf00640/minggu-04/image/mininet-5.png)<br>
![6](https://github.com/T41K41/tekn-cloud-computing/blob/935ef3e9400c5eeaec1c2cc94b691738dbf00640/minggu-04/image/mininet-6.png)<br>
Print the process list from a host process:<br>
![7](https://github.com/T41K41/tekn-cloud-computing/blob/935ef3e9400c5eeaec1c2cc94b691738dbf00640/minggu-04/image/mininet-7.png)<br>
![8](https://github.com/T41K41/tekn-cloud-computing/blob/935ef3e9400c5eeaec1c2cc94b691738dbf00640/minggu-04/image/mininet-8.png)<br>
Now, verify that you can ping from host 0 to host 1:<br>
![9](https://github.com/T41K41/tekn-cloud-computing/blob/935ef3e9400c5eeaec1c2cc94b691738dbf00640/minggu-04/image/mininet-9.png)<br>
An easier way to run this test is to use the Mininet CLI built-in pingall command, which does an all-pairs ping:<br>
![10](https://github.com/T41K41/tekn-cloud-computing/blob/935ef3e9400c5eeaec1c2cc94b691738dbf00640/minggu-04/image/mininet-10.png)<br>
Exit the CLI:<br>
![11](https://github.com/T41K41/tekn-cloud-computing/blob/935ef3e9400c5eeaec1c2cc94b691738dbf00640/minggu-04/image/mininet-11.png)<br>
If Mininet crashes for some reason, clean it up:<br>
![12](https://github.com/T41K41/tekn-cloud-computing/blob/935ef3e9400c5eeaec1c2cc94b691738dbf00640/minggu-04/image/mininet-12.png)<br>
