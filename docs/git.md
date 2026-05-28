# Git

Essa documentação informa quais os padrões e melhores práticas seguidas ao desenvolver este projeto, no que diz respeito à manipulação e uso do **git**.

## Padrões de commit

Todo commit deve seguir o seguinte padrão:

```
tipo_commit(local): breve descrição ou principal ponto do commit\n
maiores explicações do commit, ou outros pontos relevantes (opcional)
```

### Tipos de commit

- **fix**: indica que foi realizada uma correção
- **feat**: implementação de uma nova feature, por exemplo, de um novo caso de uso
- **docs**: documentação criada ou modificada
- **test**: teste criado ou modificado
- **perf**: alterações que`impactam na performance
- **refactor**: refatoração de um ou mais trechos de código
- **chore**: alterações que não foram aplicadas no código fonte
- **cleanup**: limpeza de código
- **remove**: remoção de funcionalidades, dirétórios ou arquivos não utilizados
