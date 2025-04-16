# 🧱 Common - DevTrack

Este repositório contém os componentes compartilhados entre os microsserviços do projeto **DevTrack**, como configuração, logging, contratos de eventos, middlewares e utilitários. Ele serve como base para garantir consistência, reuso e boas práticas entre os serviços da plataforma.

---

## 📦 Conteúdo

- `config/` – Leitura e parsing de variáveis de ambiente e arquivos de configuração
- `logger/` – Implementação do logger padronizado com suporte a níveis, contexto e saída estruturada
- `events/` – Definições de contratos de eventos utilizados na comunicação entre microsserviços (Kafka/SQS/etc.)
- `middleware/` – Middlewares compartilháveis entre serviços HTTP
- `errors/` – Tipos e helpers para tratamento padronizado de erros
- `utils/` – Funções utilitárias comuns

---

## 🛠️ Como usar

Adicione este repositório como um módulo Go em outros microsserviços:

```bash
go get github.com/reinasdev/devtrack-common@latest
```

Em seguida, importe os pacotes desejados nos seus serviços:

```go
import (
    "github.com/reinasdev/devtrack-common/logger"
    "github.com/reinasdev/devtrack-common/config"
    "github.com/reinasdev/devtrack-common/utils"
)
```
