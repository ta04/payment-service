--
-- PostgreSQL database dump
--

-- Dumped from database version 11.7 (Ubuntu 11.7-1.pgdg18.04+1)
-- Dumped by pg_dump version 11.7 (Ubuntu 11.7-1.pgdg18.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: payments; Type: DATABASE; Owner: sleepingnext
--

CREATE DATABASE payments;

ALTER DATABASE payments OWNER TO sleepingnext;


--
-- Name: payments; Type: TABLE; Schema: public; Owner: sleepingnext
--

\connect payments;
CREATE TABLE public.payments (
    id bigint NOT NULL,
    order_id bigint NOT NULL,
    type character(255) NOT NULL,
    picture character varying(255) NULL,
    status character varying(255) NOT NULL
);


ALTER TABLE public.payments OWNER TO sleepingnext;

--
-- Name: payments_id_seq; Type: SEQUENCE; Schema: public; Owner: sleepingnext
--

CREATE SEQUENCE public.payments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.payments_id_seq OWNER TO sleepingnext;

--
-- Name: payments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sleepingnext
--

ALTER SEQUENCE public.payments_id_seq OWNED BY public.payments.id;


--
-- Name: payments id; Type: DEFAULT; Schema: public; Owner: sleepingnext
--

ALTER TABLE ONLY public.payments ALTER COLUMN id SET DEFAULT nextval('public.payments_id_seq'::regclass);


--
-- Data for Name: payments; Type: TABLE DATA; Schema: public; Owner: sleepingnext
--

COPY public.payments (id, order_id, type, status) FROM stdin;
\.


--
-- Name: payments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sleepingnext
--

SELECT pg_catalog.setval('public.payments_id_seq', 1, false);


--
-- Name: payments payments_pkey; Type: CONSTRAINT; Schema: public; Owner: sleepingnext
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

