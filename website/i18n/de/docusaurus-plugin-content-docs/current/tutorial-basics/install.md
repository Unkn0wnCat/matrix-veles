---
sidebar_position: 1
---

# Veles Installation

:::caution

Diese Dokumentation ist für die Alpha-Version von Matrix-Veles. Veles könnte instabil und seltsam sein.

:::

:::info Software-Empfehlungen

Obwohl Veles unter Linux, Windows und macOS funktionieren sollte, **empfehlen wir Linux** zu verwenden, da Veles unter diesem entwickelt und getestet wurde! (Veles kann auf arm-basierten Minicomputern laufen)

:::

## Docker

:::tip TODO

Dieser Abschnitt wird in Kürze kommen.

:::

## Bare-Metal

:::info Vorab-Installationen

Veles verwendet *MongoDB* als Datenbank-Backend. Bitte [zuerst den MongoDB Community Server](https://www.mongodb.com/try/download/community) installieren.

:::

### Verwendung der Binärversion

Veles stellt fertige Binärdateien für die verbreitetsten Betriebssysteme und Architekturen bereit.

 1. Gehe zur [neuesten Version auf GitHub](https://github.com/Unkn0wnCat/matrix-veles/releases/latest)
 2. Navigiere unten zu "*Assets*" und finde die richtige Datei für dein Betriebssystem und deine Architektur
 3. Lade die Datei herunter (Linux/macOS: .tar.gz, Windows: .zip)
    1. (Optional) Überprüfe die MD5-Summe der heruntergeladenen Datei mit der angegebenen MD5-Summe
 4. Entpacke die Datei (Dein Betriebssystem sollte Hilfsprogramme haben, um dies zu tun)
 5. Navigiere in deinem Terminal zum entpackten Verzeichnis
 6. Führe `./matrix-veles generateConfig` aus um eine Basiskonfiguration zu erstellen<br/>(Linux: Du musst vielleicht die Datei mit `chmod +x ./matrix-veles` ausführbar machen)
 7. Bearbeite die Konfiguration in `./config.yaml` um dein Setup widerzuspiegeln
 8. Starte Matrix-Veles mit `./matrix-veles run`

Du hast jetzt eine voll funktionsfähige Installation von Veles! 🎉 Greife auf das Web-Interface unter http://127.0.0.1:8123 zu!

### Aus dem Quellcode erstellen

:::info

Erfahrung mit GoLang ist hierfür hilfreich!

:::

Um aus dem Quellcode zu bauen, stelle sicher, dass du die [neueste Version von GoLang](https://go.dev/dl/) installiert hast.

1. Öffne ein Terminal und führe `go install github.com/Unkn0wnCat/matrix-veles@latest` aus
2. Nach ein paar Minuten sollte das Build abgeschlossen sein
3. Führe `matrix-veles generateConfig` in dem Ordner aus, in dem du die Konfiguration speichern möchtest
4. Bearbeite die Konfiguration in `./config.yaml` um dein Setup widerzuspiegeln
5. Starte Matrix-Veles mit `matrix-veles run` im selben Verzeichnis wie deine Konfiguration
