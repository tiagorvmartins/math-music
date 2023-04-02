FROM golang:1.19
RUN apt-get update
RUN apt-get -qq -y install python
RUN python -m pip install --upgrade pip
RUN python -m pip install midiutil
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
COPY *.py ./
RUN go build -o /midiserver
COPY .env /app
EXPOSE 5175
CMD [ "/midiserver" ]