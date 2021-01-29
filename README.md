# onvif-viewer

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-white.svg)](https://snapcraft.io/onvif-viewer)

```sh
onvif-viewer -c /path/to/config/file/onvif-viewer.yaml
```
or just:
```
onvif-viewer
```
(use default config file: ./onvif-viewer.yaml)

or:
```sh
 onvif-viewer run -i myid -k mykey -m iothub.cloud -s www -c 60
```
-i {AccessId} -k {AccessKey} -m {MainDomain} -s {SubDomainName} -c {CheckUpdateInterval}

You can install the pre-compiled binary (in several different ways),
use Docker.

Here are the steps for each of them:

## Install the pre-compiled binary

**homebrew tap** :

```sh
$ brew install OpenIoTHub/tap/onvif-viewer
```

**homebrew** (may not be the latest version):

```sh
$ brew install onvif-viewer
```

**snapcraft**:

```sh
$ sudo snap install onvif-viewer
```
config file path: /root/snap/onvif-viewer/current/onvif-viewer.yaml

edit config file then:
```sh
sudo snap restart onvif-viewer
```

**scoop**:

```sh
$ scoop bucket add OpenIoTHub https://github.com/OpenIoTHub/scoop-bucket.git
$ scoop install onvif-viewer
```

**deb/rpm**:

Download the `.deb` or `.rpm` from the [releases page][releases] and
install with `dpkg -i` and `rpm -i` respectively.

config file path: /etc/onvif-viewer/onvif-viewer.yaml

edit config file then:
```sh
sudo systemctl restart onvif-viewer
```

**Shell script**:

```sh
$ curl -sfL https://install.goreleaser.com/github.com/mDNSService/onvif-viewer.sh | sh
```

**manually**:

Download the pre-compiled binaries from the [releases page][releases] and
copy to the desired location.

Note that the image will almost always have the last stable Go version.

[releases]: https://github.com/mDNSService/onvif-viewer/releases

