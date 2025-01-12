CREATE TABLE Users (
    ID bigserial PRIMARY KEY,
    Username varchar(255) NOT NULL UNIQUE,
    Password varchar(255) NOT NULL,
    Fullname varchar(255) NOT NULL,
    Email varchar(255) NOT NULL UNIQUE,
    PhoneNumber varchar(20) NOT NULL UNIQUE,
    Gender varchar(10) NOT NULL,
    Role varchar(10) NOT NULL,
    CreateDate timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE Posts (
    ID bigserial PRIMARY KEY,
    Title varchar(255) NOT NULL,
    Content text NOT NULL,
    CreateDate timestamp NOT NULL DEFAULT NOW(), 
    EditDate timestamp NOT NULL DEFAULT NOW(),
    UserID bigint REFERENCES Users(ID) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TRIGGER update_timestamp BEFORE UPDATE ON Posts
FOR EACH ROW EXECUTE PROCEDURE moddatetime(EditDate);

CREATE TABLE Albums (
    ID bigserial PRIMARY KEY,
    Name varchar(255) NOT NULL,
    AuthorName varchar(255) NOT NULL,
    ReleaseDate timestamp NOT NULL,
    Type varchar(255) NOT NULL,
    Description text NOT NULL,
    Link varchar(255) NOT NULL,
    CreateDate timestamp NOT NULL DEFAULT NOW(), 
    EditDate timestamp NOT NULL DEFAULT NOW(),
    UserID bigint REFERENCES Users(ID) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE TRIGGER update_timestamp BEFORE UPDATE ON Albums
FOR EACH ROW EXECUTE PROCEDURE moddatetime(EditDate);

CREATE TABLE Musics (
    ID bigserial PRIMARY KEY,
    Name varchar(255) NOT NULL,
    Type varchar(255) NOT NULL,
    -- add link
    Link varchar(255) NOT NULL,
    CreateDate timestamp NOT NULL DEFAULT NOW(), 
    EditDate timestamp NOT NULL DEFAULT NOW(),
    AlbumID bigint REFERENCES Albums(ID) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE TRIGGER update_timestamp BEFORE UPDATE ON Musics
FOR EACH ROW EXECUTE PROCEDURE moddatetime(EditDate);


CREATE TABLE MusicInfo (
    ID bigserial PRIMARY KEY,
    ArtistName varchar(255) NOT NULL,
    Role varchar(255) NOT NULL,
    Type varchar(255) NOT NULL,
    CreateDate timestamp NOT NULL DEFAULT NOW(), 
    EditDate timestamp NOT NULL DEFAULT NOW(),
    MusicID bigint REFERENCES Musics(ID) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE TRIGGER update_timestamp BEFORE UPDATE ON MusicInfo
FOR EACH ROW EXECUTE PROCEDURE moddatetime(EditDate);

CREATE TABLE Rooms (
    ID bigserial PRIMARY KEY,
    Name varchar(255) NOT NULL UNIQUE,
    Address varchar(255) NOT NULL,
    -- edit description
    Description text NOT NULL,
    -- add column
    Price double precision NOT NULL,
    CreateDate timestamp NOT NULL DEFAULT NOW(),
    EditDate timestamp NOT NULL DEFAULT NOW(),
    UserID bigint REFERENCES Users(ID) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE TRIGGER update_timestamp BEFORE UPDATE ON Rooms
FOR EACH ROW EXECUTE PROCEDURE moddatetime(EditDate);

CREATE TABLE Orders (
    ID bigserial PRIMARY KEY,
    -- FromDate timestamp NOT NULL DEFAULT NOW(),,
    -- ToDate timestamp NOT NULL DEFAULT NOW(),
    FromTo tsrange NOT NULL,
    State varchar(255) NOT NULL,
    Note text,
    CheckInDate timestamp NOT NULL,
    CheckOutDate timestamp NOT NULL,
    TotalPrice double precision NOT NULL,
    CreateDate timestamp NOT NULL DEFAULT NOW(),
    EditDate timestamp NOT NULL DEFAULT NOW(),
    UserID bigint REFERENCES Users(ID) ON DELETE CASCADE ON UPDATE CASCADE,
    RoomID bigint REFERENCES Rooms(ID) ON DELETE CASCADE ON UPDATE CASCADE,
    EXCLUDE USING GiST (ID with =, FromTo with &&)
);
CREATE TRIGGER update_timestamp BEFORE UPDATE ON Orders
FOR EACH ROW EXECUTE PROCEDURE moddatetime(EditDate);