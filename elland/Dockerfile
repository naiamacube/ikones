FROM pytorch/pytorch:1.13.0-cuda11.6-cudnn8-devel

RUN apt update

RUN pip install pytorch-lightning

RUN export PATH=/usr/local/cuda/bin:$PATH && \
    export CPATH=/usr/local/cuda/include:$CPATH && \
    export LD_LIBRARY_PATH=/usr/local/cuda/lib64:$LD_LIBRARY_PATH && \
    export DYLD_LIBRARY_PATH=/usr/local/cuda/lib:$DYLD_LIBRARY_PATH && \
    pip install torch-scatter && \
    pip install torch-sparse && \
    pip install torch-geometric && \
    pip install torch-cluster && \
    pip install torch-spline-conv
