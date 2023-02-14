


# Unisender API
This is unisender API for import contacts for Unisender service. You can find out more about the API.
  

## Informations

### Version

1.0.0

### Contact

Alexey Kondratev samarec1812@gmail.com https://github.com/samarec1812

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/x-www-form-urlencoded

### Produces
  * application/json

## All endpoints

###  operations

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /unisender/contacts | [create contact](#create-contact) |  |
  


## Paths

### <span id="create-contact"></span> create contact (*createContact*)

```
POST /unisender/contacts
```

Returns status operation

#### Consumes
  * application/x-www-form-urlencoded

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-contact-200) | OK | Successful operation |  | [schema](#create-contact-200-schema) |
| [400](#create-contact-400) | Bad Request | Operation failed |  | [schema](#create-contact-400-schema) |
| [500](#create-contact-500) | Internal Server Error | Server error |  | [schema](#create-contact-500-schema) |

#### Responses


##### <span id="create-contact-200"></span> 200 - Successful operation
Status: OK

###### <span id="create-contact-200-schema"></span> Schema

##### <span id="create-contact-400"></span> 400 - Operation failed
Status: Bad Request

###### <span id="create-contact-400-schema"></span> Schema

##### <span id="create-contact-500"></span> 500 - Server error
Status: Internal Server Error

###### <span id="create-contact-500-schema"></span> Schema

## Models

### <span id="contact"></span> Contact


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AccountId | uint64 (formatted integer)| `uint64` |  | | The AccountId of a number person who create this contact |  |
| ContactId | uint64 (formatted integer)| `uint64` |  | | The ContactId of a thing | `11225727` |
| CreatedAt | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| DeletedAt | [DeletedAt](#deleted-at)| `DeletedAt` |  | |  |  |
| Email | string| `string` |  | | The Email of a thing | `artur201@mail.ru` |
| ID | uint64 (formatted integer)| `uint64` |  | |  |  |
| Name | string| `string` |  | | The Name of a thing | `Artur Dolgih` |
| UpdatedAt | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |



### <span id="deleted-at"></span> DeletedAt


  


* composed type [NullTime](#null-time)

### <span id="model"></span> Model


> Model a basic GoLang struct which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt
It may be embedded into your model or you may build your own model without it
type User struct {
gorm.Model
}
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CreatedAt | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| DeletedAt | [DeletedAt](#deleted-at)| `DeletedAt` |  | |  |  |
| ID | uint64 (formatted integer)| `uint64` |  | |  |  |
| UpdatedAt | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |



### <span id="null-time"></span> NullTime


> NullTime implements the Scanner interface so
it can be used as a scan destination, similar to NullString.
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Time | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| Valid | boolean| `bool` |  | |  |  |


