# 🛒 ProductSorter

A modular and extensible product catalog sorting engine written in Go. This tool supports A/B testing for multiple sorting strategies (e.g., by price or performance ratio) and is designed using the Strategy Pattern and Open/Closed Principle.

---

## ✨ Features

- 📦 Sort by price (low to high)
- 📊 Sort by performance (sales/views ratio)
- 🧩 Easily extendable with new sorting strategies
- 📚 JSON file-based product source
- 📋 Pretty-printed table output for terminal display
- 🧪 Unit tested with table-driven testing

---

## 🗂️ Project Structure

```

.
├── data/               # Product data source (JSON)
│   └── products.json
├── models/             # Shared data models
│   └── product.go
├── reader/             # Responsible for reading JSON input
│   ├── reader.go
│   └── reader_test.go
├── sorter/             # All sorting logic (strategy pattern)
│   ├── performance_sorter.go
│   ├── price_sorter.go
│   ├── registry.go
│   └── sorter_test.go
├── utils/              # Utility helpers (e.g., table formatting)
│   ├── display.go
│   └── display\_test.go
├── go.mod
├── main.go             # Entry point (main)
└── README.md

````

---

## 🚀 Getting Started

### 📦 Prerequisites

- Go 1.22+
- Git

### 🔧 Installation

```bash
git clone https://github.com/yourusername/productSorter.git
cd productSorter
go mod tidy
````

### ▶️ Run

```bash
go run main.go --sort=performance
```

This will:

* Load `data/products.json`
* Sort products based on the strategy (e.g., performance)
* Print the results in a clean table format

---

## 📄 Example `.products.json`

```json
[
  {
    "id": 1,
    "name": "Alabaster Table",
    "price": 12.99,
    "created": "2019-01-04",
    "sales_count": 32,
    "views_count": 730
  },
  {
    "id": 2,
    "name": "Zebra Table",
    "price": 44.49,
    "created": "2012-01-04",
    "sales_count": 301,
    "views_count": 3279
  }
]
```

---

## 🧪 Running Tests

```bash
go test ./...
```

This runs tests for:

* Sorting logic (`sorter/`)
* Reader functionality (`reader/`)
* Table rendering (`utils/`)

---

## 📐 Design Principles

* **Strategy Pattern**: Decouples sorting logic from main control flow
* **Open/Closed Principle**: New sorters can be added without modifying core logic
* **Testability**: All modules are unit-tested

---

## 🔌 Extending with a New Sorter

1. Create a new sorter implementing the `Sorter` interface.
2. Register it in the registry:

   ```go
   registry.Register("your_sort_key", YourCustomSorter{})
   ```
3. Done — no existing logic needs to change.

