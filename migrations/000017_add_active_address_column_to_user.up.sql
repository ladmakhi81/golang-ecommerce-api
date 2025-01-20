ALTER TABLE
    _users
ADD
    active_address INT DEFAULT NULL REFERENCES _user_addresses(id)