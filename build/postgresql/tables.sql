create table public.messges
(
    id       uuid not null
        constraint messges_pk
            primary key,
    title    text,
    mask     varchar(255),
    channels text[],
    payload  bytea
);

alter table public.messges
    owner to cronos;

create table public.delay
(
    id         uuid      not null
        constraint delay_pk
            primary key,
    message_id uuid      not null
        constraint delay_messges_id_fk
            references public.messges,
    datetime   timestamp not null
);

alter table public.delay
    owner to cronos;
