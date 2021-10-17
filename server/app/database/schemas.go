package database


var insertPostPhones = `
INSERT INTO phones(country, state, countryCode, phoneNumber) VALUES($1,$2,$3,$4)
`
var selectPhonesWithNoFilters = `
SELECT * FROM phones
`
var selectPhonesWith2Filters = `
SELECT * FROM phones WHERE country = ? AND state = ?
`
var selectPhonesStateFilter = `
SELECT * FROM phones WHERE state = ?
`
var selectPhonesCountryFilter = `
SELECT * FROM phones WHERE country = ?
`

var selectRegexByCountryCode = `
SELECT * FROM regex WHERE countryCode = $1
`