FROM debian:8

RUN mkdir /authServer

ADD AuthServer /authServer/AuthServer

# set the working directory to the application directory
WORKDIR /authServer

# expose the application on port 8001.
# This should be the same as in the port used in the application
EXPOSE 8001

# set the entry point of the container to the application executable
ENTRYPOINT /authServer/AuthServer
