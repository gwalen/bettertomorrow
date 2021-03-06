create table if not exists bettertomorrow.companies(
	id BIGINT not null AUTO_INCREMENT,
	name VARCHAR(255) not null,
	tax_id VARCHAR(255),
	created_at DATETIME NOT NULL,
	street VARCHAR(255),
	house_number VARCHAR(255),
	postal_code VARCHAR(32),
	city VARCHAR(255),
	country VARCHAR(255),
	PRIMARY KEY (id)
);

create table if not exists bettertomorrow.products(
	id BIGINT not null AUTO_INCREMENT,
	name VARCHAR(255) not null,
	price DECIMAL(19, 4),
	company_id BIGINT NOT NULL,
	created_at DATETIME NOT NULL,
	PRIMARY KEY (id)
);

alter table products add constraint fk_products__companies foreign key (company_id) references companies(id);


create table if not exists bettertomorrow.customers(
	id BIGINT not null AUTO_INCREMENT,
	first_name VARCHAR(255) not null,
	last_name VARCHAR(255) not null,
	created_at DATETIME NOT NULL,
	PRIMARY KEY (id)
);

create table if not exists bettertomorrow.wallets(
	id BIGINT not null AUTO_INCREMENT,
	currency VARCHAR(5) not null,
	amount DECIMAL(19, 4),
	customer_id BIGINT NOT NULL,
	created_at DATETIME NOT NULL,
	PRIMARY KEY (id)
);

alter table wallets add constraint fk_Wallets__customers foreign key (customer_id) references customers(id);