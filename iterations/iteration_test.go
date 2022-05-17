package iteration

import (
    "testing"
    "fmt"
    "strings"
)

func TestRepeat(t *testing.T) {
    assertCorrectAnswerString := func (t testing.TB, got, want string) {
        if got != want {
            t.Errorf("expected %q but got %q", want, got)
        }
    }
    assertCorrectAnswerBool := func (t testing.TB, got, want bool) {
        if got != want {
            t.Errorf("expected %t but got %t", want, got)
        }
    }
    t.Run("say a 5 times", func (t *testing.T){
        repeated := Repeat("a", 5)
        expected := "aaaaa"
        assertCorrectAnswerString(t, repeated, expected)
    })
    t.Run("say a multiple times", func (t *testing.T){
        repeated := Repeat("a", 10)
        expected := "aaaaaaaaaa"
        assertCorrectAnswerString(t, repeated, expected)
    })
    t.Run("test string hasprefix", func (t *testing.T){
        repeated := strings.HasPrefix(Repeat("a", 10), "aaaaa")
        expected := true
        assertCorrectAnswerBool(t, repeated, expected)
    })
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 5)
    }
}

func ExampleRepeat() {
    repeated := Repeat("a", 5)
    fmt.Println(repeated)
    // Output: aaaaa
}
