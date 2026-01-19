# Change: Adopt Glamour for Markdown Rendering

## Why

当前设计中，Markdown 处理需要两个不同的功能：

1. **渲染显示** - 在终端中漂亮地显示 Markdown 内容
2. **代码块提取** - 提取代码块用于复制操作（`yc`, `yNc` 等）

当前 `pkg/markdown/` 目录只是占位符，没有实际实现。

**更好的方案：**
- 使用 [Glamour](https://github.com/charmbracelet/glamour) 进行渲染（Charmbracelet 出品，与 Bubble Tea 同一生态）
- 使用正则表达式提取代码块用于复制操作

**理由：**
- Glamour 是成熟的终端 Markdown 渲染库，已经被 Glow 等项目验证
- 与 Bubble Tea/Lipgloss 完美集成
- 支持自定义样式，可以匹配项目主题
- 代码块提取不需要完整解析，正则表达式足够且性能更好

## What Changes

- 集成 Glamour 库用于 Markdown 渲染
- 实现基于正则的代码块提取器
- 更新 `internal/chat/` 包以使用 Glamour 渲染
- 添加代码块索引用于快速跳转（`]c`, `[c`）
- 移除/简化 `pkg/markdown/` 占位符

## Impact

**Affected specs:**
- `markdown-rendering` - 新增：Markdown 渲染和代码块提取
- `chat-buffer` - 修改：使用 Glamour 渲染消息

**Affected code:**
- 更新 `internal/chat/buffer.go` 使用 Glamour
- 添加 `internal/chat/renderer.go` 封装 Glamour 渲染逻辑
- 添加 `internal/chat/extractor.go` 实现代码块提取
- 更新 `internal/chat/block.go` 与渲染集成
- 更新 `go.mod` 添加 Glamour 依赖

**Dependencies:**
- `github.com/charmbracelet/glamour` - Markdown 渲染

**Removed:**
- `pkg/markdown/` 占位符实现（如果不再需要）
