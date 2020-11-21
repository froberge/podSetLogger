# PodSet Logger

Go service that read the content of a Configuration file and display it on the web
You can fine a sample config file in the config/ folder.

## build
The application was created to run in a a docker image.  To build it use the build.sh script.

To run in docker standalone including the config file use the following command
 * ./build.sh /etc/config/

To run in a docker image and deploy inside k8s and read from a configMap use the following command
 * ./build.sh
 
## Authors
[Felix Roberge](https://github.com/froberge-cloudOps)
