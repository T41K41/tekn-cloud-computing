# 01.installasi Git pada Windows

sebelum kita menginstall Git di windows, kita sudah harus memiliki editor teks. Bisa menggunakan [Notepad++](https://notepad-plus-plus.org/) atau [Visual Studion Code](https://code.visualstudio.com/) atau [Vim](https://www.vim.org/).

1. Setelah download [Git](https://git-scm.com/downloads), lalu open file installernya yang telah di download, lalu akan muncul lisensi, selanjutnya klik **Next**.

![install-01]

2. Kemudian pilih folder untuk penyimpanan Git, Secara default Git akan tersimpan di folder C:\Program Files\Git, setelah itu klik **Next**.

![install-02]

3. Lalu klik **Next**, biarkan settingan Git di setting secara default.

![install-03]

4. Klik **Next**.

![install-04]

5. Lalu pilih editor teks yang ingin digunakan (Disini saya menggunakan editor teks Visual Code Code), lalu klik **Next**.

![install-05]

6. Checklist pada pilihan **Let Git Decide** selanjutnya klik **Next**.

![install-06]

7. Selanjutnya checklist pada pilihan **Git from the command line.....**, lalu klik **Next**.

![install-07]

8. Selanjutnya checklist pada pilihan **Use Bundled OpenSSH**, lalu klik **Next**.

![install-08]

9. Checklist pada pilihan **Use the OpenSSH library**, klik **Next**.

![install-09]

10. Checklist pada pilihan **Checkout windows-style....**, klik **Next**.

![install-10]

11. Checklist lagi pada pilihan **Use MinTTY...**, klik **Next**.

![install-11]

12. Checklist pada pilihan **Default**, klik **Next**.

![install-12]

13. Lalu pilih **Git Credential Manager**, **Next**.

![install-13]

14. Pilih **Enable file system caching**, **Next**.

![install-14]

15. Checklist pada **Enable experimental.....**, klik **Next**.

![install-15]

16. Tunggu hingga proses installasi berhasil.

![install-16]

17. Lalu klik **Finish**.

![install-17]

18. Jika berhasil terinstal, maka pada windows terdapat aplikasi **Git**.

![install-18]

19. Tampilan jika kita menggunakan aplikasi **Git-Bash**.

![install-19]

20. Tampilan jika kita menggunakan aplikasi **Git-Gui**.

![install-20]

21. Cek versi Git dengan cara ketikkan pada CMD **git --version**.

![install-21]


# 02. Konfigurasi Git Pada Windows
Konfigurasi Git pada windows dapat dilakukan dengan cara mengetikkan kode dibawah ini di CMD
```
$ git config --global user.name "AnggitaAlbiantara
$ git config --global user.email anggita.albiantara@students.utdi.ac.id
```
![konfig1]

Selanjutnya jika kalian ingin melihat hasil dari konfigurasi diatas maka kalian dapat mengetikkan code berikut di CMD
```
$ git config --lists
```
![konfig2]

Langkah ini cukup dilakukan sekali saja, kecuali jika kalian ingin melakukan perubahan nama dan email.

# 03. Membuat Repo Di Account Git Kita
  Untuk membuat repo, dapat menggunakan langkah-langkah berikut :
  1.  Klik tanda **+** pada bagian atas dekat profil Git kita, klik lalu pilih **New repository**
  
  ![repo1]
  
  2. Isikan nama owner, nama repo, keterangan, pilih settingan repo untuk private atau public, checklist untuk add README.md atau tidak, serta pilih lisensi.

![repo2]
  
  3. Klik **Create Repository**
Setelah langkah-langkah tersebut berhasil, maaka repo akan dibuat dan bisa diakses dengan menggunakan pola ```https://github.com/username/reponame```. Pada repo tersebut, hanya akan muncul 1 file, yaitu LICENSE. Jika kalian memilih membuat README pada saat langkah ke 2, juga akan muncul README.md. Ada atau tidak ada README.md tidak mempunyai efek apapun pada langkah ini.

![repo3]
    
# Clone Repo
Untuk clone repo, kalian dapat mengetikkan code dibawah pada CMD,
```git clone https://github.com/UserName/NameRepository```
![clone]

Sehingga hasil clonenya akan masuk ke penyimpanan lokal komputer kita.
![clone2]

# Mengelola Repo
  Setelah melakukan ```clone``` ke penyimpanan lokal komputer kita, semua manipulasi konten dilakukan di komputer lokal dan hasilnya akan di-**push** ke remote repo di GitHub. Dengan demikian, jangan berganti-ganti remote lokal, sekali dibuat disitu, tetap berada disitu. Jika kehilangan repo lokal, maka harus melakukan clone ulang ke direktori yang bersih (kosong) setelah itu baru lakukan pengelolaan repo. Beberapa hal yang biasanya dilakukan seperti berikut ini.

# Mengubah Isi - Push Tanpa Branching dan Merging
Perubahan isi bisa terjadi karena satu atau kombinasi beberapa hal berikut:
  1. File dihapus
  2. File diedit
  3. Membuat file / direktori baru
  4. Menghapus direktori
Untuk kasus-kasus tersebut, lakukan perubahan di komputer lokal, setelah itu push ke repo. 
![isi]
![push]

# Mengubah Isi dengan Branching and Merging
Dengan menggunakan cara ini, setiap kali akan melakukan perubaham, perubahan itu dilakukan di komputer lokal dengan membuat suatu *cabang* yang nantinya digunakan untuk menampung perubahan-perubahan tersebut. Setelah itu, cabang itu yang akan dikirim ke repo GitHub untuk dimintai review kemudian digabungkan (```merge```) ke master. Secara umum, repo yang dibuat biasanya sudah mempunyai satu branch yang disebut dengan ```master```. Cara ini lebih aman, terstruktur, terkendali, dan mempunyai history yang lebih baik. Jika perubahan yang kita buat sudah terlalu kacau dan kita menyesal, maka ada cara untuk "membersihkan" repo lokal kita. Secara umum, langkahnya adalah sebagai berikut :

1. Buat branch untuk menampung perubahan-perubahan
2. Lakukan perubahan-perubahan
3. Add dan commit perubahan-perubahan tersebut ke branch
4. Kembali ke repo master
5. Buat pull request di GitHub
6. Merge pull request di GitHub
7. Merge branch untuk menampung perubahan-perubahan tersebut ke master.
8. Selesai.

![isi1]
![isi2]
![isi3]

Setelah itu, kirim *pull request (PR)*, selanjutnya kita bisa langsung merge :
![pull]

Setelah itu, ```Confirm Merge```, branch yang kita kirimkan tadi sudah dimasukkan ke branch ```master```. Setelah itu, merge di komputer lokal kita :

![merge1]

![merge2]

![merge3]

![merge4]

# Sinkronisasi

Suatu saat, bisa saja terjadi kita menggunakan komputer lain dan mengedit repo melalui repo lokal di komputer lain, setelah itu pindah ke kamputer lain lagi. Saat itu, kita perlu melakukan sinkronisasi ke kemputer lokal. Perintah untuk sinkronisasi adalah:

```
$ git pull
```
![pull5]

Perintah ini dikerjakan di direktori tempat repo lokal kita berada.

# Membatalkan Perubahan

Praktik yang baik adalah membuat *branch* pada saat kita akan melakukan perubahan-perubahan. Jika perubahan-perubahan yang kita lakukan sudah sedemikian kacaunya, maka kita bisa membuat supaya perubahan-perubahan yang kacau tersebut hilang dan kembali ke kondisi bersih seperti semula.

![batal1]

![batal2]

# Undo Commit Terakhir

Suatu saat, mungkin kita sudah terlanjur mem-*push* perubahan ke repo GitHub, setelah itu kita baru menyadari bahwa perubahan tersebut salah. Untuk itu, kita bisa melakukan ```git revert```.

![undo1]

![undo2]

Selanjutnya, tinggal di-*push* ke repo GitHub.

![undo3]

Untuk kembali ke perubahan pada saat yang sudah lama, yang perlu dilakukan adalah melakukan ```git revert <posisi>``` kemudian mengedit secara manual kemudian push ke repo.

![revert1]

![revert2]

![revert3]

![revert4]
