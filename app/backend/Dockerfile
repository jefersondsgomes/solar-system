FROM golang

WORKDIR /app

COPY backend/go.mod .
COPY backend/go.sum .
RUN go mod download

COPY ./backend .
RUN go build -o /app/main .

CMD ["/app/main"]