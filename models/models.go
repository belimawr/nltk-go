package models

type Funcoes struct {
	Et string `db:"et"`
	Pt string `db:"pt"`
	En string `db:"en"`
}

type Formas struct {
	Et string `db:"et"`
	Pt string `db:"pt"`
	En string `db:"en"`
}

type Extras struct {
	Et string `db:"et"`
	Pt string `db:"pt"`
	En string `db:"en"`
}

type Arvores struct {
	ID         int    `db:"id"`
	Referencia string `db:"referencia"`
	N          int    `db:"n"`
	Cad        string `db:"cad"`
	Sec        string `db:"sec"`
	Sem        string `db:"sem"`
	Texto      string `db:"texto"`
	Analise    string `db:"analise"`
}

type Ramos struct {
	Id      int    `db:"id"`
	Arvore  int    `db:"arvore"`
	Palavra string `db:"palavra"`
	Lema    string `db:"lema"`
	Funcao  string `db:"funcao"`
	Forma   string `db:"forma"`
	Morfo   string `db:"morfo"`
	Pai     string `db:"pai"`
}

type RamosExtra struct {
	Ramo  int    `db:"ramo"`
	Extra string `db:"extra"`
}
