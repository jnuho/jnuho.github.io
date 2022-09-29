FROM ubuntu:18.04
MAINTAINER LGCNS
RUN \
      apt-get update && \
      apt-get install -y apache2

RUN sed -i 's/80/7080/g' /etc/apache2/ports.conf
EXPOSE 7080
CMD ["apachectl", "-D", "FOREGROUND"]
