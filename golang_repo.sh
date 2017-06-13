#!/bin/sh

# check if golang has been installed
if ! command -v go >/dev/null; then
    echo 'Please install golang first'
fi

# check if $GOPATH has been set
if [ -z ${GOPATH+x} ]
then
    echo 'Please set $GOPATH first'
    exit 1
fi

# packages to be downloaded
PKGS=('net' 'tools' 'lint' 'text' 'sys' 'crypto')

for pkg in ${PKGS[@]}
do
    if ! test -d ${pkg} 
    then
        repo="https://github.com/golang/${pkg}.git"

        echo "cloning repo ${repo} ..."
        git clone ${repo} -o ${pkg}

        echo "\n"
    fi

    # copy to destination directory
    dst="$GOPATH/src/golang.org/x/"

    if ! test -d ${dst}${pkg}
    then
        # mkdir before copying
        mkdir -p ${dst}

        echo "copying repo to ${dst}${pkg} ..."
        cp -R ${pkg} ${dst}

        echo "\n"
    fi
done

