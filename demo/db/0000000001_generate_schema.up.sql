create table public.examples
(
    id         bigserial primary key,
    name       text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

alter table public.examples
    owner to postgres;

create index idx_example_deleted_at
    on public.examples (deleted_at);