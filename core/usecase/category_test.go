package usecase

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/pkg/ltests"
    "github.com/mutannejs/luof-go/pkg/luuid"
    "github.com/google/uuid"
    "strings"
    "testing"
)

var (
    mockCategory, _ = domain.NewCategory(
        "development",
        "links about development",
        false,
    )
    mockUid = mockCategory.GetUid()
    invalidUid = uuid.New()
    updatedName = "Dev Web"
)

type CategoryRepository struct {
    Categories map[uuid.UUID]domain.Category
}

func NewMockCategoryRepository (t *testing.T, populate bool, testName string) *ltests.MockCRUDRepository[domain.Category] {
    var repo = ltests.NewMockCRUDRepository[domain.Category](populate, mockCategory)
    repo.TestRepository(t, testName)
    return repo
}

func TestCreateCategory(t *testing.T) {
    var repo = NewMockCategoryRepository(t, false, "CreateCategory")
    var cc = NewCreateCategory(repo)

    uid, err := cc.Execute(
        mockCategory.Name,
        mockCategory.Description.Content,
        mockCategory.Description.UseMarkdown,
    )

    if luuid.IsZero(uid) ||
            err != nil ||
            repo.Length() != 1 ||
            strings.Compare(mockCategory.Name, repo.GetItem(uid).Name) != 0 ||
            strings.Compare(mockCategory.Description.Content, repo.GetItem(uid).Description.Content) != 0 ||
            mockCategory.Description.UseMarkdown != repo.GetItem(uid).Description.UseMarkdown {
        ltests.PrintAndFail(t, "Insucesso na execução de CreateCategory", err)
    }
}

func TestDeleteCategory(t *testing.T) {
    var repo = NewMockCategoryRepository(t, true, "DeleteCategory")
    var dc = NewDeleteCategory(repo)

    exists, err := dc.Execute(mockUid)

    if !exists ||
            err != nil ||
            repo.Length() != 0 {
        ltests.PrintAndFail(t, "Insucesso na execução de DeleteCategory para um uid válido", err)
    }

    exists, err = dc.Execute(invalidUid)

    if exists || err != nil {
        ltests.PrintAndFail(t, "Insucesso na execução de DeleteCategory para um uid inválido", err)
    }
}

func TestGetCategoryByUid(t *testing.T) {
    var repo = NewMockCategoryRepository(t, true, "GetCategoryByUid")
    var gcbu = NewGetCategoryByUid(repo)

    category, err := gcbu.Execute(mockUid)

    if err != nil ||
            strings.Compare(mockCategory.Name, category.Name) != 0 ||
            strings.Compare(mockCategory.Description.Content, category.Description.Content) != 0 ||
            mockCategory.Description.UseMarkdown != category.Description.UseMarkdown ||
            mockCategory.CreatedAt.Compare(category.CreatedAt) != 0 ||
            mockCategory.UpdatedAt.Compare(category.UpdatedAt) != 0 {
        ltests.PrintAndFail(t, "Insucesso na execução de GetCategoryByUid para um uid válido", err)
    }

    category, err = gcbu.Execute(invalidUid)

    if (category != domain.Category{}) || err != nil {
        ltests.PrintAndFail(t, "Insucesso na execução de GetCategoryByUid para um uid inválido", err)
    }
}
