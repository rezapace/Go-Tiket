# Membuat Struktur Data Tiket
1. Buat file tiket.go di dalam folder entity.
2. Dalam file tersebut, buatlah struktur data (struct) untuk entitas tiket dengan mendefinisikan kolom-kolom yang dibutuhkan.

# Membuat Repository Tiket
1. Buat file tiket_repository.go di dalam folder repository.
2. Dalam file tersebut, buatlah interface untuk mengakses data tiket. Panggil fungsi-fungsi yang dibutuhkan dari tiket_repository.go.

# Membuat Service Tiket
1. Buat file tiket_service.go di dalam folder service.
2. Dalam file tersebut, buatlah struktur service untuk logika terkait tiket. Panggil fungsi-fungsi yang dibutuhkan dari tiket_repository.go.

# Membuat Handler Tiket
1. Buat file tiket_handler.go di dalam folder handler.
2. Dalam file tersebut, buatlah handler untuk menerima request dan menangani logika terkait tiket. Panggil fungsi-fungsi dari tiket_service.go dan sesuaikan dengan kebutuhan.

# Logika Builder
1. Tentukan apakah entitas tiket akan diatur sebagai private atau public. Buat file builderPrivieRoute.go jika private.
2. Panggil use case, repository, dan objek lain yang diperlukan di dalam file tersebut.

# Membuat Routing
1. Buka file PrivateRoute atau file yang sesuai dengan routing di aplikasi.
2. Tambahkan endpoint untuk operasi yang ingin Anda dukung menggunakan handler tiket yang sudah dibuat.

# Menjalankan Server
1. Jalankan server menggunakan perintah go run cmd/server main.go.
2. Pastikan file PrivateRoute atau file routing yang sesuai di-load dan endpoint tiket sudah terdaftar.

Dengan langkah-langkah tersebut, Anda akan memiliki struktur dasar untuk menangani entitas tiket dalam aplikasi Anda. Pastikan untuk menggantinya sesuai kebutuhan dan menyesuaikan setiap langkah dengan logika dan fitur aplikasi yang Anda buat