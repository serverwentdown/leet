#!/bin/sh

set -ex

echo "Installing dependencies"
sudo apt-get install -y git scons build-essential

echo "Cloning ws281x"
git clone https://github.com/jgarff/rpi_ws281x /tmp/rpi_ws281x

echo "Building ws281x"
cd /tmp/rpi_ws281x
scons

echo "Installing into /usr/local"
sudo cp libws2811.a /usr/local/lib
sudo cp ws2811.h pwm.h rpihw.h /usr/local/include

echo "Cleaning up"
cd /tmp
rm -rf /tmp/rpi_ws281x

