# POCs with JSON Logic using Go

## Overview
This project demonstrates how to use **JSON Logic** in a Go backend to apply dynamic business rules for a PDV (Point of Sale) system.  
The goal is to allow **multi-tenant customization** without changing code:

- Frontend sends the order payload.
- Backend evaluates JSON Logic rules.
- Frontend receives the result (discounts, restrictions, approvals).

---

## Tools
- Go
- Echo Framework
- [JSONLogic Go library](https://github.com/diegoholiveira/jsonlogic)

---

## Installation
1. Clone the repo:
```bash
git clone <repo-url>
cd json-logic-pocs
