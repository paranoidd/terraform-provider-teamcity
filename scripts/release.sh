#!/bin/bash

VERSION=$(cat VERSION)
mkdir "bin/archives/"
ARRAY=(
# OS		# ARCH		# Archive	# Ext
'darwin'	'amd64'		'tar.gz'	''
'darwin'	'386'		'tar.gz'	''
'linux'		'amd64'		'tar.gz'	''
'linux'		'386'		'tar.gz'	''
'windows'	'amd64'		'zip'		'.exe'
'windows'	'386'		'zip'		'.exe'
)

IFS=$'\t'

for ((i=0;i<$((${#ARRAY[@]}/4));i++))
do
  GOOS="${ARRAY[$(($i*4))]}"
  GOARCH="${ARRAY[$(($i*4+1))]}"
  ARCHIVE="${ARRAY[$(($i*4+2))]}"
  EXTENTION="${ARRAY[$(($i*4+3))]}"
  DIR="terraform-provider-teamcity_${VERSION}_${GOOS}_${GOARCH}"

  mkdir -p "bin/${DIR}"

  go build -o "bin/${DIR}/terraform-provider-teamcity_v${VERSION}${EXTENTION}"

  pushd "bin/$DIR"
    if [ "$ARCHIVE" == "tar.gz" ]; then
      tar -zcvf ../archives/${DIR}.${ARCHIVE} *
    fi
    if [ "$ARCHIVE" == "zip" ]; then
      zip ../archives/${DIR}.${ARCHIVE} *
    fi
  popd
done

pushd "bin/archives"
  shasum -a 256 * > sha256sums.txt
  hub release create $(echo $(echo * | xargs -n1 echo -a)) -m "v${VERSION}" "v${VERSION}"
popd

