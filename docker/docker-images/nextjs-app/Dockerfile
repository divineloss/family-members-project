FROM debian:bullseye-slim
RUN apt update -y
RUN apt install wget -y
RUN mkdir /app
WORKDIR /app
RUN wget https://deb.nodesource.com/setup_18.x -O node.sh
RUN bash -x node.sh
RUN apt install nodejs -y