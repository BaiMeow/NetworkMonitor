#!/usr/bin/env bash
set -euo pipefail

# === 路径设置（按你的仓库结构） ===
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
FRONTEND_DIR="$ROOT_DIR/frontend"
BACKEND_DIR="$ROOT_DIR/backend"
TARGET_DIR="$ROOT_DIR/NetworkMonitor_exe"
BIN_NAME="NetworkMonitor"
BIN_SRC="$BACKEND_DIR/$BIN_NAME"
BIN_DST="$TARGET_DIR/$BIN_NAME"
LOG_FILE="$TARGET_DIR/${BIN_NAME}.log"

# === 工具检查 ===
command -v pnpm >/dev/null 2>&1 || { echo "ERROR: 未安装 pnpm"; exit 1; }
command -v go   >/dev/null 2>&1 || { echo "ERROR: 未安装 Go"; exit 1; }

echo "== 1/5: 构建前端 =="
pushd "$FRONTEND_DIR" >/dev/null
# 如果没有 node_modules 就自动安装
if [ ! -d node_modules ]; then
  echo "   • 检测到缺少 node_modules，执行 pnpm install ..."
  pnpm install
fi
pnpm build
popd >/dev/null

echo "== 2/5: 构建后端（Go） =="
pushd "$BACKEND_DIR" >/dev/null
# 如需静态编译可用：CGO_ENABLED=0
go build -o "$BIN_NAME" .
popd >/dev/null

echo "== 3/5: 覆盖安装二进制 =="
mkdir -p "$TARGET_DIR"
cp -f "$BIN_SRC" "$BIN_DST"
chmod +x "$BIN_DST"

echo "== 4/5: 检测并停止已运行进程 =="


echo "== 5/5: 后台启动新进程 =="
# 日志追加到 LOG_FILE
nohup "$BIN_DST" >> "$LOG_FILE" 2>&1 &
NEW_PID=$!
echo "   • 启动完成，PID=$NEW_PID"
echo "   • 日志：$LOG_FILE"

echo "✅ 完成"
