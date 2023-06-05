CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create type news_state as enum ('draft', 'published', 'drop');

create type like_object_type as enum ('news', 'comment');

create type entities as enum ('previewImage', 'content', 'comment');

create table files (
    id uuid default uuid_generate_v4() not null primary key ,
    mime_type varchar(256) not null,
    bucket_name varchar(512) not null,
    file_name varchar(1024) not null,
    entity entities not null,
    created_at timestamp default now()             not null,
    updated_at timestamp default now(),
    deleted_at timestamp
);

create table news (
    id uuid default uuid_generate_v4() not null primary key ,
    created_at timestamp default now()             not null,
    updated_at timestamp default now(),
    deleted_at timestamp,
    title text not null,
    description text,
    preview_image uuid,
    content_file uuid references files(id),
    user_created uuid not null,
    user_updated uuid not null,
    user_deleted uuid,
    state news_state default 'draft',
    foreign key (preview_image) references files(id),
    foreign key (content_file) references files(id)
);

create table comments (
    id uuid default uuid_generate_v4() not null primary key ,
    user_created uuid not null,
    news_id uuid not null,
    comment_body text not null,
    created_at timestamp default now()             not null,
    updated_at timestamp default now(),
    deleted_at timestamp,
    parent_id uuid,
    foreign key (news_id) references news(id)
);

create table likes (
    id uuid default uuid_generate_v4() not null primary key ,
    user_created uuid not null,
    object_id uuid,
    object_type like_object_type,
    is_active boolean default true,
    foreign key (object_id) references news(id),
    foreign key (object_id) references comments(id)
);