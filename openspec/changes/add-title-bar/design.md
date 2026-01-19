# Design: Add Title Bar for Session Name Display

## Overview

本设计描述如何在界面顶部添加一个标题栏，用于显示当前会话名称，格式为 "Sessions - [会话名称]"，居中对齐。

## 视觉设计

### 当前状态（无标题栏）

```
┌──────────────┬──────────────────────────┐
│ [Sessions]   │ [Chat Buffer]            │  ← 边框颜色表示模式
│              │                          │     灰色=NORMAL
│              │                          │     绿色=INSERT
│              │                          │     蓝色=VISUAL
├──────────────┴──────────────────────────┤
│ [Type your message...]                 │  ← 输入区
└─────────────────────────────────────────┘
```

### 新状态（带标题栏）

```
┌─────────────────────────────────────────┐
│      Sessions - New Chat                │  ← 标题栏（1行，居中）
├──────────────┬──────────────────────────┤
│ [Sessions]   │ [Chat Buffer]            │  ← 内容区域
│              │                          │
│              │                          │
├──────────────┴──────────────────────────┤
│ [Type your message...]                 │  ← 输入区
└─────────────────────────────────────────┘
```

### 标题栏设计细节

- **位置**: 屏幕最顶部，横跨整个宽度
- **高度**: 1 行
- **对齐**: 居中对齐
- **内容**: "Sessions - [当前会话名称]"
- **样式**: 粗体，深色背景，浅色文字

## 实现架构

### 1. 标题栏组件 (titlebar.go)

```go
package ui

import (
    "github.com/charmbracelet/lipgloss"
)

// TitleBar renders the title bar showing session information.
type TitleBar struct {
    styles *Styles
}

// NewTitleBar creates a new title bar component.
func NewTitleBar(styles *Styles) *TitleBar {
    if styles == nil {
        styles = DefaultStyles()
    }
    return &TitleBar{
        styles: styles,
    }
}

// Render renders the title bar with the given session title.
func (t *TitleBar) Render(sessionTitle string) string {
    title := "Sessions - " + sessionTitle
    return t.styles.TitleBar.Render(title)
}

// SetWidth sets the width of the title bar.
func (t *TitleBar) SetWidth(width int) {
    t.styles.TitleBar = t.styles.TitleBar.Width(width)
}
```

### 2. 样式定义 (styles.go)

```go
type Styles struct {
    // ... existing fields ...

    // Title bar
    TitleBar lipgloss.Style
}

func DefaultStyles() *Styles {
    return &Styles{
        // ... existing styles ...

        // Title bar
        TitleBar: lipgloss.NewStyle().
            Bold(true).
            Align(lipgloss.Center).              // 居中对齐
            Foreground(lipgloss.Color("252")).   // 白色文字
            Background(lipgloss.Color("235")),   // 深色背景
    }
}
```

### 3. 布局计算 (layout.go)

```go
type Layout struct {
    Width  int
    Height int

    // NEW: Title bar layout
    TitleBar PaneLayout

    // Existing panes...
    SessionList PaneLayout
    ChatBuffer  PaneLayout
    InputArea   PaneLayout
}

func CalculateLayout(msg tea.WindowSizeMsg) Layout {
    width := msg.Width
    height := msg.Height

    titleBarHeight := 1   // NEW: 标题栏高度
    inputHeight := 5

    // Session list: 20% width
    sessionWidth := width * 20 / 100
    if sessionWidth < 20 {
        sessionWidth = 20
    }
    if sessionWidth > width-40 {
        sessionWidth = width / 3
    }

    // Chat buffer: remaining width
    chatWidth := width - sessionWidth

    // Content height (total - title bar - input)
    contentHeight := height - titleBarHeight - inputHeight

    return Layout{
        Width:  width,
        Height: height,
        TitleBar: PaneLayout{        // NEW
            X:      0,
            Y:      0,
            Width:  width,
            Height: titleBarHeight,
        },
        SessionList: PaneLayout{
            X:      0,
            Y:      titleBarHeight,  // 从标题栏下方开始
            Width:  sessionWidth,
            Height: contentHeight,
        },
        ChatBuffer: PaneLayout{
            X:      sessionWidth,
            Y:      titleBarHeight,  // 从标题栏下方开始
            Width:  chatWidth,
            Height: contentHeight,
        },
        InputArea: PaneLayout{
            X:      0,
            Y:      titleBarHeight + contentHeight,
            Width:  width,
            Height: inputHeight,
        },
    }
}
```

### 4. 渲染逻辑 (model.go)

