# Spec: Session Manager

**Status:** Active
**Version:** 1.0.0
**Owners:** vai project

## Overview

This spec defines the session manager component that handles AI conversation history persistence, session switching, and the session list pane in the UI.

## ADDED Requirements

### Requirement: Session Data Model

The session manager MUST define a structured data model for chat sessions.

#### Scenario: Session structure

**Given** a chat session
**When** the session is represented in memory
**Then** it MUST follow this structure:
```go
type Session struct {
    ID          string    // Unique session identifier (UUID)
    Title       string    // Session title (auto-generated or user-defined)
    Messages    []Message // Ordered list of messages
    CreatedAt   time.Time // Session creation timestamp
    UpdatedAt   time.Time // Last update timestamp
    Model       string    // AI model used (e.g., "gpt-4", "claude-3")
}

type Message struct {
    ID        string      // Unique message identifier
    Role      string      // "user" or "assistant"
    Content   []Block     // Structured content blocks (see chat-buffer spec)
    CreatedAt time.Time   // Message timestamp
}
```

#### Scenario: Auto-generated session title

**Given** a new session is created
**When** the user sends the first message
**Then** the session SHOULD:
- Generate a title from the first user message
- Use the first ~50 characters of the message
- Append ellipsis if truncated
- Allow the user to manually rename the session later

#### Scenario: Empty session

**Given** the application starts
**When** no previous sessions exist
**Then** a new empty session SHOULD be created
**And** the session SHOULD have:
- A unique ID
- A default title (e.g., "New Chat" or timestamp-based)
- Zero messages

---

### Requirement: Session Storage

The session manager MUST persist sessions to disk with proper serialization.

#### Scenario: Storage location

**Given** vai needs to store session data
**When** determining the storage location
**Then** sessions SHOULD be stored in:
- `~/.local/share/vai/sessions/` (following XDG Base Directory spec)
- Or `~/Library/Application Support/vai/sessions/` on macOS
- Each session as a separate JSON file: `{session-id}.json`

#### Scenario: Save session on message

**Given** the user sends a message
**When** the message is added to the current session
**Then** the session manager SHOULD:
- Append the message to the session's Messages array
- Update the UpdatedAt timestamp
- Serialize the session to JSON
- Write to disk immediately (atomic write)
- Handle write errors gracefully

#### Scenario: Load existing sessions

**Given** the application starts
**When** the session manager initializes
**Then** it SHOULD:
- Scan the sessions directory for JSON files
- Load and parse each session file
- Build an in-memory index of sessions
- Sort sessions by UpdatedAt (newest first)
- Handle corrupted files by logging and skipping

#### Scenario: Session file format

**Given** a session is serialized to JSON
**When** the file is written
**Then** the JSON format SHOULD be:
```json
{
  "id": "uuid",
  "title": "Session Title",
  "messages": [
    {
      "id": "msg-uuid",
      "role": "user",
      "blocks": [
        {"type": "text", "text": "Hello"}
      ],
      "created_at": "2024-01-15T10:30:00Z"
    }
  ],
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:35:00Z",
  "model": "gpt-4"
}
```

#### Scenario: Atomic writes

**Given** a session needs to be saved
**When** writing to disk
**Then** the session manager SHOULD:
- Write to a temporary file first: `{session-id}.json.tmp`
- Sync the temporary file to disk
- Rename the temporary file to the final name (atomic operation)
- Clean up the temporary file on error

---

### Requirement: Session List Display

The session manager MUST render a list of sessions in the left pane.

#### Scenario: List rendering

**Given** there are multiple sessions
**When** the session list is rendered
**Then** it SHOULD display:
- Session title
- Last update timestamp (e.g., "2h ago", "Yesterday", "Jan 10")
- Number of messages in the session
- Visual indicator for the currently active session

#### Scenario: List item styling

**Given** the session list is being rendered
**When** a session is displayed
**Then** the item SHOULD:
- Show the title truncated to fit the pane width
- Use subtle styling for non-active sessions
- Use bright/distinct styling for the active session
- Show a marker (e.g., `*` or `>`) for the active session

#### Scenario: Empty session list

**Given** no sessions exist
**When** the session list is rendered
**Then** it SHOULD display:
- A friendly message: "No conversations yet"
- Or "Press ? for help"

#### Scenario: List sorting

**Given** multiple sessions exist
**When** the session list is rendered
**Then** sessions SHOULD be sorted:
- Primary: UpdatedAt (most recent first)
- Secondary: CreatedAt (newer first)

---

### Requirement: Session Switching

The session manager MUST support switching between sessions.

#### Scenario: Switch from list

**Given** the user is in NORMAL mode with focus on the session list
**When** the user presses `Enter` on a session
**Then** the session manager SHOULD:
- Set the selected session as the active session
- Load the session's messages into the chat buffer
- Update the chat buffer to show the session content
- Update the visual indicator in the session list

#### Scenario: Quick switch (optional)

**Given** the user is in NORMAL mode in the chat buffer
**When** the user presses a keybinding (e.g., `Ctrl+P` for previous, `Ctrl+N` for next)
**Then** the session manager SHOULD:
- Switch to the previous/next session in the list
- Update the chat buffer content
- Show a brief indicator of the switched session

