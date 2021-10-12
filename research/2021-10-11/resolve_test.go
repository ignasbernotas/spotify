package resolve

import (
   "testing"
)

func TestResolve(t *testing.T) {
   sol, err := newResolve()
   if err != nil {
      t.Fatal(err)
   }
   if err := sol.track("eef38251727f46c28eed9284b288024e"); err != nil {
      t.Fatal(err)
   }
}
