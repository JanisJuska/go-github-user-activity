# GitHub User Activity CLI (gua)

A simple command-line tool written in Go that fetches and summarizes a GitHub user's recent activity using the GitHub Events API.

## ✨ Features

* Fetches public events for any GitHub user
* Groups activity by event type and repository
* Displays how many times each event occurred
* Clean and human-readable output

---

## 📦 Installation

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/github-user-activity.git
cd github-user-activity
```

### 2. Build the binary

```bash
go build -o gua
```

### 3. (Optional) Move binary to PATH

```bash
mv gua /usr/local/bin/
```

Now you can run `gua` from anywhere.

---

## 🚀 Usage

```bash
gua <github-username>
```

### Example

```bash
gua torvalds
```

### Sample Output

```
--- Pushed commits to a torvalds/linux: 3 times
--- Starred a repo: some/repo
--- Opened, Closed or Merged a pull request in repo: another/repo 2 times
```

---

## 🧠 How it Works

1. Fetches user events from:

   ```
   https://api.github.com/users/{username}/events
   ```
2. Parses the JSON response into Go structs
3. Aggregates events by type and repository
4. Outputs a summary to the terminal

---

## ⚠️ Notes

* Only public GitHub activity is available
* GitHub API rate limits may apply (unauthenticated requests are limited)

---

## 🛠 Requirements

* Go 1.20+ (or compatible)

---

## 📄 License

This project is open-source and available under the MIT License.

---

## 🙌 Acknowledgements

* GitHub API for providing public event data

---

## 📌 Future Improvements

* Add CLI flags (filter by event type, limit results)
* Sort output by frequency
* Improve formatting (colors, tables)
* Add caching

---

Happy hacking! 🚀
