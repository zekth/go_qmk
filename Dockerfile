# Clone Job
# FixMe: Add parameter for master/branch ?
FROM alpine/git as clone
COPY . /go_qmk
WORKDIR /
# RUN git clone https://github.com/zekth/go_qmk.git

# Cloning and pruning QMK_FIRMWARE repo
FROM alpine/git as qmk_clone
RUN git clone --branch 0.7.120 https://github.com/qmk/qmk_firmware.git
 # Removing useless stuff
RUN rm -rf qmk_firmware/docs

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
# FROM go_qmk_base_container
FROM scratch
COPY --from=gobuilder /go_qmk_api/dist/ /server/
COPY --from=gobuilder /go_qmk_api/schema.graphql /server/schema.graphql
COPY --from=nodebuilder /go_qmk_ui/dist/ /server/ui/
ENV QMK_PATH /qmk_firmware
ENV GIN_MODE release
WORKDIR /server/
EXPOSE 8080

ENTRYPOINT ["./go_qmk"]
