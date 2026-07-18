# Feat

Essa documentação informa quais passos devem ser seguidos para implementar uma nova funcionalidade, informando as pastas em que os arquivos devem ser criados, padrões de nomenclatura e outras informações.

## Core

Onde a lógica do "negócio" está implementanda, sendo dividida em três diretórios:

1. `domain`: onde as estruturas que representam os tipos de dados do projeto são declaradas.
2. `repository`: um conjunto de repositórios, onde cada repositório reune a declaração de métodos que manipulam os dados armazenados em disco relacionados ao tipo de dado representado pelo nome do repositório.
3. `usecase`: a lógica do "negócio" é implementada por casos de uso, separados em arquivos dentro deste diretório.

### Domain

Se necessário, adicione novos arquivos em `/core/domain` para representar os tipos de dados que serão utilizados para implementar a nova feature.

Cada arquivo pode conter:

- **Constantes**: mensagens de erro, por exemplo. Importante notar que não necessariamente são constantes da linguagem go (declaradas como `const`), mas sim, valores que não serão alterados durante a execução do projeto.
- **Estrutura**: uma declaração de `struct` para representar o tipo de dado. Essa struct pode ser acompanhada de métodos, e **obrigatoriamente** de um construtor. Todo construtor deve ter seu nome iniciado por `New`, por exemplo `NewLink`.

Exemplo de domain:

```go
var (
	LINK_ERROR_NEW = errors.New("error instantiate new link")
	...
)

type Link struct {
	uid uuid.UUID
	...
}

func (l Link) GetUid() uuid.UUID {
	return l.uid
}

func (l *Link) SetUid(uid uuid.UUID) {
	l.uid = uid
}

func NewLink(...) (Link, error) {
	...
}
```

Como pode-se perceber, é recomendado declarar os identificadores como privados, e implementar getters e setters para eles, principalmente se estes são do tipo `uuid.UUID`.

### Repository

Adicione novos métodos ou altere a declaração de algum método já existente para que o repositório mais apropriado possa ser utilizado pela feature a ser criada, ou se nenhum repositório representar bem a feature, crie um novo.

Ao criar um novo repositório, certifique-se de incluir apenas **interfaces** dentro dele, separando os métodos em interfaces de **leitura** e de **escrita**, e uma interface final para reunir todas as demais. Interfaces de leitura devem ter seu nome iniciado por `Read` e interfaces de escrita devem ter seu nome iniciado por `Write`, por exemplo: `ReadCategory` e `WriteCategory`. Por fim, adicione à estrutura `Repositories`, localizada no arquivo `settings_values.go`, um campo que represente o repositório criado.

Exemplo de repository:

```go
type ReadLink interface {
	Exists(uuid.UUID) (bool, error)
	GetByUid(uuid.UUID) (domain.Link, error)
}

type WriteLink interface {
	Create(domain.Link) error
	Delete(uuid.UUID) error
	Update(uuid.UUID, domain.Link) error
}

type Link interface {
	ReadLink
	WriteLink
}
```

Caso seja criado testes unitários para testar o caso de uso da feature (prática recomendada), será necessário criar um repositório mockado, caso tenha sido criado um novo repositório real, ou adicionar um método ao repositório mockado correspondente ao repositório real modificado. Os repositórios mockados ficam em `/pkg/ltests`, e possuem o nome do arquivo terminado em `mock_repository`. Normalmente esse arquivo contém uma estrutura semelhante a que segue:

```go
type CategoryMockRepository[T Identifiable] struct {
	mock.Mock
}
```

E contém a implementação de todos os métodos do repositório real, porém, utilizando métodos da biblioteca **mock** do **testify** para mockar seu retorno.

### Usecase

Uma nova feature deve incluir sempre a criação de um novo caso de uso. Cada caso de uso é separado em um pacote próprio, e é acompanhado preferencialmente por um conjunto de testes unitários.

Cada caso de uso deve conter:

1. Uma `struct`, onde cada campo é um repositório que será utilizado internamente pelo caso de uso.
2. Uma função construtora chamada apenas de `New`.
3. Um método chamado `Execute`, responsável por executar a lógica principal da feature.

O código de um caso de uso deve ser de alto nível, e tratar todos os erros que os métodos e funções usadas possam retornar (normalmente retornando o mesmo erro imediatamente).

