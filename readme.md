# betestmaxChat

Sebelum menjalankan run

```
go mod tidy
```

untuk menginstall dependecy yang di pakai

1: Struktur data terdapat di domain/data.go dan file dto berfungsi untuk struktru response data dan data itu sendiri

2: Untuk semua data terdapt di file domain/data.go dalam bentuk slice dari struct Data

3: CRUD data
```
A:Menampilakn semua data route GET:http://127.0.0.1:8080/data
```
```
B:Mengambil satu data menggunakan route GET:http://127.0.0.1:8080/data/tfx
```
```
C:Menambahkan data baru menggunakan Route POST:http://127.0.0.1:8080/data
```
request Body :
```
{
  "code": "P001",
  "name": "Smartphone",
  "description": "High-performance smartphone with 128GB storage.",
  "model": "X123",
  "tech": "5G",
  "status": "available"
}
```
```
D:MengUpdate Data menggunakan route PUT:http://127.0.0.1:8080/data/P001
```
Request Body:
```
{
  "code": "P001",
  "name": "Smartphone",
  "description": "High-performance smartphone with 128GB storage.",
  "model": "X123",
  "tech": "5G",
  "status": "available"
}
```
```
E:Delete Data mengguanakn route DELETE:http://127.0.0.1:8080/data/P001
```

4 Filter data beradasarkan data model dan tech
```
A: Filter Berdasarkan Model route GET:http://localhost:8080/filter?model=Humanoid
```
```
B: Filter Berdasarkan Tech route GET:http://localhost:8080/filter?tech=Robot
```

5 Endpoint Filter untuk Multiple Tech:http://localhost:8080/filter?tech=Robot,Cyborg