
	package main
	
	import "testing"
	
	func TestEliminarTodosConsecutivos(t *testing.T) {
		l := Desde(1, 7, 7, 7, 2)
	
		borrados := l.EliminarTodos(7)
	
		if borrados != 3 {
			t.Errorf("esperaba 3 borrados, obtuvo %d", borrados)
		}
	
		esperado := "1 -> 2 -> nil"
	
		if l.String() != esperado {
			t.Errorf("esperaba %s, obtuvo %s", esperado, l.String())
		}
	}
