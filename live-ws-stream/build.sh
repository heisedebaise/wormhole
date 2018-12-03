#!/bin/bash

yarn build
sudo rm -rf /home/tomcat/webapps/ROOT/console
sudo mv build /home/tomcat/webapps/ROOT/console