CREATE TABLE IF NOT EXISTS cards (
    cId INTEGER PRIMARY KEY, 
    name TEXT, 
    number TEXT, 
    date TEXT, 
    cvv INTEGER
);
CREATE TABLE IF NOT EXISTS logins (
    lId INTEGER PRIMARY KEY, 
    name TEXT, 
    login TEXT, 
    password TEXT
);
CREATE TABLE IF NOT EXISTS text_data (
    tId INTEGER PRIMARY KEY, 
    name TEXT, 
    data TEXT
);
CREATE TABLE IF NOT EXISTS binares_data (
    bId INTEGER PRIMARY KEY, 
    name TEXT, 
    data TEXT
);