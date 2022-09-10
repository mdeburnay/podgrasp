#Step 1 - Build React App
FROM bun:lastest as react-build

WORKDIR /app

COPY ./package.json /app/

COPY . .

RUN bun install

RUN bun run build
