# SolidE

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?logo=go)](https://go.dev/)
[![Status](https://img.shields.io/badge/status-alpha-orange)](https://github.com/horizon-tech/solide-core/releases)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

> **L'Architecture de l'Ère Post-IA.**
> Un framework backend en Go, forgé pour la transparence, la performance et l'auditabilité par les agents autonomes.

SolidE est un framework backend en langage Go conçu pour l'ère post-IA. Il rejette la complexité inutile des abstractions magiques pour embrasser une transparence radicale : **un code compris instantanément par l'IA est un code maîtrisé parfaitement par l'humain**. En combinant la performance brute du "Metal", une sécurité proactive et une architecture "AI-Native", SolidE permet de construire des systèmes intelligibles, pérennes et ultra-performants.

## Pourquoi SolidE ?

Les frameworks backend actuels sont souvent encombrés d'abstractions lourdes (réflexion, allocations dynamiques) qui obscurcissent la compréhension par la machine et pénalisent les performances. SolidE est la réponse à cette obsolescence.

**Idéal pour :**
- Les équipes qui construisent des **SaaS haute fréquence**, des **MMO** ou des systèmes de **trading** où la latence est critique.
- Les projets qui nécessitent une **sécurité native** (RBAC, JWT, protection contre les injections) sans configuration complexe.
- Les développeurs qui veulent un cadre **"Opinionated"** mais **"Zero-Compromise"** sur la qualité et la maintenabilité à long terme.

## Philosophie

- **AI-Native, pas AI-Bolted-On** : Contrairement aux outils classiques où l'IA est un plugin ajouté après coup, SolidE intègre des structures de données spécifiquement formatées pour être analysées par **SolidE Orb**.
  - *Note importante* : Orb n'est **pas** une IA temps réel qui s'exécute dans votre serveur de production (ce qui détruirait la latence). C'est un moteur d'analyse statique et sémantique qui opère en amont (CLI, Pre-commit, CI/CD) pour auditer, optimiser et sécuriser votre code avant le déploiement.
- **Performance Critique** : Conçu pour les environnements où la latence est l'ennemi. SolidE supprime les couches d'abstraction lourdes (No-Reflection, Zero-Allocation) pour offrir la vitesse du C++ avec la simplicité du Go.
- **Souveraineté & Modèle Hybride** : Une approche **Open Core** stratégique. Le moteur est open-source pour garantir l'adoption et la transparence, tandis que les modules d'intelligence (Orb Enterprise) et les plugins métiers apportent une valeur ajoutée premium.

