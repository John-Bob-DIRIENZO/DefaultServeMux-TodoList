package demoHTTP

// TodoItem
// Si je veux qu'il soit "publique" il faut que son nom
// commence avec une majuscule.
// C'est pareil avec ses attributs, si je veux qu'ils soient
// serialized en JSON, leur nom doit commencer par une majuscule
// sinon ils seront considérés comme "privés"
type TodoItem struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
