package tests

import (
	"database/sql"
	"testing"

	"github.com/mutannejs/luof-go/adapters/sqlite"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/lmigration"

	"github.com/stretchr/testify/assert"
)

func TestSetupSqlite(t *testing.T) {
    var assert = assert.New(t)

    env, err := lenv.Load(true)
    assert.NoError(err, "Tentar carregar as variáveis de ambiente não deveria retornar erro")

    db, err := sqlite.GetConnection(env)
    assert.NoError(err, "Tentar se conectar com o sqlite não deveria retornar erro")

    err = lmigration.Up(db, sqlite.GetMigration)
    assert.NoError(err, "As migrations deveriam ser executadas sem que ocorressem erros")

    err = db.QueryRow(`
            SELECT 1 FROM sqlite_master WHERE type='table' AND name='link';
        `).
        Scan(new(int))

    assert.NoError(err, "A tabela `link` deveria ter sido criada após executar as migrations")

    db.Close()
}

func TestDownSqlite(t *testing.T) {
    var assert = assert.New(t)

    env, err := lenv.Load(true)
    assert.NoError(err, "Tentar carregar as variáveis de ambiente não deveria retornar erro")

    db, err := sqlite.GetConnection(env)
    assert.NoError(err, "Tentar se conectar com o sqlite não deveria retornar erro")

    err = lmigration.Migrate(db, 0, sqlite.GetMigration)
    assert.NoError(err, "As migrations deveriam ser executadas sem que ocorressem erros")

    err = db.QueryRow(`
            SELECT 1 FROM sqlite_master WHERE type='table' AND name='link';
        `).
        Scan(new(int))

    assert.ErrorIs(sql.ErrNoRows, err, "A tabela `link` deveria ter sido dropada após executar as migrations")

    db.Close()
}
