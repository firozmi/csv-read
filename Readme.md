    dep init
    docker build -t go-csv-read .

    docker run --rm -it -p 8111:8111 -v <log-location>:/go/src/app/log/ go-csv-read

    Example:
    docker run --rm -it -p 8111:8111 -v /var/log/csv-read:/go/src/app/log/ go-csv-read