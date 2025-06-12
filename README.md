# Testa-fy

**Testa-fy** is a Go-based text transformer that injects “Damian Testa” flair into your prose, based on phonetic patterns, special-case rules, and a touch of randomness.

## Features

- **Entire phrase replacement** (e.g. “Damian Testa” → “Damian Testa AKA Pooplord 5000”)  
- **Special fringe case replacements** (e.g. “demonstration” → “damianstration”, “tickle” → “test-tickle”)  
- **Prefix rules**  
  - `trans…` → `test…`  
  - `en…` → `damien…` (words >4 letters)  
  - `in…` → `damian…` (words >6 letters)  
  - `man…` → `damian…` (words >7 letters)
- **Suffix rules**  
  - `-ction` → `-ctestation`  
  - `-ation` → `-estation`  
  - `-ment` → `-testment`  
  - `-east` → `-easta`  
  - `-est` → `-a`  
  - `-ane` → `-anian`  
  - `-ain` → `-ainian`  
  - `-ame` → `-amian`  
  - `-ess` → `-essta`  
  - `-ster` → `-sta`

- **Randomisation**: each candidate word is transformed with probability `transformProb` (default 0.7).

## Installation

```bash
git clone git@github.com:devhou-se/testa-fy.git
cd testa-fy
idk bro I forgot the git commands
