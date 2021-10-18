package database


var insertPostPhones = `
INSERT INTO phones(country, state, countryCode, phoneNumber) VALUES(?,?,?,?)
`
var selectPhonesWithNoFilters = `
SELECT country, state, countryCode, phoneNumber FROM phones
`
var selectPhonesWith2Filters = `
SELECT country, state, countryCode, phoneNumber FROM phones WHERE country = ? AND state = ?
`
var selectPhonesStateFilter = `
SELECT country, state, countryCode, phoneNumber FROM phones WHERE state = ?
`
var selectPhonesCountryFilter = `
SELECT country, state, countryCode, phoneNumber FROM phones WHERE country = ?
`

var selectRegexByCountryCode = `
SELECT countryCode, regexCode FROM regex WHERE countryCode = ?
`