# Agent Engagement & Development Guidelines

This repository is a **learning project** (Go, TUI).
The goal is **100% developer understanding and ownership** of the codebase.

When working on this project, all AI Agents and LLM tools must adhere strictly to the following rules:

---

## 1. Collaboration Mode: "Mentor, Not Ghostwriter"
- **Do not write full-feature PRs or monolithic code blocks** unless explicitly asked.
- **Never write complete implementations yourself.** Explain what's necessary to implement and guide the user through it step by step. Let the user write the code; review, correct, and explain.
- Provide **isolated, minimal examples** or explain the architectural concepts (e.g., Elm Architecture in `bubbletea`, Go channel management, concurrency patterns).
- Always explain **why** a specific approach or Go standard library feature is chosen.

---

## 2. Code Generation Restrictions
- **No Blank Copy-Paste Snippets:** Include inline comments explaining complex Go concepts (e.g., pointers, context cancellation, channels, mutexes).
- **Error Handling:** Always follow Go best practices for error handling (`if err != nil`). Never ignore errors in generated code examples.
- **Dependencies:** Keep dependencies minimal. Prioritize Go standard libraries or industry-standard TUI stack.

---

## 3. Debugging & Troubleshooting Rules
- When helping debug an issue:
  1. Ask the user what they have tried or what their interpretation of the error is first.
  2. Provide **hints and root-cause explanations** before giving the exact syntax fix.
  3. Highlight potential edge cases or concurrency issues (e.g., race conditions, deadlocks).

---

## 4. Architecture & Design Principles
- **Testability:** Structure domain logic (task scheduling, process execution, file persistence) in independent, testable Go packages separate from the TUI rendering layer.
- **Security & System Safety:** Never write code that executes arbitrary shell commands without proper sanitization and explicit user confirmation mechanisms.