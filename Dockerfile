# Clone Job
# FixMe: Add parameter for master/branch ?
FROM alpine/git as clone
COPY . /go_qmk
WORKDIR /
# RUN git clone https://github.com/zekth/go_qmk.git

# API Build job base on makefile
FROM golang as gobuilder
COPY --from=clone /go_qmk/api /go_qmk_api
WORKDIR /go_qmk_api
RUN make ci-build

# UI Build job base on yarn
FROM node:10 as nodebuilder
COPY --from=clone /go_qmk/ui /go_qmk_ui
WORKDIR /go_qmk_ui
RUN yarn
RUN yarn build

# Runtime
# FixMe: Add QMK
FROM scratch
COPY --from=gobuilder /go_qmk_api/dist/ /server/
COPY --from=nodebuilder /go_qmk_ui/dist/ /server/ui/
WORKDIR /server/
EXPOSE 8080

ENTRYPOINT ["./go_qmk"]
