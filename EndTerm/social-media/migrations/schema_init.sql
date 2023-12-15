
CREATE TABLE posts
(
    id           BIGSERIAL PRIMARY KEY,
    user_id      BIGINT not null,
    author       varchar(100) default '',
    title        VARCHAR(100) default '',
    text         VARCHAR(2000) default '',
    image        varchar(100) default '',
    date_created timestamp not null default now()
);

CREATE INDEX posts_id_idx on posts(user_id);