CREATE TABLE servers
(
    id            serial       not null unique,
    name          text         not null
);

CREATE TABLE bandwidths
(
    server_id       integer references servers (id) on delete cascade       not null,
    value           float                                                   not null,
    interface_name  text                                                    unique,
    value_id        serial                                                  not null unique
);