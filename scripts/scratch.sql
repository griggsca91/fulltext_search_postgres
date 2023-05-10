-- testing out tsvector and tsquery
select colors, text, 
to_tsvector(colors || ' ' || text) @@ to_tsquery('combat & damage'), 
ts_rank(to_tsvector(colors || ' ' || text), to_tsquery('combat & damage'))
from cards limit 10


-- adding uuid to the vector 996b9acb-5654-44a6-a44d-22790e77311d
select 
name,
colors, 
text, 
ts_rank(to_tsvector(name || ' ' || colors || ' ' || text), to_tsquery('wormfang & crab')) as rank
from cards
where to_tsvector(name || ' ' || colors || ' ' || text) @@ to_tsquery('wormfang & crab')
order by rank desc
limit 10;
