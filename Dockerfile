FROM scratch

COPY app/main /main

EXPOSE 8082

CMD ["/main"]
