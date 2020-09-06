FROM alpine

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY ./bin/gservice /usr/src/app
EXPOSE 3000

CMD [ "/usr/src/app/gservice" ]