Exemplo de caso de uso:

```go
type GetAllRootCategories struct {
	Repo repository.Category
}

func New(repo repository.Category) GetAllRootCategories {
	return GetAllRootCategories{repo}
}

func (garcUseCase *GetAllRootCategories) Execute() (
	categories []domain.Category,
	err error,
) {
	categories, err = garcUseCase.Repo.GetAllRootCategories()
	return
}
```

Os casos de uso normalmente seguem alguns padrões, que variam conforme seu objetivo:

* **Criar um registro**: recebe como entrada todos os dados do registro, com exceção de identificadores únicos e de valores calculados (como data de criação), estes devem ser criados pelo caso de uso. Se existir um dado de domínio que represente o registro, este deve ser construido e passado na chamada do repositório. Por fim, deve ser retornado o identificador criado e possíveis erros.
* **Deletar um registro**: recebe como parâmetro a chave do registro, verifica se o registro existe retornando um erro apropriado caso não exista, faz outras verificações quando necessárias, e por fim passa a chave para a chamada do repositório, retornando possíveis erros.

#### Usecase Test

Um teste unitário de um caso de uso, deve essencialmente, testar se o caso de uso se comporta como esperado para determinadas saídas dos métodos chamados internamente. Por exemplo, se esperado que seja retornado um erro quando o método interno `Exists` dê como saída o valor `false`, um teste unitário poderia testar esse comportamento mockando a saída de `Exists` e comparando o retorno da função `Execute` com o erro esperado.

Os testes unitários devem servir também como documentação, então uma mensagem clara de erro deve ser passada em cada `assert`, quando cabível. É importante distinguir o que é trabalho da lógica do négocio do que é trabalho do repositório, para não incluir testes unitários que testam outro escopo (fora do core). Da mesma forma, todo teste deve considerar que sua entrada está sempre correta, pois se não estivesse, seria trabalho do `handler` da requisição parar a execução antes de chamar o caso de uso.

Um teste de caso de uso será organizado normalmente da seguinte maneira:

```go
var assert = assert.New(t)

// instanciação do(s) repositório(s) mockado(s) usado(s) pelo caso de uso
var repo = repository.New...MockRepository()
// instanciação do caso de uso (normalmente o nome da variável é a junção das inicias do seu nome)
vat cdu = New(repo)

// mock dos métodos usados internamente
repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, nil)

// execução do caso de uso
valor, err := cdu.Execute(mockUid)

// testes
assert.NoError(
	err,
	"buscar tal valor não deveria retornar erro")
...
```

## Adapters

Todo código que acessa sistemas externos (que não fazem parte da lógica do negócio) deve ser escrito como um adaptador. Devem ser modelados de forma que os sistemas externos possas ser substituídos por outros, sem que o código do `core` seja alterado.

### Repositórios

Esse é o caso dos repositórios, responsáveis por acessar o banco de dados. Como exemplo, temos no projeto o diretório `/adapters/sqlite` que implementa os repositórios declarados no `core` para uso com um banco de dados SQLite. Para adicionar um respositório, dois passos são necessários:

1. Implementar o repositório em um novo pacote dentro de `sqlite`.
1. Instanciar o repositório em `GetRepositories`.

Cada repositório é implementado em um pacote separado, e é dividido em pelo menos três arquivos, além do arquivo de testes unitários, um para declaração da estrutura que representa o repositório e do seu construtor, um para implementação dos métodos de leitura (tem seu nome iniciado por `read`) e outro para implementação dos métodos de escrita (tem seu nome iniciado por `write`). A organização do repositório é semelhante à definição escrita em `/core/repository`.

#### Estrutra e construtor

Exemplo do arquivo que declara a estrutura e o construtor:

```go
type Category struct {
	DB *sql.DB
}

func New(db *sql.DB) *Category {
	return &Category{db}
}
```

#### Leitura

Basicamente, métodos de leitura podem ser de três tipos:

1. Retorna um valor booleano, por exemplo: verifica se existe um registro com a chave informada.
2. Retorna um elemento do domínio, por exemplo: retorna a categoria que possui a chave informada.
3. Retorna um array de elementos, por exemplo: um array de categorias

Um método de exists normalmente vai seguir o seguinte padrão:

