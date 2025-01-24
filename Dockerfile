# デプロイ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.23.5-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Dockerfileの「COPY . .」は、現在のディレクトリのすべてのファイルとフォルダをDockerイメージの作業ディレクトリにコピー
COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# ---

# デプロイ用コンテナ
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# --

# ローカル開発環境で利用するホットリロード環境（ホットリロードを実現するOSSがairだそう）
FROM golang:1.23.5 as dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
