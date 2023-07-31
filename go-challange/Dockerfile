# stage de build
FROM golang:alpine AS build

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go

# stage imagem final
FROM scratch

WORKDIR /app

COPY --from=build /app/api ./

CMD [ "./api" ]

# commands
## docker build --tag mascalmeida/fullcycle .
## docker run mascalmeida/fullcycle