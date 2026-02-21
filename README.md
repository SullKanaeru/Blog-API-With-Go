# Akses link ini untuk testing api yang sudah dipublic: https://drive.google.com/file/d/1V_U_F3XLIKpFLXkGr--COAxLLOHVkuZ-/view?usp=sharing

# 🚀 Golang Blog API Server

RESTful API untuk platform blogging yang dibangun menggunakan **Go (Golang)**, **Fiber framework**, dan **PostgreSQL**. Proyek ini dilengkapi dengan sistem autentikasi berbasis *Cookie*, *Role-Based Access Control* (RBAC), serta keamanan level *Resource Ownership*.

## ✨ Fitur Utama
* **Keamanan JWT Berbasis Cookie:** Token JWT disimpan dengan aman di *HTTP-Only Cookie*, melindunginya dari serangan XSS.
* **Role-Based Access Control (RBAC):** Memisahkan hak akses antara `author` (penulis) dan `viewer` (pembaca).
* **Resource Ownership:** Sistem cerdas yang memastikan seorang `author` hanya dapat mengedit atau menghapus artikel miliknya sendiri.
* **Auto-Slug Generation:** Judul artikel secara otomatis diubah menjadi URL *slug* yang ramah SEO.
* **Relasi Database Lengkap:** Memanfaatkan GORM `.Preload()` untuk menyajikan data artikel beserta profil penulisnya dalam satu *response* JSON yang bersarang (*nested*).

## 🛠️ Teknologi yang Digunakan
* **Bahasa:** Go (Golang)
* **Web Framework:** Go Fiber (v2)
* **ORM:** GORM
* **Database:** PostgreSQL (Neon.tech)
* **Keamanan:** Bcrypt (Hashing Password) & Golang-JWT
* **Migrasi:** Golang-Migrate

---

## ⚙️ Cara Instalasi & Menjalankan Server

1. **Clone repositori ini:**
   ```bash
   git clone <url-repo-kamu>
   cd blog_api

```

2. **Install semua *dependencies*:**
```bash
go mod tidy

```


3. **Siapkan *Environment Variables*:**
Buat file bernama `.env` di *root* folder dan isi dengan konfigurasi berikut:
```env
PORT=3000
DATABASE_URL=postgres://<user>:<password>@<host>/<dbname>?sslmode=require
JWT_SECRET=rahasia_super_kuat_ganti_ini

```


4. **Jalankan Server (Migrasi akan berjalan otomatis):**
```bash
go run cmd/api/main.go

```



---

## 📡 Dokumentasi Endpoint API

Berikut adalah daftar rute yang tersedia untuk dihubungkan ke antarmuka aplikasi.

### 👤 Manajemen User & Autentikasi

| Method | Endpoint | Keterangan | Akses |
| --- | --- | --- | --- |
| `POST` | `/api/users/register` | Mendaftar akun baru (Bisa kirim `"role": "author"` atau `"viewer"`) | Publik |
| `POST` | `/api/users/login` | Login dan mendapatkan HTTP-Only Cookie JWT | Publik |
| `POST` | `/api/users/logout` | Menghapus Cookie sesi saat ini | Butuh Login |
| `GET` | `/api/users` | Mengambil semua user (Bisa filter via *query* `?role=author`) | Publik |
| `GET` | `/api/users/me` | Mengambil detail profil akun yang sedang login | Butuh Login |

### 📝 Manajemen Artikel (Blog Posts)

| Method | Endpoint | Keterangan | Akses |
| --- | --- | --- | --- |
| `GET` | `/api/posts` | Mengambil seluruh artikel beserta profil penulisnya | Publik |
| `GET` | `/api/posts/:slug` | Mengambil satu artikel secara spesifik berdasarkan *slug* | Publik |
| `POST` | `/api/posts` | Membuat artikel baru | Khusus `author` |
| `PUT` | `/api/posts/:id` | Mengedit artikel | Khusus Pemilik |
| `DELETE` | `/api/posts/:id` | Menghapus artikel | Khusus Pemilik |

---
