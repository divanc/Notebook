class User {
  name: string;
  email: string;
  age: number;

  constructor(name, email, age) {
    this.name = name;
    this.email = email;
    this.age = age;
    console.log("Thank You," + name + "!");
  }
}

let John = new User("John Doe", "yuraist@icloud.com", 70);
