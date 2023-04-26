FROM --platform=linux/x86_64 golang:latest AS build
RUN mkdir /core
WORKDIR /core
COPY go.mod .
COPY go.sum .
RUN go mod download  
COPY . ./
RUN go build -o botcore main.go



FROM --platform=linux/x86_64 golang:latest
RUN mkdir /app
WORKDIR /app/
COPY --from=build /core/botcore ./
COPY --from=build /core/core.db ./
COPY --from=build /core/app.env ./
CMD ["/app/botcore"]