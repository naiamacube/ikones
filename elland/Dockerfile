FROM pytorch/pytorch:1.13.0-cuda11.6-cudnn8-devel

RUN apt update

RUN pip install pytorch-lightning

RUN pip install torch-scatter && \
    pip install torch-sparse && \
    pip install torch-geometric && \
    pip install torch-cluster && \
    pip install torch-spline-conv