```go
func (cr *Category) Exists(uid uuid.UUID) (exists bool, err error) {
	err = cr.DB.QueryRow(
		`SELECT name FROM category WHERE uid_category = ?`,
		uid).Scan(new(string))

	exists = err != sql.ErrNoRows

	if err == sql.ErrNoRows {
		err = nil
	}

	return
}
```

* Repare que `Exists` não retorna erro se o dado buscado não existir

Já um método que retorna um único dado do domínio, segue o padrão:

```go
func (cr *Category) GetByUid(uid uuid.UUID) (l domain.Category, err error) {
	err = cr.DB.QueryRow(
			`SELECT
				name,
				...
			FROM category WHERE uid_category = ?`,
			uid).
		Scan(
			&l.Name,
			...)

	if err == nil {
		l.SetUid(uid)
	}

	return
}
```

* Não há a necessidade de buscar o `uid` no banco para o caso de exemplo.

E para métodos que podem retornar vários itens:

```go
func (cr *Category) GetAllRootCategories() (
	categories []domain.Category, err error,
) {
	var rows *sql.Rows

	rows, err = cr.DB.Query(
		`
			SELECT
				uid_category,
				name,
				...
			FROM category
			WHERE uid_father IS NULL
		`)

	if err != nil {
		return
	}

	categories = make([]domain.Category, 0, 0)
	var c domain.Category
	var categoryUid uuid.UUID
	var categoryUidStr string
	var loopErr error

	for rows.Next() {
		loopErr = rows.Scan(
			&categoryUidStr,
			&c.Name,
			...)

		if loopErr != nil {
			continue
		}

		categoryUid, loopErr = uuid.Parse(categoryUidStr)

		if loopErr != nil {
			continue
		}

		c.SetUid(categoryUid)
		categories = append(categories, c)
	}

	err = rows.Err()
	rows.Close()

	return
}
```

* Erros dentro do loop não devem impedir que as próximas linhas sejam lidas, apenas pular a leitura da linha atual.
* Como o `uid`'vêm do banco em formato de string, é necessário converter ele para `uuid.UUID`.
* Obrigatoriamente, antes de iniciar o loop, deve ser criado um array vazio e atribuído para a variável que será retornada pelo método, caso contrário, se a busca não retornar registros será retornado `nil` em vez de um array vazio.

#### Escrita

Os métodos de escrita são basicamente resolvidos em duas operações, uma chamada a `DB.Exec` e `return`:

```go
func (cr *Category) DeleteSubcategory(
	childUid uuid.UUID,
) (err error) {
	_, err = cr.DB.Exec(
		`
			UPDATE category
			SET uid_father = null
			WHERE uid_category = ?
		`,
		childUid)

	return
}
```

#### Teste em repositórios

Os testes unitários de um repositório utilizam um banco de dados de teste, mockado durante a execução dos testes e limpo no final. Para que esse comportamento seja respeitado, alguns métodos devem ser definidos e corretamente implementados:

- Em `SetupSuite` o banco de dados em ambiente de teste é carregado na versão necessária para seu funcionamento.
- Pode também inserir dados na base, caso estes não impactem na previsibilidade dos testes.

**Exemplo**:

```go
var (
	categoryTableMigration uint = 1765719599
...
func (ts *TestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	db, _ := sqlite.GetConnection(env)
	lmigration.Migrate(db, categoryTableMigration, sqlite.GetMigration)
	cr := New(db)

	ts.env = env
	ts.db = db
	ts.cr = cr
}
```

- `TearDownTest` deve limpar o banco para que cada teste seja executado em condições previsíveis.

**Exemplo**:

```go
func (ts *TestSuite) TearDownTest() {
	ts.db.Exec("DELETE FROM category")
}
```

- `TearDownSuite` deve limpar o banco completamente.

**Exemplo**:

```go
func (ts *TestSuite) TearDownSuite() {
	lmigration.Drop(ts.db, sqlite.GetMigration)
	ts.db.Close()
}
```

## Cmd

Em `cmd` estão localizadas todas as interfaces do projeto, podendo ser, por exemplo, uma API ou uma CLI.

Cada interface tem sua própria `main`.

### API

Baseada no padrão REST, a API utiliza o framework **echo** para implementar um servidor HTTP.

Além da `main`, há alguns subdiretórios na pasta, sendo eles:

> ToDo

- **custom**:
- **handler**:
- **interface**:
- **middleware**:
- **route**:
