# Rangkuman Software as a service

# Perbedaan antara Iaas,Saas dan Paas
Infrastructure as a Service (IaaS) 
 Ini hanya menyediakan infrastruktur dasar (Mesin virtual, Software Define Network, Storage terlampir). Pengguna akhir harus mengonfigurasi dan mengelola platform dan lingkungan, menyebarkan aplikasi di dalamnya. Ini berarti menjalankan aplikasi di cloud publik. Pengguna menggunakan aplikasi ini melalui Internet. Aplikasi ini dikelola oleh Penyedia Layanan. Beberapa misalnya, dari Penyedia Layanan, adalah SalesForce, Microsoft (Office 365), Oracle, Google (Google Apps), dll.

Platform as a Service (PaaS) :
Provider ini menyediakan platform yang memungkinkan pengguna akhir untuk mengembangkan, menjalankan, dan mengelola aplikasi tanpa kerumitan membangun dan memelihara infrastruktur. Ini menyediakan lingkungan bagi pengembang untuk membangun aplikasi yang dapat digunakan pengguna. IaaS meliputi: -
1.Pengguna membuat mesin virtual (VM) sesuai permintaan.
2.Dari perpustakaan gambar VM.


Software as a Service (SaaS) : 
Kadang-kadang disebut sebagai "perangkat lunak sesuai permintaan". Biasanya diakses oleh pengguna menggunakan thin client melalui browser web. Di SaaS semuanya dapat dikelola oleh vendor: aplikasi, runtime, data, middleware, OS, virtualisasi, server, penyimpanan dan jaringan, Pengguna akhir harus menggunakannya. Ini agak mirip dengan IaaS tetapi perbedaannya adalah:
1.Pengembang menyediakan aplikasi yang dijalankan platform.
2.Mereka tidak secara langsung membuat VM.


