use jamitizedb;

const collections = db.getCollectionNames();
collections.forEach(function(collName) {
    db[collName].drop();
    print("Dropped collection: " + collName);
});

db.createCollection("users");
db.createCollection("sessions");
db.createCollection("messages");


db.users.createIndex({ email: 1 }, { unique: true });

db.users.insertMany([
    {
        name: "Vishesh Modi",
        email: "modivishesh8@gmail.com",
        instrument: "Guitar",
        genres: ["Indie", "Pop"],
        city: "Sydney",
        experience: 3,
        availableForJam: true,
        bio: "Looking to jam on weekends.",
        socials: {
            instagram: "@vjam",
            youtube: "youtube.com/vjam"
        }
    },
    {
        name: "Alex Johnson",
        instrument: "Drums",
        genres: ["Rock", "Jazz"],
        city: "Melbourne",
        experience: 5,
        availableForJam: false
    }
]);

print("Setup complete!");
