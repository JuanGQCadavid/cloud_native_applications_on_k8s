FROM registry.hpc.ut.ee/mirror/library/alpine

RUN apk add --no-cache python3 py3-pip

RUN python3  -m venv /venv

RUN /venv/bin/pip3 install flask

COPY server.py /opt/server.py

#Announcing that container will listen to port 5000 
EXPOSE 5000

CMD /venv/bin/python3 /opt/server.py