# Saas(Software as a Service)
Dengan model ini, satu versi aplikasi, dengan satu konfigurasi digunakan untuk semua pelanggan. Aplikasi ini diinstal pada beberapa mesin untuk mendukung skalabilitas (disebut penskalaan horizontal). Dalam beberapa kasus, versi kedua aplikasi disiapkan untuk menawarkan sekelompok pelanggan tertentu dengan akses ke versi pra-rilis aplikasi untuk tujuan pengujian. Dalam model tradisional ini, setiap versi aplikasi didasarkan pada kode unik. Meskipun pengecualian, beberapa solusi SaaS tidak menggunakan multitenancy, untuk mengelola sejumlah besar pelanggan secara hemat biaya. Apakah multitenancy adalah komponen yang diperlukan untuk perangkat lunak sebagai layanan adalah topik kontroversi. 

 ![saas](https://github.com/T41K41/tekn-cloud-computing/blob/03a58ff9399808b60338b6e8b5b7e7b8ccaee089/minggu-02/saas.webp)
 Ada dua varietas utama SaaS: 
1.	Vertical SaaS A Software 
    yang menjawab kebutuhan industri tertentu (misalnya, perangkat lunak untuk industri perawatan kesehatan, pertanian, real estat, keuangan). 
3.	SaaS Horisontal
    Produk yang berfokus pada kategori perangkat lunak (pemasaran, penjualan, alat pengembang, dan perancangan)
    
![hor-ver](https://github.com/T41K41/tekn-cloud-computing/blob/03a58ff9399808b60338b6e8b5b7e7b8ccaee089/minggu-02/ver-hor.jpg)

# Kelebihan & Kekurangan
kelebihan
•	Kesederhanaan aplikasi Perangkat Lunak Aplikasi Arsitektur SaaS yang dirancang sebagai solusi SaaS biasanya diakses melalui web melalui berbagai jenis perangkat. 

•	Kemajuan dalam bahasa pemrograman sisi klien seperti JavaScript telah menghasilkan antarmuka web yang lebih intuitif dan, dengan demikian, membuat penggunaan aplikasi yang dikirimkan melalui internet semudah digunakan seperti rekan desktop mereka. 

•	Nilai Ekonomis Model pembayaran biaya subskrip bulanan atau tahunan memudahkan bisnis untuk menganggarkan, memasangkannya dengan biaya penyiapan infrastruktur nol, dan mudah untuk melihat bagaimana memilih untuk menggunakan solusi SaaS dapat menghemat uang bisnis. 

•	Keamanan Keamanan adalah aspek penting dari solusi pengembangan perangkat lunak dan platform SaaS tidak berbeda.  

•	Sebagai konsumen aplikasi yang dirancang menggunakan SaaS, Anda tidak perlu khawatir dengan bagaimana data Anda dikunci.  Itu disimpan dengan aman di cloud! Kompatibilitas Dengan instalasi perangkat lunak tradisional .

Kekurangan
•	Kesederhanaan aplikasi Perangkat Lunak Aplikasi Arsitektur SaaS yang dirancang sebagai solusi SaaS biasanya diakses melalui web melalui berbagai jenis perangkat. 

•	Kemajuan dalam bahasa pemrograman sisi klien seperti JavaScript telah menghasilkan antarmuka web yang lebih intuitif dan, dengan demikian, membuat penggunaan aplikasi yang dikirimkan melalui internet semudah digunakan seperti rekan desktop mereka. 

•	Nilai Ekonomis Model pembayaran biaya subskrip bulanan atau tahunan memudahkan bisnis untuk menganggarkan, memasangkannya dengan biaya penyiapan infrastruktur nol, dan mudah untuk melihat bagaimana memilih untuk menggunakan solusi SaaS dapat menghemat uang bisnis. 

•	Kontrol Anda telah mengurangi kontrol sistem TI Anda dan sepenuhnya bergantung pada penyedia layanan. Keamanan data karena sumber daya dibagi di antara entitas yang berbeda di cloud publik, ada risiko pelanggaran yang lebih tinggi. Kunci Anda mungkin merasa sangat sulit dan tidak hemat biaya untuk berpindah dari satu penyedia cloud ke penyedia cloud lainnya jika Anda tidak puas dengan layanannya. Biaya yang dikeluarkan selama periode waktu yang lebih lama (katakanlah 3 tahun atau lebih) mungkin tidak kurang jika dibandingkan dengan memiliki sistem TI di rumah.

•	Keamanan Keamanan adalah aspek penting dari solusi pengembangan perangkat lunak dan platform SaaS tidak berbeda, Sebagai konsumen aplikasi yang dirancang menggunakan SaaS, Anda tidak perlu khawatir dengan bagaimana data Anda dikunci.  Itu disimpan dengan aman di cloud! Kompatibilitas Dengan instalasi perangkat lunak tradisional .

# Cara Membangun Aplikasi SaaS berbasis Cloud

1. Menggunakan bahasa pemograman yang tepat seperti Python karena lebih fleksibel dalam banyak kasus

2. Menggunakan database yang tepat seperti Database Document Oriented seperti MongoDB

3. Menggunakan sistem antrian yang tepat atau untuk sebuah protokol komunikasi asinkron yang memungkinkan pengirim dan penerima berinteraksi disaat yang sama atau dikenal dengan MSMQ, seperti RabbitMq

Dengan Python, MongoDB dan RabbitMQ dapat mencukupi dalam dasar pembangunan SaaS. Mungkin akan ada penambahan lain seperti kebutuhan software pemantauan dan analitik yang tepat.

# Refferensi:
1. [What is the difference between IaaS, SaaS, and PaaS?](https://www.quora.com/What-is-the-difference-between-IaaS-SaaS-and-Paas). 
2. [SaaS Platform Architecture](https://hackernoon.com/saas-software-as-a-service-platform-architecture-757a432270f5).
3. [SaaS (Software as a Service) Platform  itecture](https://www.devteam.space/blog/saas-software-as-a-service-platform-architecture/).
4. [How to build a cloud-based SaaS Application](https://usersnap.com/blog/cloud-based-saas-architecture-fundamentals/).

