# 使用基本的 Ubuntu 镜像作为基础
FROM mcr.microsoft.com/devcontainers/base:jammy

# Prepare apt
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive

# Install Go
ENV PATH="/usr/local/go/bin:$PATH"
RUN export GO_VERSION="1.23.3" \
    && wget "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" \
    && tar -C /usr/local -xzf "go${GO_VERSION}.linux-amd64.tar.gz" \
    && rm "go${GO_VERSION}.linux-amd64.tar.gz"

# 设置默认工作目录
WORKDIR /workspace
