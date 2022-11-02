FROM ubuntu:22.04

LABEL com.naiama.image.author="mlops@naiama.com"

SHELL ["/bin/bash", "-c"]

RUN apt-get update && apt-get install -y build-essential curl vim

WORKDIR /tmp
ENV MINICONDA Miniconda3-py39_4.12.0-Linux-x86_64.sh
ENV MINICONDA_CHECKSUM 78f39f9bae971ec1ae7969f0516017f2413f17796670f7040725dd83fcff5689
RUN curl -O https://repo.anaconda.com/miniconda/$MINICONDA && \
    shasum -a256 $MINICONDA | awk -v ac=$MINICONDA_CHECKSUM '$1==ac{print checksum-ok}'
RUN bash $MINICONDA -b -p /opt/miniconda
ENV PATH /opt/miniconda/bin:$PATH
RUN eval "$(conda shell.bash hook)"

RUN conda update -n base -c defaults conda

# Python 3.9.*
COPY ./environment.yml ./
RUN conda env create -f environment.yml
RUN conda init
SHELL ["conda", "run", "-n", "mlops", "/bin/bash", "-c"]

# PyTorch 1.13.1 + pre-requisites
RUN conda install astunparse numpy ninja pyyaml setuptools cmake cffi typing_extensions future six requests dataclasses mkl mkl-include
RUN conda install pytorch torchvision torchaudio cpuonly -c pytorch

# Lightning 1.7.7
RUN conda install pytorch-lightning -c conda-forge

CMD ["/bin/bash"]
