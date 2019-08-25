FROM alpine:latest

ARG ENV
# Install app dependencies
COPY .env.local /src/
COPY *.pem /src/
COPY go-archetype /src/
RUN if [ "x$ENV" = "xLOCAL" ] ; then cp -f /src/.env.local /src/.env ; else rm -f /src/.env.local ; fi
EXPOSE 8000

WORKDIR /src
CMD [ "./go-archetype" ]