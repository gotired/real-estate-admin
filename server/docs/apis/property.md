# Property API documentation

1. **List Properties**

     `Method` : `GET`

     `Endpoint` : /api/v1/properties

     `Parameters` :

     | Key     | Type    | Detail                                  | Mandatory |
     | ------- | ------- | --------------------------------------- | --------- |
     | sec     | integer | data section `default = 1`              | N         |
     | limit   | integer | limit data in response `default = 10`   | N         |
     | keyword | string  | property title keyword `default = None` | N         |
     | type    | string  | property type `default = None`          | N         |

     `Response Schema` :

     | Key         | Type    | Detail                             |
     | ----------- | ------- | ---------------------------------- |
     | status_code | integer | api response status code           |
     | status      | enum    | response status                    |
     | data        | object  | response data                      |
     | detail      | string  | response detail when error happens |

     - `data`

     | Key        | Type          | Detail                        |
     | ---------- | ------------- | ----------------------------- |
     | name       | string        | user name                     |
     | email      | string        | user email                    |
     | avatar     | string        | user avatar                   |
     | properties | array[string] | all properties that user owns |

2. **Get Property Detail**

     `Method` : `GET`

     `Endpoint` : /api/v1/properties/{id}

     `Parameter` :

     - `id` `(uuid)` : property id in uuid format

3. **Create Property**

     `Method` : `POST`

     `Endpoint` : /api/v1/properties

4. **Delete Property**

     `Method` : `DELETE`

     `Endpoint` : /api/v1/properties/{id}

     `Parameter` :

     - `id` `(uuid)` : property id in uuid format

5. **Update Property**

     `Method` : `PATCH`

     `Endpoint` : /api/v1/properties/{id}

     `Parameter` :

     - `id` `(uuid)` : property id in uuid format
