truncate table users;
truncate table userpayments;
truncate table usercurrencies;

# 1
insert into users(username, status) values('payera', 0);
# 2
insert into users(username, status) values('payerb', 1);
# 3
insert into users(username, status) values('payerc', 2);
# 4
insert into users(username, status) values('payerd', 3);
# 5
insert into users(username, status) values('payeea', 0);
# 6
insert into users(username, status) values('payeeb', 1);
# 7
insert into users(username, status) values('payeec', 2);
# 8
insert into users(username, status) values('payeed', 3);
# 9
insert into users(username, status) values('payeee', 2);
# 10
insert into users(username, status) values('payeef', 2);

insert into usercurrencies(userid, currency, balance) values(2, 'USD', 10000);
insert into usercurrencies(userid, currency, balance) values(3, 'USD', 10000);
insert into usercurrencies(userid, currency, balance) values(3, 'EUR', 10000);
insert into usercurrencies(userid, currency, balance) values(7, 'USD', 0);
insert into usercurrencies(userid, currency, balance) values(7, 'EUR', 1);
insert into usercurrencies(userid, currency, balance) values(9, 'USD', 1);
insert into usercurrencies(userid, currency, balance) values(9, 'EUR', -1);