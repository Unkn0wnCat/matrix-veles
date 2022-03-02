---
sidebar_position: 1
---

# Prometheus mit Veles verwenden

Die Verwendung von Prometheus mit Veles ist sehr einfach, da Veles standardmäßig einen kompatiblen Metrikendpunkt bereitstellt.

Um die Metriken zu verwenden, fügen Sie einfach einen Scrape-Job zu Prometheus hinzu, um die Metriken von `http://[dein-server-hier]:8123/metrics` zu ziehen.

Und das war's auch schon, du hast jetzt Zugriff auf ein paar relevante Metriken um die Gesundheit deines Veles zu überwachen! 🎉