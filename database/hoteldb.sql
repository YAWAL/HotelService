--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.7
-- Dumped by pg_dump version 9.6.7

-- Started on 2018-02-18 16:22:41

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE hoteldb;
--
-- TOC entry 2139 (class 1262 OID 16384)
-- Name: hoteldb; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE hoteldb WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'Ukrainian_Ukraine.1251' LC_CTYPE = 'Ukrainian_Ukraine.1251';


ALTER DATABASE hoteldb OWNER TO postgres;

\connect hoteldb

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 1 (class 3079 OID 12387)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2142 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 185 (class 1259 OID 16385)
-- Name: hotel_rooms; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE hotel_rooms (
    number integer NOT NULL,
    room_quantity integer,
    is_free boolean
);


ALTER TABLE hotel_rooms OWNER TO postgres;

--
-- TOC entry 188 (class 1259 OID 24600)
-- Name: rents; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE rents (
    id integer NOT NULL,
    hotel_room_num integer,
    tenant_id integer
);


ALTER TABLE rents OWNER TO postgres;

--
-- TOC entry 187 (class 1259 OID 24598)
-- Name: rents_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE rents_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE rents_id_seq OWNER TO postgres;

--
-- TOC entry 2143 (class 0 OID 0)
-- Dependencies: 187
-- Name: rents_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE rents_id_seq OWNED BY rents.id;


--
-- TOC entry 186 (class 1259 OID 16390)
-- Name: tenants; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tenants (
    id integer NOT NULL,
    name character varying(20),
    last_name character varying(30)
);


ALTER TABLE tenants OWNER TO postgres;

--
-- TOC entry 2009 (class 2604 OID 24603)
-- Name: rents id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY rents ALTER COLUMN id SET DEFAULT nextval('rents_id_seq'::regclass);


--
-- TOC entry 2131 (class 0 OID 16385)
-- Dependencies: 185
-- Data for Name: hotel_rooms; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO hotel_rooms (number, room_quantity, is_free) VALUES (2, 1, true);
INSERT INTO hotel_rooms (number, room_quantity, is_free) VALUES (3, 1, true);
INSERT INTO hotel_rooms (number, room_quantity, is_free) VALUES (5, 2, true);
INSERT INTO hotel_rooms (number, room_quantity, is_free) VALUES (6, 2, true);
INSERT INTO hotel_rooms (number, room_quantity, is_free) VALUES (8, 3, true);
INSERT INTO hotel_rooms (number, room_quantity, is_free) VALUES (1, 1, false);
INSERT INTO hotel_rooms (number, room_quantity, is_free) VALUES (4, 1, false);
INSERT INTO hotel_rooms (number, room_quantity, is_free) VALUES (7, 2, false);
INSERT INTO hotel_rooms (number, room_quantity, is_free) VALUES (9, 3, false);
INSERT INTO hotel_rooms (number, room_quantity, is_free) VALUES (10, 4, false);


--
-- TOC entry 2134 (class 0 OID 24600)
-- Dependencies: 188
-- Data for Name: rents; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO rents (id, hotel_room_num, tenant_id) VALUES (1, 1, 1);
INSERT INTO rents (id, hotel_room_num, tenant_id) VALUES (2, 10, 4);
INSERT INTO rents (id, hotel_room_num, tenant_id) VALUES (3, 9, 3);
INSERT INTO rents (id, hotel_room_num, tenant_id) VALUES (4, 7, 2);
INSERT INTO rents (id, hotel_room_num, tenant_id) VALUES (5, 4, 5);
INSERT INTO rents (id, hotel_room_num, tenant_id) VALUES (6, 5, 3);


--
-- TOC entry 2144 (class 0 OID 0)
-- Dependencies: 187
-- Name: rents_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('rents_id_seq', 6, true);


--
-- TOC entry 2132 (class 0 OID 16390)
-- Dependencies: 186
-- Data for Name: tenants; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO tenants (id, name, last_name) VALUES (1, 'PETRO', 'PETROVICH');
INSERT INTO tenants (id, name, last_name) VALUES (2, 'IVAN', 'IVANOVICH');
INSERT INTO tenants (id, name, last_name) VALUES (3, 'VASYL', 'VASYLOVYCH');
INSERT INTO tenants (id, name, last_name) VALUES (4, 'KARIM', 'KARIMOVYCH');
INSERT INTO tenants (id, name, last_name) VALUES (5, 'DIMA', 'DIMOVYCH');


--
-- TOC entry 2011 (class 2606 OID 16389)
-- Name: hotel_rooms hotel_rooms_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY hotel_rooms
    ADD CONSTRAINT hotel_rooms_pkey PRIMARY KEY (number);


--
-- TOC entry 2013 (class 2606 OID 16394)
-- Name: tenants tenants_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tenants
    ADD CONSTRAINT tenants_pkey PRIMARY KEY (id);


--
-- TOC entry 2141 (class 0 OID 0)
-- Dependencies: 6
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2018-02-18 16:22:42

--
-- PostgreSQL database dump complete
--

