-- testing out tsvector and tsquery
select colors, text, to_tsvector(colors || ' ' || text) @@ to_tsquery('b') from cards limit 10;
