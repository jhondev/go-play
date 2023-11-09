CREATE TABLE Locations (
    ID VARCHAR(255) PRIMARY KEY,
    City VARCHAR(255),
    Country VARCHAR(255),
    Address VARCHAR(255),
    ZipCode VARCHAR(255)
);
CREATE TABLE Owners (
    Email VARCHAR(255) PRIMARY KEY,
    FirstName VARCHAR(255),
    LastName VARCHAR(255),
    Phone VARCHAR(255),
    LocationID INT,
    FOREIGN KEY (LocationID) REFERENCES Locations(ID)
);
CREATE TABLE Companies (
    TaxNumber VARCHAR(255) PRIMARY KEY,
    OwnerID INT,
    Name VARCHAR(255),
    LocationID INT,
    FOREIGN KEY (OwnerID) REFERENCES Owners(Email),
    FOREIGN KEY (LocationID) REFERENCES Locations(ID)
);
CREATE TABLE Franchise (
    ID VARCHAR(255) PRIMARY KEY,
    URL VARCHAR(255),
    Name VARCHAR(255),
    LocationID INT,
    CompanyID INT,
    FOREIGN KEY (CompanyID) REFERENCES Companies(TaxNumber),
    FOREIGN KEY (LocationID) REFERENCES Locations(ID)
);