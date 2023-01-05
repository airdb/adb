FROM airdb/base:latest

WORKDIR /srv

ADD ./output /srv/output
ADD ./.env.* /srv/output
ADD ./.env.* /srv/
ADD ./.env /srv/output
ADD ./.env /srv/
RUN apk add lego

ENTRYPOINT ["./output/adb"]
CMD ["serve"]
#CMD ["config", "./config/config.yml", "--additional-files", "./config/config_yummy.yml"]
