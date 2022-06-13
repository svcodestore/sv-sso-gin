drop schema if exists sv_sso;
create schema sv_sso;
use sv_sso;

create table users
(
    id         bigint unsigned  not null,
    uuid       binary(16)       not null,
    login_id   varchar(16)      not null,
    password   varchar(1024),
    name       varchar(32),
    alias      varchar(32),
    phone      varchar(16),
    email      varchar(1024),
    lang       char(5) collate utf8mb4_0900_as_cs default 'zh_CN',
    status     tinyint unsigned not null          default 1,
    created_at datetime(6)      not null          default current_timestamp(6),
    created_by bigint unsigned  not null,
    updated_at datetime(6)      not null          default current_timestamp(6) on update current_timestamp(6),
    updated_by bigint unsigned  not null,
    constraint users_chk_status check ( status = 0 or status = 1),
    constraint users_chk_lang check ( lang = 'zh_CN' or lang = 'zh_TW' or lang = 'zh_HK' or lang = 'en_US'),
    constraint users_fk_created_by foreign key (created_by) references `users` (id) on update cascade on delete restrict,
    constraint users_fk_updated_by foreign key (updated_by) references `users` (id) on update cascade on delete restrict,
    unique key `user_uuid_uindex` (`uuid`),
    unique key `user_login_id_uindex` (`login_id`),
    index `user_uuid_index` (`uuid`),
    primary key (id)
) engine = InnoDB;

create table organizations
(
    id         bigint unsigned  not null,
    code       varchar(64)      not null,
    name       varchar(255)     not null,
    status     tinyint unsigned not null default 1,
    created_at datetime(6)      not null default current_timestamp(6),
    created_by bigint unsigned  not null,
    updated_at datetime(6)      not null default current_timestamp(6) on update current_timestamp(6),
    updated_by bigint unsigned  not null,
    primary key (id),
    constraint organizations_chk_status check ( status = 0 or status = 1),
    constraint organizations_fk_created_by foreign key (created_by) references `users` (id) on update cascade on delete restrict,
    constraint organizations_fk_updated_by foreign key (updated_by) references `users` (id) on update cascade on delete restrict,
    constraint unique organizations_code_uindex (code),
    constraint unique organizations_name_uindex (name)
) engine = InnoDB;

create table applications
(
    id            bigint unsigned  not null,
    code          varchar(64)      not null,
    name          varchar(255)     not null,
    internal_url  varchar(255),
    homepage_url  varchar(255),
    status        tinyint unsigned not null default 1,
    client_id     varchar(255)     not null,
    client_secret varchar(255),
    redirect_uris varchar(255)     not null,
    login_uris    varchar(255)     not null,
    token_format  varchar(100)              default 'JWT',
    created_at    datetime(6)      not null default current_timestamp(6),
    created_by    bigint unsigned  not null,
    updated_at    datetime(6)      not null default current_timestamp(6) on update current_timestamp(6),
    updated_by    bigint unsigned  not null,
    primary key (id),
    constraint applications_chk_status check (
            status = 0
            or status = 1
        ),
    constraint applications_fk_created_by foreign key (created_by) references `users` (id) on update cascade on delete restrict,
    constraint applications_fk_updated_by foreign key (updated_by) references `users` (id) on update cascade on delete restrict,
    constraint unique applications_code_unique_index_code (code),
    constraint unique applications_name_unique_index_code (name),
    constraint unique applications_code_unique_index_client_id (client_id)
) engine = InnoDB;

create table organization_application
(
    organization_id bigint unsigned  not null,
    application_id  bigint unsigned  not null,
    status          tinyint unsigned not null default 1,
    created_at      datetime(6)      not null default current_timestamp(6),
    created_by      bigint unsigned  not null,
    updated_at      datetime(6)      not null default current_timestamp(6) on update current_timestamp(6),
    updated_by      bigint unsigned  not null,
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
    application_id bigint unsigned  not null,
    user_id        bigint unsigned  not null,
    status         tinyint unsigned not null default 1,
    created_at     datetime(6)      not null default current_timestamp(6),
    created_by     bigint unsigned  not null,
    updated_at     datetime(6)      not null default current_timestamp(6) on update current_timestamp(6),
    updated_by     bigint unsigned  not null,
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
                         homepage_url,
                         client_id,
                         client_secret,
                         redirect_uris,
                         login_uris,
                         created_by,
                         updated_by) VALUE (
                                            0,
                                            'SSO',
                                            'SV SSO',
                                            'http://localhost:8000',
                                            '0ef9d7b504019278e740',
                                            '42fbdcb2b910024594c9be51463bbe4861f5b44a',
                                            'http://localhost:8000/callback',
                                            'http://localhost:3000/oauth2.0/authorize?',
                                            0,
                                            0
    ), (
        1515870566593069056,
        'ERP_WEB',
        'SV ERP WEB',
        'http://localhost:3100',
        '60f9bd80d01913d3c74e',
        '6ec3749d9bc70dbacaa58ed378243bb01c655ed3',
        'http://localhost:3100/callback',
        'http://localhost:3000/oauth2.0/authorize?',
        0,
        0
    ), (
        1518381705071689728,
        'SSO_LOGIN',
        'SSO LOGIN',
        'http://localhost:3000',
        'b4a970346a91d6467f47',
        'e5e009b4abdec1195235f1b918d40dd90740dbe6',
        'http://localhost:3000/callback',
        'http://localhost:3000/oauth2.0/authorize?',
        0,
        0
    ), (
        1518512539921547264,
        'SSO_AUTH',
        'SSO AUTHORIZATION',
        'http://localhost:8100',
        'fa29064eafbe1dcbfd29',
        'd882d54a53bfcbe4d8f5d6bf11b617f762810c4e',
        'http://localhost:8100/callback',
        'http://localhost:3000/oauth2.0/authorize?',
        0,
        0
    ), (
        1531087548372221952,
        'SV_RESOURCE',
        'SV RESOURCE',
        'http://localhost:3200',
        '7c23f4aa74bbe21a3834',
        'fbcf8d0b3194d513b553fd4dcf119f3e96910160',
        'http://localhost:3200/callback',
        'http://localhost:3000/oauth2.0/authorize?',
        0,
        0
    ), (
        1531101724570288128,
        'SV_IM',
        'SV IM',
        'http://localhost:3300',
        '771de5728d8cf2bc60c3',
        'cdceef55ee1d4df787672caba0b92effc845a4a1',
        'http://localhost:3300/callback',
        'http://localhost:3000/oauth2.0/authorize?',
        0,
        0
    );

insert into application_user(application_id, user_id, created_by, updated_by)
VALUES (0, 0, 0, 0);

insert into organization_application(organization_id, application_id, created_by, updated_by)
VALUES (0, 0, 0, 0);
