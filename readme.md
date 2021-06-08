# POST COLLECTION
```bash
    #you can import post collection with this link
    https://www.getpostman.com/collections/ff501dfdf0a0997f939a

```
# Cara Menggunakan unit testing
```bash
    - Jalankan Terlebih dahulu main.go
    - Masuk kedalam folder test 
    - kemudian Ketikan perintah berikut :
        - go test -v -run=TestTambahBerat
        - go test -v -run=TestUpdateBerat
        - go test -v -run=TestDetailBerat
        - go test -v -run=TestIndexBerat
        - go test -v -run=TestHapusBerat

```


# Api Spesification

## Tambah Berat

Request :
- Method : POST
- Endpoint : `/tambah_berat`
- Body :
    Form-Data
```json 
{
    params[tanggal]:2019-08-06
    params[max_berat]:50
    params[min_berat]:48
}
```

Response :

```json 
{
    {
    "message": "Berhasil Disimpan",
    "response": {
        "Tanggal": "2019-08-06",
        "Max": 50,
        "Min": 48,
        "Perbedaan": 2
    },
    "status": 200
    }
}
```

## Update Berat

Request :
- Method : POST
- Endpoint : `/update_berat`
- Body :
    Form-Data
```json 
{
    params[tanggal]:2019-08-06
    params[max_berat]:55
    params[min_berat]:48
}
```

Response :

```json 
{
    {
    "message": "Berhasil Di ubah",
    "response": {
        "Tanggal": "2019-08-05",
        "Max": 55,
        "Min": 48,
        "Perbedaan": 7
    },
    "status": 200
    }
}
```

## Detail Berat

Request :
- Method : POST
- Endpoint : `/detail_berat`
- Body :
    Form-Data
```json 
{
    params[tanggal]:2019-08-06
}
```

Response :

```json 
{
    {
    "response": [
        {
            "Tanggal": "2019-08-06",
            "Max": 50,
            "Min": 48,
            "Perbedaan": 2
        }
    ],
    "status": 200
    }
}
```

## Index Berat

Request :
- Method : POST
- Endpoint : `/index_berat`
- Body :

Response :

```json 
{
    {
    "response": {
        "BeratBadan": [
            {
                "Tanggal": "2019-08-23",
                "Max": 50,
                "Min": 49,
                "Perbedaan": 1
            },
            {
                "Tanggal": "2019-08-06",
                "Max": 50,
                "Min": 48,
                "Perbedaan": 2
            },
            {
                "Tanggal": "2019-08-05",
                "Max": 55,
                "Min": 48,
                "Perbedaan": 7
            }
        ],
        "RataRataBerat": {
            "Max": 51,
            "Min": 48,
            "Perbedaan": 3
        }
    },
    "status": 200
    }
}
```

## Hapus Berat

Request :
- Method : POST
- Endpoint : `/hapus_berat`
- Body :
    Form-Data
```json 
{
    params[tanggal]:2019-08-06
}
```

Response :

```json 
{
    {
    "message": "Berhasil Di hapus",
    "response": {
        "Tanggal": "2019-08-06"
    },
    "status": 200
    }
}
```