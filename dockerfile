FROM golang:1.22rc2-bookworm

##buat folder APP
RUN mkdir /be-online-store

##set direktori utama
WORKDIR /be-online-store

##copy seluruh file ke completedep
ADD . /be-online-store

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["./main"]