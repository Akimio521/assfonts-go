FROM goreleaser/goreleaser-cross:v1.24.4

RUN apt-get update && \
    apt-get install -y pkg-config ca-certificates

RUN git clone https://github.com/Microsoft/vcpkg.git /vcpkg && \
    /vcpkg/bootstrap-vcpkg.sh

RUN apt-get install -y wget && \
    source /etc/os-release && \
    wget -q https://packages.microsoft.com/config/debian/$VERSION_ID/packages-microsoft-prod.deb && \
    dpkg -i packages-microsoft-prod.deb && \
    rm packages-microsoft-prod.deb && \
    apt-get update && \
    apt-get install -y powershell

RUN apt-get install -y clang llvm libxml2-dev uuid-dev libssl-dev bash patch make tar xz-utils bzip2 gzip sed cpio unzip rsync git libbz2-dev && \
git clone https://github.com/tpoechtrager/osxcross.git && \
cd osxcross
