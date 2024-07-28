FROM golang

WORKDIR /app

COPY . .

RUN go build -o movies.exe

EXPOSE 8080

CMD [ "./movies.exe" ]