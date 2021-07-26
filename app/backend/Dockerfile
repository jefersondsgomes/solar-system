FROM golang

WORKDIR /app

COPY backend/go.mod .
COPY backend/go.sum .
RUN go mod download

COPY . .

RUN go build -o /solar-system-api

CMD [ "/solar-system-api" ]