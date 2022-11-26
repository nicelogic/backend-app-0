
FROM golang

ENV APP_PATH=/opt/service/user
ENV GOPROXY https://goproxy.cn

WORKDIR $APP_PATH

COPY ./src/. $APP_PATH

RUN cd $APP_PATH
RUN go mod tidy
RUN go build .
CMD [ "./user" ]

EXPOSE 80