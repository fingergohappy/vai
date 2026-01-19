# Change: Implement Vim-style TUI AI Chat Framework

## Why

网页端 AI 聊天工具存在以下痛点：
- 代码块选择和复制体验差，多段回答难以精准操作
- 鼠标操作打断工程师的键盘工作流
- 需要频繁在终端和浏览器之间切换上下文

工程师的真实工作环境在终端（Vim/tmux/shell），需要一个**键盘优先、以高效查看和精准复制为核心价值**的 TUI AI 聊天工具。

## What Changes

- 实现基于 Bubble Tea 的 TUI 框架基础架构
- 构建三区域布局：左侧对话历史列表、中间聊天内容缓冲区、底部输入区
- 实现 Vim 风格的导航模式系统（NORMAL/INSERT/VISUAL）
- 实现结构化的消息渲染系统（支持文本块和代码块）
- 实现焦点切换和区域间导航

## Impact

**Affected specs:**
- `tui-framework` - 新增：TUI 框架核心架构
- `chat-buffer` - 新增：聊天内容显示和交互
- `session-manager` - 新增：会话管理和历史记录
- `vim-navigation` - 新增：Vim 风格导航和模式系统

**Affected code:**
- 新增 Go 模块结构和 Bubble Tea 应用入口
- 新增内部 UI 组件包（tui/）
- 新增状态管理模型
- 新增键盘事件处理和路由系统

**Dependencies:**
- `charmbracelet/bubbletea` - TUI 框架
- `charmbracelet/lipgloss` - 样式系统
- `charmbracelet/bubbles` - 预构建组件（textarea, viewport 等）
