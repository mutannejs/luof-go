# feat

Essa documentação informa quais passos devem ser seguidos para implementar uma nova funcionalidade, informando as pastas em que os arquivos devem ser criados, padrões de nomenclatura e outras informações.

## Core

Onde a lógica do "negócio" está implementanda, sendo dividida em três diretórios:

1. `domain`: onde as estruturas que representam os tipos de dados do projeto são declaradas.
2. `repository`: um conjunto de repositórios, onde cada repositório reune a declaração de métodos que manipulam os dados armazenados em disco relacionados ao tipo de dado representando pelo nome do repositório.
3. `usecase`: a lógica do "negócio" é implementada por casos de uso, separados em arquivos dentro deste diretório.

### domain

Se necessário, adicione novos arquivos em `core/domain` para representar os tipos de dados que serão utilizados para implementar a nova feature.

Cada arquivo pode conter:

- **Constantes**: como por exemplo, mensagens de erro
- **Estrutura**: uma declaração de `struct` para representar o tipo de dado. Essa struct pode ser acompanhada de métodos, e **obrigatoriamente** de um construtor. Todo construtor deve ter seu nome iniciado por `New`, por exemplo `NewLink`

### repository

Adicione novos métodos ou altere a declaração de algum método já existente para que o repositório mais apropriado possa ser utilizado pela feature a ser criada, ou se nenhum repositório representar bem a feature, crie um novo.

Ao criar um novo repositório, certifique-se de incluir apenas **interfaces** dentro dele, separando os métodos em interfaces de **leitura** e de **escrita**, e uma interface final para reunir todas as demais. Interfaces de leitura devem ter seu nome iniciado por `Read` e interfaces de escrita devem ter seu nome iniciado por `Write`, por exemplo: `ReadCategory` e `WriteCategory`. Por fim, adicione à estrutura `Repositories`, localizada no arquivo `settings_values.go`, um campo que represente o repositório criado.

### usecase

Uma nova feature deve incluir sempre a criação de um novo caso de uso. Cada caso de uso é separado em um pacote próprio, e é acompanhado preferencialmente por um conjunto de testes unitários.

Cada caso de uso deve conter:

1. Uma `struct`, onde cada campo é um repositório que será utilizado internamente pelo caso de uso
2. Uma função construtora chamada apenas de `New`
3. Um método chamado `Execute`, responsável por executar a lógica principal da feature

O código de um caso de uso deve ser de alto nível, e tratar todos os erros que os métodos e funções usadas possam retornar (normalmente retornando o mesmo erro imediatamente).
