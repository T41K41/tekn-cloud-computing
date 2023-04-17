# Apache OFBiz Installation

### Install Packet yang dibutuhkan

Install paket Java dan Gradle

```bash
t41k41@t41k41kom:~$ sudo apt-get install apt-transport-https ca-certificates wget dirmngr gnupg software-properties-common unzip -y
t41k41@t41k41kom:~$ sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 8AC3B29174885C03
t41k41@t41k41kom:~$ sudo add-apt-repository --yes https://adoptopenjdk.jfrog.io/adoptopenjdk/deb/
t41k41@t41k41kom:~$ sudo apt-get install adoptopenjdk-8-hotspot -y
t41k41@t41k41kom:~$ sudo apt-get install gradle
```

Setelah itu cek apakah sudah terinstall dengan baik.<br>
![1](image\installasi\install-01.png)<br>
![2](image\installasi\install-02.png)<br>
![3](image\installasi\install-03.png)<br>
![4](image\installasi\install-04.png)<br>

### Instalasi OFBiz

Download OFBiz, kemudian extract file yang telah didownload.

```bash
t41k41@t41k41kom:~$ wget https://dlcdn.apache.org/ofbiz/apache-ofbiz-18.12.07.zip
t41k41@t41k41kom:~$ unzip apache-ofbiz-18.12.07
t41k41@t41k41kom:~$ cd apache-ofbiz-18.12.07
t41k41@t41k41kom:~/apache-ofbiz-18.12.07$ sudo ./gradlew cleanAll loadDefault
```

Tunggu sampai proses instalasi selesai.<br>
![5](image\installasi\install-05.png)<br>

### Akses OFBiz

Jalankan service OFBiz, kemudian akses melalui browser.

```bash
t41k41@t41k41kom:~/apache-ofbiz-18.12.07$ sudo ./gradlew "ofbiz --load-data readers=seed"
t41k41@t41k41kom:~/apache-ofbiz-18.12.07$sudo ./gradlew "ofbiz --load-data readers=seed,seed-initial,ext"
t41k41@t41k41kom:~/apache-ofbiz-18.12.07$ sudo ./gradlew ofbiz
```

Untuk login, user/password: admin/ofbiz<br>
Default dashboard: https://SERVER_IP:8443/ordermgr/control/main<br>
Catalog Manager: https://SERVER_IP:8443/catalog<br>
E-Commerce: https://SERVER_IP:8443/ecommerce<br>
WebTools: https://SERVER_IP:8443/webtools<br>

![6](image\installasi\akses.png)<br>
![7](image\installasi\akses-1.png)<br><br>
![8](image\installasi\akses-2.png)<br>
![9](image\installasi\akses-4.png)<br>
