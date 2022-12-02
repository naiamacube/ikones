FROM ubuntu:22.04

LABEL com.naiama.image.author="mlops@naiama.com"

SHELL ["/bin/bash", "-c"]

RUN apt-get update && \
    apt-get install -y build-essential lsb-release wget && \
    apt-get clean all

RUN wget -O- https://apt.releases.hashicorp.com/gpg \
      | gpg --dearmor \
      | tee /usr/share/keyrings/hashicorp-archive-keyring.gpg && \
    echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" \
      | tee /etc/apt/sources.list.d/hashicorp.list && \
    apt-get update && \
    apt-get install vault=1.12.2-1

CMD ["/bin/bash"]
