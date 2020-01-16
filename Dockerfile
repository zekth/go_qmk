FROM golang

COPY dist/ /server/
WORKDIR /server/

RUN /server/go-qmk