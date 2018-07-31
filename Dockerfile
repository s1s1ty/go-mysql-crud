FROM alpine

COPY dist/go-mysql-crud /bin/

EXPOSE 5001

ENTRYPOINT [ "/bin/go-mysql-crud" ]
