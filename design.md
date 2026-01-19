# 终端 Vim 风格 AI 聊天工具 —— 设计思路整理

> 本文不是 MVP 方案，而是对当前**完整设计思路的系统化整理**，用于后续交给 AI / 人类实现、评估复杂度、拆解模块。

---

## 1. 设计动机（Why）

### 1.1 现有问题

- **网页端 AI 聊天复制体验差**
  - 代码块选择困难
  - 多段回答难以精准复制
  - 鼠标操作打断键盘流

- **工程师真实工作环境在终端**
  - Vim / tmux / shell
  - 键盘优先
  - 不希望频繁切换到浏览器

- API Key 模式存在问题
  - 成本敏感
  - 能力与网页版本存在体感差距

### 1.2 目标

构建一个：

- **终端内使用的 AI 聊天工具（TUI）**
- **强 Vim 操作范式**
- **以“高效查看 + 精准复制”为核心价值**

不是聊天玩具，而是**工程工具**。

---

## 2. 核心设计原则（Principles）

1. **键盘优先（Keyboard First）**
2. **Vim 思维，而不是 Vim 模仿**
3. **状态显式、行为可预测**
4. **复制是第一等公民**
5. **模型解耦、状态清晰**

---

## 3. 整体 UI 布局设计

### 3.1 总体布局

```
+--------------------------------------------------+
| 状态栏： NORMAL | INSERT | VISUAL | COPY        |
+----------------------+---------------------------+
| 对话列表（历史）     | 当前对话内容（Buffer）   |
|                      |                           |
|                      |                           |
+----------------------+---------------------------+
| 输入区（Prompt）                                 |
+--------------------------------------------------+
```

### 3.2 区域说明

#### 左栏：对话列表（Session / History）

- 显示历史会话
- 支持上下选择
- 可切换当前对话

#### 右栏：对话 Buffer

- 类似 Vim buffer / less / pager
- 只读为主
- 支持滚动、搜索、选择、复制

#### 底部：输入区

- 插入模式输入
- 回车发送

#### 顶部：状态栏

- 显示 Vim Mode
- 显示当前焦点区域

---

## 4. Vim 模式设计

### 4.1 模式划分

| 模式 | 用途 |
|----|----|
| NORMAL | 浏览 / 移动 / 切换焦点 |
| INSERT | 输入 Prompt |
| VISUAL | 选择文本 |
| COPY | 针对代码块 /消息复制 |

> **不是所有区域支持所有模式**

### 4.2 区域 × 模式矩阵

| 区域 | NORMAL | INSERT | VISUAL |
|----|----|----|----|
| 对话列表 | ✅ | ❌ | ❌ |
| 对话 Buffer | ✅ | ❌ | ✅ |
| 输入区 | ❌ | ✅ | ❌ |

---

## 5. 焦点与窗口切换

### 5.1 焦点区域

- `HistoryPane`
- `ChatBufferPane`
- `InputPane`

### 5.2 切换方式（类 tmux / vim）

- `Ctrl-w h / l / j / k`
- 焦点只影响 **键盘事件路由**，不直接影响 Mode

---

## 6. ChatBuffer 设计（重点）

### 6.1 内容结构

```go
type Message struct {
    Role string // user / ai
    Blocks []Block
}

type Block interface {
    Kind() BlockType
}

type TextBlock struct { Text string }

type CodeBlock struct {
    Lang string
    Lines []string
}
```

**结论：ChatBuffer 必须是结构化数据，而不是纯字符串**。

---

### 6.2 多代码块问题的解决方案

一个回答中可能有多个代码块：

#### 解决策略：

1. **代码块编号**

```
[1] ```go
    ...
    ```

[2] ```bash
    ...
    ```
```

2. Vim 风格跳转

- `]c` / `[c`：跳转下一个 / 上一个 code block
- `yc`：复制当前 code block
- `y2c`：复制第 2 个 code block

3. VISUAL 模式选择

- `v` → 移动 → `y`

---

## 7. 复制模型（核心价值）

### 7.1 复制对象

- 当前行
- 选中区域
- 整个代码块
- 整条消息

### 7.2 实现原则

- **复制不依赖终端鼠标选择**
- 内部维护 selection range
- 使用系统剪贴板（pbcopy / wl-copy / xclip）

---

## 8. Bubble Tea 架构映射

### 8.1 顶层 Model

```go
type Model struct {
    Mode   vim.Mode
    Focus  ui.Focus

    History history.Model
    Buffer  chatbuffer.Model
    Input   input.Model
}
```

### 8.2 Update 分发模型

- 顶层 Model 只做：
  - Msg 分发
  - 子 Model 合并
  - Cmd 聚合

子 Model：
- 自治
- 无直接依赖
- 通过 Msg 间接通信

---

## 9. 性能与长对话处理

### 9.1 潜在问题

- 对话记录极长
- 全量渲染会卡顿

### 9.2 应对策略

- ChatBuffer 使用 **viewport / windowed rendering**
- 只渲染可见行
- 历史消息分页
- 结构化数据避免反复解析 markdown

---

## 10. 关于 Markdown 渲染

### 10.1 结论

- **Bubble Tea / TUI 不适合完整 Markdown 渲染**
- 但适合：
  - 标题弱化
  - code block 强化
  - 缩进 / 高亮

### 10.2 实践建议

- Markdown → AST（一次）
- AST → Block
- TUI 渲染 Block

---

## 11. 产品价值判断（严格）

### 11.1 适合人群

- 重度终端用户
- Vim 用户
- 工程师 / 运维 / 后端

### 11.2 核心不可替代点

- 键盘驱动的复制体验
- 精准代码块操作
- 无浏览器上下文切换

> **价值不是“能聊天”，而是“能高效用 AI 工作”**

---

## 12. 定位总结

一句话：

> **这是一个 Vim 思维的 AI 对话与代码消费工具，而不是 ChatGPT 的终端壳。**

---

（文档结束）

