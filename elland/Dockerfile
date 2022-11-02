FROM ubuntu:22.04

LABEL com.naiama.image.author="mlops@naiama.com"

RUN apt update && apt install -y build-essential curl

ENV FILENAME Anaconda3-2022.10-Linux-x86_64.sh
RUN curl -O https://repo.anaconda.com/archive/$FILENAME && shasum -a256 $FILENAME | awk '$1=="e7ecbccbc197ebd7e1f211c59df2e37bc6959d081f2235d387e08c9026666acd"{print checksum-ok}' && \
    bash $FILENAME -bf && eval "$($HOME/anaconda3/bin/conda shell.bash hook)" && conda init

# There is no maintained binary for PyTorch CPU-only, so we have to compile from source
RUN $HOME/anaconda3/bin/conda install pytorch torchvision torchaudio cpuonly -c pytorch

RUN $HOME/anaconda3/bin/conda install pytorch-lightning -c conda-forge

# RUN pip install torch-scatter && \
    # pip install torch-sparse && \
    # pip install torch-geometric && \
    # pip install torch-cluster && \
    # pip install torch-spline-conv

CMD ["/bin/bash"]
