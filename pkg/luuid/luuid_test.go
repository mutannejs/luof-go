package luuid

import (
    "fmt"
    "github.com/google/uuid"
    "testing"
)

func TestIsZero(t *testing.T) {
    var uid uuid.UUID
    if !IsZero(uid) {
        fmt.Println("Zero não reconhecido")
        t.Fail()
    }
    uid, _ = New()
    if IsZero(uid) {
        fmt.Println("Valor diferente de zero reconhecido como zero")
        t.Fail()
    }
}

func TestZero(t *testing.T) {
    var zero = Zero()
    if !IsZero(zero) {
        fmt.Println("Zero não retornou corretamente")
        t.Fail()
    }
}
