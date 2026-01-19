# Design: Remove Status Bar, Use Border Colors for Mode Indication

## Overview

本设计描述如何移除状态栏并通过边框颜色来区分不同的 Vim 模式（NORMAL、INSERT、VISUAL）。

## 视觉设计

### 当前状态（带状态栏）

```
┌─────────────────────────────────────────┐
│ NORMAL | Buffer                        │  ← 状态栏（1行）
├──────────────┬──────────────────────────┤
│ [Sessions]   │ [Chat Buffer]            │  ← 内容区域
│              │                          │
│              │                          │
├──────────────┴──────────────────────────┤
│ [Type your message...]                 │  ← 输入区
└─────────────────────────────────────────┘
```

### 新状态（无边框栏，模式通过边框颜色区分）

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

### 焦点指示（粗边框 + 青色）

无论在哪个模式下，当前焦点区域都使用：
- **ThickBorder**（粗边框）
- **Cyan color 151**（青色）

## 模式颜色方案

| 模式   | 边框颜色 | Lipgloss Color | 说明     |
| ------ | -------- | -------------- | -------- |
| NORMAL | 灰色     | 240            | 默认模式 |
| INSERT | 绿色     | 142            | 输入模式 |
| VISUAL | 蓝色     | 33             | 选择模式 |

### 颜色选择理由

- **NORMAL (灰色 240)**: 低对比度，表示"静默"状态
- **INSERT (绿色 142)**: 高对比度，表示"活动"状态（类似 Vim 的模式提示）
- **VISUAL (蓝色 33)**: 中等对比度，表示"选择"状态

## 实现架构

### 1. 样式定义 (styles.go)

```go
type Styles struct {
    // Mode colors (保留，用于边框)
    NormalMode lipgloss.Color
    InsertMode lipgloss.Color
    VisualMode lipgloss.Color

    // 移除: StatusBar lipgloss.Style

    // Pane styles - 添加模式变体
    SessionList      lipgloss.Style
    ChatBuffer       lipgloss.Style
    InputArea        lipgloss.Style
    FocusedBorder    lipgloss.Style  // 焦点时使用

    // 新增：模式边框样式
    NormalModeBorder  lipgloss.Style  // NORMAL 模式的边框
    InsertModeBorder  lipgloss.Style  // INSERT 模式的边框
    VisualModeBorder  lipgloss.Style  // VISUAL 模式的边框
}
```

### 2. 渲染逻辑 (model.go)

```go
func (m Model) View() string {
    // 不再渲染状态栏

    // 直接渲染三个区域
    sessionPane := m.renderSessionPane()
    chatPane := m.renderChatPane()
    inputPane := m.renderInputPane()

    // 合并（不再包含状态栏）
    topSection := lipgloss.JoinHorizontal(lipgloss.Top, sessionPane, chatPane)
    mainContent := lipgloss.JoinVertical(lipgloss.Left, topSection, inputPane)

    return mainContent
}

func (m Model) renderChatPane() string {
    // 根据模式和焦点选择样式
    style := m.getPaneStyle(m.Focus == ui.FocusBuffer)
    return style.Width(m.Layout.ChatBuffer.Width).Height(...).Render(content)
}

func (m Model) getPaneStyle(isFocused bool) lipgloss.Style {
    if isFocused {
        // 焦点区域：粗边框 + 青色
        return m.Styles.FocusedBorder
    }

    // 非焦点区域：根据模式选择边框颜色
    switch m.Mode {
    case vim.ModeNormal:
        return m.Styles.NormalModeBorder
    case vim.ModeInsert:
        return m.Styles.InsertModeBorder
    case vim.ModeVisual:
        return m.Styles.VisualModeBorder
    default:
        return m.Styles.NormalModeBorder
    }
}
```

### 3. 布局计算 (layout.go)

```go
type Layout struct {
    // 移除: StatusBarHeight int
    SessionList Rect
    ChatBuffer  Rect
    InputArea   Rect
}

func CalculateLayout(msg tea.WindowSizeMsg) Layout {
    width, height := msg.Width, msg.Height

    // 移除: statusBarHeight := 1

    inputHeight := 5
    contentHeight := height - inputHeight  // 不再减去 status bar

    sessionWidth := width / 5
    chatWidth := width - sessionWidth

    return Layout{
        // 不再设置 StatusBarHeight
        SessionList: Rect{X: 0, Y: 0, Width: sessionWidth, Height: contentHeight},
        ChatBuffer:  Rect{X: sessionWidth, Y: 0, Width: chatWidth, Height: contentHeight},
        InputArea:   Rect{X: 0, Y: contentHeight, Width: width, Height: inputHeight},
    }
}
```

## 边框行为

### 焦点 + 模式组合

| 状态                | 边框样式      | 边框颜色 |
| ------------------- | ------------- | -------- |
| NORMAL + 非焦点     | NormalBorder  | 灰色 240 |
| NORMAL + 焦点       | FocusedBorder | 青色 151 |
| INSERT + 非焦点     | InsertModeBorder  | 绿色 142 |
| INSERT + 焦点（输入）| FocusedBorder | 青色 151 |
| VISUAL + 非焦点     | VisualModeBorder  | 蓝色 33  |
| VISUAL + 焦点       | FocusedBorder | 青色 151 |

**关键设计决策**：
- 焦点始终使用 `FocusedBorder`（粗边框 + 青色），无论当前模式
- 非焦点区域使用模式边框颜色
- 这样用户可以同时看到：当前模式（非焦点区域颜色）+ 当前焦点（粗边框）

## 迁移路径

### 阶段 1：删除状态栏
1. 从 `model.go` 移除 `StatusBar.Render()` 调用
2. 从 `View()` 中移除 `statusBar` 变量
3. 更新 `JoinVertical` 不包含状态栏

### 阶段 2：添加模式边框
1. 在 `styles.go` 中添加 `NormalModeBorder`、`InsertModeBorder`、`VisualModeBorder`
2. 创建 `getPaneStyle()` 辅助方法
3. 更新 `renderSessionPane()`、`renderChatPane()`、`renderInputPane()` 使用新模式样式

### 阶段 3：更新布局
1. 从 `layout.go` 移除 `statusBarHeight`
2. 更新高度计算：`contentHeight = height - inputHeight`
3. 验证所有内容正确显示

### 阶段 4：清理
1. 删除 `internal/ui/statusbar.go`
2. 从 `styles.go` 移除 `StatusBar` 样式
3. 从 `model.go` 移除 `StatusBar` 字段
4. 验证构建成功

## 验证标准

变更后的 UI 应该：
1. ✅ 没有状态栏显示
2. ✅ NORMAL 模式：非焦点区域边框为灰色
3. ✅ INSERT 模式：非焦点区域边框为绿色
4. ✅ VISUAL 模式：非焦点区域边框为蓝色
5. ✅ 当前焦点区域：粗边框 + 青色
6. ✅ 内容区域高度增加 1 行（原本被状态栏占用）
7. ✅ 所有内容正确显示，无溢出

## 优势

| 方面         | 有状态栏 | 无状态栏 |
| ------------ | -------- | -------- |
| 屏幕利用率   | 1 行浪费 | +1 行内容 |
| 模式可见性   | 文字     | 颜色     |
| 视觉噪音     | 多       | 少       |
| 学习曲线     | 需阅读   | 直观     |
| 屏幕空间     | 固定占用 | 更大内容 |
