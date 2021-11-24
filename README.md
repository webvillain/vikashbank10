sql query and statement which are used by me into this appliation

........


created by vikash parashar on 23 nov 2021 , 07:38


// creating table statment
`create table if not exists users(id integer primary key , firstname text not null , lastname  text not null ,email text not null );`



// insert user into database
 `insert into users (firstname , lastname , email)values(?,?,?);`



// get all users query
`select * from users ;`




// get single user
`select * from users where id = ?`

// delete a user from database
`delete from users where id = ?;`
``
// update a user into database
`update users set firstname = ?,lastname = ?,email = ? where id = ?;`

3 out of 5 are working

update user is working 
create user is working
delete user is working