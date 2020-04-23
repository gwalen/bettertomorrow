create table if not exists bettertomorrow.employees(
	id BIGINT not null AUTO_INCREMENT,
	first_name VARCHAR(255),
	last_name VARCHAR(255),
	created_at DATETIME NOT NULL,
	PRIMARY KEY (id)
);

create table if not exists bettertomorrow.roles(
	id BIGINT not null AUTO_INCREMENT,
	name VARCHAR(255) not null,
	department_name VARCHAR(255),
	created_at DATETIME NOT NULL,
	started_at DATETIME NOT NULL,
	finished_at DATETIME,
	employee_id BIGINT NOT NULL,
	PRIMARY KEY (id)
);

alter table roles add constraint fk_roles__employees foreign key (employee_id) references employees(id);