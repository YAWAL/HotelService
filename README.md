# HotelService

Used libraries:

ORM - GORM; 

Routing - Gorilla Mux Toolkit;

Dependency management - dep;

Database - PostgreSQL 9.6;

To run this app you need:

1) install PostgreSQL 9.6 and create test database using script *https://github.com/YAWAL/HotelService/blob/v0.0.1/database/hoteldb.sql*
Database name - hoteldb; password - root;
initial database has some test data;

2) download code from this branch and run it. Default url: *localhost:8000/api*;
3) use following URLs and methods:
- to display list of ordered rooms with tenant info:
*localhost:8000/api/rents* **GET**;
- to add rent item:
*localhost:8000/api/rents/{hotelNum:[0-9]+}/{tenId:[0-9]+}* **POST**,
where hotelNum - available hotel room number, tenId - id of tenant in tenants table (table tenants should be filled before creating rent);
- to update rent item:
*localhost:8000/api/rents/{hotelNum:[0-9]+}/{tenId:[0-9]+}* **PUT**,
where hotelNum - hotel room number, tenId - correct id number of tenant, instead of wrong filled before;
- to delete rent item:
*localhost:8000/api/rents/{hotelNum}* **DELETE**,
where hotelNum - hotel room number
