
if [ "$#" -lt 2 ]; then
  echo "You must enter at least 2 parameter, which are the image name and version.  Only the 3rd parameter is optional"
  exit 1;
fi

echo "The images name: " $1
IMAGENAME=$1

echo "The version: " $2
VERSION=$2

docker build --build-arg confFilePah=$3 . -t $IMAGENAME:$VERSION 

