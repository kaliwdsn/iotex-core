FROM golang:1.10.2-stretch

RUN apt-get install -y --no-install-recommends make && \
    mkdir -p $GOPATH/src/github.com/iotexproject/iotex-core/

COPY ./ $GOPATH/src/github.com/iotexproject/iotex-core/

ARG SKIP_DEP=false

RUN if [ "$SKIP_DEP" != true ] ; \
    then \
        cd $GOPATH/src/github.com/iotexproject/iotex-core/ && \
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
        dep ensure ; \
    fi


RUN cd $GOPATH/src/github.com/iotexproject/iotex-core/ && \
    make clean build && \
    ln -s $GOPATH/src/github.com/iotexproject/iotex-core/bin/server /usr/local/bin/iotex-server  && \
    ln -s $GOPATH/src/github.com/iotexproject/iotex-core/bin/actioninjector /usr/local/bin/iotex-actioninjector && \
    mkdir -p /usr/local/lib/iotex/ && \
    ln -s $GOPATH/src/github.com/iotexproject/iotex-core/bin/addrgen /usr/local/bin/iotex-addrgen && \
    mkdir -p /usr/local/lib/iotex/ && \
    cp $GOPATH/src/github.com/iotexproject/iotex-core/crypto/lib/libsect283k1_ubuntu.so /usr/lib/ && \
    mkdir -p /etc/iotex/ && \
    ln -s $GOPATH/src/github.com/iotexproject/iotex-core/config.yaml /etc/iotex/config.yaml && \
    ln -s $GOPATH/src/github.com/iotexproject/iotex-core/blockchain/testnet_actions.yaml /etc/iotex/testnet_actions.yaml && \
    sed -i '/inMemTest/a \    genesisActionsPath: \/etc\/iotex\/testnet_actions\.yaml' /etc/iotex/config.yaml && \
    mkdir -p /var/log/iotex/

CMD ["iotex-server", "-log-path=/var/log/iotex/server.log", "-config=/etc/iotex/config.yaml"]