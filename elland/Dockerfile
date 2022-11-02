FROM ubuntu:22.04

LABEL com.naiama.image.author="mlops@naiama.com"

RUN apt-get update && apt-get install -y build-essential curl python3.11 python3.11-dev

ENV ANACONDA Anaconda3-2022.10-Linux-x86_64.sh
ENV ANACONDA_CHECKSUM e7ecbccbc197ebd7e1f211c59df2e37bc6959d081f2235d387e08c9026666acd
RUN curl -O https://repo.anaconda.com/archive/$ANACONDA && \
    shasum -a256 $ANACONDA | awk -v ac=$ANACONDA_CHECKSUM '$1==ac{print checksum-ok}' && \
    bash $ANACONDA -bf && eval "$($HOME/anaconda3/bin/conda shell.bash hook)" && conda init

# There is no maintained binary for PyTorch CPU-only, so we have to compile from source
RUN $HOME/anaconda3/bin/conda install pytorch torchvision torchaudio cpuonly -c pytorch

RUN $HOME/anaconda3/bin/conda install pytorch-lightning -c conda-forge

RUN python3 -m pip install torch-scatter torch-sparse torch-geometric

# RUN pip install torch-scatter && \
    # pip install torch-sparse && \
    # pip install torch-geometric && \
    # pip install torch-cluster && \
    # pip install torch-spline-conv

CMD ["/bin/bash"]
