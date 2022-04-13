drop schema if exists sv_sso;
create schema sv_sso;
use sv_sso;

create table users
(
    id         bigint      not null,
    uuid       binary(16)  not null,
    login_id   varchar(16) not null,
    password   varchar(1024),
    name       varchar(32),
    alias      varchar(32),
    phone      varchar(16),
    email      varchar(1024),
    lang       char(5) collate utf8mb4_0900_as_cs default 'zh_CN',
    status     tinyint(1)  not null               default 1,
    created_at datetime(6) not null               default current_timestamp(6),
    created_by bigint      not null,
    updated_at datetime(6) not null               default current_timestamp(6) on update current_timestamp(6),
    updated_by bigint      not null,
    constraint users_chk_status check ( status = 0 or status = 1),
    constraint users_chk_lang check ( lang = 'zh_CN' or lang = 'zh_TW' or lang = 'zh_HK' or lang = 'en_US'),
    constraint users_fk_created_by foreign key (created_by) references `users` (id) on update cascade on delete restrict,
    constraint users_fk_updated_by foreign key (updated_by) references `users` (id) on update cascade on delete restrict,
    unique key `user_login_id_uindex` (`login_id`),
    index `user_uuid_index` (`uuid`),
    primary key (id)
) engine = InnoDB;

create table organizations
(
    id         bigint      not null,
    code       varchar(64) not null,
    name       varchar(255),
    status     tinyint(1)  not null default 1,
    created_at datetime(6) not null default current_timestamp(6),
    created_by bigint      not null,
    updated_at datetime(6) not null default current_timestamp(6) on update current_timestamp(6),
    updated_by bigint      not null,
    primary key (id),
    constraint organizations_chk_status check ( status = 0 or status = 1),
    constraint organizations_fk_created_by foreign key (created_by) references `users` (id) on update cascade on delete restrict,
    constraint organizations_fk_updated_by foreign key (updated_by) references `users` (id) on update cascade on delete restrict,
    constraint unique organizations_code_uindex (code)
) engine = InnoDB;

create table applications
(
    id            bigint      not null,
    code          varchar(64) not null,
    name          varchar(255),
    internal_url  varchar(255),
    homepage_url  varchar(255),
    status        tinyint(1)  not null default 1,
    client_id     varchar(255),
    client_secret varchar(255),
    redirect_uris varchar(255),
    token_format  varchar(100)         default 'JWT',
    created_at    datetime(6) not null default current_timestamp(6),
    created_by    bigint      not null,
    updated_at    datetime(6) not null default current_timestamp(6) on update current_timestamp(6),
    updated_by    bigint      not null,
    primary key (id),
    constraint applications_chk_status check (
                status = 0
            or status = 1
        ),
    constraint applications_fk_created_by foreign key (created_by) references `users` (id) on update cascade on delete restrict,
    constraint applications_fk_updated_by foreign key (updated_by) references `users` (id) on update cascade on delete restrict,
    constraint unique applications_code_uindex (code)
) engine = InnoDB;

create table organization_application
(
    organization_id bigint      not null,
    application_id  bigint      not null,
    status          tinyint(1)  not null default 1,
    created_at      datetime(6) not null default current_timestamp(6),
    created_by      bigint      not null,
    updated_at      datetime(6) not null default current_timestamp(6) on update current_timestamp(6),
    updated_by      bigint      not null,
    primary key (organization_id, application_id),
    constraint organization_application_chk_status check ( status = 0 or status = 1),
    constraint organization_application_fk_organization foreign key (organization_id) references `organizations` (id) on update cascade on delete cascade,
    constraint organization_application_fk_application foreign key (application_id) references `applications` (id) on update cascade on delete cascade,
    constraint organization_application_fk_created_by foreign key (created_by) references `users` (id) on update cascade on delete restrict,
    constraint organization_application_fk_updated_by foreign key (updated_by) references `users` (id) on update cascade on delete restrict,
    index organization_application_organization_id_index (organization_id),
    index organization_application_application_id_index (application_id)
) engine = InnoDB;


create table application_user
(
    application_id bigint      not null,
    user_id        bigint      not null,
    status         tinyint(1)  not null default 1,
    created_at     datetime(6) not null default current_timestamp(6),
    created_by     bigint      not null,
    updated_at     datetime(6) not null default current_timestamp(6) on update current_timestamp(6),
    updated_by     bigint      not null,
    primary key (application_id, user_id),
    constraint application_user_chk_status check ( status = 0 or status = 1),
    constraint application_user_fk_application foreign key (application_id) references `applications` (id) on update cascade on delete cascade,
    constraint application_user_fk_user foreign key (user_id) references `users` (id) on update cascade on delete cascade,
    constraint application_user_fk_created_by foreign key (created_by) references `users` (id) on update cascade on delete restrict,
    constraint application_user_fk_updated_by foreign key (updated_by) references `users` (id) on update cascade on delete restrict,
    index application_user_application_id_index (application_id),
    index application_user_user_id_index (user_id)
) engine = InnoDB;

insert into users(id,
                  uuid,
                  login_id,
                  password,
                  name,
                  created_by,
                  updated_by) VALUE (
                                     0,
                                     uuid_to_bin("00000000-0000-0000-0000-000000000000", true),
                                     'root',
                                     '$2y$10$3f6E142T20Aob3cGyJM/keBUPMZ3P1qsIrCysVngGDUZ2deezJUKG',
                                     'Root',
                                     0,
                                     0
    ), (
        1508366740931739648,
        uuid_to_bin("12f7b33c-4b86-4df8-a2bc-304c2b97b4ff", true),
        'admin',
        '$2y$10$3f6E142T20Aob3cGyJM/keBUPMZ3P1qsIrCysVngGDUZ2deezJUKG',
        'Admin',
        0,
        0
    );

insert into organizations(id, code, name, created_by, updated_by) VALUE (0, 'ROOT', 'root organization', 0, 0);

insert into applications(id,
                         code,
                         name,
                         client_id,
                         client_secret,
                         token_format,
                         created_by,
                         updated_by) VALUE (
                                            0,
                                            'SSO',
                                            'SV SSO',
                                            '0ef9d7b504019278e740',
                                            '42fbdcb2b910024594c9be51463bbe4861f5b44a',
                                            '',
                                            0,
                                            0
    );

insert into application_user(application_id, user_id, created_by, updated_by)
VALUES (0, 0, 0, 0);

insert into organization_application(organization_id, application_id, created_by, updated_by)
VALUES (0, 0, 0, 0);
