FROM golang:1.23.2 as builder
WORKDIR /home/sachin19183/project/client

COPY go.mod ./
RUN go mod tidy
COPY . .

RUN go build -o /home/sachin19183/project/client/bin/client . || { echo 'Go build failed'; exit 1; }
RUN ls -l /home/sachin19183/project/client/bin/   # Check the contents of the bin directory
#RUN file /home/sachin19183/server/bin/server  # Check the binary type
FROM golang:1.23.2
COPY --from=builder /home/sachin19183/project/client/bin /home/sachin19183/project/client/bin
# Make sure the binary has execute permissions
RUN chmod +x /home/sachin19183/project/client/bin/client
EXPOSE 8080
ENTRYPOINT ["/home/sachin19183/project/client/bin/client"]
