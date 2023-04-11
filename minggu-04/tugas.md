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
![1]()<br>
The default topology is the minimal topology, which includes one OpenFlow kernel switch connected to two hosts, plus the OpenFlow reference controller.

### Step 3: Interact with Hosts and Switches

Display Mininet CLI commands:<br>
![2]()<br>
Display nodes and links:<br>
![3]()<br>
Dump information about all nodes:<br>
![4]()<br>
Run a command on a host process:<br>
![5]()<br>
![6]()<br>
Print the process list from a host process:<br>
![7]()<br>
![8]()<br>
Now, verify that you can ping from host 0 to host 1:<br>
![9]()<br>
An easier way to run this test is to use the Mininet CLI built-in pingall command, which does an all-pairs ping:<br>
![10]()<br>
Exit the CLI:<br>
![11]()<br>
If Mininet crashes for some reason, clean it up:<br>
![12]()<br>
