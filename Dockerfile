FROM alpine:latest
COPY --chown=200:200 --chmod=555 server /app/
CMD ["/app/server"]
