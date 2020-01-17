# Clone Job
# FixMe: Add parameter for master/branch ?
FROM alpine/git as clone
WORKDIR /
RUN git clone https://github.com/zekth/go_qmk.git

# Build job base on makefile
FROM golang as gobuilder
COPY --from=clone /go_qmk/api /go_qmk_api
WORKDIR /go_qmk_api
RUN make ci-build

# Runtime
# FixMe: Add QMK
FROM scratch
COPY --from=gobuilder /go_qmk_api/dist/ /server/
WORKDIR /server/

EXPOSE 8080

CMD ["/server/go_qmk"]
