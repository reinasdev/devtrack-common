# 🛣️ Roadmap DevTrack SaaS

## 🔹 Fase 1 – Arquitetura e Fundamentos *(Semana 1-2)*

| Objetivo | Aprender a estruturar microserviços e a arquitetura do projeto |
|----------|----------------------------------------------------------------|
| **Tarefas**: |
| - Criar o diagrama da arquitetura (Excalidraw)  |
| - Definir os repositórios separados para os serviços  |
| - Criar o repositório `devtrack-common` (contendo logger, config, definições de eventos etc.) |
| - Criar o esqueleto do primeiro serviço (`auth-service`), incluindo um README documentando o serviço e suas responsabilidades |

🧠 _Aprendizado_: boas práticas em Go, separação de domínios e contratos entre serviços, importância da documentação inicial.

## 🔹 Fase 2 – Desenvolvimento Local *(Semana 3-4)*

| Objetivo | Ter os serviços rodando localmente com Docker e começar a integrar as camadas de negócio |
|----------|---------------------------------------------------------------------|
| **Tarefas**: |
| - Configurar Docker Compose com Postgres, Redis e os microsserviços básicos |
| - Implementar o `auth-service`: login com OAuth2 (Google/GitHub) e emissão de JWT |
| - Criar e documentar o `user-service` com operações CRUD para os perfis |
| - Documentar, no README de cada repositório, o funcionamento interno e as interfaces do serviço |
| - Atualizar os diagramas no Excalidraw conforme novas camadas e integrações são adicionadas |
| - Implementar a comunicação entre serviços, por exemplo, por meio de um sistema de mensageria (NATS ou Kafka em ambiente local) |

🧠 _Aprendizado_: integração dos serviços, mensageria, documentação contínua e evolução dos diagramas.

## 🔹 Fase 3 – Infraestrutura e CI/CD *(Semana 5-6)*

| Objetivo | Automatizar o build, os testes e o deploy dos microsserviços |
|----------|----------------------------------------------------------------|
| **Tarefas**: |
| - Criar o repositório `devtrack-infra` com os manifests (Helm/ArgoCD) e configurações de infraestrutura |
| - Configurar pipeline de CI com GitHub Actions para cada repositório: build, lint, testes e geração de artefatos |
| - Implantar CD usando ArgoCD (inicialmente em um cluster local via Minikube ou kind) |
| - Criar Helm charts e atualizar os diagramas no Excalidraw para refletir a implantação via CI/CD |
| - Garantir que cada repositório, desde o início, contenha um README que documente seu setup, seu propósito e suas dependências |

🧠 _Aprendizado_: versionamento e automação de infraestrutura, GitOps, integração contínua e documentação sincronizada.

## 🔹 Fase 4 – Migração para AWS *(Semana 7-8)*

| Objetivo | Rodar o SaaS em produção na AWS utilizando serviços nativos e práticas de escalabilidade |
|----------|-----------------------------------------------------------------------|
| **Tarefas**: |
| - Provisionar um cluster EKS (via Terraform ou Console AWS) |
| - Migrar o banco de dados para o RDS (PostgreSQL) |
| - Configurar o ElastiCache (Redis) para caching e rate limiting |
| - Implantar a camada de mensageria usando MSK (Kafka) ou SQS |
| - Configurar o Ingress com ALB e integrar o Cert Manager para HTTPS |
| - Instalar o ArgoCD no EKS e configurar o deploy automático dos microsserviços |
| - Integrar o AWS Secrets Manager para gerenciamento seguro de configurações |
| - Atualizar os diagramas do Excalidraw para ilustrar a arquitetura na AWS |

🧠 _Aprendizado_: orquestração em nuvem com AWS, integração de serviços nativos, escalabilidade e segurança em ambientes de produção.

## 🔹 Fase 5 – Documentação, Testes e Refinamento *(Semana 9+)*

| Objetivo | Garantir que o sistema esteja bem documentado, testado e pronto para iterações futuras |
|----------|-------------------------------------------------------------------|
| **Tarefas**: |
| - Garantir que cada repositório possua um README atualizado e detalhado, descrevendo: |
|   - Configuração do ambiente |
|   - Instruções de build e deploy |
|   - Dependências e interações com outros microsserviços |
| - Manter os diagramas no Excalidraw atualizados conforme novas integrações ou alterações na arquitetura |
| - Implementar testes integrados, contract testing e um plano de rollback |
| - Revisar periodicamente a estratégia de CI/CD e a documentação para garantir aderência às melhores práticas |

🧠 _Aprendizado_: a importância da manutenção contínua e da comunicação transparente através da documentação técnica.