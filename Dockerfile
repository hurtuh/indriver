FROM golang:1.13-alpine
EXPOSE 8080
COPY service-indriver .
CMD ["./service-indriver"]