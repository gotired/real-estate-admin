# Property API documentation

1. **List Properties**

   `Method` : `GET`

   `Endpoint` : /api/v1/properties

   `Parameters` :

   | Key      | Type    | Detail                                  | Mandatory |
   | -------- | ------- | --------------------------------------- | :-------: |
   | skip     | integer | data section `default = 0`              |     N     |
   | limit    | integer | limit data in response `default = 10`   |     N     |
   | keyword  | string  | property title keyword `default = null` |     N     |
   | sort_key | string  | sorting key `default = title`           |     N     |
   | order    | integer | sorting order `default = 1`             |     N     |
   | type     | string  | property type `default = null`          |     N     |

   `Response Schema` :

   | Key         | Type          | Detail                             |
   | ----------- | ------------- | ---------------------------------- |
   | status_code | integer       | api response status code           |
   | status      | enum          | response status                    |
   | data        | array(object) | response data                      |
   | detail      | string        | response detail when error happens |

   - `data`

   | Key         | Type   | Detail                |
   | ----------- | ------ | --------------------- |
   | id          | uuid   | property id           |
   | owner       | object | property owner detail |
   | description | string | property description  |
   | type        | string | property type         |
   | location    | string | property location     |
   | price       | float  | property price        |
   | photo       | string | property photo        |

   - `owner`

   | Key    | Type   | Detail       |
   | ------ | ------ | ------------ |
   | id     | uuid   | owner id     |
   | name   | string | owner name   |
   | email  | string | owner email  |
   | avatar | string | owner avatar |

2. **Get Property Detail**

   `Method` : `GET`

   `Endpoint` : /api/v1/properties/{id}

   `Parameter` :

   - `id` `(uuid)` : property id in uuid format

3. **Create Property**

   `Method` : `POST`

   `Endpoint` : /api/v1/properties

   `Request Schema` :

   | Key         | Type   | Detail               | Mandatory |
   | ----------- | ------ | -------------------- | :-------: |
   | email       | string | user email           |     Y     |
   | title       | enum   | property title       |     Y     |
   | description | string | property description |     Y     |
   | type        | string | property type        |     Y     |
   | location    | string | property location    |     Y     |
   | price       | string | property price       |     Y     |
   | photo       | string | property photo       |     Y     |

4. **Delete Property**

   `Method` : `DELETE`

   `Endpoint` : /api/v1/properties/{id}

   `Parameter` :

   - `id` `(uuid)` : property id in uuid format

   `Response Schema` :

   | Key         | Type    | Detail          |
   | ----------- | ------- | --------------- |
   | status_code | integer | api status code |
   | status      | enum    | response status |
   | detail      | string  | response detail |

5. **Update Property**

   `Method` : `PATCH`

   `Endpoint` : /api/v1/properties/{id}

   `Parameter` :

   - `id` `(uuid)` : property id in uuid format

   `Request Schema` :

   | Key         | Type   | Detail               | Mandatory |
   | ----------- | ------ | -------------------- | :-------: |
   | owner       | uuid   | user uuid            |     N     |
   | title       | enum   | property title       |     N     |
   | description | string | property description |     N     |
   | type        | string | property type        |     N     |
   | location    | string | property location    |     N     |
   | price       | string | property price       |     N     |
   | photo       | string | property photo       |     N     |

   `Response Schema` :

   | Key         | Type    | Detail          |
   | ----------- | ------- | --------------- |
   | status_code | integer | api status code |
   | status      | enum    | response status |
   | detail      | string  | response detail |
