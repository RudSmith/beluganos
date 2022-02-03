mkdir $HOME/go/src

FILES=$HOME/go/pkg/mod/*

for dirs in $FILES
do
    mkdir $HOME/go/src/$(echo $dirs|cut -d '/' -f 7)
done

for str in github.com gitlab.com golang.org
do
    FILES=$HOME/go/pkg/mod/$str/*
    for dirs in $FILES
    do
	mkdir $HOME/go/src/$str/$(echo $dirs|cut -d '/' -f 8)
    done
done

for str in google.golang.org gopkg.in
do
    FILES=$HOME/go/pkg/mod/$str/*
    for dirs in $FILES
    do
	FILE=$HOME/go/src/$str/$(echo $dirs|cut -d '/' -f 8|cut -d '@' -f 1)
	if [ -d $FILE ]; then
	    echo "$FILE already exists"
	else
	    cp -r $dirs $FILE
	fi
    done
done

cp -r $HOME/go/pkg/mod/github.com/vishvananda/netlink@v1.1.0 $HOME/go/src/github.com/vishvananda/netlink
cp -r $HOME/go/pkg/mod/golang.org/x/sys@v0.0.0-20200602100848-8d3cce7afc34 $HOME/go/src/golang.org/x/sys
cp -r $HOME/go/pkg/mod/github.com/golang/protobuf@v1.5.2 $HOME/go/src/github.com/golang/protobuf

cp -r $HOME/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1 $HOME/go/src/github.com/sirupsen/logrus
sudo rm -r ~/go/src/gopkg.in/yaml.v2
cp -r $HOME/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0 $HOME/go/src/gopkg.in/yaml.v2
cp -r $HOME/go/pkg/mod/github.com/spf13/viper@v1.8.1 $HOME/go/src/github.com/spf13/viper
cp -r $HOME/go/pkg/mod/golang.org/x/net@v0.0.0-20190620200207-3b0461eec859 $HOME/go/src/golang.org/x/net

sudo apt install snapd
sudo snap install go --classic --channel=1.16

go get github.com/mdlayher/ndp@17ab9e3

cp -r $HOME/go/pkg/mod/github.com/mdlayher/ndp@v0.0.0-20200602162440-17ab9e3e5567 $HOME/go/src/github.com/mdlayher/ndp


for str in github.com gitlab.com golang.org
do
    FILES=$HOME/go/pkg/mod/$str/*
    for dirs in $FILES
    do
	for subdirs in $dirs/*
	do
	    FILE=$HOME/go/src/$str/$(echo $dirs|cut -d '/' -f 8)/$(echo $subdirs|cut -d '/' -f 9|cut -d '@' -f 1)
	    if [ -d $FILE ]; then
		echo "$FILE already exists"
	    else
		cp -r $subdirs $FILE
	    fi
	done
    done
done