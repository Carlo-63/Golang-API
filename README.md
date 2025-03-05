# 📌 Backend Golang API

Un'API REST sviluppata con **Golang** e il framework **Gin**. 

## 🚀 Setup

1. **Clona il repository**

2. **Installa le dipendenze**
   ```sh
   go mod tidy
   ```
3. **Compila il config** (modifica `config/config.json` con le credenziali DB e la "secretKey")

4. **Crea il database con le tabelle necessarie (todos e users)** (seguendo i modelli presenti in `models/`)

5. **Avvia il server**
   ```sh
   go run main.go
   ```

---
## 📡 API Endpoints

### 🟢 Autenticazione
- **POST** `/register` → Registra un nuovo utente.

### 🟠 TODOs
- **POST** `/addTodo` → Aggiunge un nuovo TODO (autenticazione richiesta).
- **GET** `/getTodos` → Restituisce tutti i TODO (autenticazione richiesta).
- **GET** `/getTodosOfUser` → Restituisce i TODO di un utente specifico (autenticazione richiesta).

### 🔵 Altri Endpoints
- **GET** `/saluta` → Endpoint di test che restituisce un saluto.

---
## 🔐 Sicurezza
- Usa **JWT** per l'autenticazione (necessario per i TODO).
- Configura **CORS** per proteggere l'accesso cross-origin.
- Gestisce le connessioni DB con un **connection pool**.

---
## 🛠 Tecnologie
- **Golang + Gin** (Framework API)
- **GORM** (Gestione database)
- **JWT** (Autenticazione)

📌 **Autore:** Carlo 63

