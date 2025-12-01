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

## Developer Guide — Setup lengkap & menjalankan aplikasi

Dokumentasi ini menjelaskan semua yang diperlukan untuk menjalankan aplikasi secara lokal (development) maupun production-like. Ikuti langkah di bawah secara berurutan.

### 1) Prasyarat

- Go 1.20+ terinstall (https://go.dev/dl/)
- PostgreSQL (atau DB sesuai `DATABASE_URL`) — versi modern (12+ recommended)
- Git

Opsional (untuk auto-reload saat develop):

- `air` (rekomendasi) atau `CompileDaemon` / `reflex` jika Anda ingin auto-restart saat kode berubah.

Contoh instalasi (macOS Homebrew):

```bash
brew install go postgresql
go install github.com/air-verse/air@latest   # jika ingin air (module path baru)
# atau: go install github.com/githubnemo/CompileDaemon@latest
```

### 2) Dapatkan kode dan persiapan awal

1. Clone repo dan pindah ke folder proyek:
```bash
git clone git@github.com:career-rizkiramdhani/learn-go.git
cd learn-go
```
2. Download dependency dan siapkan module:
```bash
go mod tidy
```

3. Buat file environment: salin `.env.example` menjadi `.env` dan ubah nilainya sesuai environment Anda.
```bash
cp .env.example .env
# lalu edit `.env` sesuai kebutuhan (DB, JWT secret, PORT dsb)
```

Keterangan variabel penting di `.env`:

- `DATABASE_URL` — DSN Postgres lengkap (contoh: `host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable`). Jika tidak diset, app akan fallback ke komponen `POSTGRES_*`.
- `JWT_SECRET` — secret untuk menandatangani JWT.
- `JWT_EXPIRED` — masa berlaku token dalam jam (default: 72).
- `PORT` — port aplikasi (default project mungkin 8080/8085).
- `USE_TLS` — `true`/`false`. Jika `true`, aplikasi akan memastikan `CERT_FILE` & `KEY_FILE` ada (dapat dibuat self-signed secara otomatis).

### 3) Database: migrasi & seeder

Project menyediakan helper CLI untuk migrasi dan seeding via `main.go`.

- Jalankan migrate (create tables):
```bash
go run main.go migrate
```

- Jalankan migrate fresh (rollback -> migrate -> seed):
```bash
go run main.go migrate:fresh
```

- Hapus (rollback semua tables yang dikelola migration):
```bash
go run main.go rollback
```

Catatan: seeder membuat satu admin default kecuali sudah ada user.

### 4) Menjalankan aplikasi — Development (dengan auto-reload)

Rekomendasi: gunakan `air` (auto-rebuild & restart). Repo sudah berisi konfigurasi `.air.toml` dan `Makefile` dengan target `dev`.

Instal `air` jika belum:
```bash
go install github.com/air-verse/air@latest
# pastikan $GOPATH/bin ada di PATH
```

Kemudian jalankan:
```bash
cd /path/to/learn-go
make dev
```

atau langsung:
```bash
air
```

Perilaku: `air` akan membangun binary ke `tmp/main` lalu menjalankan; saat file `.go` diubah, `air` otomatis rebuild dan restart server.

Alternatif ringan: gunakan `CompileDaemon`:
```bash
CompileDaemon -command='go run main.go'
```

### 5) Menjalankan aplikasi — Production / run once

Build binary dan jalankan:
```bash
go build -o app
./app
```

Atau jalankan langsung (tidak direkomendasikan untuk production):
```bash
go run main.go
```

Jika ingin menjalankan di background (nohup):
```bash
nohup go run main.go > server.log 2>&1 & echo $! > server.pid
```

Lihat logs:
```bash
tail -f server.log
```

Berhentikan proses:
```bash
kill $(cat server.pid) && rm server.pid
```

### 6) Testing endpoints (curl)

- Sign-in (POST):
```bash
curl -X POST http://localhost:8080/api/signin \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'
```

- List users (GET) dengan pagination (page & size optional):
```bash
curl "http://localhost:8080/api/users?page=1&size=10"
```

- Create user (POST):
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"username":"newuser","password":"secret123","email":"new@example.com"}'
```

### 7) Response format (auth)

Sign-in sekarang mengembalikan response terstruktur seperti ini:

```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "result": {
      "user": { "id": 123, "username": "johndoe", "email": "john@example.com" },
      "token": "<jwt-token>"
    }
  }
}
```

### 8) Best practices & troubleshooting

- Jangan commit file sensitif (mis. `.env`) ke repo. Gunakan `.env.example` sebagai template.
- `server.pid` tidak seharusnya dicommit pada umumnya — jika Anda ingin saya hapus `server.pid` dari repo dan tambahkan ke `.gitignore`, beri tahu.
- Jika build error terkait dependensi, jalankan `go mod tidy` dan pastikan Go version sesuai.
- Jika menggunakan TLS dan self-signed cert, di Postman matikan `SSL certificate verification` atau tambahkan cert ke trust store.

### 9) CI / Deployment notes

- Untuk CI: jalankan `go test ./...` (tambahkan unit tests), `go vet`, dan `go build` sebagai pipeline steps.
- Untuk deployment: bangun binary (`go build -o app`), jalankan dengan process manager (`systemd`, `supervisord`, containerize with Docker).

---

Jika Anda ingin, saya bisa:
- Menambahkan `docker-compose.yml` untuk Postgres + app (dev compose),
- Menambahkan contoh `systemd` service file untuk menjalankan binary di server, atau
- Membuat script `scripts/dev.sh` yang meng-setup environment dan menjalankan `air`.

Beritahu saya langkah mana yang mau saya lanjutkan.
