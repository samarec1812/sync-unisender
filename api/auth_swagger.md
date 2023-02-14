


# auth.proto
  

## Informations

### Version

version not set

## Tags

  ### <span id="tag-authorization-amo"></span>AuthorizationAmo

  ### <span id="tag-authorization-unisender"></span>AuthorizationUnisender

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  authorization_amo

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /token/get | [authorization amo get token](#authorization-amo-get-token) |  |
  


###  authorization_unisender

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /utoken/save | [authorization unisender save token](#authorization-unisender-save-token) |  |
  


## Paths

### <span id="authorization-amo-get-token"></span> authorization amo get token (*AuthorizationAmo_GetToken*)

```
POST /token/get
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [AuthAuthAmoRequest](#auth-auth-amo-request) | `models.AuthAuthAmoRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#authorization-amo-get-token-200) | OK | A successful response. |  | [schema](#authorization-amo-get-token-200-schema) |
| [default](#authorization-amo-get-token-default) | | An unexpected error response. |  | [schema](#authorization-amo-get-token-default-schema) |

#### Responses


##### <span id="authorization-amo-get-token-200"></span> 200 - A successful response.
Status: OK

###### <span id="authorization-amo-get-token-200-schema"></span> Schema
   
  

[AuthAuthAmoResponse](#auth-auth-amo-response)

##### <span id="authorization-amo-get-token-default"></span> Default Response
An unexpected error response.

###### <span id="authorization-amo-get-token-default-schema"></span> Schema

  

[RPCStatus](#rpc-status)

### <span id="authorization-unisender-save-token"></span> authorization unisender save token (*AuthorizationUnisender_SaveToken*)

```
POST /utoken/save
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [AuthAuthUnisenderRequest](#auth-auth-unisender-request) | `models.AuthAuthUnisenderRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#authorization-unisender-save-token-200) | OK | A successful response. |  | [schema](#authorization-unisender-save-token-200-schema) |
| [default](#authorization-unisender-save-token-default) | | An unexpected error response. |  | [schema](#authorization-unisender-save-token-default-schema) |

#### Responses


##### <span id="authorization-unisender-save-token-200"></span> 200 - A successful response.
Status: OK

###### <span id="authorization-unisender-save-token-200-schema"></span> Schema
   
  

[AuthAuthUnisenderResponse](#auth-auth-unisender-response)

##### <span id="authorization-unisender-save-token-default"></span> Default Response
An unexpected error response.

###### <span id="authorization-unisender-save-token-default-schema"></span> Schema

  

[RPCStatus](#rpc-status)

## Models

### <span id="auth-auth-amo-request"></span> authAuthAmoRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| clientId | string| `string` |  | |  |  |
| clientSecret | string| `string` |  | |  |  |
| code | string| `string` |  | |  |  |
| grantType | string| `string` |  | |  |  |
| redirectUri | string| `string` |  | |  |  |



### <span id="auth-auth-amo-response"></span> authAuthAmoResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| authCode | string| `string` |  | |  |  |



### <span id="auth-auth-unisender-request"></span> authAuthUnisenderRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| accountId | uint64 (formatted string)| `string` |  | |  |  |
| token | string| `string` |  | |  |  |
| uname | string| `string` |  | |  |  |



### <span id="auth-auth-unisender-response"></span> authAuthUnisenderResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| authCode | string| `string` |  | |  |  |



### <span id="protobuf-any"></span> protobufAny


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| @type | string| `string` |  | |  |  |



### <span id="rpc-status"></span> rpcStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| code | int32 (formatted integer)| `int32` |  | |  |  |
| details | [][ProtobufAny](#protobuf-any)| `[]*ProtobufAny` |  | |  |  |
| message | string| `string` |  | |  |  |


