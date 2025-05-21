FROM golang:1.24.3 AS build

ENV PATH=/usr/local/go/bin:$PATH

ENV GOLANG_VERSION=1.24.3

WORKDIR /app

RUN mkdir /app/sentbon

# Download Go modules
#COPY go.mod go.sum backend/main.go ./

#RUN go mod download

# Copy the source code. 
#COPY backend/api/*.go backend/getIP/*.go backend/utils/*.go backend/main.go ./

COPY backend ./sentbon/backend

RUN go mod init

RUN go build main.go

FROM node:alpine3.20

WORKDIR /app

COPY --from=build /app/main ./main

RUN chmod +x ./main

COPY ./frontend ./app/

RUN npm i frontend/

EXPOSE 3000

EXPOSE 8080

CMD [ "./main", "npm serve -s ./frontend/build" ]