FROM oven/bun as builder

WORKDIR /app
COPY . .

RUN bun install
RUN bun run build

FROM oven/bun

WORKDIR /app

COPY --from=builder /app/.next ./.next
COPY --from=builder /app/package.json ./package.json
COPY --from=builder /app/public ./public

RUN bun install --production

EXPOSE 3000

CMD ["bun", "run", "start"]