```go
type Model struct {
    // ... existing fields ...

    // UI components
    TitleBar  *ui.TitleBar  // NEW: Title bar component
    Layout    ui.Layout
    Styles    *ui.Styles
}

func NewModel(cfg config.Config) Model {
    styles := ui.DefaultStyles()
    titleBar := ui.NewTitleBar(styles)

    return Model{
        // ... existing fields ...
        Styles:   styles,
        TitleBar: titleBar,
    }
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.Layout = ui.CalculateLayout(msg)
        m.TitleBar.SetWidth(msg.Width)  // NEW: 设置标题栏宽度
        m.ready = true
        m.Input.SetSize(m.Layout.InputArea.Width, m.Layout.InputArea.Height)
    }
    // ... rest of update logic ...
}

func (m Model) View() string {
    if !m.ready {
        return "Initializing vai..."
    }

    // NEW: 渲染标题栏
    titleBar := m.renderTitleBar()

    sessionPane := m.renderSessionPane()
    chatPane := m.renderChatPane()
    inputPane := m.renderInputPane()

    topSection := lipgloss.JoinHorizontal(
        lipgloss.Left,
        sessionPane,
        chatPane,
    )

    // NEW: 将标题栏加入垂直布局
    mainContent := lipgloss.JoinVertical(
        lipgloss.Top,
        titleBar,    // 标题栏在顶部
        topSection,
        inputPane,
    )

    return mainContent
}

// NEW: 渲染标题栏
func (m Model) renderTitleBar() string {
    // 获取当前会话标题（从 Session 子模型）
    currentTitle := m.Session.GetCurrentTitle()
    if currentTitle == "" {
        currentTitle = "New Chat"
    }
    return m.TitleBar.Render(currentTitle)
}
```

### 5. Session Model 更新

需要在 `internal/session/` 中添加获取当前会话标题的方法：

```go
// GetCurrentTitle returns the title of the current session.
func (m Model) GetCurrentTitle() string {
    if len(m.Sessions) == 0 {
        return "New Chat"
    }
    if m.currentIndex >= 0 && m.currentIndex < len(m.Sessions) {
        return m.Sessions[m.currentIndex].Title
    }
    return "New Chat"
}
```

## 布局计算详解

### 高度计算

```
总高度 = titleBarHeight + contentHeight + inputHeight
       = 1          + (H - 1 - 5)    + 5
       = H
```

### Y 坐标分配

| 组件       | Y 坐标 | 说明                     |
| ---------- | ------ | ------------------------ |
| TitleBar   | 0      | 从顶部开始               |
| SessionList| 1      | 标题栏下方               |
| ChatBuffer | 1      | 标题栏下方               |
| InputArea  | 1 + contentHeight | 内容区域下方 |

## 颜色方案

| 元素     | 背景颜色 | 前景颜色 | 说明           |
| -------- | -------- | -------- | -------------- |
| TitleBar | 235      | 252      | 深灰背景，白色文字 |
| Title文本 | -        | 252      | 粗体，居中     |

## 长会话名称处理

标题栏使用 `Align(lipgloss.Center)` 居中对齐：
- 如果会话名称过长，会自然扩展
- Lipgloss 会自动处理文本渲染
- 不会截断文本，确保完整显示

## 迁移路径

### 阶段 1：创建标题栏组件
1. 创建 `internal/ui/titlebar.go`
2. 在 `styles.go` 中添加 `TitleBar` 样式
3. 实现 `NewTitleBar()` 和 `Render()` 方法

### 阶段 2：更新布局
1. 在 `layout.go` 中添加 `TitleBar PaneLayout`
2. 添加 `titleBarHeight := 1`
3. 调整 contentHeight 计算
4. 更新所有区域的 Y 坐标

### 阶段 3：集成到主视图
1. 在 `Model` 结构体中添加 `TitleBar` 字段
2. 在 `NewModel()` 中初始化标题栏
3. 在 `Update()` 中设置标题栏宽度
4. 添加 `renderTitleBar()` 方法
5. 在 `View()` 中将标题栏加入布局

### 阶段 4：获取会话标题
1. 在 `session.Model` 中添加 `GetCurrentTitle()` 方法
2. 在 `renderTitleBar()` 中调用此方法获取标题

### 阶段 5：验证
1. 构建项目
2. 运行应用
3. 验证标题栏正确显示
4. 验证布局正确
5. 测试长会话名称显示

## 验证标准

变更后的 UI 应该：
1. ✅ 顶部显示标题栏
2. ✅ 标题栏居中显示 "Sessions - [会话名称]"
3. ✅ 标题栏高度为 1 行
4. ✅ 内容区域从标题栏下方开始
5. ✅ 长会话名称完整显示，不被截断
6. ✅ 与现有的模式边框颜色兼容
7. ✅ 终端大小时标题栏宽度自动调整

## 优势

| 方面       | 无标题栏 | 有标题栏 |
| ---------- | -------- | -------- |
| 会话可见性 | 不清晰   | 清晰     |
| 视觉层次   | 扁平     | 分层     |
| 长标题支持 | 可能溢出 | 完整显示 |
| 用户认知   | 需猜测   | 明确     |
