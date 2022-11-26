FROM docker:20.10.21-git

LABEL com.naiama.image.author="mlops@naiama.com"

# Install Go
ENV GOLANG go1.19.3.linux-amd64.tar.gz
ENV GOLANG_SHA256 "74b9640724fd4e6bb0ed2a1bc44ae813a03f1e72a4c76253e2d5c015494430ba"
RUN wget -O $GOLANG https://go.dev/dl/$GOLANG && \
    sha256sum $GOLANG | awk -v checksum=$GOLANG_SHA256 '$1==checksum{print"checksum-ok"}' && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf $GOLANG
ENV PATH $PATH:/user/local/go/bin

CMD ["/bin/bash"]
