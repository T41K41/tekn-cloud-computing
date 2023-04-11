# OpenStack Deployment on Ubuntu with DevStack

### Step 1: Update Ubuntu System

Login, update, upgrade and reboot your Ubuntu system.

```bash
t41k41@t41k41kom:~$ sudo apt-get update
t41k41@t41k41kom:~$ sudo apt -y upgrade
t41k41@t41k41kom:~$ sudo apt -y dist-upgrade
t41k41@t41k41kom:~$ sudo reboot
```

### Step 2: Add Stack User

For other installations of Ubuntu 20.04.5, run the commands below to create DevStack deployment user.

```bash
t41k41@t41k41kom:~$ sudo useradd -s /bin/bash -d /opt/stack -m stack
```

Enable sudo privileges for this user without need for a password.

```bash
pindo@pindokom:~$ echo "stack ALL=(ALL) NOPASSWD: ALL" | sudo tee /etc/sudoers.d/stack
stack ALL=(ALL) NOPASSWD: ALL
```

Switch to stack user, test the internet connection and check version of Python.

```bash
t41k41@t41k41kom:~$ sudo su - stack
stack@t41k41kom:~$
```

![1](image\switchstack.png)

### Step 3: Download DevStack

Clone Destack deployment code from Github.

```bash
stack@t41k41kom:~$ git clone https://git.openstack.org/openstack-dev/devstack
```

![2](image\git.png)
Create a local.conf file with 4 passwords and Host IP address.

```bash
stack@t41k41kom:~$ cd devstack
stack@t41k41kom:~/devstack$ nano local.conf
```

![3](image\localconf.png)

### Step 4: Start Openstack Deployment on Ubuntu 20.04.5 with DevStack

Now that you’ve configured the minimum required config to get started with DevStack, start the installation of Openstack.

```bash
stack@t41k41kom:~/devstack$ FORCE=yes ./stack.sh
```

DevStack will install;
This will take a 15 – 20 minutes, largely depending on the speed of your internet connection. At the end of the installation process, you should see output like this:

```bash
This is your host IP address: 10.0.2.15
This is your host IPv6 address: 2401:ee00:1040:100:2c8d:5afa:b7e:c0a0
Horizon is now available at http://10.0.2.15/dashboard
Keystone is serving at http://10.0.2.15/identity/

The default users are: admin and demo
The password: StrongAdminSecret
```

![4](image\gagalinstall.png)
