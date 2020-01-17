# Clone Job
# FixMe: Add parameter for master/branch ?
FROM alpine/git as clone
WORKDIR /
RUN git clone https://github.com/zekth/go_qmk.git

# Build job base on makefile
FROM golang as builder
COPY --from=clone /go_qmk /go_qmk
WORKDIR /go_qmk
RUN make ci-build

# Runtime
# FixMe: Add QMK
FROM scratch
COPY --from=builder /go_qmk/dist/ /server/
WORKDIR /server/

EXPOSE 8080

CMD ["/server/go_qmk"] 
