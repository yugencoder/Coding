USE address_service;

select cc.name AS name, rcu.earnings as earnings
FROM cities cc
         JOIN
     (
         select c.id, SUM(fare) as earnings
         from users u
                  JOIN rides r on u.id = r.user_id
                  JOIN cities c on c.id = u.city_id
         GROUP BY c.id
     ) AS rcu ON cc.id = rcu.id;


select c.name, SUM(r.fare)
from cities c
         JOIN users u ON u.city_id = c.id
         JOIN rides r ON r.user_id = u.id
GROUP BY  c.id,c.name;


