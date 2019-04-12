#!/bin/sh

set -ex

echo "Installing dependencies"
sudo apt-get install -y git scons build-essential

echo "Cloning ws281x"
git clone https://github.com/jgarff/rpi_ws281x /tmp/rpi_ws281x

echo "Building ws281x"
cd /tmp/rpi_ws281x
scons deb

echo "Installing package"
sudo dpkg --purge libws2811 || true
sudo dpkg -i libws2811*.deb

echo "Cleaning up"
cd /tmp
rm -rf /tmp/rpi_ws281x

