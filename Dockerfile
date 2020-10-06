# Build the static executable /entrypoint
FROM golang:alpine as builder
WORKDIR /app/
ADD entrypoint.go /app/
RUN go build -o /entrypoint .

FROM alpine

# Add rsync and a user for the dockerfile
RUN apk add --no-cache rsync
RUN addgroup -S mirror && adduser -S mirror -G mirror

# Create a volume at /data/ that is chowned by mirror
RUN mkdir /data/ && chown -R mirror:mirror /data/
VOLUME /data/

# Copy over the entrypoint
COPY --from=builder /entrypoint /entrypoint

ENV FOLDER /data/
ENV DELAY 24h

USER mirror:mirror
EXPOSE 8080
CMD [ "/entrypoint" ]