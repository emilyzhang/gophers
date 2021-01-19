CREATE ROLE root WITH LOGIN SUPERUSER PASSWORD 'plaintext-secrets-are-safe';

CREATE TABLE IF NOT EXISTS events (
  id BIGSERIAL primary key,
  name varchar (256) not null,
  location varchar (256) not null,
  cents integer default 0,
  disabled boolean default FALSE
);

CREATE TABLE IF NOT EXISTS attendees (
  id BIGSERIAL primary key,
  name varchar (256) not null,
  email varchar (256) not null,
  disabled boolean default FALSE
);

CREATE TABLE IF NOT EXISTS tickets (
  id BIGSERIAL primary key,
  event_id bigint references events (id),
  attendee_id bigint references attendees (id),
  status varchar (256) not null
);
