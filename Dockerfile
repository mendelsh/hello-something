FROM scratch

COPY app/main /

EXPOSE 8082

CMD ["/main"]
