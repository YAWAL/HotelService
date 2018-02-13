-- queries for creating tables
CREATE TABLE IF NOT EXISTS hotel_rooms (
  number INT,
  room_quantity INT,
  is_free BOOLEAN
);
CREATE TABLE IF NOT EXISTS tenants (
  id INT,
  name VARCHAR(20),
  last_name VARCHAR(30)
);
CREATE TABLE IF NOT EXISTS rents(
  id INT,
  hotel_room_num INT,
  tenant_id INT,
  start_day TIMESTAMP,
  end_day TIMESTAMP
);

-- adding primary keys into tables
ALTER TABLE hotel_rooms ADD PRIMARY KEY (number);
ALTER TABLE tenants ADD PRIMARY KEY (id);
ALTER TABLE rents ADD PRIMARY KEY (id);

