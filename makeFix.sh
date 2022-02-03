#! /bin/bash

for dir in src/*
do
    if [ -d $dir ]; then
    cp -r $dir ~/go/src/
    fi
done

mv ~/go/src/github.com/\!burnt\!sushi/ ~/go/src/github.com/BurntSushi
mv ~/go/src/github.com/\!promon\!logicalis/ ~/go/src/github.com/PromonLogicalis
mv ~/go/src/gopkg.in/robfig/cron.v2@v2.0.0-20150107220207-be2e0b0deed5/ ~/go/src/gopkg.in/robfig/cron.v2

sudo apt install golang-goprotobuf-dev

