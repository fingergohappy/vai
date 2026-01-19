# Tasks: Adopt Glamour for Markdown Rendering

This document outlines the tasks for integrating Glamour for Markdown rendering and implementing regex-based code block extraction.

## Task Summary

- Add Glamour dependency
- Implement regex-based code block extractor
- Implement Glamour renderer wrapper
- Update chat buffer to use both systems
- Implement code block navigation and copying
- Clean up placeholder code

---

## Phase 1: Dependency Setup

### 1.1 Add Glamour dependency
- [ ] Run `go get github.com/charmbracelet/glamour@latest`
- [ ] Run `go mod tidy`

**Validation:** Glamour is in go.mod and builds successfully

### 1.2 Verify Glamour basic usage
- [ ] Create simple test program that renders Markdown with Glamour
- [ ] Verify ANSI output works in terminal

**Validation:** Can render "```go\nfmt.Println(\"Hello\")\n```" with syntax highlighting

---

## Phase 2: Code Block Extractor (Regex)

### 2.1 Create extractor structure
- [ ] Create `internal/chat/extractor.go`
- [ ] Define `CodeBlockInfo` struct (Number, Language, Content, Start, End)
- [ ] Define `CodeBlockIndex` struct (Blocks, Source)

**Validation:** Types compile successfully

### 2.2 Implement regex extraction
- [ ] Compile regex pattern for fenced code blocks: `^```(\w*)\n([\s\S]*?)\n```$`
- [ ] Implement `Extract(source string) *CodeBlockIndex` function
- [ ] Handle code blocks without language identifier
- [ ] Ignore malformed/incomplete code blocks

**Validation:** Can extract code blocks from test Markdown

### 2.3 Implement index queries
- [ ] Implement `GetBlock(n int) *CodeBlockInfo` method
- [ ] Implement `FindByPosition(pos int) *CodeBlockInfo` method
- [ ] Implement `Next(current int) *CodeBlockInfo` method
- [ ] Implement `Prev(current int) *CodeBlockInfo` method

**Validation:** Can query code blocks by number and position

### 2.4 Add tests
- [ ] Test extraction with multiple code blocks
- [ ] Test extraction with no language identifier
- [ ] Test extraction with malformed blocks
- [ ] Test boundary conditions (empty, no blocks)

**Validation:** All tests pass

---

## Phase 3: Glamour Renderer Wrapper

### 3.1 Create renderer structure
- [ ] Create `internal/chat/renderer.go`
- [ ] Define `MarkdownRenderer` struct
- [ ] Define `RenderedContent` struct (Text, Lines, Height)

**Validation:** Types compile successfully

### 3.2 Implement Glamour initialization
- [ ] Implement `NewRenderer()` function
- [ ] Create Glamour TermRenderer with default style
- [ ] Implement error handling (fall back to plain text)

**Validation:** Can create renderer without errors

### 3.3 Implement render function
- [ ] Implement `Render(source string) *RenderedContent` function
- [ ] Handle word wrap based on width
- [ ] Split rendered text into lines for viewport

**Validation:** Can render test Markdown to ANSI text

### 3.4 Add style customization
- [ ] Implement `SetStyle(style string)` function
- [ ] Map project theme colors to Glamour styles
- [ ] Support "dark", "light", and "auto" styles

**Validation:** Can switch between dark and light themes

### 3.5 Add caching
- [ ] Implement simple render cache (map[string]*RenderedContent)
- [ ] Add cache size limit
- [ ] Implement cache clearing

**Validation:** Repeated renders use cache

---

## Phase 4: Message Structure Update

### 4.1 Update Message type
- [ ] Add `RawMarkdown string` field
- [ ] Add `Rendered *RenderedContent` field
- [ ] Add `BlockIndex *CodeBlockIndex` field
- [ ] Keep existing `Blocks []Block` for compatibility or remove if not needed

**Validation:** Message type compiles

### 4.2 Update Message creation
- [ ] Update `NewMessage()` to accept Markdown string
- [ ] Trigger extraction and rendering on creation
- [ ] Handle errors gracefully

**Validation:** Can create messages from Markdown

---

## Phase 5: Chat Buffer Integration

### 5.1 Update buffer rendering
- [ ] Modify `chat.Model.View()` to use `Rendered.Text`
- [ ] Ensure ANSI codes are preserved in output
- [ ] Handle messages without rendering (fallback)

**Validation:** Buffer displays rendered Markdown

### 5.2 Add extraction trigger
- [ ] Call extractor when adding messages
- [ ] Store BlockIndex in message
- [ ] Update message count display

**Validation:** Messages have valid BlockIndex

