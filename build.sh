set -ex

#SET NEEDED VARIABLE
#HOSTNAME=gcr.io
#PROJECTID=tools-poc
IMAGENAME=golang-helloworld
VERSION=0.0.1

docker build --build-arg confFilePah=$1 . -t $IMAGENAME:$VERSION 
#docker build -t $HOSTNAME/$PROJECTID/$IMAGENAME:$VERSION .
