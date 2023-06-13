# Diagram - Docker Compose

Diagram:<br>
<div align="center"><img src="gambar/tugas/diagram.jpg"></div>

Penjelasan:<br>
- Docker image: Image yang dapat kita build atau pull melalui Docker Regitery (Docker Hub).
- Container: Container dapat membawa image tadi dan digunakan untuk membantu development. Satu Image dapat digunakan banyak Container.
- Dockerd: berisikan Images serta Container.
- Docker Client: Sebagai antarmuka sisi client, dapat melakukan perintah melalui terminal untuk menjalankan, membuat dan lain-lain pada docker.
- Docker Compose: Dapat menjalankan lebih dari satu contener sekali jalan dan untuk memulai hanya perlu up down tanpa remove container atau dapat juga untuk menajalnkan service di swarm.
- Docker Swarm: Sebagai bentuk orkestrasi docker paling sederhana untuk meenjalankan aplikasi di atas cluster. Semisal php html biasa hanya bisa jalan di 1 host. Jika menggunakan pake docker swarm bisa kita duplikat dan saling load balance ke banyak host.
