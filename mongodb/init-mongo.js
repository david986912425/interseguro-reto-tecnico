db = db.getSiblingDB('intersegurodb')

db.createUser({
    user: "david",
    pwd: "secret",
    roles: [
        {
            role: 'readWrite',
            db: 'intersegurodb'
        },
    ],
});

db.createCollection('users');

db.users.insertOne({
    name: 'David',
    lastname: 'Mamani',
    email: 'dm7659746@gmail.com',
    password: '$2a$12$s13hbY2sF5Zouhrw.Gri9.DdzsStJJldJSQ4AeMKBwzVroNPjfGaG',
    role: 'admin'
});