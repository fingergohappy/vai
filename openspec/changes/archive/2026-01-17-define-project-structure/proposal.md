# Change: Define Project Directory Structure

## Why

当前项目只有 `go.mod` 文件，需要定义清晰的目录结构来支持 TUI 应用的开发。基于设计文档中的 Bubble Tea 架构和组件划分，需要一个符合 Go 最佳实践的目录结构。

设计文档中提到的关键组件：
- 顶层 Model（Mode, Focus, 子 Model 组合）
- 子 Model（History, Buffer, Input）
- Vim 模式系统
- UI 组件
- 会话管理
- 配置系统

## What Changes

定义完整的项目目录结构，包括：
- 入口点 (`cmd/vai/`)
- 内部包 (`internal/`)
- 按功能域划分的子包
- 测试目录结构
- 配置和资源文件位置

## Impact

**Affected specs:**
- `project-structure` - 新增：项目目录结构和组织规范

**Affected code:**
- 创建所有必需的目录
- 添加每个包的骨架代码
- 设置 go.work 等工具配置

**Dependencies:**
- 无新增依赖
- 确保目录结构符合 Go 标准项目布局
