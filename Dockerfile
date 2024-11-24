FROM golang:1.23.2
FROM node:latest

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN air
RUN npm install -D tailwindcss
RUN make css
RUN make

COPY . .
COPY package*.json ./

EXPOSE 3000

CMD [ "npm", "start" ]

RUN go mod tidy