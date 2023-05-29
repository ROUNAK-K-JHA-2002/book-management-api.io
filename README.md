
# Book Management API using Golang 



![App Screenshot](https://ondemand.bannerbear.com/signedurl/9K5qxXae32jEAGRDkj/image.jpg?modifications=W3sibmFtZSI6InJlcG8iLCJ0ZXh0IjoiUk9VTkFLLUstSkhBLTIwMDIgLyAqYm9vay1tYW5hZ2VtZW50LWFwaS5pbyoifSx7Im5hbWUiOiJkZXNjIiwidGV4dCI6IlRoaXMgaXMgYSBiZWdpbm5lci1mcmllbmRseSBDUlVEIEFQSSBmb3IgbWFuYWdpbmcgYm9va3MsIGRldmVsb3BlZCB1c2luZyBHb2xhbmcgd2l0aCBNeVNRTCBhcyB0aGUgRGF0YWJhc2UuIEl0IGVmZmljaWVudGx5IGhhbmRsZXMgQVBJIHJvdXRlcyBmb3IgR0VULCBQT1NULCBVUERBVEUsIGFuZCBERUxFVEUgb3BlcmF0aW9ucy4ifSx7Im5hbWUiOiJhdmF0YXI1IiwiaGlkZSI6dHJ1ZX0seyJuYW1lIjoiYXZhdGFyNCIsImhpZGUiOnRydWV9LHsibmFtZSI6ImF2YXRhcjMiLCJoaWRlIjp0cnVlfSx7Im5hbWUiOiJhdmF0YXIyIiwiaGlkZSI6dHJ1ZX0seyJuYW1lIjoiYXZhdGFyMSIsImltYWdlX3VybCI6Imh0dHBzOi8vYXZhdGFycy5naXRodWJ1c2VyY29udGVudC5jb20vdS85ODE0NTIzMz92PTQifSx7Im5hbWUiOiJjb250cmlidXRvcnMiLCJ0ZXh0IjoiUk9VTkFLLUstSkhBLTIwMDIifSx7Im5hbWUiOiJzdGFycyIsInRleHQiOiIxIn1d&s=e2e0e4b3ac1ac1b4f04922088a4f584bd711a9a42128f4d9871f4ffb66f00540)







# 💻 Tech Stack:
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
## API Reference

#### --Get all Books
```http
  GET /book/
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `nil` | `nil` | nil |


#### --Add a book to DB
```http
  POST /book/
```
Body to be sent in raw JSON format
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `Name`      | `string` | **Required**. Name of book to be stored |
| `Author`      | `string` | **Required**. Author of book to be stored |
| `Publication`      | `string` | **Required**. Publication of book to be stored |


#### --Get a specific book
```http
  GET /book/${bookId}/
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |


#### --Update a specific book
```http
  PUT /book/${bookId}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to Update |

Body to be sent in raw JSON format
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `Name`      | `string` | **Required**. Name of book to be stored |
| `Author`      | `string` | **Required**. Author of book to be stored |
| `Publication`      | `string` | **Required**. Publication of book to be stored |


#### --Delete a specific book
```http
  DELETE /book/${bookId}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to delete |


