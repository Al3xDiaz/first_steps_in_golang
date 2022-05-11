FROM golang:1.18-alpine3.14 as build
WORKDIR /usr/lib/app
COPY . .
RUN go build main.go

FROM scratch as runner
WORKDIR /
COPY --from=build /usr/lib/app/main .
CMD [ "/main" ]



