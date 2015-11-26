FROM ubuntu

RUN apt-get -y install wget
RUN wget http://influxdb.s3.amazonaws.com/influxdb_0.9.5.1_amd64.deb
RUN dpkg -i influxdb_0.9.5.1_amd64.deb

EXPOSE 8083 8086

# rodar esse comando assim que startar o docker
# $service influxdb start
