FROM golang:1-alpine3.10 as build
WORKDIR /usr/lib/app
COPY . .
RUN go build src/main.go

FROM scratch as runner
WORKDIR /
COPY --from=build /usr/lib/app/main .
CMD [ "/main" ]