### 5.3 Implement lazy rendering
- [ ] Only render visible messages
- [ ] Render on-demand when scrolling
- [ ] Use LRU cache for rendered content

**Validation:** Performance is good with 100+ messages

---

## Phase 6: Code Block Navigation

### 6.1 Track current code block
- [ ] Add `currentCodeBlock int` field to buffer Model
- [ ] Update when user navigates
- [ ] Display in UI (status or inline)

**Validation:** Current code block is visible

### 6.2 Implement ]c (next code block)
- [ ] Implement `jumpToNextCodeBlock()` method
- [ ] Calculate line position from BlockIndex
- [ ] Scroll viewport to show code block
- [ ] Handle wrap-around (last to first)

**Validation:** ]c jumps to next code block

### 6.3 Implement [c (previous code block)
- [ ] Implement `jumpToPrevCodeBlock()` method
- [ ] Calculate line position from BlockIndex
- [ ] Scroll viewport to show code block
- [ ] Handle wrap-around (first to last)

**Validation:** [c jumps to previous code block

### 6.4 Implement N]c (jump to specific)
- [ ] Parse count from keybinding
- [ ] Jump to Nth code block
- [ ] Clamp to valid range

**Validation:** 3]c jumps to 3rd code block

---

## Phase 7: Code Block Copying

### 7.1 Implement yc (copy current)
- [ ] Implement `copyCurrentCodeBlock()` method
- [ ] Get content from BlockIndex (not Rendered)
- [ ] Copy to clipboard using clipboard package
- [ ] Display confirmation message

**Validation:** yc copies code block to clipboard

### 7.2 Implement yNc (copy by number)
- [ ] Parse count from keybinding
- [ ] Get Nth block from BlockIndex
- [ ] Copy to clipboard
- [ ] Display confirmation with block number

**Validation:** y2c copies 2nd code block

### 7.3 Implement ym (copy entire message)
- [ ] Implement `copyEntireMessage()` method
- [ ] Copy all code blocks or full message
- [ ] Display confirmation

**Validation:** ym copies full message

### 7.4 Handle clipboard errors
- [ ] Display error if clipboard not available
- [ ] Don't crash on copy failure
- [ ] Suggest installing pbcopy/xclip/wl-copy

**Validation:** Graceful error handling

---

## Phase 8: Cleanup

### 8.1 Remove or update placeholder code
- [ ] Update `pkg/markdown/parser.go` to note Glamour usage
- [ ] OR remove `pkg/markdown/` if not needed
- [ ] Update any references in docs

**Validation:** No placeholder code remains

### 8.2 Update documentation
- [ ] Update README.md to mention Glamour
- [ ] Update architecture.md with rendering approach
- [ ] Document regex pattern for extraction

**Validation:** Documentation is accurate

---

## Phase 9: Testing

### 9.1 Unit tests
- [ ] Test extractor with various Markdown formats
- [ ] Test renderer with Glamour
- [ ] Test message integration
- [ ] Test navigation logic
- [ ] Test copy operations

**Validation:** All unit tests pass

### 9.2 Integration tests
- [ ] Test full pipeline (Markdown → Extract + Render)
- [ ] Test with real AI responses
- [ ] Test edge cases (empty, no blocks, malformed)

**Validation:** Integration tests pass

### 9.3 Manual testing
- [ ] Test with various AI responses
- [ ] Test code block navigation
- [ ] Test copy operations
- [ ] Test theme switching

**Validation:** Manual testing successful

---

## Dependencies

- **Phase 1** must be completed first (need Glamour)
- **Phase 2** and **Phase 3** can be done in parallel
- **Phase 4** depends on Phase 2 and 3
- **Phase 5** depends on Phase 4
- **Phase 6** and **Phase 7** depend on Phase 5
- **Phase 8** and **Phase 9** can be done in parallel with implementation

## Estimated Time

- **Phase 1:** 30 minutes
- **Phase 2:** 2-3 hours (extractor + tests)
- **Phase 3:** 2-3 hours (renderer + caching)
- **Phase 4:** 1 hour
- **Phase 5:** 2-3 hours
- **Phase 6:** 2 hours
- **Phase 7:** 1-2 hours
- **Phase 8:** 1 hour
- **Phase 9:** 2-3 hours

**Total:** ~15-20 hours

## Definition of Done

This change is complete when:
1. Glamour renders Markdown in chat buffer
2. Code blocks are extracted with regex
3. `]c` / `[c` navigation works
4. `yc` / `yNc` / `ym` copying works
5. Copied code is plain text (no ANSI)
6. Performance is acceptable (100+ messages)
7. All tests pass
8. Documentation is updated
