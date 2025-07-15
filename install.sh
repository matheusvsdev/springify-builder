#!/bin/bash
set -e

BINARY_NAME="springify"
REPO_OWNER="matheusvaldevino"
REPO_NAME="springify"

echo "Iniciando instalação do $BINARY_NAME..."

# Detectar sistema operacional
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Ajustar arquitetura para Go
if [ "$ARCH" = "x86_64" ]; then
  ARCH="amd64"
elif [ "$ARCH" = "arm64" ] || [ "$ARCH" = "aarch64" ]; then
  ARCH="arm64"
else
  echo "Arquitetura não suportada: $ARCH"
  exit 1
fi

# Montar nome do arquivo
ARCHIVE_NAME="${BINARY_NAME}-${OS}-${ARCH}"
TAR_FILE="${ARCHIVE_NAME}.tar.gz"
DOWNLOAD_URL="https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/latest/download/${TAR_FILE}"

echo "⬇Baixando ${TAR_FILE}..."
curl -LO "$DOWNLOAD_URL"

# Extrair, tornar executável e instalar
tar -xzf "$TAR_FILE"
chmod +x "$BINARY_NAME"
sudo mv "$BINARY_NAME" /usr/local/bin/$BINARY_NAME

# Validar instalação
echo "Verificando instalação..."
if command -v $BINARY_NAME >/dev/null 2>&1; then
  echo "Instalação concluída com sucesso!"
  $BINARY_NAME --help
else
  echo "O binário não foi detectado após a instalação."
  exit 1
fi
