# Go-Tiket
 
## Description
Merupakan aplikasi backend dengan menggunakan bahasa pemerograman Golang yang berfungsi untuk melakukan pembelian tiket kereta api. Aplikasi ini dibuat dengan menggunakan konsep Clean Architecture dan menggunakan database PostgreSQL sebagai penyimpanan data.

## How to Run
1. Clone repository ini

```bash
git clone https://github.com/rezapace/go-tiket
```

2. Masuk ke dalam direktori repository

```bash
cd go-tiket
```

3. Cek file env dan sesuaikan dengan konfigurasi database yang digunakan

4. Jalankan perintah untuk membuat database dan lakukan migrate
    
```bash
go run cmd/migrate/main.go
```

5. Jalankan aplikasi

```bash
go mod tidy
go run cmd/server/main.go
```

## Detail Proyek

**Tema Proyek:** Aplikasi Ticketing Berbasis Web dengan nama "Depublic."

**Deskripsi:** Platform ini digunakan untuk jual-beli tiket konser dan event. Pengguna dapat mendaftar sebagai pembeli dan mencari jadwal konser sesuai kebutuhan mereka, serta membayar tiket secara online. Platform juga menyediakan informasi akurat tentang event yang sedang berlangsung.

**Fitur Wajib yang Harus Diimplementasikan:**

- Registrasi User.
- Implementasi In-App Notification (notifikasi dalam aplikasi, bukan push notification).
- Profil User.
- Histori Transaksi.
- Pencarian dan Filter.

**Teknologi yang Digunakan:**

- Bahasa pemrograman: Golang.
- Penggunaan bahasa pemrograman lain atau library pihak ketiga tidak diperbolehkan.

**Arsitektur:**

- Menggunakan arsitektur dasar Model-View-Controller (MVC) dengan dua lapisan (backend dan frontend).
- Implementasi fitur pencarian, pembagian halaman, dan penyortiran.

**Peran yang Tersedia:**

- Buyer (pembeli).
- Admin.

**Pengembangan dan Deployment:**

- Pengembangan menggunakan Vercel untuk tahap staging dan produksi.
- Git workflow dengan dua cabang utama: "master" untuk produksi dan "develop/staging" untuk pengujian.
- Deliverables meliputi API yang dapat digunakan oleh Front End, dokumentasi Swagger, dan repository proyek.


## API Documentation
API Documentation dapat diakses pada file
output\Backend Ticketing.postman_collection.json
```bash
https://github.com/rezapace/Go-Tiket/blob/main/output/Backend%20Ticketing.postman_collection.json
```

<table>
  <tr>
    <td><img src="https://github.com/rezapace/Go-Tiket/blob/main/Materi/flow%201.png?raw=true" alt="Figma 1"></td>
    <td><img src="https://github.com/rezapace/Go-Tiket/blob/main/Materi/flow%202.png?raw=true" alt="Figma 2"></td>
  </tr>
</table>
<table>
  <tr>
    <td><img src="https://github.com/rezapace/Go-Tiket/blob/main/Materi/postman.jpg?raw=true" alt="Figma 1"></td>
    <td><img src="https://github.com/rezapace/Go-Tiket/blob/main/Materi/postman.jpg?raw=true" alt="Figma 2"></td>
  </tr>
</table>
