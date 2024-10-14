# User API documentation

1.  **List Users**

    `Method` : `GET`

    `Endpoint` : /api/v1/users

    `Response Schema` :

    | Key         | Type          | Detail                             |
    | ----------- | ------------- | ---------------------------------- |
    | status_code | integer       | api response status code           |
    | status      | enum          | response status                    |
    | data        | array[object] | response data                      |
    | detail      | string        | response detail when error happens |

    - `data`

    | Key    | Type   | Detail      |
    | ------ | ------ | ----------- |
    | id     | uuid   | user id     |
    | name   | string | user name   |
    | email  | string | user email  |
    | avatar | string | user avatar |

2.  **Create User**

    `Method` : `POST`

    `Endpoint` : /api/v1/users

    `Request Schema` :

    | Key    | Type   | Detail      | Mandatory |
    | ------ | ------ | ----------- | :-------: |
    | name   | string | user name   |     Y     |
    | email  | string | user email  |     Y     |
    | avatar | string | user avatar |     Y     |

    `Response Schema` :

    | Key         | Type    | Detail                             |
    | ----------- | ------- | ---------------------------------- |
    | status_code | integer | api response status code           |
    | status      | enum    | response status                    |
    | detail      | string  | response detail when error happens |

3.  **Get User by ID**

    `Method` : `GET`

    `Endpoint` : /api/v1/users/{id}

    `Parameter` :

    - `id` `(uuid)` : user id in uuid format

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
    | id         | uuid          | user id                       |
    | name       | string        | user name                     |
    | email      | string        | user email                    |
    | avatar     | string        | user avatar                   |
    | properties | array[object] | all properties that user owns |

    - `properties`

    | Key         | Type    | Detail               |
    | ----------- | ------- | -------------------- |
    | id          | uuid    | property id          |
    | title       | string  | property title       |
    | description | string  | property description |
    | type        | string  | property type        |
    | location    | string  | property location    |
    | price       | integer | property price       |
    | photo       | string  | property photo       |
