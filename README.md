practice:
1. Buat endpoint:
    a. GET /student -> untuk ambil semua data student
        - Response berbentuk array
    b. GET /student/:id -> untuk ambil data student berdasarkan id 
        - Response json

    c. POST /student -> untuk menambahkan 1 data student
        Gunakan body param dengan contoh format seperti ini:
        {
            "first_name": "John",
            "last_name": "Doe"
        }
        
        Dengan response:
        {
            "status": "success",
        }

    d. PUT /student -> untuk update 1 data student
        Gunakan body param dengan contoh format seperti ini:
        {
            "id": 1,
            "first_name": "John",
            "last_name": "Doe"
        }

        Dengan response:
        {
            "status": "success",
        }
    e. DELETE /student -> untuk delete data student berdasarkan id
        Format response:
        {
            "status": "success",
        }
2. Untuk datanya, silahkan buka file students.json
3. Untuk practice ini, kita tidak menggunakan database untuk save data. Hanya menggunakan json file yang terlampir
4. untuk mengambil & mengolah data, bisa dilihat contohnya pada file main.go