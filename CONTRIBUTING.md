# Contributing to Keystone

Thank you for your interest in contributing to Keystone!  
We welcome all kinds of contributions — whether it's features, bug fixes, documentation updates, or even design ideas.

---

## How to Contribute

### 1. Check Existing Issues

Before starting work, check the [Issues](../../issues) tab to see if your idea is already tracked.

We use the following issue types:

- 🛠️ **Feature**: New capabilities or improvements.
- 🧹 **Task**: Setup work, refactoring, technical housekeeping.
- 🐞 **Bug**: Issues where functionality is broken.

### 2. Open an Issue (If Needed)

If your contribution isn't already tracked, please open a new issue using the relevant template.

### 3. Fork and Branch

- Fork the repository.
- Create a new branch with a clear name:
  - For features: `feat/short-description`
  - For bugs: `bug/short-description`
  - For tasks: `task/short-description`

Example:
feat/add-refresh-token-support
bug/fix-password-hashing
task/update-dockerfile

### 4. Make Your Changes

- Follow the project's existing coding style.
- Write clear, maintainable, and minimalistic code.
- Keep commits small and focused.

### 5. Write Tests (When Applicable)

- Unit tests are highly encouraged, especially for core logic.
- Testing adapters (DB/HTTP) is optional for early MVP.

### 6. Submit a Pull Request

- Reference the related issue number in your pull request.
- Provide a clear description of what you changed and why.

---

## Code Style Guidelines

- Follow Go idioms (`gofmt`, `goimports`).
- Favor simplicity over cleverness.
- Keep public functions and interfaces documented (short GoDoc comments).

---

## Communication

- Keep discussions polite, technical, and focused on improving Keystone.
- We value security, simplicity, and maintainability above "shiny" complexity.

---

Thank you for helping build something small, strong, and foundational!
