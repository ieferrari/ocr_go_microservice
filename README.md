# OCR GO Microservice
Api endpoints to call tesseract ocr functions

## Getting started

With docker:

    docker run -p5005:5005 ieferrari/ocr_go_microservice

Get the text from an image url:

    curl -X POST \
         -d '{"msg": "https://pbs.twimg.com/media/EH-Pvo9WwAEKFwc?format=jpg&name=small"}' \
         -H "Content-Type: application/json" \
         http://127.0.0.1:5005/ocr_from_url


***
## Basic installation 

[how to install tesseract](https://github.com/tesseract-ocr/tessdoc/blob/main/Compiling-%E2%80%93-GitInstallation.md)

apt-get install automake ca-certificates g++ git libtool libleptonica-dev make pkg-config

git clone https://github.com/tesseract-ocr/tesseract.git

cd tesseract
    ./autogen.sh
    ./configure
    make
    sudo make install
    sudo ldconfig

https://notesalexp.org/tesseract-ocr/#tesseract_5.x
sudo apt-get install tesseract-ocr

go get github.com/otiai10/gosseract/v2

wget https://github.com/tesseract-ocr/tessdata/raw/4.00/spa.traineddata
tesseract --tessdata-dir . example.png  outputbase -l spa --psm 3




***
## Other languages alternatives

Depending on your architecture, it may be more efficient to call tesseract from a wrapper in your preferred language.
This container is an alternative if your team is having troubles installing the tesseract components for a specific language,
or if you want a centralized ocr implementation in the first place.


**Python example:**

```python
import pytesseract
from PIL import Image
pytesseract.pytesseract.tesseract_cmd ='/usr/local/bin/tesseract'
print(pytesseract.image_to_string(Image.open('./example.png'), lang='spa').replace("º", 'o'))
```


