db = db.getSiblingDB("jamitizedb");

db.createUser({
  user: "jamitize-user",
  pwd: "jamitize-user",
  roles: [
    { role: "readWrite", db: "jamitizedb" }
  ]
});