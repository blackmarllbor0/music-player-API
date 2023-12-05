create table if not exists users
(
  id serial primary key,
  username varchar(50) not null,
  password_hash varchar(255) not null,
  email varchar(100) not null,
  is_deleted boolean default false,
  registered_data timestamp default current_timestamp,
  create_at timestamp default current_timestamp,
  update_at timestamp
);

create table if not exists authors
(
  id serial primary key,
  author_name varchar(100) not null,
  is_deleted boolean default false,
  create_at timestamp default current_timestamp,
  update_at timestamp
  -- default on update current_timestamp
);

create table if not exists genres
(
  id serial primary key,
  genre_name varchar(50) not null,
  is_deleted boolean default false,
  create_at timestamp default current_timestamp,
  update_at timestamp
);

create table if not exists albums
(
  id serial primary key,
  album_name varchar(100) not null,
  author_id integer not null,
  release_year year,
  is_deleted boolean default false,
  create_at timestamp default current_timestamp,
  update_at timestamp
);

create table if not exists tracks
(
  id serial primary key,
  track_name varchar(100) not null,
  author_id integer not null,
  genre_id integer not null,
  album_id integer not null,
  duration interval,
  is_deleted boolean default false,
  release_date date,
  create_at timestamp default current_timestamp,
  update_at timestamp,
  -- default on update current_timestamp

  foreign key (author_id) references authors(id),
  foreign key (genre_id) references genres(id),
  foreign key (album_id) references albums(id)
);

create table if not exists collections
(
  id serial primary key,
  collection_name varchar(50) not null,
  is_deleted boolean default false,
  create_at timestamp default current_timestamp,
  update_at timestamp,
  -- default on update current_timestamp
);