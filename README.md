# SOSMED_API

## REQUIREMENTS
+ docker
+ docker-compose

## RUNNIG APP
```
    docker-compose up
```
## ENDPOINTS

### BASE URL = http://localhost:3000/api
<br/>

### <b> User Register <b> 
``` 
    METHOD  : POST
    PATH    : /v1/

    BODY    : Username STRING
              Password String

```
### <b> User Login <b> 
``` 
    METHOD  : POST
    PATH    : /v1/login

    BODY    : Username STRING
              Password String

```

### <b> User Upload Short Video <b> 
``` 
    METHOD  : POST
    PATH    : /v1/upload/:accessToken

    BODY    : Category  STRING
              Caption   String

```

### <b> Get All Sort Video <b> 
``` 
    METHOD  : GET
    PATH    : /v1/

    BODY    : None

```

### <b> Stream Short Video <b> 
``` 
    METHOD  : GET
    PATH    : /v1/:video-file-name

    BODY    : None

```