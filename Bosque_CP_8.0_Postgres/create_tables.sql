create table metadados(
	nome varchar(30) default 'Floresta Sintá(c)tica',
	autor varchar(40) default 'Rui Vilela - Linguateca, pólo de Braga',
        email varchar(40) default 'ruivilela@di.uminho.pt',
        url varchar(40) default 'http://linguateca.di.uminho.pt',
	data_sql timestamp default now(),
        data_script_sql timestamp default '2008-10-13',
	fonte varchar(30) default 'CetemPúblico',
        versao varchar(4) default '8.0',
	formato varchar(80) default 'SQL (PostgreSQL) (Importado do Tiger-XML)',
	adicional text default 'Versão SQL em testes'
);

insert into metadados default values;

create table arvores(
	id int,
	referencia varchar(10),
	n int,
	cad varchar(20),
	sec varchar(10),
	sem varchar(10),
	texto text,
	analise int
);

create table extras(
        et varchar(10),
	pt varchar(200),
	en varchar(200)
);

create table funcoes(
        et varchar(10),
	pt varchar(200),
	en varchar(200)
);

create table formas(
        et varchar(10),
	pt varchar(200),
	en varchar(200)
);

create table ramos(
	id int,
	arvore int not null,
	palavra varchar(120),
	lema varchar(120),
	funcao varchar(10),
	forma varchar(10),
	morfo varchar(20),
	pai int
);

create table ramos_extras(    /* A arvore pode ter vários parametro extra */
	ramo int not null,
	extra varchar(10) not null
);

