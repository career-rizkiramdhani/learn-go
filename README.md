# learn-go

Proyek contoh API Go (Echo + GORM) dengan autentikasi JWT, migrasi, seeder, dan utilitas TLS.

## Ringkasan
- Bahasa: Go
- Framework web: Echo
- ORM: GORM (Postgres)
- Autentikasi: JWT (HS256)

Project ini menyediakan beberapa perintah custom mirip Laravel artisan:

- `go run main.go` : jalankan server (HTTP/HTTPS sesuai konfigurasi `.env`).
- `go run main.go migrate` : jalankan migrasi (GORM AutoMigrate).
- `go run main.go rollback` : hapus tabel yang dikelola migrasi (rollback semua tabel yang terdaftar).
- `go run main.go migrate:fresh` : alias `rollback` lalu `migrate` lalu `seeder` (reset database + seed).
- `go build` / `go build -o app.exe` : build binary aplikasi.

> Catatan: perintah-perintah di atas dieksekusi dari root project.

## Struktur penting

- `main.go` — entry point; juga menangani perintah CLI (migrate, rollback, migrate:fresh).
- `config/` — konfigurasi aplikasi dan inisialisasi DB.
- `models/` — definisi model GORM (User dll.).
- `migration/` — migrasi helper `RunMigrations()` dan `RollbackAll()`.
- `seeder/` — seed data awal (admin user).
- `service/` — layanan bisnis, termasuk `auth_service.go` yang membuat JWT.
- `router/` — definisi route API (semua berada di `/api` group).
- `tlsutil/` — utilitas untuk membuat self-signed certificate (jika diperlukan).

## .env / konfigurasi
Salin contoh file `.env.example` menjadi `.env` lalu sesuaikan nilai-nilainya.

Variabel penting:
- `DATABASE_URL` — DSN Postgres (contoh: `host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable`).
  - Jika tidak diset, aplikasi akan memakai `POSTGRES_HOST`, `POSTGRES_PORT`, `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`, `POSTGRES_SSLMODE` untuk membangun DSN.
- `JWT_SECRET` — secret untuk menandatangani JWT.
- `JWT_EXPIRED` — masa berlaku token (dalam jam). Default: `72`.
- `PORT` — port aplikasi (default: `8085` pada setup ini).
- `USE_TLS` — `true`|`false`. Jika `true`, server akan memakai TLS. Aplikasi akan otomatis menghasilkan `cert.pem` dan `key.pem` jika tidak ada (local dev).
- `CERT_FILE`, `KEY_FILE` — path ke file sertifikat dan key (default: `cert.pem`, `key.pem`).

Contoh minimal `.env`:

```powershell
# DATABASE_URL atau komponen Postgres
DATABASE_URL=
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=dbgotest
POSTGRES_SSLMODE=disable

JWT_SECRET=supersecretkey
JWT_EXPIRED=72
PORT=8085
USE_TLS=true
CERT_FILE=cert.pem
KEY_FILE=key.pem
```

> Jika `USE_TLS=true`, server akan menghasilkan `cert.pem`/`key.pem` secara otomatis menggunakan `tlsutil.EnsureCert` (self-signed) saat startup.

## Setup local (langkah demi langkah)

1. Pastikan Go (>=1.20+) terinstall.
2. Siapkan Postgres dan buat database/user sesuai `.env`.
3. Salin `.env.example` menjadi `.env` dan isi nilai yang sesuai.
4. Install dependencies & build (opsional):

```powershell
go mod tidy
go build ./...
```

5. Untuk membuat schema & seed data (migrate fresh):

```powershell
go run main.go migrate:fresh
```

6. Jalankan server:

```powershell
go run main.go
```

Server akan berjalan pada `http://localhost:<PORT>` (atau `https://` jika `USE_TLS=true`).

## Testing login (Postman / curl)
- Endpoint sign-in: `POST /api/signin`
- Body JSON:
```json
{
  "username": "admin",
  "password": "password"
}
```

Jika `USE_TLS=true` dan Anda menggunakan self-signed cert, di Postman matikan "SSL certificate verification" atau impor `cert.pem` ke trust store.

## Rollback / migrate:fresh
- Rollback: menghapus tabel yang didefinisikan di `migration.RollbackAll()`.
  ```powershell
  go run main.go rollback
  ```
- Migrate fresh (rollback + migrate + seed):
  ```powershell
  go run main.go migrate:fresh
  ```

Catatan penting: rollback menghapus data. Gunakan dengan hati-hati.

## Menambahkan custom commands
`main.go` sudah membaca `os.Args[1]` untuk perintah sederhana. Jika Anda ingin menambah perintah baru:

1. Tambahkan case baru pada bagian CLI di `main.go`.
2. Implementasikan logika di package yang sesuai (mis. `migration`, `seeder`, `tools`).

Contoh menambah perintah `export:users`:

```go
case "export:users":
    // panggil fungsi export di package tools
    tools.ExportUsersCSV()
    return
```

## Safety & Next steps
- Untuk environment production, jangan gunakan self-signed cert. Gunakan cert resmi (Let's Encrypt / CA) dan set `USE_TLS=true` dengan file cert/key yang sah.
- Pertimbangkan menambah konfirmasi interaktif sebelum menjalankan `rollback` di environment non-dev.
- Jika ingin, saya bisa menambahkan skrip PowerShell di `scripts/` untuk mempermudah perintah biasa (`start`, `migrate_fresh`, `rollback`).

---

Jika Anda mau, saya bisa: menambahkan skrip `scripts/migrate_fresh.ps1` & `scripts/rollback.ps1`, atau menambahkan konfirmasi interaktif sebelum rollback. Pilih yang Anda inginkan.
