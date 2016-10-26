# Base image
FROM python:3.5.2-alpine

ARG project_dir=/web/socket/

# install apk and git
RUN apk update
RUN apk add git

# install dependencies
WORKDIR $project_dir
ADD requirements.txt .
RUN pip install -r requirements.txt

# install Flask-SocketIO from remote repos at github.com.
RUN git clone https://github.com/miguelgrinberg/Flask-SocketIO.git Flask-SocketIO
WORKDIR $project_dir/Flask-SocketIO/example
