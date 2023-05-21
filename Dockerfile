FROM golang:1.20-alpine3.18 as build
WORKDIR /build
COPY go.* .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o firetest
FROM alpine:3.18
WORKDIR /usr/bin
COPY --from=build /build/firetest firetest
RUN apk add git curl nmap git
# Download ffuf
RUN wget https://github.com/ffuf/ffuf/releases/download/v2.0.0/ffuf_2.0.0_linux_amd64.tar.gz -O ffuf.tar.gz; \
    tar xvfz ffuf.tar.gz ffuf

# Download Wordlists
WORKDIR /wordlists
RUN wget https://github.com/danielmiessler/SecLists/raw/master/Passwords/Leaked-Databases/rockyou.txt.tar.gz
RUN wget http://ffuf.me/wordlist/common.txt
RUN wget http://ffuf.me/wordlist/parameters.txt
RUN wget http://ffuf.me/wordlist/subdomains.txt

COPY docker-entrypoint.sh /
ENTRYPOINT [ "/docker-entrypoint.sh" ]