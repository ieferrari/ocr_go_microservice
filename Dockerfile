FROM golang:latest

# tesseract installation
RUN apt-get update -qq
RUN apt-get install -y -qq libtesseract-dev libleptonica-dev

# In case you face TESSDATA_PREFIX error, you minght need to set env vars
# to specify the directory where "tessdata" is located.
ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr/4.00/tessdata/

# Load languages.
# These {lang}.traineddata would b located under ${TESSDATA_PREFIX}/tessdata.
RUN apt-get install -y -qq \
  tesseract-ocr-eng \
  tesseract-ocr-spa
# See https://github.com/tesseract-ocr/tessdata for the list of available languages.
# If you want to download these traineddata via `wget`, don't forget to locate
# downloaded traineddata under ${TESSDATA_PREFIX}/tessdata.


RUN cd ${GOPATH}
COPY ./app ${GOPATH}/app
WORKDIR ${GOPATH}/app
RUN go get -t github.com/otiai10/gosseract
RUN go mod tidy
RUN go build main.go ocr.go
CMD [ "./main" ]

# BUILD DOCKER IMAGE
# docker build -f Dockerfile . -t ieferrari/ocr_go_microservice
#   successfully build [some_hash]
# PUSH TO DOCKER-HUB
# docker tag f0b1d3686100 ieferrari/ocr_go_microservice:0.0.1
# docker push ieferrari/ocr_go_microservice:0.0.1
# RUN THE CONTAINER
# docker run -p5005:5005 ieferrari/ocr_go_microservice -d
