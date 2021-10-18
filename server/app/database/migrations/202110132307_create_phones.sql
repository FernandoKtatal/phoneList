CREATE TABLE IF NOT EXISTS phones
(
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    country TEXT,
    state bool,
    countryCode INT,
    phoneNumber INT
)