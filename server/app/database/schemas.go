package database


var insertPostSchema = `
INSERT INTO phones(country, state, countryCode, phoneNumber) VALUES($1,$2,$3,$4)
`
var selectPostedSchema = `
SELECT * FROM phones
`