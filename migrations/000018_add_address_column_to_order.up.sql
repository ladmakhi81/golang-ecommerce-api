ALTER TABLE
    _orders
ADD
    address_id INT NOT NULL REFERENCES _user_addresses(id)