# **MINI PROJECT**

## API usage documentation.
#
# **Gobal Response**

### This response are applied on all **POST**, **PUT**, **DELETE** API

* **Response (401 - Unauthorized) :**
#### **IF EITHER OF THE USERNAME OR PASSWORD FOR BASIC AUTH IS WRONG**  
```json
{
    "result": "Username atau Password salah"
}
```
#### **IF EITHER OF THE USERNAME OR PASSWORD FOR BASIC AUTH IS EMPTY**
```json
{
    "result": "Username atau Password tidak boleh kosong"
}
```

#### **IF THERE IS NO BASIC AUTH USED** 
```json
{
    "result": "Unauthorized"
}
```

# **Category**
### **1.GET /categories**

#### Get list of all exist categories


* **Request Header :**
```json
NONE
```

* **Request Body :**
```json
NONE
```

* **Response (200 - OK) :**
#### *IF EXIST*
```json
{
    "result": [
        {
            "id": 4,
            "name": "SciFi",
            "created_at": "2023-03-26T23:05:26Z",
            "updated_at": "2023-03-26T23:05:26Z"
        },
        {
            "id": 5,
            "name": "Fantasy",
            "created_at": "2023-03-26T23:12:42Z",
            "updated_at": "2023-03-26T23:12:42Z"
        },
        {
            "id": 6,
            "name": "Horror",
            "created_at": "2023-03-27T04:28:10Z",
            "updated_at": "2023-03-27T04:28:10Z"
        },
        {
            "id": 7,
            "name": "Crime",
            "created_at": "2023-03-27T04:44:16Z",
            "updated_at": "2023-03-27T04:44:16Z"
        },
        {
            "id": 8,
            "name": "Mystery",
            "created_at": "2023-03-27T05:03:52Z",
            "updated_at": "2023-03-27T05:03:52Z"
        },
        {
            "id": 9,
            "name": "Comedy",
            "created_at": "2023-03-27T05:04:09Z",
            "updated_at": "2023-03-27T05:04:09Z"
        }
    ]
}
```
#### *IF THERE IS NO DATA EXIST*
```json
{
    "result": null
}
```

### **2.POST /categories**

#### Create New Categories


* **Authorization:**
```bash
Basic Auth
```

* **Request Body :**
```json
{
    "Name": "<category name>"
}
```

* **Response (201 - Created) :**
```json
{
    "result": "Success Insert Category"
}
```
### **3.PUT /categories/:id <category.ID>**

#### Edit Categories


* **Authorization:**
```bash
Basic Auth
```

* **Request Body :**
```json
{
    "Name": "<category name>"
}
```

* **Response (200 - OK) :**
```json
{
    "result": "Success Update Category"
}
```
### **4.DELETE /categories/:id <category.ID>**

#### DELETE Categories


* **Authorization:**
```bash
Basic Auth
```

* **Request Body :**
```json
NONE
```

* **Response (200 - OK) :**
```json
{
    "result": "Success Delete Category"
}
```
### **3.GET /categories/:id<category.ID>/books**

#### Get Books by categories ID


* **Authorization:**
```bash
NONE
```

* **Request Body :**
```json
NONE
```

* **Response (200 - OK) :**
#### *IF EXIST*
```json
{
    "result": [
        {
            "id": 2,
            "title": "Naruto",
            "description": "Komik",
            "image_url": "https://lol.com",
            "release_year": 1980,
            "price": "20000",
            "total_page": 201,
            "thickness": "tebal",
            "created_at": "2023-03-27T00:38:35Z",
            "updated_at": "2023-03-27T05:09:40Z",
            "category_id": 5
        },
        {
            "id": 4,
            "title": "Twilight",
            "description": "NOVEL",
            "image_url": "https://lol.com",
            "release_year": 1980,
            "price": "20000",
            "total_page": 200,
            "thickness": "sedang",
            "created_at": "2023-03-27T05:07:58Z",
            "updated_at": "2023-03-27T05:07:58Z",
            "category_id": 5
        },
        {
            "id": 6,
            "title": "Solo Leveling",
            "description": "KOMIK",
            "image_url": "https://lol.com",
            "release_year": 1980,
            "price": "20000",
            "total_page": 200,
            "thickness": "sedang",
            "created_at": "2023-03-27T05:09:06Z",
            "updated_at": "2023-03-27T05:09:06Z",
            "category_id": 5
        }
    ]
}
```
#### *IF THERE IS NO DATA EXIST*
```json
{
    "result": null
}
```



# **Book**
### **1.GET /books**

#### Get list of all exist books


* **Request Header :**
```json
NONE
```

* **Request Body :**
```json
NONE
```

* **Response (200 - OK) :**
#### *IF EXIST*
```json
{
    "result": [
        {
            "id": 2,
            "title": "Naruto",
            "description": "Komik",
            "image_url": "https://lol.com",
            "release_year": 1980,
            "price": "20000",
            "total_page": 201,
            "thickness": "tebal",
            "created_at": "2023-03-27T00:38:35Z",
            "updated_at": "2023-03-27T05:09:40Z",
            "category_id": 5
        },
        {
            "id": 4,
            "title": "Twilight",
            "description": "NOVEL",
            "image_url": "https://lol.com",
            "release_year": 1980,
            "price": "20000",
            "total_page": 200,
            "thickness": "sedang",
            "created_at": "2023-03-27T05:07:58Z",
            "updated_at": "2023-03-27T05:07:58Z",
            "category_id": 5
        },
        {
            "id": 5,
            "title": "Barakamon",
            "description": "KOMIK",
            "image_url": "https://lol.com",
            "release_year": 1980,
            "price": "20000",
            "total_page": 200,
            "thickness": "sedang",
            "created_at": "2023-03-27T05:08:23Z",
            "updated_at": "2023-03-27T05:08:23Z",
            "category_id": 9
        },
        {
            "id": 6,
            "title": "Solo Leveling",
            "description": "KOMIK",
            "image_url": "https://lol.com",
            "release_year": 1980,
            "price": "20000",
            "total_page": 200,
            "thickness": "sedang",
            "created_at": "2023-03-27T05:09:06Z",
            "updated_at": "2023-03-27T05:09:06Z",
            "category_id": 5
        }
    ]
}
```
#### *IF THERE IS NO DATA EXIST*
```json
{
    "result": null
}
```

### **2.POST /books**

#### Create New Book


* **Authorization:**
```bash
Basic Auth
```

* **Request Body :**
```json
{
    "Title": "Solo Leveling",
    "Description": "KOMIK",
    "Image_url": "https://lol.com",
    "Release_year": 1980,
    "Total_page": 200,
    "Category_id": 5,
    "Price": "20000"
}
```

* **Response (201 - Created) :**
```json
{
    "result": "Success Insert Book"
}
```
### **3.PUT /books/:id <book.ID>**

#### Edit Book


* **Authorization:**
```bash
Basic Auth
```

* **Request Body :**
```json
{
    "Title": "Solo Leveling",
    "Description": "KOMIK",
    "Image_url": "https://lol.com",
    "Release_year": 1980,
    "Total_page": 200,
    "Category_id": 5,
    "Price": "20000"
}
```

* **Response (200 - OK) :**
```json
{
    "result": "Success Update Book"
}
```
### **4.DELETE /books/:id <book.ID>**

#### DELETE Book


* **Authorization:**
```bash
Basic Auth
```

* **Request Body :**
```json
NONE
```

* **Response (200 - OK) :**
```json
{
    "result": "Success Delete Book"
}
```