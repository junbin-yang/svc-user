#! /bin/sh
set -eu

BASE_DIR=$(cd `dirname $0`; pwd)

function show_usage
{
    echo "Usage: $0 <label>"
    echo "Example:"
    echo "    sh $0 1.0.0.001"
    exit 1
}

while getopts ":" optname
do
    case $optname in
        *)
            show_usage
            exit 0
            ;;
    esac
done

if [ $# -lt 1 ]
then
    show_usage
    exit
fi

VERSION=$1
BasicName=svc-user
#BuildRoot="/usr/src/packages"
BuildRoot="/root/rpmbuild"
SPECS="$BuildRoot/SPECS"
SOURCES="$BuildRoot/SOURCES"
SRPMS="$BuildRoot/SRPMS"

cd $BASE_DIR

/usr/local/protoc/bin/protoc -I proto --go_out=plugins=grpc:. ./proto/user.proto --doc_out=./docs --doc_opt=html,index.html
/usr/local/protoc/bin/protoc --php_out=sdk/php ./proto/user.proto
swag init --exclude proto

sed -i s/0.0.0.000/$VERSION/g  $BASE_DIR/config/config.go
sed -i s/0000-00-00/$(date "+%Y-%m-%d")/g $BASE_DIR/config/config.go

mkdir -p ${SOURCES}/${BasicName}.${VERSION}/docs
mkdir -p ${SOURCES}/${BasicName}.${VERSION}/web/dist

go build
mv ${BasicName} ${SOURCES}/${BasicName}.${VERSION}/
cp -ar docs/* ${SOURCES}/${BasicName}.${VERSION}/docs/
cp -ar web/dist/* ${SOURCES}/${BasicName}.${VERSION}/web/dist/
cp -ar rpmbuild/${BasicName}.yml ${SOURCES}/
cp -ar rpmbuild/${BasicName}.service ${SOURCES}/

cd ${SOURCES}
tar -cvf ${BasicName}.${VERSION}.tar.gz ${BasicName}.${VERSION}

sVERSION=${VERSION%.*}
release=${VERSION##*.}
cp $BASE_DIR/rpmbuild/${BasicName}.spec $SPECS/${BasicName}.spec
sed -i "s/_VERSION_/$sVERSION/g" $SPECS/${BasicName}.spec
sed -i "s/_RELEASE_/$release/g" $SPECS/${BasicName}.spec

rpmbuild --buildroot=/usr/src/packages -bb $SPECS/${BasicName}.spec
#rm -rf ${SOURCES}/${BasicName}.*
#rm -rf ${SRPMS}/${BasicName}.*
