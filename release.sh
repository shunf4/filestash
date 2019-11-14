#!/bin/bash
cp -R config/* ./dist/data/state/config/
sudo chmod -R o-r-w-x- ./dist/
mv dist filestash
tar -cf filestash_linux-amd64.tar ./filestash
mv filestash dist
sudo mv filestash_linux-amd64.tar /media/datadisk/syncthing/gossa/
