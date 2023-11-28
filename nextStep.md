# Membuat Sturktur Data Event
1. Buat file event.go di dalam folder entity
2. Dalam file tersebut, buatlah struktur data (struct) untuk entitas event dengan definisi kolom-kolom yang dibutuhkan

# Membuat Repository Event
1. Buat file event.repository.go di dalam folder repository
2. Dalam file tersebut, buat interface untuk mengakses data Event.

# Membuat Service Event
1. Buat file event.service.go di dalam folder service
2. Dalam file tersebut, buat logikan terkait event dengan memanggil fungsi eventRepository

# Membuat Handler Event
1. Buat file event.handler.go di dalam folder handler
2. Dalam file tersebut, membuat logika terkait event dengan memanggil fungi event service

# Membuat Routing
1. Tambahkan endpoint handler yang telah dibuat pada private route (mengapa? karena agar user bisa login/ register terlebih dahulu dengan melewati middleware)

# Menambahkan Logika ke Builder
1. Buka file builder.go di folder builder
2. panggil Use Case, repository, dan objek lain yang diperlukan di dalam file tersebut

# Menjalankan Server
1. Jalankan server menggunakan perintah go run cmd/server/main.go (cek kembali dimana direktori anda berada agar tidak terjadi kesalahan dalam running)

# Cara menjalankan file
1. Git clone https://github.com/Capstone-Project-Mikti-Group-2/Depublic-BE.git
2. cd Depublic-BE
3. buat branch baru dengan nama kalian exp: (git checkout nama)
4. coba edit file readme.md exp: (edit by nama kalian)
5. git add .
6. git commit -m “edit readme by nama kalian”
7. go mod tidy 
8. (ketik dalam terminal) migrate -database postgres://postgres/password@localhost:5432/depublic-db?sslmode=disable -path db/migrations up
9. go run cmd/server/main.go
