API QUiZ

Link URL Public : https://quiz-production-4fe7.up.railway.app/
Link Repo : https://github.com/yogiansyahhdygm/quiz.git

Middleware Basic Auth untuk mengamankan endpoint.
Untuk Basic Auth yang dipakai :
- Username : admin
- Password : admin
*note : semua endpoint wajib sudah authorization

Endpoint API :
1. Kategori :
  - Menampilkan semua kategori            : https://quiz-production-4fe7.up.railway.app/api/categories            GET
  - Menampilkan kategori berdasarkan ID   : https://quiz-production-4fe7.up.railway.app/api/categories/:id        GET
  - Menambahkan kategori baru             : https://quiz-production-4fe7.up.railway.app/api/categories            POST
  - Mengubah kategori                     : https://quiz-production-4fe7.up.railway.app/api/categories/:id        PUT
  - Menghapus kategori by id              : https://quiz-production-4fe7.up.railway.app/api/categories/:id        DELETE

2. Buku :
  - Menampilkan semua buku                : https://quiz-production-4fe7.up.railway.app/api/books                 GET
  - Menampilkan detail buku berdasarkan ID: https://quiz-production-4fe7.up.railway.app/api/books:id              GET
  - Menampilkan buku berdasarkan kategori : https://quiz-production-4fe7.up.railway.app/api/categories/:id/books  GET
  - Menambahkan buku baru                 : https://quiz-production-4fe7.up.railway.app/api/books                 POST
  - Mengubah buku                         : https://quiz-production-4fe7.up.railway.app/api/books:id              PUT
  - Menghapus buku by id                  : https://quiz-production-4fe7.up.railway.app/api/books:id              DELETE


Contoh Postman :
a. POST Buku (otomatis konversi thicknes) :
{
  "title": "Belajar Golang",
  "description": "Panduan lengkap Golang",
  "image_url": "https://example.com/golang.jpg",
  "release_year": 2024,
  "price": 75000,
  "total_page": 120,
  "category_id": 1,
  "created_by": "admin"
} 

b. POST Kategori :
{
  "name": "Magic"
}

c. PUT buku: sama dengan contoh post buku
d. PUT kategori : sama dengan contoh post kategori


Terima kasih
