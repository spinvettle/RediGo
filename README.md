
# 🚀 RediGo

> **A Redis-compatible in-memory data store, written from scratch in Go.**  
> Not just a toy — it’s your playground to understand how Redis *really* works.

[![Go Version](https://img.shields.io/badge/Go-1.22%2B-blue?logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)


---

## 🔍 What is RediGo?

**RediGo** is a lightweight, educational implementation of the [Redis](https://redis.io) protocol and core data structures — built entirely in **Go**, with no external dependencies.

It supports:
- ✅ RESP (REdis Serialization Protocol) v2 parsing
- ✅ Core commands: `SET`, `GET`, `DEL`, `EXISTS`, `KEYS`, `FLUSHDB`, etc.
- ✅ Multiple data types: Strings (with TTL), soon: Lists, Hashes, Sets...
- ✅ Concurrent-safe storage using Go's concurrency primitives
- ✅ TCP server listening on port 6380 (to avoid conflict with real Redis)

> ⚠️ **Note**: This is **not** production-ready. It’s designed for learning, hacking, and fun!

---

## 🛠️ Quick Start

### Prerequisites
- Go 1.22+

### Run from source
```bash
git clone https://github.com/your-username/RediGo.git
cd RediGo
go run main.go
