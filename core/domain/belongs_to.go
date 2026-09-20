package domain

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	ALREADY_BELONGS = &i18n.Message{
		ID: "ALREADY_BELONGS",
		Description: "Retornado quando requisitado inserir um link na categoria que ele já pertence",
		Other: "the link already belongs to the category",
	}
	HAS_LINKS = &i18n.Message{
		ID: "HAS_LINKS",
		Description: "Retornado quando requisitado deletar uma categoria que possui links",
		Other: "the category cannot be deleted because it has links",
	}
	NOT_BELONGS = &i18n.Message{
		ID: "NOT_BELONGS",
		Description: "Retornado quando requisitado remover um link de uma categoria, sendo que ele não pertence a ela",
		Other: "the link does not belong to the category",
	}
)
