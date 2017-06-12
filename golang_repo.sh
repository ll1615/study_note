#!/bin/sh

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

    # move directory
    dst="$GOPATH/src/golang.org/x/${pkg}"
    # mkdir before moving
    mkdir -p ${dst}

    echo "moving repo to ${dst}"
    mv -f ${pkg} ${dst}

    echo "\n"
done