#### Scenario: Create new session

**Given** the user is viewing an existing session
**When** the user presses a keybinding for new session (e.g., `Ctrl+T` or `:new`)
**Then** the session manager SHOULD:
- Create a new empty session
- Set it as the active session
- Clear the chat buffer
- Add the new session to the session list

---

### Requirement: Session Deletion

The session manager MUST support deleting sessions.

#### Scenario: Delete from list

**Given** the user is in NORMAL mode with focus on the session list
**When** the user presses `dd` or `:delete`
**And** the current session is NOT the active one
**Then** the session manager SHOULD:
- Delete the session file from disk
- Remove the session from the in-memory index
- Update the session list display

#### Scenario: Delete active session

**Given** the user tries to delete the active session
**When** the user presses `dd` or `:delete`
**Then** the session manager SHOULD:
- Prompt for confirmation: "Delete active session? (y/n)"
- If confirmed:
  - Delete the session file
  - Switch to the most recent other session (or create new)
  - Update the chat buffer
- If not confirmed:
  - Return to NORMAL mode without deleting

#### Scenario: Delete all sessions (optional)

**Given** the user wants to clear all history
**When** the user runs a command (e.g., `:clearall`)
**Then** the session manager SHOULD:
- Prompt for strong confirmation: "Delete ALL sessions? This cannot be undone. Type 'yes' to confirm:"
- If confirmed:
  - Delete all session files
  - Create a new empty session
  - Clear the session list
- If not confirmed:
  - Return to NORMAL mode

---

### Requirement: Session Renaming

The session manager MUST support renaming sessions.

#### Scenario: Rename from list

**Given** the user is in NORMAL mode with focus on the session list
**When** the user presses `r` or `:rename`
**Then** the session manager SHOULD:
- Switch focus to a prompt input
- Display the current title as default
- Allow the user to type a new title
- Save the new title on Enter
- Cancel on Esc

#### Scenario: Rename active session

**Given** the user is viewing a session
**When** the user renames it from anywhere
**Then** the session manager SHOULD:
- Update the session's Title field
- Save the session to disk
- Update the session list display
- Keep the session as active

---

### Requirement: Session Search

The session manager MUST support searching through sessions.

#### Scenario: Search sessions

**Given** the user is in NORMAL mode with focus on the session list
**When** the user presses `/` to search
**Then** the session manager SHOULD:
- Display a search prompt
- Filter the session list as the user types
- Match against session titles
- Highlight matching text
- Allow Esc to cancel search

#### Scenario: Search navigation

**Given** the user has searched for sessions
**When** multiple sessions match
**Then** the user SHOULD be able to:
- Press `n` to jump to next match
- Press `N` to jump to previous match
- Press Enter to select the current match

#### Scenario: No search results

**Given** the user searches for sessions
**When** no sessions match the query
**Then** the session list SHOULD display:
- "No matching sessions"
- Maintain the full list (not filtered)

---

### Requirement: Session Import/Export

The session manager MUST support importing and exporting sessions.

#### Scenario: Export session

**Given** the user wants to share a conversation
**When** the user runs a command (e.g., `:export` or `:e {filename}`)
**Then** the session manager SHOULD:
- Export the current session to a JSON file
- Use the specified filename or prompt for one
- Display confirmation: "Exported to {filename}"

#### Scenario: Export format

**Given** a session is exported
**When** the export file is written
**Then** the format SHOULD be:
- JSON for machine readability
- Include all session data (id, title, messages, timestamps)
- Be valid for re-importing

#### Scenario: Import session

**Given** the user has a session JSON file
**When** the user runs a command (e.g., `:import {filename}`)
**Then** the session manager SHOULD:
- Load and parse the JSON file
- Validate the format
- Generate a new ID (to avoid conflicts)
- Add the session to the list
- Display confirmation: "Imported session: {title}"

#### Scenario: Import error handling

**Given** the user tries to import an invalid file
**When** the import command runs
**Then** the session manager SHOULD:
- Display an error message: "Failed to import: {reason}"
- NOT add any session
- NOT crash or throw unhandled errors

---

### Requirement: Session Metadata

The session manager MUST track and display session metadata.

#### Scenario: Message count

**Given** a session has multiple messages
**When** the session is displayed in the list
**Then** it SHOULD show the message count
**And** the format SHOULD be: "{count} msgs" or just "{count}"

#### Scenario: Time formatting

**Given** a session's UpdatedAt timestamp
**When** it is displayed in the session list
**Then** it SHOULD use relative formatting:
- "< 1m" for less than a minute
- "2h" for hours
- "3d" for days
- "Jan 15" for older dates
- "Jan 15, 2024" for last year

#### Scenario: Model indicator

**Given** a session was created with a specific AI model
**When** the session is displayed
**Then** it MAY show:
- The model name as a badge
- A shortened version (e.g., "GPT-4" instead of "gpt-4-turbo-preview")
- Color-coded by model provider

---

## Cross-References

- **tui-framework**: Defines the top-level Model that hosts this session-manager sub-model
- **vim-navigation**: Defines the NORMAL mode operations for session list navigation
- **chat-buffer**: Consumes the session data to display messages in the chat buffer
