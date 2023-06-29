FROM debian:bullseye-slim
RUN apt update -y
RUN echo "deb http://ftp.de.debian.org/debian bullseye-backports main" > /etc/apt/sources.list.d/debian-backports.list
RUN apt update -y
RUN apt install golang-1.19* -y
RUN cp /usr/lib/go-1.19/bin/go /usr/lib/go-1.19/bin/gofmt /usr/bin
RUN mkdir /api
WORKDIR /api