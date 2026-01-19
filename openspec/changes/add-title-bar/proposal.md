# Change: Add Title Bar for Session Name Display

## Why

当前没有专门的标题栏来显示会话信息，导致：
- 会话名称显示位置不明确
- 长会话名称可能会被截断或溢出
- 用户无法清楚看到当前正在操作的会话名称

**用户体验改进：**
- 清晰显示当前会话名称
- 长会话名称不会被截断
- 提供更好的视觉层次结构
- 屏幕顶部有明确的标题指示

## What Changes

1. **添加标题栏组件**
   - 创建 `internal/ui/titlebar.go` 用于标题栏渲染
   - 在 `internal/ui/styles.go` 中添加 TitleBar 样式

2. **添加标题栏布局**
   - 在 `internal/ui/layout.go` 中添加 `TitleBar` 布局字段
   - 设置标题栏高度为 1 行
   - 调整其他区域的 Y 位置，从标题栏下方开始

3. **更新主视图渲染**
   - 在 `internal/app/model.go` 中添加 `renderTitleBar()` 方法
   - 在 `View()` 方法中将标题栏添加到垂直布局的顶部

4. **显示会话名称**
   - 标题栏居中显示 "Sessions - [当前会话名称]"
   - 支持长会话名称的显示（不截断）

## Impact

**Affected specs:**
- `basic-ui` - 修改：添加标题栏，更新布局计算

**Affected code:**
- 新增 `internal/ui/titlebar.go` - 标题栏组件
- 修改 `internal/app/model.go` - 添加标题栏渲染
- 修改 `internal/ui/layout.go` - 添加标题栏布局计算
- 修改 `internal/ui/styles.go` - 添加标题栏样式

**不涉及：**
- 不修改模式切换逻辑
- 不修改焦点切换逻辑
- 不修改输入区域

## 依赖

- 依赖 `remove-statusbar-mode-border-colors` 变更的完成状态
- 与现有的边框颜色模式指示兼容
