# teo-service
Pembangunan aplikasi web dari proyek KESAT ITB dalam bahasa Golang. Website ini juga disertai API interface.


# End Point

| Kode | Method | URL | Input | Output | Deskripsi |
| --- | --- | --- | --- | --- | --- |
| REST-1 | GET | / | GetRequest | Html page | Menampilkan halaman utama tergantung session |
| REST-2 | GET | /register | GetRequest | Html page | Menampilkan halaman register |
| REST-3 | POST | /register | Username stringPassword stringEmail string | Create user Record | Membuat user baru |
| REST-4 | GET | /login | GetRequest | Html page | Menampilkan halaman login |
| REST-5 | POST | /login | Username StringPassword string | Session Set | Mengeset session apabila berhasil meredirect kembali apabila gagal |
| REST-6 | GET | /report | GetRequest | Html page | Menampilkan halaman laporan |
| REST-7 | POST | /report | Idreporter intLink\_photo stringDetail stringCategory int | CREATE report | Menyimpan laporan baru pada database |
| REST-8 | DELETE | /report?id={{id}} | Bind id int | DELETE report | Menghapus report pada database |
| REST-9 | GET | /report?id={{id}} | Bind id int | Html page | Mendapatkan halaman report yang diinginkan |
| REST-10 | GET | /timeline | GetRequest | Html page | Mendapatkan seluruh lini masa |
| REST-11 | GET | /api/label?id={{id}} | Bind id int | JSON | Mendapatkan json label id tertentu |
| REST-12 | GET | /api/label | GetRequest | JSON | Mendapatkan json semua label |
| REST-13 | GET | /api/report?id={{id}} | Bind id int | JSON | Mendapatkan json report sesuai id |
| REST-14 | GET | /api/report | GetRequest | JSON | Mendapatkan json semua report yang ada |
