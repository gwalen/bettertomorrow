create table if not exists companies(
	id BIGINT not null,
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