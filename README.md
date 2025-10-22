# luof-go

Implementação em **Go** do Back-end do **Luof**.

### Arquitetura

O projeto é baseado na **arquitetura hexagonal**, conhecida também como portas e adaptadores.

### Modelagem

O **core** do projeto é exemplificado no seguinte diagrama:

```mermaid
erDiagram
    direction LR
    Link o{--o| BelongsTo : belong
    BelongsTo o|--o{ Category : belong
    Category o{--o| SubCategory : belong
    Link {
        string uid_link "PK"
        string url
        string name "NOT NULL"
        string description
        bool use_markdown
        datetime created_at "NOT NULL"
        datetime updated_at "NOT NULL"
    }
    BelongsTo {
        string uid_link "FK PK"
        string uid_category "FK PK"
        datetime iserted_at "NOT NULL"
        bool is_main
    }
    Category {
        string uid_category "PK"
        string name "NOT NULL"
        string description
        bool use_markdown
        datetime created_at "NOT NULL"
        datetime updated_at "NOT NULL"
    }
    SubCategory {
        string uid_child "FK PK"
        string uid_father "FK PK"
        datetime iserted_at "NOT NULL"
    }
```

Essa modelagem deve servir como base para que as operações abaixo sejam realizadas:

* Operações de CRUD sobre **links**
* Operações de CRUD sobre **categorias**
* Operações e agrupamentos simples de **links** por **categorias**

### Próximas etapas

- [X] Testes para link e category use cases
- [ ] Testes para belongs to use cases
- [ ] Adaptador SQLite
- [ ] API

### Projeto inicial

Esse repositório é parte de uma reformulação do projeto de mesmo nome disponibilizado no repositório https://github.com/mutannejs/luof, projeto implementando em **C** e somente com CLI.

### Principais referências

* Ports & Adapters Examples: https://github.com/nrjohnstone/ports-adapters-examples
* Hexagonal Architecture in Go: https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3
