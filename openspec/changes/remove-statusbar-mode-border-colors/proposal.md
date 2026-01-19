# Change: Remove Status Bar, Use Border Colors for Mode Indication

## Why

当前状态栏占用屏幕空间，但提供的价值有限：
- 状态栏显示模式和焦点信息，但占用一行宝贵的屏幕空间
- 模式信息可以通过更直观的方式传递（边框颜色）
- 移除状态栏后，内容区域可以增加一行，提升可用性

**用户体验改进：**
- 更多的内容显示空间
- 模式切换更直观（通过颜色而非文字）
- 简化界面，减少视觉噪音

## What Changes

1. **移除状态栏组件**
   - 删除 `internal/ui/statusbar.go`
   - 从 `internal/app/model.go` 中移除状态栏渲染
   - 从 `internal/ui/styles.go` 中移除 StatusBar 样式

2. **通过边框颜色区分模式**
   - NORMAL 模式：默认边框颜色（灰色 240）
   - INSERT 模式：绿色边框（142）
   - VISUAL 模式：蓝色边框（33）

3. **更新布局计算**
   - `internal/ui/layout.go` 中移除 `statusBarHeight`
   - 所有高度计算不再减去状态栏高度

4. **保留焦点指示**
   - 当前焦点区域使用粗边框（ThickBorder + 青色 151）
   - 非焦点区域使用普通边框

## Impact

**Affected specs:**
- `basic-ui` - 修改：移除状态栏，添加模式边框颜色

**Affected code:**
- 删除 `internal/ui/statusbar.go`
- 修改 `internal/app/model.go` - 移除状态栏渲染，添加模式边框颜色
- 修改 `internal/ui/layout.go` - 移除状态栏高度计算
- 修改 `internal/ui/styles.go` - 添加模式边框样式，移除 StatusBar

**不涉及：**
- 不修改模式切换逻辑
- 不修改焦点切换逻辑
- 不修改子模型（session、chat、input）

## 依赖

- 依赖 `implementation-mvp` 变更的当前状态
- 与现有功能兼容
