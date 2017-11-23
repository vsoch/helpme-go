FROM asciinema/asciinema

#############################################
# Build and Usage
#############################################

# Build
# docker build -t vanessa/helpme .
#
# Database
# mkdir -p /tmp/data
# docker run --name helpme-postgres --env POSTGRES_PASSWORD=helpme \
#                                   --env POSTGRES_USER=helpme  \
#                                   --env POSTGRES_DB=db --publish 5432:5432 \
#                                      -v /tmp/data:/var/run/postgresql \
#                                      -d postgres
#
# Run (Docker)
# docker run -p 80:80 -v /tmp/data:/tmp/data vanessa/helpme -db-connect "host=/tmp/data user=helpme dbname=db sslmode=disable"
#     (Local)
# helpme -db-connect "host=/tmp/data user=helpme dbname=db sslmode=disable"

#############################################
# Environment
#############################################

ENV ASCIINEMA_API_URL https://www.asciinema.org
ENV VERSION 1.9.2
ENV OS linux
ENV ARCH amd64
ENV ASCIINEMA_CONFIG_HOME /opt
ENV PATH ${PATH}:/usr/local/go/bin

#ADD config /opt #TODO: add asciinema config?

#############################################
# Install Go
#############################################

RUN apt-get update && apt-get install -y wget git
WORKDIR /tmp
RUN wget https://redirector.gvt1.com/edgedl/go/go${VERSION}.${OS}-${ARCH}.tar.gz && \
         tar -C /usr/local -xzf go${VERSION}.${OS}-${ARCH}.tar.gz && \
         rm go${VERSION}.${OS}-${ARCH}.tar.gz

ENV HELPME_HOME /usr/local/go/src/github.com/vsoch/helpme
RUN mkdir -p ${HELPME_HOME}
WORKDIR ${HELPME_HOME}
ADD . ${HELPME_HOME}
RUN go get

RUN apt-get clean

ENTRYPOINT ["helpme"]
