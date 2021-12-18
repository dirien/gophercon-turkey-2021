FROM alpine:3.15.0
COPY dist/gophercon-turkey-2021-linux-amd64 /app/gophercon-turkey-2021-linux-amd64
RUN adduser -D simple
RUN chown -R simple:simple /app
USER simple
CMD ["/app/gophercon-turkey-2021-linux-amd64"]