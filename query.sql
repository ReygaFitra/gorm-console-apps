create table products(
	product_code varchar(10) not null primary key,
	product_name varchar(200) not null,
	stock smallint,
	price integer
);

create table orders (
	id_order serial primary key,
	order_date timestamp,
	payment_method varchar(100),
	order_total integer
);

create table order_details(
	id_order_detail serial primary key not null,
	id_order integer not null,
	product_code varchar(100) not null,
	qty smallint,
	constraint id_order_foreign foreign key (id_order) references orders(id_order) on update cascade on delete restrict,
	constraint product_code_foreign foreign key (product_code) references products(product_code) on update cascade on delete restrict
)