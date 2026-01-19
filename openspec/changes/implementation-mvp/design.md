# Design: Incremental Implementation Strategy

## Overview

本文档定义 vai 项目的增量实现策略，采用"先画框，再填充"的方法。

## 实现阶段

### 阶段 1：UI 框架（当前变更）

**目标：** 画出 UI 轮廓，验证布局和交互基础

**实现内容：**
```
┌─────────────────────────────────────────┐
│ NORMAL | Buffer                        │  ← 状态栏
├──────────────┬──────────────────────────┤
│ [Sessions]   │ [Chat Buffer]            │  ← 三区域
│              │                          │     框架
│ - Session 1  │ - Welcome to vai!        │
│ - Session 2  │ - This is a placeholder │
│              │                          │
├──────────────┴──────────────────────────┤
│ [Type your message...]                 │  ← 输入区占位符
└─────────────────────────────────────────┘
```

**功能清单：**
- ✅ 三区域布局渲染
- ✅ 状态栏显示（模式 + 焦点）
- ✅ 焦点切换（`Ctrl+w h/j/k/l`）
- ✅ 模式显示（NORMAL/INSERT）
- ✅ 静态占位符内容
- ❌ 真实输入
- ❌ 真实消息
- ❌ 真实会话

**验收标准：**
- 启动后能看到完整 UI 框架
- 可以在三个区域间切换焦点
- 状态栏正确显示当前焦点
- 按 `Ctrl+q` 可以退出

---

### 阶段 2：真实输入

**目标：** 可以输入和"发送"消息

**实现内容：**
- 真实的输入框（bubbles.TextArea）
- Enter 发送消息
- 消息显示在聊天区（纯文本，无 Markdown）
- 输入区清空

**新增功能：**
- ✅ 真实文本输入
- ✅ Enter 发送
- ✅ 消息历史存储（内存中）
- ❌ Markdown 渲染
- ❌ 会话持久化

---

### 阶段 3：基础导航

**目标：** 实现模式系统和基础滚动

**实现内容：**
- NORMAL/INSERT 模式切换
- 聊天区滚动（`j`/`k`）
- 模式键路由基础

**新增功能：**
- ✅ `i`/`a` 进入 INSERT 模式
- ✅ `Esc` 返回 NORMAL 模式
- ✅ `j`/`k` 滚动聊天区
- ❌ 代码块导航
- ❌ VISUAL 模式

---

### 阶段 4：会话管理

**目标：** 多会话支持

**实现内容：**
- 会话列表真实数据
- 创建新会话
- 切换会话
- 会话持久化（JSON 文件）

**新增功能：**
- ✅ 会话列表显示真实会话
- ✅ `Ctrl+t` 创建新会话
- ✅ `Enter` 切换会话
- ❌ 会话重命名
- ❌ 会话删除

---

### 阶段 5：Markdown 渲染

**目标：** 美化消息显示

**实现内容：**
- 集成 Glamour
- 渲染 Markdown
- 代码块语法高亮

**新增功能：**
- ✅ Markdown 渲染
- ✅ 代码块高亮
- ❌ 代码块导航
- ❌ 代码块复制

---

### 阶段 6：代码块功能

**目标：** 代码块操作

**实现内容：**
- 正则提取代码块
- `]c`/`[c` 跳转
- `yc`/`yNc` 复制
- VISUAL 模式选择

**新增功能：**
- ✅ 代码块索引
- ✅ `]c`/`[c` 跳转
- ✅ `yc` 复制代码块
- ✅ VISUAL 模式
- ✅ `y` 复制选择

---

## 架构演进

### 阶段 1 架构（当前）

```
Model (top-level)
├── Mode (enum)
├── Focus (enum)
├── SessionList (placeholder Model)
├── ChatBuffer (placeholder Model)
└── InputArea (placeholder Model)
```

### 最终架构

```
Model (top-level)
├── Mode (enum)
├── Focus (enum)
├── SessionList (real Model)
│   ├── Sessions []Session
│   └── Storage
├── ChatBuffer (real Model)
│   ├── Messages []Message
│   ├── Viewport
│   ├── Renderer (Glamour)
│   └── Extractor (regex)
└── InputArea (real Model)
    ├── TextArea (bubbles)
    └── VimMotion
```

## 占位符策略

### 阶段 1 占位符实现

```go
// SessionList placeholder
func (m SessionList) View() string {
    return `
┌──────────────┐
│ [Sessions]   │
│              │
│ - Session 1  │
│ - Session 2  │
│              │
│ (TODO: impl)│
└──────────────┘`
`
}

// ChatBuffer placeholder
func (m ChatBuffer) View() string {
    return `
┌──────────────────────────┐
│ [Chat Buffer]            │
│                          │
│ Welcome to vai!          │
│                          │
│ This is a placeholder.   │
│ Real implementation      │
│ coming soon...           │
└──────────────────────────┘`
`
}
```

### 逐步替换占位符

```go
// 阶段 2：真实输入
type InputArea struct {
    textarea textarea.Model  // 真实组件
}

// 阶段 4：真实会话
type SessionList struct {
    sessions []Session       // 真实数据
    selected  int
}

// 阶段 5：真实渲染
type ChatBuffer struct {
    messages  []Message
    renderer  *glamour.TermRenderer
}
```

## 验证策略

### 每个阶段独立验证

| 阶段 | 可运行 | 测试内容 | 验证方式 |
|------|--------|----------|----------|
| 1 | ✅ | UI 框架显示 | 运行 `vai` 看到界面 |
| 2 | ✅ | 输入消息 | 输入 "hello" 看到显示 |
| 3 | ✅ | 滚动导航 | `j`/`k` 滚动，`i` 进入输入 |
| 4 | ✅ | 多会话 | `Ctrl+t` 新会话，切换 |
| 5 | ✅ | 美化显示 | 发送 Markdown 看高亮 |
| 6 | ✅ | 代码操作 | `]c` 跳转，`yc` 复制 |

### 持续集成

每个阶段完成后：
1. 运行 `go build ./...` 确保编译
2. 运行 `go test ./...` 确保测试通过
3. 手动测试新功能
4. Git 提交（标记为 `phase-N`）

## 优势

| 方面 | 一次性实现 | 增量实现 |
|------|------------|----------|
| 验证时间 | 133 个任务完成 | 6 次验证（每阶段） |
| 调试难度 | 难以定位问题 | 每次变更少，易调试 |
| 反馈速度 | 看到结果需要很久 | 每阶段都有可见结果 |
| 风险控制 | 一个 bug 阻塞全部 | 问题隔离在单个阶段 |
| 心理负担 | 压力大 | 每个小胜利 |

## 实现顺序建议

**推荐顺序：**
1. **阶段 1（UI 框架）** - 最重要，建立基础
2. **阶段 2（真实输入）** - 体验核心功能
3. **阶段 3（基础导航）** - 交互基础
4. **阶段 4（会话管理）** - 完整体验
5. **阶段 5（Markdown）** - 美化显示
6. **阶段 6（代码块）** - 核心价值

**可并行：**
- 阶段 4 和 5 可以互换
- 阶段 6 可分两步：先提取，再复制

## 下一步

当前提案专注于**阶段 1**，后续阶段通过独立变更提案实现。
