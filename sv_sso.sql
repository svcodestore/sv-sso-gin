drop schema if exists sv_sso;
create schema sv_sso;
use sv_sso;
create table users (
    id bigint not null,
    uuid binary(16) not null,
    login_id varchar(16) not null,
    password varchar(1024),
    name varchar(32),
    alias varchar(32),
    phone varchar(16),
    email varchar(1024),
    lang char(5) collate utf8mb4_0900_as_cs default 'zh_CN',
    status tinyint(1) not null default 1,
    created_at datetime(6) not null default current_timestamp(6),
    created_by bigint not null,
    updated_at datetime(6) not null default current_timestamp(6) on update current_timestamp(6),
    updated_by bigint not null,
    constraint users_chk_status check (
        status = 0
        or status = 1
    ),
    constraint users_chk_lang check (
        lang = 'zh_CN'
        or lang = 'zh_TW'
        or lang = 'zh_HK'
        or lang = 'en_US'
    ),
    unique key `user_login_id_uindex` (`login_id`),
    index `user_uuid_index` (`uuid`),
    primary key (id)
) engine = InnoDB;
insert into users(
        id,
        uuid,
        login_id,
        password,
        name,
        created_by,
        updated_by
    ) VALUE (
        0,
        uuid_to_bin("00000000-0000-0000-0000-000000000000", true),
        'root',
        '$2y$10$3f6E142T20Aob3cGyJM/keBUPMZ3P1qsIrCysVngGDUZ2deezJUKG',
        'Root',
        0,
        0
    );
insert into users(
        id,
        uuid,
        login_id,
        password,
        name,
        alias,
        created_by,
        updated_by
    ) VALUE (
        1508366740931739648,
        uuid_to_bin("12f7b33c-4b86-4df8-a2bc-304c2b97b4ff", true),
        'admin',
        '$2y$10$3f6E142T20Aob3cGyJM/keBUPMZ3P1qsIrCysVngGDUZ2deezJUKG',
        'Admin',
        'Administrator',
        0,
        0
    );
create table organizations (
    id bigint not null,
    code varchar(64) not null,
    name varchar(255),
    status tinyint(1) not null default 1,
    created_at datetime(6) not null default current_timestamp(6),
    created_by bigint not null,
    updated_at datetime(6) not null default current_timestamp(6) on update current_timestamp(6),
    updated_by bigint not null,
    primary key (id),
    constraint organizations_chk_status check (
        status = 0
        or status = 1
    ),
    constraint unique organizations_code_uindex (code)
) engine = InnoDB;
insert into organizations(id, code, name, created_by, updated_by) VALUE (0, 'ROOT', 'root organization', 0, 0);
create table applications (
    id bigint not null,
    code varchar(64) not null,
    name varchar(255),
    internal_url varchar(255),
    homepage_url varchar(255),
    status tinyint(1) not null default 1,
    client_id varchar(255),
    client_secret varchar(255),
    redirect_uris varchar(255),
    token_format varchar(100) default 'JWT',
    created_at datetime(6) not null default current_timestamp(6),
    created_by bigint not null,
    updated_at datetime(6) not null default current_timestamp(6) on update current_timestamp(6),
    updated_by bigint not null,
    primary key (id),
    constraint applications_chk_status check (
        status = 0
        or status = 1
    ),
    constraint unique applications_code_uindex (code)
) engine = InnoDB;
insert into applications(
        id,
        code,
        name,
        client_id,
        client_secret,
        token_format,
        created_by,
        updated_by
    ) VALUE (
        0,
        'SSO',
        'SV SSO',
        '0ef9d7b504019278e740',
        '42fbdcb2b910024594c9be51463bbe4861f5b44a',
        '',
        0,
        0
    );