FROM golang as build

RUN mkdir -p /usr/build
WORKDIR /usr/build
COPY . .
RUN make build

FROM alpine

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY --from=build /usr/build/bin/gservice /usr/src/app
EXPOSE 3000

CMD [ "/usr/src/app/gservice" ]
