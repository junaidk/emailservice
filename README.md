#### Description
A microservice for sending email
can be used in knative framework.

#### API

POST /api/v1/send



#### Build docker image

```
docker build -t emailservice .
```

#### Run docker image

```
docker run -d \
-e USER_NAME="<username of smtp server>" \
-e PASSWORD="<password of smtp server>" \
-e F_ADDR="<email address which will be sending email> \
--name emailservice \
emailservice
```