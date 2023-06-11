CREATE TABLE post(
    post_id serial PRIMARY KEY,
    post_title VARCHAR(50) NOT NULL,
    post_content TEXT,
    post_img TEXT,
    email VARCHAR(45) not null,
    constraint email foreign key (email) references blogger
    (email) on delete cascade on update cascade
);