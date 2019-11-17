FROM debian

RUN adduser --disabled-password --gecos "" app
USER app
WORKDIR /home/app

COPY portal .
COPY web .

ENTRYPOINT ["./portal"]
CMD ["--config", "/etc/config/portal.yml"]
