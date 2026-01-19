# Change: Refine Keybinding Scope and Remove Command Mode

## Why

当前设计中存在以下需要精简的问题：

1. **代码块快捷键作用域不明确** - 代码块相关快捷键（`yc`, `yNc`, `ym`, `]c`, `[c`）应该仅在聊天缓冲区（chat buffer）中生效，避免与其他区域的快捷键冲突
2. **输入区缺少 Vim 移动方式** - 输入区应该支持 Vim 风格的光标移动（如 `h`/`l` 左右移动，`w`/`b` 单词移动等），提升编辑体验
3. **命令模式过于复杂** - 初期实现不需要 Ex 风格的命令模式（`:`），可以通过快捷键和配置文件完成所有操作

## What Changes

- 限制代码块快捷键仅在聊天缓冲区 NORMAL/VISUAL 模式下生效
- 为输入区 INSERT 模式添加 Vim 风格的移动和编辑快捷键
- 移除命令模式（`:`）相关功能
- 通过快捷键替代原命令模式功能：
  - `Ctrl+t` - 新建会话
  - `Ctrl+q` - 退出应用
  - `?` - 显示帮助

## Impact

**Affected specs:**
- `vim-navigation` - 修改：移除命令模式，添加 INSERT 模式 Vim 移动支持，明确快捷键作用域

**Affected code:**
- 简化键盘事件路由逻辑
- 减少需要实现的模式状态
- 输入区组件需要支持更多 Vim 快捷键

**Removed dependencies:**
- 无需实现命令提示符组件
- 无需实现命令解析器

**Simplified scope:**
- 减少约 30% 的初始实现工作量
- 更符合"简单优先"的设计原则
