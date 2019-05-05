# Mongo DB

1) Download or do brew
2) It's kinda JSONy

## Beginning

1) You can launch in the background or foreground:

Foreground:

```console
mongod --config /usr/local/etc/mongod.conf
```

Background:

```console 
brew services start mongodb
```

2. Then launch it using `mongo`

What you can do:

* `show dbs`
* `use DBNAME` to create db
* `db` to whereami
* `db.createUser({})`

3. So let's create user

### Users

```js
db.createUser({
  user: "NAME",
  pwd: "PASS",
  role: ["readWrite" , "dbAdmin"]
});
```


4. and create collection...

### Collections

* `db.createCollection("COLLECTIONNAME")` is like a table
* `db.COLLECTION.insert({})` 
* `db.COLLECTION.find()`

and it is scalable!

### Collection edit

#### Write

For example:

`db.customers.insert([{first_name: "Regina", last_name: "Valetova"},{first_name: "Ivan", last_name: "Dorow", age: 13}])`

Or to make prettier:  `db.COLLECTION.find().pretty();`

#### Edit

5. In order to Match and update:

For example:

`db.customers.update({first_name: "Regina"},{first_name: "Regina", last_name: "Waletou", age: 14})`

##### $set

However, if we want to change just one cell, we can do:

`db.customers.update({first_name:"Regina"},{$set:{hasSpouse: true}})`

##### $inc

`db.customers.update({last_name: "Valetova"},{$inc:{age:1}})`


##### $set

Deleting an option:

`db.customers.update({first_name:"Regina"},{$unset:{hasSpouse: true}})`

#### upsert

If we want to match something and edit, but it isn't there, it won't make anything, yet we can create new object, in case none was found:

`db.customers.update({first_name:"Regina"},{$set:{hasSpouse: true},{upsert: true}})`

#### rename

if we want to rename the key:

`db.customers.update({first_name: "Brad"},{$rename: {"age": "old"}})`

#### remove

`db.customers.remove({first_name: "Darth"})`

#### $or

`db.customers.find({$or:[{first_name: "Luke"},{age: 22}]})`

#### $gt $lt

`db.customers.find({age:{$gt:14}})`

#### inner objects

`db.customers.find(["adress.city": "Boston"]);

### Sorting

`db.cutomers.find().sort({last_name: 1});`

Would sort alphabetically by last name, or `-1` if reversed

### Count

`db.customers.find().count()`

`db.customers.find().limit(4)`

### forEach

`db.customers.find().forEach(function(doc){print("Customer Name: " + doc.first_name)});`
