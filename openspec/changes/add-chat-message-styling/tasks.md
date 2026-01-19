# Tasks: Add Chat Message Styling

Ordered implementation plan for chat message styling feature.

## Phase 1: Create Message Renderer Component

- [ ] Create `internal/chat/message_view.go` file
- [ ] Define `ChatMessage` struct with style configuration
- [ ] Implement `NewChatMessage()` constructor
- [ ] Implement `Render(msg Message, maxWidth int) string` method
- [ ] Implement `renderUserMessage()` private method
  - [ ] Green border style (color 142)
  - [ ] Right alignment
  - [ ] "You:" label in green
  - [ ] Gray content text (color 244)
- [ ] Implement `renderAssistantMessage()` private method
  - [ ] Blue border style (color 33)
  - [ ] Left alignment
  - [ ] "AI:" label in blue
  - [ ] White content text (color 252)

## Phase 2: Update Styles

- [ ] Add `UserMessageBorder` style to `ui.Styles` struct
- [ ] Add `AssistantMessageBorder` style to `ui.Styles` struct
- [ ] Add `UserLabel` style to `ui.Styles` struct
- [ ] Add `AssistantLabel` style to `ui.Styles` struct
- [ ] Add `UserMessageContent` style to `ui.Styles` struct
- [ ] Add `AssistantMessageContent` style to `ui.Styles` struct
- [ ] Initialize new styles in `DefaultStyles()` function
- [ ] Verify color values match design spec (green: 142, blue: 33, gray: 244, white: 252)

## Phase 3: Integrate into Chat Buffer

- [ ] Add `messageRenderer *ChatMessage` field to `chat.Model` struct
- [ ] Initialize message renderer in `NewModel()` constructor
- [ ] Update `Model.View()` to use message renderer
- [ ] Implement empty state handling in `View()`
  - [ ] Display "No messages yet." placeholder
  - [ ] Display "Type a message to start chatting..." hint
- [ ] Implement message iteration in `View()`
- [ ] Add vertical spacing (2 empty lines) between messages
- [ ] Handle edge case of nil/empty content blocks

## Phase 4: Width Calculation

- [ ] Implement `calculateMessageWidth(maxWidth int) int` helper
- [ ] Account for border width (4 chars total)
- [ ] Account for padding (2 chars total)
- [ ] Implement text wrapping for long content
- [ ] Test with various terminal widths

## Phase 5: Testing and Validation

- [ ] Test single user message display
- [ ] Test single AI message display
- [ ] Test alternating user/AI messages
- [ ] Test multiple consecutive user messages
- [ ] Test multiple consecutive AI messages
- [ ] Test very long message content
- [ ] Test empty message content
- [ ] Test terminal resize behavior
- [ ] Verify color accessibility (contrast ratios)
- [ ] Verify empty state placeholder
- [ ] Verify message spacing consistency

## Phase 6: Documentation

- [ ] Add package documentation to `message_view.go`
- [ ] Document color scheme in code comments
- [ ] Update relevant design documentation if needed
