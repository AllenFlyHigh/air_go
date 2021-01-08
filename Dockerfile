FROM golang:alpine
WORKDIR .
COPY main .
COPY conf/app.ini ./conf/
COPY templates/*.html ./templates/
EXPOSE 8098
CMD ["./main"]