# Klipa

**Klipa** is a simple and efficient CLI clipboard manager that allows you to save, retrieve, copy, and manage clipboard entries directly from your terminal.

![License](https://img.shields.io/badge/license-MIT-blue.svg)

---

## ✨ Features

- Save clipboard values or raw CLI input  
- List all saved clipboard entries  
- Retrieve values using keys  
- Copy stored values back to the system clipboard  
- Delete specific entries or flush the entire clipboard store  
- Simple key-based organization for fast access  

---

## 📦 Installation

You can download the latest release binaries from the [Releases](https://github.com/Humiditii/klipa/releases) page or install via Homebrew:

```bash
brew install Humiditii/homebrew-tap/klipa
```

> ✅ Or, use the simplified method by tapping the repo first:

```bash
brew tap Humiditii/homebrew-tap
brew install klipa
```

---

## 🧠 Usage

Klipa supports several useful commands to manage your clipboard data efficiently.

### ✅ `save`

- **Description**: Save clipboard entry or a value from the CLI.
- **Usage**:
  ```bash
  klipa save                         # Saves value from current system clipboard
  klipa save -v "my text" -k note1  # Saves "my text" under the key 'note1'
  ```

---

### ✅ `list`

- **Description**: List all saved clipboard entries by their keys.
- **Usage**:
  ```bash
  klipa list
  ```

---

### ✅ `get`

- **Description**: Retrieve a saved item using its key.
- **Usage**:
  ```bash
  klipa get keyName
  ```

---

### ✅ `copy`

- **Description**: Copy a saved value back to your system clipboard.
- **Usage**:
  ```bash
  klipa copy keyName
  ```

---

### ✅ `clear`

- **Description**: Delete a specific clipboard entry by its key.
- **Usage**:
  ```bash
  klipa clear keyName
  ```

---

### ✅ `flush`

- **Description**: Completely clear all saved clipboard entries.
- **Usage**:
  ```bash
  klipa flush
  ```

---

## 🔍 Examples

```bash
# Save a clipboard item
klipa save

# Save a specific value under a key
klipa save -v "secret-token" -k token

# List all saved items
klipa list

# Retrieve a value
klipa get token

# Copy value back to clipboard
klipa copy token

# Remove an item
klipa clear token

# Remove everything
klipa flush
```

---

## 📝 License

Klipa is released under the [MIT License](LICENSE).
