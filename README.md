# PodSet Logger

Go service that read the content of a Configuration file and display it on the web
You can fine a sample config file in the config/ folder.

## Build
The application was created to run in a docker image.  To build it use the build.sh script with the define paramets

Name  | Required |  Description 
--------------- | ---------| -----------------
IMAGE_NAME | Y |The name of the image you want to give
VERSION | Y|  The version of the images. Use latest if you don't want to version yet.
CONFIG_FILE | N |  The config file path if needed.


To build a standalone docker images that including the config file use the following command.
 * ./build.sh [IMAGE_NAME] [VERSION] /etc/config/
 * _ex: ./build.sh podset-logger latest /etc/config/_   

To build a docker images that will be deploy in a kubernetes cluster using a ConfigMap use the following command.
 * ./build.sh [IMAGE_NAME] [VERSION]
 * _ex: ./build.sh podset-logger latest_
 
## Authors
[Felix Roberge](https://github.com/roberge.felix@gmail.com)
"
