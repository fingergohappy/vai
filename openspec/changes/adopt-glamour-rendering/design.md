# Design: Markdown Rendering with Glamour

## Overview

本文档说明如何使用 Glamour 进行 Markdown 渲染，同时用正则表达式提取代码块用于复制操作。

## 架构设计

### 双轨处理

```
AI Markdown 响应
       │
       ▼
┌──────────────┐
│ 代码块提取器 │ ← 正则表达式，用于复制操作
│  (Regex)     │
└──────┬───────┘
       │
       ├─────────────┬──────────────┐
       ▼             ▼              ▼
  代码块索引      Glamour渲染    原始文本
  (用于跳转)      (用于显示)    (用于搜索)
```

### 为什么分离？

1. **Glamour 渲染** - 产生 ANSI 彩色文本，用于显示
   - 语法高亮
   - 代码块着色
   - 格式美化

2. **正则提取** - 产生纯文本代码块，用于复制
   - 不需要渲染
   - 只需要原始代码内容
   - 性能更好

## 组件设计

### 1. 代码块提取器 (CodeBlockExtractor)

**位置:** `internal/chat/extractor.go`

**职责:**
- 使用正则表达式提取 fenced code blocks
- 维护代码块索引（位置、语言、内容）
- 支持代码块编号

**正则表达式:**
```go
var codeBlockRE = regexp.MustCompile(
    "^```([a-zA-Z0-9+_-]*)?\n" +  // 语言标识 (可选)
    "([\\s\\S]*?)" +              // 代码内容
    "\n```$"                       // 结束标记
)
```

**数据结构:**
```go
type CodeBlockInfo struct {
    Number   int      // 代码块编号 (从1开始)
    Language string   // 语言标识
    Content  string   // 原始代码内容
    Start    int      // 在原始文本中的起始位置
    End      int      // 在原始文本中的结束位置
}

type CodeBlockIndex struct {
    Blocks []CodeBlockInfo
    Source string // 原始 Markdown 文本
}
```

**功能:**
- `Extract(source string) *CodeBlockIndex` - 提取所有代码块
- `GetBlock(n int) *CodeBlockInfo` - 获取第 n 个代码块
- `FindByPosition(pos int) *CodeBlockInfo` - 根据文本位置查找代码块

### 2. Glamour 渲染器 (MarkdownRenderer)

**位置:** `internal/chat/renderer.go`

**职责:**
- 封装 Glamour 渲染逻辑
- 应用自定义样式
- 缓存渲染结果

**数据结构:**
```go
type MarkdownRenderer struct {
    glamour glamour.TermRenderer
    style  glamour.Stylesheet
}

type RenderedContent struct {
    Text      string    // 渲染后的 ANSI 文本
    Lines     []string  // 按行分割的文本
    Height    int       // 渲染高度（行数）
}
```

**功能:**
- `Render(source string) *RenderedContent` - 渲染 Markdown
- `SetStyle(style string)` - 设置样式主题
- `ClearCache()` - 清除缓存

### 3. 消息渲染集成

**位置:** `internal/chat/buffer.go`

**集成方式:**
```go
type Message struct {
    ID          string
    Role        Role
    RawMarkdown string          // 原始 Markdown
    Rendered    *RenderedContent // Glamour 渲染结果
    BlockIndex  *CodeBlockIndex  // 代码块索引
    CreatedAt   time.Time
}
```

**渲染流程:**
1. 接收 AI 响应（Markdown 格式）
2. 并行处理：
   - Glamour 渲染 → `Rendered`
   - 正则提取 → `BlockIndex`
3. 显示时使用 `Rendered.Text`
4. 复制时使用 `BlockIndex.GetBlock(n).Content`

## Glamour 配置

### 样式选择

Glamour 提供多种内置样式：
- `dark` - 暗色主题
- `light` - 亮色主题
- `notty` - 无样式（纯文本）

### 自定义样式

```go
// 创建 Glamour 渲染器
renderer, _ := glamour.NewTermRenderer(
    glamour.WithStyles(glamourStyles),
    glamour.WithWordWrap(width),
)
```

### 与项目主题集成

```go
// 将项目的 Lipgloss 颜色映射到 Glamour 样式
glamourStyles := glamour.StylesheetConfig{
    Document: glamour.Style{
        StylePrimitive: style Primitive{
            BlockPrefix: "",
            BlockSuffix: "",
        },
    },
    CodeBlock: glamour.Style{
        // 使用项目的代码块颜色
        StylePrimitive: style.Primitive{
            Background: projectConfig.Colors.CodeBlockBg,
            Color:      projectConfig.Colors.CodeBlockFg,
        },
    },
    // ...
}
```

## 代码块操作实现

### 跳转到代码块 (]c / [c)

```go
func (m *Model) jumpToNextCodeBlock() {
    msg := m.currentMessage()
    if msg == nil || msg.BlockIndex == nil {
        return
    }

    next := msg.BlockIndex.Next(m.currentCodeBlock)
    if next != nil {
        m.currentCodeBlock = next.Number
        m.scrollToLine(next.Start) // 滚动到代码块位置
    }
}
```

### 复制代码块 (yc)

```go
func (m *Model) copyCurrentCodeBlock() error {
    msg := m.currentMessage()
    if msg == nil || msg.BlockIndex == nil {
        return fmt.Errorf("no code blocks")
    }

    block := msg.BlockIndex.GetBlock(m.currentCodeBlock)
    if block == nil {
        return fmt.Errorf("code block not found")
    }

    return m.clipboard.Copy(block.Content)
}
```

## 性能考虑

### 渲染缓存

```go
type RenderCache struct {
    cache map[string]*RenderedContent
    maxSize int
}

func (c *RenderCache) Get(source string) *RenderedContent {
    if cached, ok := c.cache[source]; ok {
        return cached
    }
    // 渲染并缓存
}
```

### 惰性渲染

- 只渲染可见消息
- 滚动时按需渲染
- LRU 缓存策略

### 正则性能

正则表达式提取代码块非常快：
- O(n) 时间复杂度
- 单次扫描
- 无需构建完整 AST

## 错误处理

### 格式错误的 Markdown

```go
func Extract(source string) *CodeBlockIndex {
    // 正则会自动处理格式错误
    // 未闭合的代码块会被忽略
    // 不会抛出错误
}
```

### Glamour 渲染错误

```go
renderer, err := glamour.NewTermRenderer(...)
if err != nil {
    // 回退到纯文本
    return &RenderedContent{Text: source}
}
```

## 依赖更新

```go
// go.mod
require (
    github.com/charmbracelet/glamour v0.x.x
    // ...
)
```

## 迁移路径

1. **Phase 1:** 添加 Glamour 依赖
2. **Phase 2:** 实现 `CodeBlockExtractor`（正则）
3. **Phase 3:** 实现 `MarkdownRenderer`（Glamour 封装）
4. **Phase 4:** 更新 `Message` 结构
5. **Phase 5:** 集成到 `chat/buffer.go`
6. **Phase 6:** 实现代码块操作
7. **Phase 7:** 清理 `pkg/markdown/` 占位符

## 优势总结

| 方面 | 方案 | 优势 |
|------|------|------|
| 渲染 | Glamour | 成熟、美观、与生态集成 |
| 提取 | 正则表达式 | 简单、快速、满足需求 |
| 架构 | 双轨分离 | 职责清晰、易于维护 |
