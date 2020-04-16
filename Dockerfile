FROM golang as builder

ADD . /twqrp-api
WORKDIR /twqrp-api
RUN go mod download && go build

FROM ronmi/mingo
COPY --from=builder /twqrp-api/twqrp-api /twqrp-api
