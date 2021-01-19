INSERT INTO events
  (id, name, location, cents)
VALUES
  (1, 'party at netlify', 'san francisco, ca, usa', 0),
  (2, 'yoyo mama tambien', 'oakland, ca, usa', 1000),
  (3, 'festival of lights', 'london, UK', 5033),
  (4, 'oktoberfest', 'munich, germany', 32718)
;

INSERT INTO attendees
  (id, name, email)
VALUES
  (1, 'ryan neal', 'ryan@netlify.com'),
  (2, 'ryan neal', 'ryan@netlify.com'),
  (3, 'bruce wayne', 'bruce@wayne.com'),
  (4, 'barry allen', 'barry@jl.com'),
  (5, 'dinah lance', 'canary@jl.com'),
  (6, 'shiera hall', 'hawk@jl.com'),
  (7, 'pamela isley', 'ivy@villians.com'),
  (8, 'selina kyle', 'cat@villians.com'),
  (9, 'waylon jones', 'croc@villians.com'),
  (10, 'roman sionis', 'mask@villians.com'),
  (11, 'harvey dent', 'manyfaces@gotham.com'),
  (12, 'dorrance', 'bane@villians.com')
;


INSERT INTO tickets
  (event_id, attendee_id, status)
VALUES
  (1, 1, 'paid'),
  (1, 4, 'pending'),
  (2, 2, 'paid'),
  (3, 5, 'pending'),
  (2, 10, 'paid'),
  (4, 12, 'pending'),
  (2, 3, 'paid'),
  (2, 6, 'pending'),
  (2, 7, 'paid'),
  (4, 8, 'pending'),
  (4, 9, 'pending'),
  (4, 10, 'paid'),
  (4, 11, 'pending'),
  (4, 12, 'paid'),
  (4, 1, 'paid'),
  (3, 10, 'paid')
;

ALTER SEQUENCE  events_id_seq RESTART WITH 5;
ALTER SEQUENCE  attendees_id_seq RESTART WITH 13;
ALTER SEQUENCE  tickets_id_seq RESTART WITH 17;
