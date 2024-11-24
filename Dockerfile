FROM golang:1.23.2

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN air
RUN npm install -D tailwindcss

COPY . .
RUN go mod tidy