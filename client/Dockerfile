FROM jarredsumner/bun:edge

WORKDIR /usr/src/app

COPY *.package.json ./
COPY . .

RUN bun install

COPY . .

RUN bun build

EXPOSE 3030
