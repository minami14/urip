FROM golang AS build

WORKDIR $GOPATH/src/github.com/minami14/urip

COPY . .

RUN CGO_ENABLED=0 go build -o /urip main.go


FROM scratch

COPY --from=build /urip /urip

CMD ["/urip"]
