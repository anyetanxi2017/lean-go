```
FROM golang AS build-env
MAINTAINER www.983132370@qq.com
RUN mkdir /server
COPY . /server

WORKDIR /server
FROM build-env
CMD ["/server/gin-vue-admin"]
```
