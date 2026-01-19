# Change: Implementation MVP - Basic UI Shell First

## Why

当前 `add-tui-framework` 变更包含 133 个任务，范围太大，难以快速验证设计和架构。采用**增量实现**策略更符合开发实践：

**问题：**
- `add-tui-framework` 任务过多，难以开始
- 无法快速看到 UI 效果
- 调试困难（所有功能一起实现）
- 无法渐进式验证设计决策

**解决方案：分阶段实现**

**阶段 1：UI 框架（当前变更）**
- 画框：创建三区域布局
- 占位符：用静态内容填充各区域
- 验证：确认布局、焦点切换、模式显示正确

**阶段 2：输入功能**
- 实现真实的输入区
- 可以输入和"发送"消息
- 消息显示在聊天区

**阶段 3：导航功能**
- 实现模式切换
- 实现焦点切换
- 实现基础滚动

**阶段 4：会话功能**
- 实现会话列表
- 实现会话切换
- 实现会话持久化

**阶段 5：代码块功能**
- 实现代码块提取
- 实现代码块跳转
- 实现代码块复制

## What Changes

本变更专注于**阶段 1：UI 框架**：
- 实现三区域布局的视觉框架
- 实现状态栏（显示模式和焦点）
- 实现焦点切换（`Ctrl+w h/j/k/l`）
- 实现基础模式显示（NORMAL/INSERT）
- 用静态占位符内容填充各区域
- **不实现**真实的输入、渲染、会话等功能

## Impact

**Affected specs:**
- `basic-ui` - 新增：基础 UI 框架

**Affected code:**
- 更新 `cmd/vai/main.go` 启动 Bubble Tea
- 实现 `internal/app/model.go` 的基本 View
- 实现 `internal/ui/layout.go` 三区域布局计算
- 实现各区域的占位符 View

**不涉及：**
- 不修改 session/chat/input 等业务逻辑
- 不实现 Markdown 渲染
- 不实现代码块操作

## 依赖

- 依赖 `project-structure` spec（已归档）
- 与 `add-tui-framework` 并行存在（可选择性应用其部分内容）
