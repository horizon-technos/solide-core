# Guide de Contribution à SolidE

Tout d'abord, merci de votre intérêt pour SolidE ! 

Nous construisons un framework backend pour l'ère Post-IA. Notre philosophie est simple : **un code compris instantanément par l'IA est un code maîtrisé parfaitement par l'humain**. En contribuant à SolidE, vous ne rejoignez pas seulement un projet open-source ; vous participez à la création d'un standard industriel où la performance, la sécurité et la lisibilité ne sont pas des options, mais des exigences natives.

Ce guide décrit notre processus de développement, nos standards de code et la façon dont vous pouvez soumettre vos contributions.

## Code de Conduite

En participant à ce projet, vous vous engagez à respecter notre Code de Conduite (basé sur le Contributor Covenant). Nous attendons de tous les contributeurs un comportement professionnel, respectueux et constructif. Les comportements toxiques, le harcèlement ou les attaques personnelles ne seront pas tolérés.

## Comment Contribuer ?

### 1. Signaler un Bug
- Vérifiez d'abord que le bug n'a pas déjà été signalé dans les [Issues](https://github.com/horizon-tech/solide-core/issues).
- Si ce n'est pas le cas, ouvrez une nouvelle issue en utilisant le template **"Bug Report"**.
- Soyez aussi précis que possible : version de SolidE, version de Go, OS, extraits de code, et logs d'erreur.

### 2. Proposer une Fonctionnalité (RFC)
Chez Horizon, nous ne codons pas à l'aveugle. **Aucune ligne de code majeure n'est écrite sans un document de design préalable.** 
- Si vous souhaitez ajouter une fonctionnalité importante ou modifier l'architecture, vous devez rédiger une **RFC (Request For Comments)**.
- Ouvrez une issue avec le tag `[RFC]` en détaillant : le problème, la solution proposée, les impacts sur la performance et les alternatives envisagées.
- La communauté et l'équipe Core ont **48 heures** pour commenter. La décision finale est actée par le CTO.

### 3. Soumettre une Pull Request (PR)
- Les PRs doivent cibler la branche `main`.
- Une PR doit résoudre une issue existante. Mentionnez le numéro de l'issue (ex: `Fixes #42`).
- Assurez-vous que votre branche est à jour avec `main` avant de soumettre.

## Le Processus de Review (The Horizon Standard)

Une Pull Request chez Horizon doit franchir deux barrières strictes. Aucune dérogation n'est possible.

```mermaid
flowchart LR
    A[Pull Request] --> B{The Automated Gate}
    B -->|Linter, Tests, Benchmarks| C{Échec ?}
    C -->|Oui| D[Rejet & Feedback]
    C -->|Non| E{The Wisdom Gate}
    E -->|Revue Humaine| F{Validé ?}
    F -->|Non| D
    F -->|Oui| G[Merge sur Main]
    
    style B fill:#e1f5fe,stroke:#0288d1
    style E fill:#fff3e0,stroke:#f57c00
    style G fill:#e8f5e9,stroke:#388e3c