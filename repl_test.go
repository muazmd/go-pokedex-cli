package main

import "testing"

func TestCleanInput(t *testing.T){
  cases := []struct{
    input  string
    expected []string
  }{
    {
      input: "hello World",
      expected: []string{
        "hello",
        "world",
      },

    },
  }
  for _, cs := range cases {
    actual := cleanInput(cs.input)
    expected := cs.expected

    if len(actual) != len(expected){
      t.Errorf("Length is not matching %v vs %v",len(actual), len(expected))

    continue
    }

    for i := range actual{
      actualWord := actual[i]
      expectedWord := expected[i]
      if actualWord != expectedWord{
        t.Errorf("Words are not the same %v vs %v",actualWord, expectedWord)
      }
    }

  }
}
