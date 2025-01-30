CREATE TABLE Users
(
    id serial not null unique,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    middle_name varchar(255) not null
);

CREATE TABLE Tickets
(
    id serial not null unique,
    point_start varchar(255) not null,
    point_end varchar(255) not null,
    order_number varchar(255) not null,
    service_provider varchar(255) not null,
    date_start date not null,
    date_finish date not null,
    created_at date not null
);

CREATE TABLE Documents
(
    id serial not null unique,
    type_document varchar(255) not null,
    number varchar(255) not null,
    user_id int references Users(id) on delete cascade not null
);

CREATE TABLE TicketUsers
(
    id serial not null unique,
    ticket_id int references Tickets(id) on delete cascade not null,
    user_id int references Users(id) on delete cascade not null,
    done boolean not null
);