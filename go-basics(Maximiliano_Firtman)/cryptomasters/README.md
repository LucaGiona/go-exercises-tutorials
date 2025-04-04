# CryptoMasters – Go-basierte Anwendung zur Kursabfrage  
[See English version below](#cryptomasters--go-based-console-application)

## Projektbeschreibung

CryptoMasters ist ein Lernprojekt in Go zur Echtzeitabfrage von Kryptowährungskursen über eine externe REST-API.  
Es zeigt, wie man Daten aus dem Web abruft, verarbeitet und in der Konsole ausgibt. Später soll das Projekt die Daten auch selbst als Schnittstelle bereitstellen.

---

## Ziel & Fokus

- Umsetzung als reine Konsolenanwendung
- Erlernen, wie man Daten aus Web-APIs konsumiert
- Ziel: diese Daten später selbst bereitstellen (z. B. über eigene REST-API)
- Verwendung von Goroutines zur parallelen Abfrage
- API-Datenquelle: [CEX.io](https://cex.io)
- Hinweis: Der verwendete JSON-to-Go-Formatter nutzt standardmäßig `float64` für numerische Felder

---

## Projektstruktur

```txt
cryptomasters/
├── api/                 # API-Zugriff (HTTP + JSON-Verarbeitung)
│   ├── cex.go           # Abfrage der Kursdaten
│   └── responses.go     # JSON-Strukturen für API-Antwort
├── datatypes/           # Eigene Typdefinitionen
│   └── data.go          # Typ "Rate" (Preis + Währung)
├── api_test.go          # Tests zur API-Validierung
├── main.go              # CLI-Startpunkt mit Goroutines
└── go.mod               # Go Moduldefinition
```

---

## Beispielausgabe

```bash
The rate for BTC is 72850.00 
The rate for ETH is 3620.45 
The rate for BCH is 432.19 
```

---

## Test ausführen

```bash
go test ./...
```

Beispiel-Test:

```go
func TestAPICall(t *testing.T) {
    _, err := api.GetRate("")
    if err == nil {
        t.Error("expected error for empty currency, got nil")
    }
}
```

---

## Lizenz

MIT – frei nutzbar für Lern- und Demonstrationszwecke.

---
<a id="cryptomasters--go-based-console-application"></a>

# CryptoMasters – Go-based Console Application

## Project Description

CryptoMasters is a Go learning project that fetches real-time cryptocurrency rates from an external REST API.  
It demonstrates how to consume data from the web, process it, and print results to the terminal. Later on, the project will aim to serve this data itself via an API.

---

## Goals & Focus

- Designed as a console-only application
- Learn how to consume web data from external APIs
- Goal: eventually serve the data via own API or service
- Uses Goroutines for concurrent execution
- Data source: [CEX.io](https://cex.io)
- Note: JSON-to-Go formatter uses `float64` by default for numeric types

---

## Project Structure

```txt
cryptomasters/
├── api/                 # API logic (HTTP + JSON)
│   ├── cex.go           # Performs external request
│   └── responses.go     # Data structures for JSON response
├── datatypes/           # Custom Go types
│   └── data.go          # Type "Rate" (price + currency)
├── api_test.go          # Unit tests for the API
├── main.go              # Entry point with CLI and concurrency
└── go.mod               # Go module definition
```

---

## Example Output

```bash
The rate for BTC is 72850.00 
The rate for ETH is 3620.45 
The rate for BCH is 432.19 
```

---

## Run Tests

```bash
go test ./...
```

Example test:

```go
func TestAPICall(t *testing.T) {
    _, err := api.GetRate("")
    if err == nil {
        t.Error("expected error for empty currency, got nil")
    }
}
```

---

## License

MIT – free for educational and demonstration purposes.
