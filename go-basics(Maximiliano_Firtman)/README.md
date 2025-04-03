# <img src="https://blog.golang.org/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Go Logo" width="200"/>
### Eine Einführung für mein Lernen  
([For English version, click here](#go--learning-project-english-version))



---

## Nützliche Links zum [frontendmasters.com](https://frontendmasters.com/courses/go-basics/)   Kurs

- [Slides (GitHub)](https://github.com/firtman/go-fundamentals/blob/main/slides.pdf)  
- [Go Tour – Interaktive Lernplattform](https://go.dev/tour/welcome/1)  
- [frontendmasters.com Kursseite](https://frontendmasters.com/courses/go-basics/)

---

## Go Lernprojekt: Frontend Masters Course

Dieses Projekt ist Teil meines Lernfortschritts mit der Programmiersprache **Go (Golang)**. Es enthält einfache, aber praxisnahe Beispiele zu grundlegenden Konzepten wie Structs, Methoden, Interfaces, Goroutines und Channels.

---

### Lerninhalte & Konzepte

#### Structs und Methoden
- Eigene Typen wie `Instructor`, `Course` und `Workshop`
- Eingebettete Structs (Composition statt Vererbung)
- Methodenbindung wie `.Print()` oder `.SignUp()`
- `String()`-Methoden für saubere Konsolenausgabe mit `fmt.Println()`

#### Factory-Funktionen
- Z.  B. `NewInstructor(name, lastname)` zur Initialisierung
- Erlaubt Validierung und Defaults

#### Interfaces (Polymorphismus)
- Interface `Signable` mit Methode `SignUp() bool`
- Ermöglicht polymorphe Slices `[]Signable`

#### Concurrency mit Goroutines & Channels
- Starten paralleler Abläufe mit `go`
- Kommunikation via Channels `chan string`
- Verwendung von `make(chan T, size)`, `close()`, `<-channel`, `range channel`

---

## Projektstruktur

```txt
.
├── main.go                 // Einstiegspunkt mit Beispielen
├── data/
│   ├── instructor.go       // Instructor-Typ + Factory + Methoden
│   ├── course.go           // Course-Typ + Stringer
│   ├── workshop.go         // Workshop mit eingebetteten Structs
│   └── interfaces.go       // Interface `Signable`
```

---
<a id="go--learning-project-english-version"></a>

<img src="https://blog.golang.org/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Go Logo" width="200"/>

##  Learning Project (English Version)

This repository contains practical examples created while learning the [frontendmasters.com](https://frontendmasters.com/courses/go-basics/)  

**Go Basics Course**

## Useful Links

- [Course Slides on GitHub](https://github.com/firtman/go-fundamentals/blob/main/slides.pdf)  
- [Go Tour – Official Interactive Tutorial](https://go.dev/tour/welcome/1)  
- [Course Page on Frontend Masters](https://frontendmasters.com/courses/go-basics/)

## Topics Covered

- Structs and methods (Instructor, Course, Workshop)
- Embedded structs (composition over inheritance)
- Factory functions for clean initialization
- Interfaces and polymorphism via `Signable`
- Concurrency using Goroutines and Channels



