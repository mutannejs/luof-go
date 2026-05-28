# feat

Essa documentação informa quais passos devem ser seguidos para implementar uma nova funcionalidade, informando as pastas em que os arquivos devem ser criados, padrões de nomenclatura e outras informações.

## Core

Onde a lógica do "negócio" está implementanda, sendo dividida em três diretórios:

1. `domain`: onde as estruturas que representam os tipos de dados do projeto são declaradas.
2. `repository`: um conjunto de repositórios, onde cada repositório reune a declaração de métodos que manipulam os dados armazenados em disco relacionados ao tipo de dado representando pelo nome do repositório.
3. `usecase`: a lógica do "negócio" é implementada por casos de uso, separados em arquivos dentro deste diretório.

### domain

Se necessário, adicione novos arquivos em `/core/domain` para representar os tipos de dados que serão utilizados para implementar a nova feature.

Cada arquivo pode conter:

- **Constantes**: mensagens de erro, por exemplo. Importante notar que não necessariamente são constantes da linguagem go (declaradas como `const`), mas sim, valores que não serão alterados durante a execução do projeto
- **Estrutura**: uma declaração de `struct` para representar o tipo de dado. Essa struct pode ser acompanhada de métodos, e **obrigatoriamente** de um construtor. Todo construtor deve ter seu nome iniciado por `New`, por exemplo `NewLink`

### repository

Adicione novos métodos ou altere a declaração de algum método já existente para que o repositório mais apropriado possa ser utilizado pela feature a ser criada, ou se nenhum repositório representar bem a feature, crie um novo.

Ao criar um novo repositório, certifique-se de incluir apenas **interfaces** dentro dele, separando os métodos em interfaces de **leitura** e de **escrita**, e uma interface final para reunir todas as demais. Interfaces de leitura devem ter seu nome iniciado por `Read` e interfaces de escrita devem ter seu nome iniciado por `Write`, por exemplo: `ReadCategory` e `WriteCategory`. Por fim, adicione à estrutura `Repositories`, localizada no arquivo `settings_values.go`, um campo que represente o repositório criado.

Caso seja criado testes unitários para testar o caso de uso (recomendado) da feature, será necessário criar um repositório mockado, caso tenha sido criado um novo repositório real, ou adicionar um método ao repositório mockado correspondente ao repositório real modificado. Os repositórios mockados ficam em `/pkg/ltests`, e possuem o nome do arquivo terminado em `mock_repository`. Normalmente esse arquivo contém uma estrutura semelhante a que segue:

```
type CategoryMockRepository[T Identifiable] struct {
	mock.Mock
}
```

E contém a implementação de todos os métodos do repositório real, porém, utilizando métodos da biblioteca **mock** do **testify** para mockar seu retorno.

### usecase

Uma nova feature deve incluir sempre a criação de um novo caso de uso. Cada caso de uso é separado em um pacote próprio, e é acompanhado preferencialmente por um conjunto de testes unitários.

Cada caso de uso deve conter:

1. Uma `struct`, onde cada campo é um repositório que será utilizado internamente pelo caso de uso
2. Uma função construtora chamada apenas de `New`
3. Um método chamado `Execute`, responsável por executar a lógica principal da feature

O código de um caso de uso deve ser de alto nível, e tratar todos os erros que os métodos e funções usadas possam retornar (normalmente retornando o mesmo erro imediatamente).

#### usecase test

Um teste unitário de um caso de uso, deve essencialmente, testar se o caso de uso se comporta como esperado para determinadas saídas dos métodos chamados internamente. Por exemplo, se esperado que seja retornado um erro quando o método interno `Exists` dê como saída o valor `false`, um teste unitário poderia testar esse comportamento mockando a saída de `Exists` e comparando o retorno da função `Execute` com o erro esperado.

Os testes unitários devem servir também como documentação, então uma mensagem clara de erro deve ser passada em cada `assert` quando cabível. É importante distinguir o que é trabalho da lógica do négocio do que é trabalho do repositório, para não incluir testes unitários que testam outro escopo (fora do core). Da mesma forma, todo teste deve considerar que sua entrada está sempre correta, pois se não estivesse, seria trabalho do `handler` da requisição para a execução antes de chamar o caso de uso.

Um teste de caso de uso será organizado normalmente da seguinte maneira:

```
var assert = assert.New(t)

// instanciação do(s) repositório(s) mockado(s) usado(s) pelo caso de uso
var repo = repository.New...MockRepository()
// instanciação do caso de uso (normalmente o nome da variável é a junção das inicias do seu nome)
vat cdu = New(repo)

// mock dos métodos usado internamente
repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, nil)

// execução do caso de uso
valor, err := cdu.Execute(mockUid)

// testes
assert.NoError(
	err,
	"buscar tal valor não deveria retornar erro")
...
```
